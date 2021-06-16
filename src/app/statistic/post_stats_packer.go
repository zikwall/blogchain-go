package statistic

import (
	"context"
	"fmt"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/log"
	"sync"
	"time"
)

type Packaging interface {
	flatten() []interface{}
}

type WriteAPI interface {
	Write(Packaging)
	Close()
}

type Packer struct {
	Clickhouse *clickhouse.Clickhouse
	context    context.Context
}

func CreatePostStatisticPacker(ctx context.Context, clickhouse *clickhouse.Clickhouse) *Packer {
	ps := &Packer{}
	ps.context = ctx
	ps.Clickhouse = clickhouse

	return ps
}

func (pk *Packer) WriteAPI(table clickhouse.Table) WriteAPI {
	api := &writeApi{
		table:   table,
		context: pk.context,
		ch:      pk.Clickhouse,
		mu:      sync.RWMutex{},
	}

	defer func() {
		go api.schedulePackets()
	}()

	return api
}

type writeApi struct {
	table   clickhouse.Table
	context context.Context
	ch      *clickhouse.Clickhouse
	mu      sync.RWMutex
	packets [][]interface{}
}

func (wp *writeApi) Write(packer Packaging) {
	wp.mu.Lock()
	wp.packets = append(wp.packets, packer.flatten())
	wp.mu.Unlock()
}

func (wp *writeApi) Close() {
	if batches := wp.lookPackets(); len(batches) > 0 {
		_, _ = wp.ch.InsertWithMetrics(wp.table, batches)
	}
}

func (wp *writeApi) lookPackets() [][]interface{} {
	wp.mu.RLock()
	batchesSnapshot := wp.packets
	wp.mu.RUnlock()

	return batchesSnapshot
}

func (wp *writeApi) flushPackets() {
	wp.mu.Lock()
	wp.packets = wp.packets[:0]
	wp.mu.Unlock()
}

func (wp *writeApi) schedulePackets() {
	defer func() {
		wp.Close()
		log.Info(fmt.Sprintf("packer (%s) scheduler is stopping", wp.table.Name))
	}()

schedule:
	select {
	case <-wp.context.Done():
		return
	case <-time.After(5 * time.Second):
	}

	log.Info(fmt.Sprintf("packer (%s) scheduler is running", wp.table.Name))

	if batches := wp.lookPackets(); len(batches) > 0 {
		wp.flushPackets()

		if affected, err := wp.ch.InsertWithMetrics(wp.table, batches); err != nil {
			log.Warning(err)
		} else {
			log.Info(fmt.Sprintf(
				"packer (%s) count records: %d | affected records: %d", wp.table.Name, len(batches), affected),
			)
		}
	} else {
		log.Info(fmt.Sprintf("packer (%s) nothing to write", wp.table.Name))
	}

	goto schedule
}

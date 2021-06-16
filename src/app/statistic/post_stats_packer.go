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

type Packer struct {
	Clickhouse *clickhouse.Clickhouse
	context    context.Context
	mu         sync.RWMutex
	packets    [][]interface{}
}

func CreatePostStatisticPacker(ctx context.Context, clickhouse *clickhouse.Clickhouse) *Packer {
	ps := &Packer{}
	ps.mu = sync.RWMutex{}
	ps.context = ctx
	ps.Clickhouse = clickhouse

	go ps.schedulePackets()

	return ps
}

func (ps *Packer) NonBlockingWritePackets(packers ...Packaging) {
	packs := make([][]interface{}, 0, len(packers))

	for _, pkt := range packers {
		packs = append(packs, pkt.flatten())
	}

	ps.mu.Lock()
	ps.packets = append(ps.packets, packs...)
	ps.mu.Unlock()
}

func (ps *Packer) lookPackets() [][]interface{} {
	ps.mu.RLock()
	batchesSnapshot := ps.packets
	ps.mu.RUnlock()

	return batchesSnapshot
}

func (ps *Packer) flushPackets() {
	ps.mu.Lock()
	ps.packets = ps.packets[:0]
	ps.mu.Unlock()
}

func (ps *Packer) schedulePackets() {
	defer func() {
		log.Info("Stop post statistic packer scheduler")
	}()

schedule:
	select {
	case <-ps.context.Done():
		return
	case <-time.After(5 * time.Second):
	}

	log.Info("The post statistic packer scheduler is running")

	if batches := ps.lookPackets(); len(batches) > 0 {
		ps.flushPackets()

		if affected, err := ps.Clickhouse.InsertWithMetrics(postStatsTable, batches); err != nil {
			log.Warning(err)
		} else {
			log.Info(fmt.Sprintf("[CLICKHOUSE] Count records: %d | Affected records: %d", len(batches), affected))
		}
	} else {
		log.Info("[CLICKHOUSE] Nothing to send to Clickhouse")
	}

	goto schedule
}

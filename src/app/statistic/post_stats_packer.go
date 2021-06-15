package statistic

import (
	"context"
	"fmt"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/log"
	"sync"
	"time"
)

type PostStatisticPacker struct {
	Clickhouse *clickhouse.Clickhouse
	context    context.Context
	mu         sync.RWMutex
	batches    []PostStats
}

func CreatePostStatisticPacker(ctx context.Context, clickhouse *clickhouse.Clickhouse) *PostStatisticPacker {
	ps := &PostStatisticPacker{}
	ps.mu = sync.RWMutex{}
	ps.batches = []PostStats{}
	ps.context = ctx
	ps.Clickhouse = clickhouse

	go ps.schedule()

	return ps
}

func (ps *PostStatisticPacker) AppendRecords(stats ...PostStats) {
	ps.mu.Lock()
	ps.batches = append(ps.batches, stats...)
	ps.mu.Unlock()
}

func (ps *PostStatisticPacker) all() []PostStats {
	ps.mu.RLock()
	batchesSnapshot := ps.batches
	ps.mu.RUnlock()

	return batchesSnapshot
}

func (ps *PostStatisticPacker) flush() {
	ps.mu.Lock()
	ps.batches = ps.batches[:0]
	ps.mu.Unlock()
}

func (ps *PostStatisticPacker) schedule() {
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

	if batches := ps.all(); len(batches) > 0 {
		ps.flush()

		rows := make([][]interface{}, 0, len(batches))

		for _, batch := range batches {
			rows = append(rows, batch.flatten())
		}

		if affected, err := ps.Clickhouse.InsertWithMetrics(postStatsTable, rows); err != nil {
			log.Warning(err)
		} else {
			log.Info(fmt.Sprintf("[CLICKHOUSE] Count records: %d | Affected records: %d", len(rows), affected))
		}
	} else {
		log.Info("[CLICKHOUSE] Nothing to send to Clickhouse")
	}

	goto schedule
}

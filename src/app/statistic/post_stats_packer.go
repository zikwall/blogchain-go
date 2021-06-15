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
	cb := &PostStatisticPacker{}
	cb.mu = sync.RWMutex{}
	cb.batches = []PostStats{}
	cb.context = ctx
	cb.Clickhouse = clickhouse

	go cb.schedule()

	return cb
}

func (cb *PostStatisticPacker) AppendRecords(stats ...PostStats) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.batches = append(cb.batches, stats...)
}

func (cb *PostStatisticPacker) all() []PostStats {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	return cb.batches
}

func (cb *PostStatisticPacker) flush() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.batches = []PostStats{}
}

func (cb *PostStatisticPacker) schedule() {
	defer func() {
		log.Info("Stop post statistic packer scheduler")
	}()

schedule:
	select {
	case <-cb.context.Done():
		return
	case <-time.After(5 * time.Second):
	}

	log.Info("The post statistic packer scheduler is running")

	batches := cb.all()
	if len(batches) > 0 {
		cb.flush()

		rows := make([][]interface{}, 0, len(batches))

		for _, batch := range batches {
			rows = append(rows, batch.flatten())
		}

		if affected, err := cb.Clickhouse.InsertWithMetrics(postStatsTable, rows); err != nil {
			log.Warning(err)
		} else {
			log.Info(fmt.Sprintf("[CLICKHOUSE SCHEDULER] Count records: %d | Affected records: %d", len(rows), affected))
		}
	} else {
		log.Info("[CLICKHOUSE SCHEDULER] Nothing to send to Clickhouse")
	}

	goto schedule
}

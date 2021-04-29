package statistic

import (
	"context"
	"fmt"
	"github.com/zikwall/blogchain/src/platform/clickhouse"
	"github.com/zikwall/blogchain/src/platform/log"
	"sync"
	"time"
)

type (
	ClickhouseBatcher struct {
		clickhouse *clickhouse.Clickhouse
		context    context.Context
		mu         sync.RWMutex
		batches    []PostStats
	}
)

func CreateClickhouseBatcher(ctx context.Context, clickhouse *clickhouse.Clickhouse) *ClickhouseBatcher {
	cb := &ClickhouseBatcher{}
	cb.mu = sync.RWMutex{}
	cb.batches = []PostStats{}
	cb.context = ctx
	cb.clickhouse = clickhouse

	go cb.schedule()

	return cb
}

func (cb *ClickhouseBatcher) AppendRecords(stats ...PostStats) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.batches = append(cb.batches, stats...)
}

func (cb *ClickhouseBatcher) all() []PostStats {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	return cb.batches
}

func (cb *ClickhouseBatcher) flush() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.batches = []PostStats{}
}

func (cb *ClickhouseBatcher) schedule() {
	defer func() {
		log.Info("Stop clickhouse batcher scheduler")
	}()

schedule:
	select {
	case <-cb.context.Done():
		return
	case <-time.After(5 * time.Second):
	}

	log.Info("The clickhouse batcher scheduler is running")

	batches := cb.all()
	if len(batches) > 0 {
		cb.flush()

		rows := make([][]interface{}, 0, len(batches))

		for _, batch := range batches {
			rows = append(rows, batch.flatten())
		}

		if affected, err := cb.clickhouse.InsertWithMetrics(postStatsTable, rows); err != nil {
			log.Warning(err)
		} else {
			log.Info(fmt.Sprintf("[CLICKHOUSE SCHEDULER] Count records: %d | Affected records: %d", len(rows), affected))
		}
	} else {
		log.Info("[CLICKHOUSE SCHEDULER] Nothing to send to Clickhouse")
	}

	goto schedule
}

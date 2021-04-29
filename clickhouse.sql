CREATE TABLE post_stats (
    post_id UInt64,
    owner_id UInt64,
    hostname String,
    os String,
    browser String,
    platform String,
    ip String,
    country String,
    region String,
    insert_ts DateTime,
    date Date
) ENGINE MergeTree() PARTITION BY toYYYYMM(date) ORDER BY (post_id, date) SETTINGS index_granularity=8192

-- ENGINE = ReplicatedMergeTree('/clickhouse/tables/blogchain/{shard}/post_stats_sharded', '{replica}', post_id, (owner_id, date), 8192)
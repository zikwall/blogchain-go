CREATE TABLE post_stats
(
    post_id   UInt64,
    owner_id  UInt64,
    hostname  String,
    os        String,
    browser   String,
    platform  String,
    ip        String,
    country   String,
    region    String,
    insert_ts DateTime,
    date      Date
) ENGINE MergeTree() PARTITION BY toYYYYMM(date) ORDER BY (post_id, date) SETTINGS index_granularity=8192

-- ENGINE = ReplicatedMergeTree('/clickhouse/tables/blogchain/{shard}/post_stats_sharded', '{replica}', post_id, (owner_id, date), 8192)

CREATE
MATERIALIZED VIEW post_stats_consumer TO post_stats_views
    AS
SELECT post_id, count() as views
FROM post_stats
GROUP BY post_id;

CREATE TABLE post_stats_views
(
    post_id UInt64,
    views   UInt64
) ENGINE = SummingMergeTree()
ORDER BY post_id

-- INSERT INTO post_stats_views Values(10,1),(11,2),(10,1)
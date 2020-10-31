
-- +migrate Up
alter table content modify column content MEDIUMTEXT not null;

-- +migrate Down
alter table content modify column content TEXT not null;
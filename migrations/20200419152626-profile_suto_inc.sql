
-- +migrate Up
alter table profile
add column `id` int NOT NULL primary key AUTO_INCREMENT first;

-- +migrate Down


-- +migrate Up
alter table content modify column annotation text not null;

-- +migrate Down
alter table content modify column annotation varchar(255) not null;
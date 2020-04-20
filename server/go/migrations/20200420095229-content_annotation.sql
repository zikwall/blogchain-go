
-- +migrate Up
alter table content add column annotation varchar(255) not null after title;

-- +migrate Down
alter table content drop column annotation;
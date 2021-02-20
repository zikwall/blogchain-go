
-- +migrate Up
alter table content add column uuid varchar(36) not null after id;
update content set uuid=uuid() where uuid="";

-- +migrate Down
alter table content drop column uuid;
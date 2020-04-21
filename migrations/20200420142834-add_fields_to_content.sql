
-- +migrate Up
alter table content
add column created_at int(11) not null,
add column updated_at int(11) null,
add column image varchar(255) null;

update content set created_at=UNIX_TIMESTAMP();

-- +migrate Down
alter table content
drop column created_at,
drop column updated_at,
drop column image;
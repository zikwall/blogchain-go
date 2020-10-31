
-- +migrate Up
alter table profile
add column location varchar(255) null,
add column status varchar(50) null,
add column description varchar(255) null;

-- +migrate Down
alter table profile
drop column location;
drop column status,
drop column description;
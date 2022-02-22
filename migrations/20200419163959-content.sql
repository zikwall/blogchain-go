
-- +migrate Up
create table content (
    id int primary key not null AUTO_INCREMENT,
    user_id int not null,
    title varchar(200) not null,
    content text not null
);

-- +migrate Down
drop table content;

-- +migrate Up
create table tags (
    id int(11) primary key not null AUTO_INCREMENT,
    name varchar(35) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table content_tag (
    id int(11) primary key not null AUTO_INCREMENT,
    content_id int(11) not null,
    tag_id int(11) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

insert into tags (name)
values
('Программирование'), ('Алгоритмы'), ('Обработка изображений'), ('Алгоритмы'),
('Управление проектами'), ('Робототехника'), ('Управление разработкой'), ('Системное администрирование'),
('Open source'), ('Интернет вещей'), ('Графичекий дизайн'), ('Компьютерное железо'), ('Сетевые технологии');

-- +migrate Down
drop table tags;
drop table content_tag;
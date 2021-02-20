
-- +migrate Up
alter table tags add column `label` varchar(35) not null;
update tags set `label`='unlabeled';

-- +migrate Down
alter table tags drop column `label`;
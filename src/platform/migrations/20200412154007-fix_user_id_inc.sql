
-- +migrate Up
ALTER TABLE `user` MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

-- +migrate Down

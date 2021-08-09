use backend_db;

CREATE TABLE `people` (
	`id` SERIAL NOT NULL,
	`datetime` DATETIME NOT NULL,
	`people_count` INT NOT NULL,
	`holiday` TINYINT NOT NULL,
	PRIMARY KEY (`id`)
);
-- id(int primary key), 日時(datetime), 人数(int), 祝日かどうか(tidyint)
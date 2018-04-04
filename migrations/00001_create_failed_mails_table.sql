-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `failed_mails` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `action` varchar(80) NOT NULL,
  `payload` json NOT NULL,
  `reason` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `failed_mails`;

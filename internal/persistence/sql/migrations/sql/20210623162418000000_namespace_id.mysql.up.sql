CREATE TABLE `keto_namespace_ids` (
`id` char(36) NOT NULL,
`nid` char(36) NOT NULL,
`serial_id` INTEGER NOT NULL,
`created_at` DATETIME NOT NULL,
`updated_at` DATETIME NOT NULL,
PRIMARY KEY(`nid`, `id`)
) ENGINE=InnoDB;
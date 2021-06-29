CREATE TABLE `keto_networks` (
`network_id` char(36) NOT NULL,
PRIMARY KEY(`network_id`),
`created_at` DATETIME NOT NULL,
`updated_at` DATETIME NOT NULL
) ENGINE=InnoDB;
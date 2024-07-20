-- Create "users" table
CREATE TABLE `users` (`id` varchar(36) NOT NULL DEFAULT "UUID()", `username` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

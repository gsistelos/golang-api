-- Create "posts" table
CREATE TABLE `posts` (`id` varchar(36) NOT NULL DEFAULT "UUID()", `content` varchar(255) NOT NULL, `visibility` varchar(255) NOT NULL DEFAULT "public", `user_id` varchar(36) NOT NULL, PRIMARY KEY (`id`), INDEX `user_id` (`user_id`), CONSTRAINT `user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

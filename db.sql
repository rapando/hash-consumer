CREATE TABLE `md5` (
  `raw_password` varbinary(512) NOT NULL UNIQUE,
  `hashed_password` varbinary(512) NOT NULL,
  `created` datetime(6) NOT NULL DEFAULT current_timestamp(6)
);
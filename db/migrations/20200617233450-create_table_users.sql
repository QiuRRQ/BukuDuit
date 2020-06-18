
-- +migrate Up
CREATE TABLE IF NOT  EXISTS "users" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "mobile_phone" varchar(12) NOT NULL,
  "pin" varchar(128),
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

-- +migrate Down

DROP TABLE IF EXISTS  "users";

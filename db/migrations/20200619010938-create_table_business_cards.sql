
-- +migrate Up
CREATE TABLE IF NOT EXISTS "business_cards" (
                                  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
                                  "full_name" varchar(50) NOT NULL,
                                  "book_name" varchar(50),
                                  "mobile_phone" varchar(12),
                                  "tag_line" varchar(100),
                                  "address" text,
                                  "email" varchar(50),
                                  "avatar" char(36),
                                  "user_id" char(36) NOT NULL,
                                  "created_at" timestamp,
                                  "updated_at" timestamp,
                                  "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "business_cards";
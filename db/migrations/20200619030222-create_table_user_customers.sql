
-- +migrate Up

CREATE TABLE IF NOT EXISTS "user_customers" (
                                  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
                                  "full_name" varchar(50),
                                  "mobile_phone" varchar(20),
                                  "business_id" char(36) NOT NULL,
                                  "debt" int4 DEFAULT 0,
                                  "payment_date" date,
                                  "created_at" timestamp,
                                  "updated_at" timestamp,
                                  "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS  "user_customers";
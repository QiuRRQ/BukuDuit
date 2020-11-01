CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "authors"(
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(255) NULL,
  "address" VARCHAR(255) NULL,
  "city" VARCHAR(255) NULL,
  "province" VARCHAR(255) NULL,
  "postal_code" VARCHAR(255) NULL,
  "no_telp" VARCHAR(255) NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
  
  DROP TABLE IF EXISTS "books" ;

CREATE TABLE IF NOT EXISTS "books" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "title" VARCHAR(255) NULL,
  "publisher_id" char(36) NOT NULL,
  "authors_id" char(36) NOT NULL,
  "book_img" VARCHAR(1000) NULL,
  "stock" INT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
	
	DROP TABLE IF EXISTS "borrows_card" ;

CREATE TABLE IF NOT EXISTS "borrows_card" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "members_id" char(36) NOT NULL,
  "trans_date" VARCHAR(45) NULL,
  "trans_month" VARCHAR(45) NULL,
  "trans_year" VARCHAR(45) NULL,
  "status" VARCHAR(45) NULL,
  "books_id" char(36) NOT NULL,
  "jumlah" INT NULL,
"borrow_done" char(1) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
	
	DROP TABLE IF EXISTS "buku_has_category" ;

CREATE TABLE IF NOT EXISTS "buku_has_category" (
"id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "buku_id" char(36) NOT NULL,
  "category_id" char(36) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
	
	DROP TABLE IF EXISTS "category" ;

CREATE TABLE IF NOT EXISTS "category" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(255) NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
  
  DROP TABLE IF EXISTS "members" ;

CREATE TABLE IF NOT EXISTS "members" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "no_member" VARCHAR(45) NULL,
  "name" VARCHAR(45) NULL,
  "no_telp" VARCHAR(45) NULL,
  "address" VARCHAR(45) NULL,
  "city" VARCHAR(45) NULL,
  "province" VARCHAR(45) NULL,
  "member_img" VARCHAR(1000) NULL,
  "gender" CHAR(1) NULL,
  "birth_date" VARCHAR(45) NULL,
  "birth_month" VARCHAR(45) NULL,
  "birth_year" VARCHAR(45) NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
  
  DROP TABLE IF EXISTS "publishers" ;

CREATE TABLE IF NOT EXISTS "publishers" (
  "id" char(36) PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "name" VARCHAR(255) NULL,
  "address" VARCHAR(255) NULL,
  "city" VARCHAR(255) NULL,
  "province" VARCHAR(255) NULL,
  "postal_code" VARCHAR(255) NULL,
  "no_telp" VARCHAR(255) NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp);
  
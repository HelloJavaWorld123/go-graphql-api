/*
 Navicat Premium Data Transfer

 Source Server         : local-postgresql
 Source Server Type    : PostgreSQL
 Source Server Version : 120002
 Source Host           : 172.16.20.195:5432
 Source Catalog        : postgres
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 120002
 File Encoding         : 65001

 Date: 21/04/2020 10:24:12
*/


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "age" int4,
  "profession" varchar(50) COLLATE "pg_catalog"."default",
  "friendly" bool
)
;
COMMENT ON COLUMN "public"."users"."id" IS '主键';
COMMENT ON COLUMN "public"."users"."name" IS '名称';
COMMENT ON COLUMN "public"."users"."age" IS '年龄';
COMMENT ON COLUMN "public"."users"."profession" IS '职业';

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 'kevin', 35, 'waiter', 't');
INSERT INTO "public"."users" VALUES (2, 'angela', 21, 'concierge', 't');
INSERT INTO "public"."users" VALUES (3, 'alex', 26, 'zoo keeper', 'f');
INSERT INTO "public"."users" VALUES (4, 'becky', 67, 'retired', 'f');
INSERT INTO "public"."users" VALUES (5, 'kevin', 15, 'in school', 't');
INSERT INTO "public"."users" VALUES (6, 'frankie', 45, 'teller', 't');

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

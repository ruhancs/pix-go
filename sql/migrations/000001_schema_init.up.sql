CREATE TABLE "accounts" (
  "id" varchar PRIMARY KEY,
  "owner_name" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "number" varchar NOT NULL,
  "bank_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "banks" (
  "id" varchar PRIMARY KEY,
  "code" varchar NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" varchar PRIMARY KEY,
  "account_from_id" varchar NOT NULL,
  "pix_key_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "status" varchar NOT NULL,
  "description" varchar NOT NULL,
  "cancel_description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pix_keys" (
  "id" varchar PRIMARY KEY,
  "kind" varchar NOT NULL,
  "key" varchar NOT NULL,
  "status" varchar NOT NULL,
  "account_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("bank_id") REFERENCES "banks" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("pix_key_id") REFERENCES "pix_keys" ("id");
ALTER TABLE "transactions" ADD FOREIGN KEY ("account_from_id") REFERENCES "accounts" ("id");

ALTER TABLE "pix_keys" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
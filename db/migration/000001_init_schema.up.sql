CREATE TABLE account (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "country_code" int NOT NULL
);

CREATE TABLE entries (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "ammount" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE transfers (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "ammount" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);



CREATE INDEX ON account ("owner");

CREATE INDEX ON entries ("account_id");

CREATE INDEX ON transfers ("from_account_id");

CREATE INDEX ON transfers ("to_account_id");

CREATE INDEX ON transfers ("from_account_id", "to_account_id");

COMMENT ON COLUMN Entries."ammount" IS 'this can be negative or positive';

COMMENT ON COLUMN Transfers."ammount" IS 'this should always be positive';

ALTER TABLE entries ADD FOREIGN KEY ("account_id") REFERENCES account ("id");

ALTER TABLE transfers ADD FOREIGN KEY ("from_account_id") REFERENCES account ("id");

ALTER TABLE transfers ADD FOREIGN KEY ("to_account_id") REFERENCES account ("id");

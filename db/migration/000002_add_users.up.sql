CREATE TABLE users (
  "username" varchar PRIMARY KEY,
  "hashedpassword" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001.01.01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE account ADD FOREIGN KEY ("owner") REFERENCES users ("username");
ALTER TABLE account ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");
CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "posts" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "author" BIGSERIAL  NOT NULL
);

CREATE INDEX ON "accounts" ("owner");

ALTER TABLE "posts" ADD FOREIGN KEY ("author") REFERENCES "accounts" ("id");
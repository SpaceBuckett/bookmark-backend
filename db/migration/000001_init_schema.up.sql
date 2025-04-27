CREATE TABLE "userprofile" (
                        "id" bigserial PRIMARY KEY,
                        "username" varchar NOT NULL,
                        "email" varchar UNIQUE NOT NULL,
                        "hashed_password" varchar NOT NULL,
                        "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookmark" (
                            "id" bigserial PRIMARY KEY,
                            "owner_id" bigserial NOT NULL,
                            "title" varchar NOT NULL,
                            "url" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "bookmark" ("owner_id");

CREATE UNIQUE INDEX ON "bookmark" ("owner_id", "title");

ALTER TABLE "bookmark" ADD FOREIGN KEY ("owner_id") REFERENCES "userprofile" ("id") ON DELETE CASCADE;
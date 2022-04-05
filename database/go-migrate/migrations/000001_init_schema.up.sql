CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "brands" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);
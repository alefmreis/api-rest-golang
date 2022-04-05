-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE "brands" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS brands;
-- +goose StatementEnd

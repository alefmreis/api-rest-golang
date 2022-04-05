-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE brands ADD COLUMN status int DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE brands DROP COLUMN IF EXISTS status;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN public_key VARCHAR(255) DEFAULT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN public_key;

-- +goose StatementEnd

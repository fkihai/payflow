-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE IF EXISTS transactions
ADD COLUMN qr_url TEXT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS transactions
DROP COLUMN qr_url;
-- +goose StatementEnd

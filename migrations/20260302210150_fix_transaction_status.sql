-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions
DROP CONSTRAINT transactions_status_check;

ALTER TABLE transactions
ADD CONSTRAINT transactions_status_check
CHECK (status IN ('pending','settlement','expire','refund'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
DROP CONSTRAINT transactions_status_check;

ALTER TABLE transactions
ADD CONSTRAINT transactions_status_check
CHECK (status IN ('pending','settlement','expired','refund'));
-- +goose StatementEnd

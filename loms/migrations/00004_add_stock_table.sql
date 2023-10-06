-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock
(
  sku bigserial PRIMARY KEY,
  count bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock CASCADE;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock 
(
  sku bigint,
  count bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock CASCADE;
-- +goose StatementEnd

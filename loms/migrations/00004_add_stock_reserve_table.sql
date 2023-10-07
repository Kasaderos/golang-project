-- +goose Up
-- +goose StatementBegin
CREATE TABLE reserved_stock 
(
  user_id bigint,
  sku bigint,
  count bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reserved_stock CASCADE;
-- +goose StatementEnd

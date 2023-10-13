-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_item
(
  order_id bigint REFERENCES orders(id),
  sku bigint,
  count int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_item CASCADE;
-- +goose StatementEnd

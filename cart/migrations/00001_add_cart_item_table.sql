-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_item
(
  user_id     bigint,
  sku         bigint,
  price       int,
  amount      int,
  created_at  timestamptz default now() not null,
  constraint  id primary key (user_id, sku)
);

CREATE INDEX idx_user_id ON cart_item(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_item CASCADE;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock 
(
  sku            bigint           not null,
  available      bigint           not null   check (available >= 0),
  total_reserved bigint default 0 not null   check (total_reserved >= 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock CASCADE;
-- +goose StatementEnd

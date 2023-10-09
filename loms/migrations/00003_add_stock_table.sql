-- +goose Up
-- +goose StatementBegin
CREATE TABLE stock 
(
  sku      bigint not null,
  count    bigint not null,
  reserved bigint default 0 not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stock CASCADE;
-- +goose StatementEnd

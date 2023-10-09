-- name: CreateOrder :one
INSERT INTO orders (
    user_id,
    status_id
) VALUES ($1, $2)
RETURNING id;

-- name: AddOrderItem :exec
INSERT INTO order_item (
    order_id,
    sku,
    count
) VALUES ($1, $2, $3);

-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: SetStatus :exec
update orders 
set status_id = $1
where id = $2;

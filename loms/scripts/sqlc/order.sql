-- name: CreateOrder :one
INSERT INTO user_order (
    user_id,
    status
) VALUES ($1, $2)
RETURNING id;

-- name: AddOrderItem :exec
INSERT INTO order_item (
    order_id,
    sku,
    count
) VALUES ($1, $2, $3);

-- name: GetOrderByID :one
SELECT * FROM user_order WHERE id = $1;

-- name: SetStatus :exec
update user_order 
set status = $1
where id = $2;
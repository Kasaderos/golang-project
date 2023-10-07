-- name: AddCartItem :exec
insert into cart_item (
    user_id,
    sku,
    name,
    price,
    count
) VALUES ($1, $2, $3, $4, $5);

-- name: DeleteItem :exec
delete from cart_item 
where user_id = $1 and sku = $2;

-- name: DeleteItemByUser :exec
delete from cart_item 
where user_id = $1;

-- name: GetItemsByUserID :many
select * from cart_item 
where user_id = $1;


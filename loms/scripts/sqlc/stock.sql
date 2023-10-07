-- name: ReserveStock :exec
insert into reserved_stock (
    user_id,
    sku,
    count
) VALUES ($1, $2, $3);

-- name: GetReservedStockByUsedID :many
select * from reserved_stock
where user_id = $1;

-- name: RemoveStocks :exec
update stock
set count = count - $1
where sku = $2;

-- name: DeleteReservedStockByUserID :exec
delete from reserved_stock 
where user_id = $1;

-- name: CountStocksBySKU :one
select count from stock
where sku = $1;

-- name: CountReservedStocksBySKU :one
select sum(count) from reserved_stock
where sku = $1;
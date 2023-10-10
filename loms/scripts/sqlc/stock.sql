-- name: ReserveStock :one
update stock 
set available = available - $1,
    total_reserved = total_reserved + $1
where sku = $2
returning available;

-- name: ReserveRemove :exec
update stock
set total_reserved = total_reserved - $1
where sku = $2;

-- name: ReserveCancel :exec
update stock 
set available = available + $1,
    total_reserved = total_reserved - $1
where sku = $2;

-- name: GetBySKU :one
select available - total_reserved from stock
where sku = $1;

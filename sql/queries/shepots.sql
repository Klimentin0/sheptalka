-- name: CreateShepot :one
INSERT INTO shepots(id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteAllShepots :exec
DELETE FROM users;

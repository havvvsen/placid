-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    email = $1
LIMIT 1;

-- name: AddUser :exec
INSERT INTO users (email, password_hash)
    VALUES ($1, $2);

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;

-- name: UpdateUserPasswordHash :exec
UPDATE
    users
SET
    password_hash = $2
WHERE
    email = $1;

-- name: UpdateUserEmail :exec
UPDATE
    users
SET
    email = $2
WHERE
    uuid = $1;

-- name: UpdateUserAccountLevel :exec
UPDATE
    users
SET
    is_admin = $2
WHERE
    uuid = $1;

-- name: UpdateUserSubscription :exec
UPDATE
    users
SET
    is_premium = $2
WHERE
    uuid = $1;


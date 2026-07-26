-- name: AddNewsletterSubsciber :exec
INSERT INTO newsletter (email)
    VALUES ($1);

-- name: GetNewsletterSubscribers :many
SELECT
    *
FROM
    newsletter;

-- name: GetNewsletterSubscriber :one
SELECT
    *
FROM
    newsletter
WHERE
    email = $1;

-- name: DeleteNewsletterSubscriber :exec
DELETE FROM newsletter
WHERE email = $1;


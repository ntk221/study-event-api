-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY date;

-- name: CreateEvent :one
INSERT INTO events (
  name, date
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateEvent :one
UPDATE events
SET name = $2, date = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;

-- name: ListEventsByDateRange :many
SELECT * FROM events
WHERE date BETWEEN $1 AND $2
ORDER BY date;

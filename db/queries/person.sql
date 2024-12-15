-- name: CreatePerson :one
INSERT INTO people (name)
VALUES ($1)
RETURNING *;

-- name: GetPerson :one
SELECT *
FROM people
WHERE id = $1;

-- name: ListPeople :many
SELECT *
FROM people;

-- name: UpdatePerson :exec
UPDATE people
SET name = $2
WHERE id = $1;

-- name: DeletePerson :exec
DELETE FROM people
WHERE id = $1;

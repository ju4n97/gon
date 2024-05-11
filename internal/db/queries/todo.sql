-- name: GetTodo :one
SELECT *
FROM todo
where id = $1 
LIMIT 1;

-- name: ListTodos :many
SELECT *
FROM todo
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTodo :one
INSERT INTO todo
(title, is_completed)
VALUES
($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todo
SET title = $1, is_completed = $2
WHERE id = $3
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todo
WHERE id = $1
RETURNING *;

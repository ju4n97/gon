-- name: GetTodo :one
SELECT *
FROM todos 
where id = $1 
LIMIT 1;

-- name: ListTodos :many
SELECT *
FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTodo :one
INSERT INTO todos
(title, completed)
VALUES
($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET title = $1, completed = $2
WHERE id = $3
RETURNING *;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;

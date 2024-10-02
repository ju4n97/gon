-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo;
-- +goose StatementEnd

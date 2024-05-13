## Prerequisites

* [Go 1.22](https://go.dev/dl/)
* [Docker](https://www.docker.com/get-started/)
* [Task](https://taskfile.dev/installation/)

> [!NOTE]
> For the `db:docs` task, which generates database documentation automatically, you'll need [Node.js](https://nodejs.org/) or [Bun](https://bun.sh/). The `@dbml/cli` package is required for this task, but there isn't an official Docker image available yet. Once one is released, this project will be updated to remove the dependency on a JS runtime. [Progress here](https://github.com/dbml/cli/issues/123).

## Getting started

1. Clone repository

## Getting started (Contributing to this project)

1. Fork repository

## Migrations workflow

This project uses [Goose](https://github.com/pressly/goose) for migrations. The most commonly used commands are set up using [Task](https://taskfile.dev/installation/) for ease of use.

1. **Create a new migration file**: Use the task `migrate:create` to generate a new migration under `./internal/db/migrations` and pass the name of the migration as an argument of the task. Replace `your_migration_name` with the name of your migration.

    ```sh
    task migrate:create -- your_migration_name
    ```

2. **Write the migration and rollback SQL code:** In the newly generated SQL file, write the SQL code to alter the database, as well as the corresponding SQL code to rollback those changes. For example:

    ```sql
    -- +goose Up
    -- +goose StatementBegin
    CREATE TABLE my_new_table (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )
    -- +goose StatementEnd

    -- +goose Down
    -- +goose StatementBegin
    DROP TABLE my_new_table;
    -- +goose StatementEnd
    ```

    > [!TIP]
    > After creating your migration file, you might want to make sure the file is not malformed. You can use the task `migrate:validate` to check the file without running it.

3. **Run the migration:** After creating the migration file and writing the SQL code, you can run the migration using the following command:

    ```sh
    task migrate:up
    ```

    > [!NOTE]
    > Alternatively, you can simply run the project with `task dev`, and all the available migrations will be applied automatically before starting the HTTP server.

4. **Generate Go code with SQLC** Once the database schema is updated, you can use SQLC to generate Go code from the migrations directory. This will allow you to interact with your database using strong types.

    ```sh
    task sqlc:generate
    ```

5. **Rollback the migration if needed**: if you need to rollback the changes made by the migration, you can use the following command:

    ```sh
    task migrate:down
    ```

## License

[MIT](./LICENSE)

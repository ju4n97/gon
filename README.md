> [!IMPORTANT]
> This project is currently in a very early stage of development. As such, features and documentation are subject to change.

## Contributing

<details>
    <summary>Local development</summary>

- Clone this repository.
- Install [Go 1.22 or later](https://go.dev/dl/).
- Install [Task](https://taskfile.dev/installation/).
- Create a `.env` file based on the `env.example` file and fill in the variables.
- Install the project's dependencies with `task install`.
- Run:
  - `task dev` to start the development HTTP server.
  - `task test` to run the unit tests.
  - `task lint` to run the linter.  

</details>

Refer to the [contributor's guide](CONTRIBUTING.md) for more in-depth
information.

## License

[MIT](LICENSE)

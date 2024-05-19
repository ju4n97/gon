<!-- markdownlint-disable MD041 -->
<!-- markdownlint-disable MD013 -->

<!-- markdownlint-disable MD033 -->
<a href="https://gon.github.io">
  <img src="./.github/assets/banner.png" alt="Gon banner" width="100%" height="200" style="object-fit: cover; margin-bottom: 1rem" />
</a>

<center>
   <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/jm2097/gon/ci.yaml">
   <img alt="GitHub License" src="https://img.shields.io/github/license/jm2097/gon">

   <table>
      <tbody>
         <tr>
            <td>
               <a href="https://gon.github.io">📚 Documentation</a>
            </td>
            <td>
               <a href="https://gon.github.io">🗂️ Open API spec</a>
            </td>
            <td>
               <a href="https://gon.github.io">🛢️ Database documentation</a>
            </td>
         </tr>
      </tbody>
   </table>
</center>
<!-- markdownlint-enable MD033 -->

## What is Gon?

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

## Prerequisites for development

Before you begin, ensure you have installed the following:

- [Go 1.22 or later](https://go.dev/dl/): Go is used to implement the logic of the project.
- [Task](https://taskfile.dev/installation/): Task is a cross-platform alternative to [Make](https://www.gnu.org/software/make/manual/make.html) for managing project tasks. It's used to easily run the project, generate migrations, execute linters, and more. ([See available tasks](./Taskfile.yaml)).
- [Bun](https://bun.sh/): Bun is used to run the documentation server. Furthermore, some tasks, such as `db:docs`, might require a JavaScript runtime, as they do not yet have official Docker images.

### Optional dependencies

- [Docker](https://www.docker.com/get-started/): If you choose to use Docker, the project relies on the [`docker-compose.yaml`](./docker-compose.yaml) file at the root of the project directory to manage services. This file defines the required services, such as PostgreSQL, caching, and others. If you prefer to install dependencies directly on your system, you can still refer to this file to know the required services for installation.

## Setup development environment

1. Clone the repository to your local machine and change to the project's root directory:

   ```sh
   git clone https://github.com/jm2097/gon && cd gon
   ```

2. Create a `.env` file based on the `env.example` file and fill in the required variables:

   ```sh
   cp .env.example .env
   ```

3. Install all the dependencies and development tools for the project:

   ```sh
   task
   ```

4. Spin up Docker containers for services that the project relies upon:

   ```sh
   docker compose up -d
   ```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute.

## License

[MIT](./LICENSE)

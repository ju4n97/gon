[![Cybernetically enhanced web apps: kit](https://sveltejs.github.io/assets/banner.png)](https://svelte.dev)

[![license](https://img.shields.io/npm/l/svelte.svg?style=flat-square)](LICENSE.md)
![GitHub package.json version](https://img.shields.io/github/package-json/v/MesaTechLabs/kitten?style=flat-square)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/MesaTechLabs/kitten/ci.yml?style=flat-square&label=CI)

- [🌐 Visit live demo](#)
- [📚 Read the docs](#)

## What is this project?

This is a decidedly opinionated SvelteKit template, packed with modern technologies and industry best practices. It offers a comprehensive array of features with the main goal of saving a lot of time when configuring a brand new project, removing all the tedious boilerplate that comes at the start and allowing developers to focus on building the core features of their projects. Its flexible design and modular approach make it easy to customize for various project requirements.

## Technologies & Features

- 🧡 [SvelteKit](https://kit.svelte.dev) and [Svelte 5](https://svelte.dev/blog/runes).
- 🍞 [Bun](https://bun.sh) as the all-in-one runtime, package manager, and test runner.
- 🌐 [ParaglideJS](https://inlang.com/m/gerre34r/library-inlang-paraglideJs) for internationalization and localization (see [i18n section](#i18n)).
- 🛢️ [DrizzleORM](https://orm.drizzle.team) as the database bridge, supporting multiple SQL systems out of the box:
  - [PostgreSQL (default)](src/lib/server/db/postgres) – [see how to change](#database)
  - [MySQL](src/lib/server/db/mysql)
  - [Turso](src/lib/server/db/turso)
- 🔐 [Lucia Auth 3](https://lucia-auth.com/) for easy and secure authentication, with ready-to-use [OAuth 2.0](https://oauth.net/2) provider implementations:
  - [GitHub](src/routes/api/v1/oauth/github)
  - [Google](src/routes/api/v1/oauth/google)
  > _[PRs are welcome](./CONTRIBUTING.md) to support additional OAuth providers ([Arctic Providers](https://arctic.js.org/))_
- 📨 [svelte-email](https://github.com/carstenlebek/svelte-email) for creating and managing customizable and easy email templates.
- 📬 Generic email adapters for popular providers, plus [MailHog](https://github.com/mailhog/MailHog) for local and E2E email testing:
  - [MailHog adapter](src/lib/email/adapters/mailhog.ts)
  - [Nodemailer adapter](src/lib/email/adapters/nodemailer.ts)
  - [Resend adapter](src/lib/email/adapters/resend.ts)
  - [SendGrid adapter](src/lib/email/adapters/send-grid.ts)
  > _[PRs are welcome](./CONTRIBUTING.md) to support additional email providers_
- 🧩 [shadcn-svelte](https://www.shadcn-svelte.com) for a collection of accessible and customizable UI components.
- 💅 [UnoCSS](https://github.com/antfu/unocss) as the utility-first atomic CSS engine, with default presets:
  - [PresetUno](https://unocss.dev/presets/uno)
  - [PresetTypography](https://unocss.dev/presets/typography)
  - [PresetIcons](https://unocss.dev/presets/icons)
  - [PresetWebFonts](https://unocss.dev/presets/web-fonts)
- 🎨 Modern and customizable [themes](src/styles/themes) built with UnoCSS and shadcn-svelte:
  - [Default](src/styles/themes/high-contrast.css)
  - [High Contrast](src/styles/themes/high-contrast.css)
  - [Coffee](src/styles/themes/coffee.css)
  - [Aqua](src/styles/themes/aqua.css)
  - [Mountain](src/styles/themes/mountain.css)
  - [Valentine](src/styles/themes/valentine.css)
  - [Hacker](src/styles/themes/hacker.css)
  - [Synthwave](src/styles/themes/synthwave.css)
  - [Space Cowboy](src/styles/themes/spacecowboy.css)  
  > _[PRs are welcome](./CONTRIBUTING.md) to implement additional themes_
- 📝 [Superforms](https://superforms.rocks/get-started/valibot) and [Formsnap](https://formsnap.dev/) for powerful and accessible forms.
- ✅ [Valibot](https://valibot.dev/) for modular, lightweight and fast schema validation.
- 😸 Over 200,000 open-source vector icons from popular sets via [Pure CSS Icons](https://unocss.dev/presets/icons) and [Iconify](https://iconify.design/). 🔎 You can refer to [Icônes](https://icones.js.org/) or [Iconify](https://icon-sets.iconify.design/) to see all the icon-sets available.
- 🧪 Unit testing with the [bun test runner](https://bun.sh/docs/cli/test).
- 🧑‍💻 E2E testing with [Playwright](https://playwright.dev/), following best practices from [Playwright best practices](https://playwright.dev/docs/best-practices) and using the [Page Object Model pattern](https://playwright.dev/docs/pom). Includes [accessibility tests](tests/a11y) and [lighthouse metrics tests](tests/lighthouse).
- 🐋 [Dockerfile](./Dockerfile) for containerizing the app and [docker-compose](./docker-compose.yaml) for managing app dependencies.
- 🪝 [Lefthook](https://github.com/evilmartians/lefthook) for efficient Git hooks management, allowing automated tasks to run at specific points in the Git workflow.
- 🔀 Ready-to-use [GitHub Actions Workflows](./github/workflows) for automation:
  - [CI checks](./github/workflows/ci.yaml) for continuous integration and testing
  - [Lint translations](./github/workflows/ci.yaml) for ensuring consistent translation files
- 🗂️ Pre-configured [GitHub issue templates](.github/ISSUE_TEMPLATE) using the [new forms syntax](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/syntax-for-issue-forms).
- 📦 [Changesets](https://github.com/changesets/changesets) for streamlined package versioning.
- 👮 [BiomeJS](https://biomejs.dev) for linting and formatting, ensuring consistent code style (with [Prettier](https://prettier.io) for HTML formatting until [Biome supports it](https://biomejs.dev/internals/language-support/)).
- ⌨️ [hotkeys-js](https://github.com/jaywcjlove/hotkeys-js) for customizable OS-based keyboard shortcuts, improving accessibility and productivity.
- 📈 Top-notch SEO with [svelte-meta-tags](https://www.npmjs.com/package/svelte-meta-tags), best practices like [JSON-LD](https://developers.google.com/search/docs/appearance/structured-data/intro-structured-data), [schema.org markup](https://schema.org/), [open graph meta tags](https://ogp.me/), [Twitter card meta tags](https://developer.twitter.com/en/docs/twitter-for-websites/cards/overview/abouts-cards), and [automatic sitemap generation](https://www.npmjs.com/package/svelte-sitemap) (see [SEO section](#seo)).
- 🏎️ Great out-of-the-box performance and optimized [Core Web Vitals](https://web.dev/explore/learn-core-web-vitals) through:
  - Optimized [Largest Contentful Paint (LCP)](https://web.dev/optimize-lcp/) with techniques like code-splitting, lazy-loading, and optimizing critical rendering paths.
  - Improved [First Input Delay (FID)](https://web.dev/optimize-fid/) by minimizing main thread work and leveraging techniques like Web Workers and code-splitting.
  - Optimized [Cumulative Layout Shift (CLS)](https://web.dev/optimize-cls/) by using size containment and avoiding layout shifts.
  - Leveraging [Svelte runes](https://svelte.dev/blog/runes) and [ParaglideJS treeshakeable messages](https://inlang.com/m/gerre34r/library-inlang-paraglideJs) for efficient code and reduced bundle sizes.
- ♿ Great out-of-the-box accessibility through:
  - [shadcn-svelte](https://www.shadcn-svelte.com/) components, which are built on top of [Melt UI](https://melt-ui.com/) and follow the [WAI-ARIA design patterns](https://www.w3.org/WAI/ARIA/apg/patterns/).
  - High contrast themes for improved visibility and readability, both light and dark.
  - [OpenDyslexic font](https://opendyslexic.org/) option for users with dyslexia or reading difficulties.
  - Keyboard navigation support and focus management for enhanced accessibility.
  - Proper semantic markup and ARIA attributes for better screen reader compatibility.
  - Accessibility testing with [axe](https://www.deque.com/axe/) and [Pa11y](https://pa11y.org/) to identify and resolve accessibility issues.

## Prerequisites

- Use [Bun](https://bun.sh/) for running commands and install dependencies. Or use [Node](https://nodejs.org)/[Deno](https://deno.com) by removing `bun.lockb`, installing the corresponding SvelteKit adapter, updating `svelte.config.js`, and removing `svelte-adapter-bun`.

- The project uses external systems: [PostgreSQL](https://www.postgresql.org/) for the database and [MailHog](https://github.com/mailhog/MailHog) for local email testing (more can be added later). Traditionally, these require local installation, but you can use [Docker containers](https://www.docker.com/resources/what-container/) and the included [docker-compose](https://docs.docker.com/compose) (`docker-compose.yaml`) to run dependencies in containers (see [Setting up the project locally](#setting-up-the-project-locally)).

### Summary

- Install [Bun](https://bun.sh/), or [Node](https://nodejs.org), or [Deno](https://deno.com).
- (Optional) Install [Docker](https://www.docker.com) and [Docker Compose](https://docs.docker.com/compose).

## Creating a new project from this template

You have multiple options to start a new project based on this template:

- [Create a GitHub repo from this template](https://github.com/MesaTechLabs/kitten/generate).
- [Download the zip](https://github.com/MesaTechLabs/kitten/archive/refs/heads/master.zip) from the GitHub repository, which excludes the `.git` directory.
- Scaffold the project using [degit](https://github.com/Rich-Harris/degit).

  ```sh
  bunx degit MesaTechLabs/kitten my-new-project
  ```

Alternatively, you can fork it on StackBlitz, an online IDE:  

[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/fork/github/jm2097/angular-boilerplate)

## Setting up the project locally

1. Clone the project using one of the options from the "Creating a new project from this template" section.
2. Install the project dependencies:

    ```sh
    bun install
    ```

3. Configure environment variables by copying the example file and filling in all required variables:

    ```sh
    cp .env.example .env
    ```

4. Start the PostgreSQL and MailHog Docker containers:

    ```sh
    docker-compose up -d
    ```

5. Create a local database instance by running the database migrations:

    ```sh
    bun run postgres:push
    ```

6. Run the project:

    ```sh
    bun run dev
    ```

## Contributing

Please see the [Contributing Guide](./CONTRIBUTING.md) for information on contributing to this project.

## i18n

> [!NOTE]
> ParaglideJS messages are automatically generated at `src/paraglide/messages` during `bun run dev` or `bun run build`

## SEO

[Search Engine Optimization section]

## Authentication

[Authentication section]

## Email

[Email section]

## E2E testing

[End-to-End testing section]

## License

[MIT](./LICENSE)

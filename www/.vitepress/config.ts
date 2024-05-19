import { type DefaultTheme, defineConfig } from "vitepress";

const nav = [
  {
    text: "Guide",
    link: "/guide",
  },
  {
    text: "Integrations",
    items: [
      {
        text: "Headless CMS",
        link: "/headless-cms",
      },
      {
        text: "Strapi",
        link: "/strapi",
      },
    ],
  },
  {
    text: "API",
    items: [
      {
        text: "REST",
        link: "/api/rest",
      },
      {
        text: "GraphQL",
        link: "/api/graphql",
      },
      {
        text: "tRPC",
        link: "/api/trpc",
      },
    ],
  },
  {
    text: "v0.0.0",
    items: [
      {
        text: "Releases",
        link: "https://github.com/jm2097/gon/releases",
      },
      {
        text: "Contributing guide",
        link: "https://github.com/jm2097/gon/blob/main/CONTRIBUTING.md",
      },
    ],
  },
] satisfies DefaultTheme.NavItem[];

const sidebar = {
  "/guide/": [
    { text: "What is Gon?", link: "/guide/what-is-gon" },
    { text: "Quick start", link: "/guide/quick-start" },
    { text: "Anatomy", link: "/guide/anatomy" },
    {
      text: "Configuration",
      collapsed: false,
      items: [
        { text: "Configuration", link: "/guide/config/" },
        {
          text: "Loader - Environment",
          link: "/guide/config/loader-environment",
        },
        { text: "Loader - JSON", link: "/guide/config/loader-json" },
        { text: "Loader - YAML", link: "/guide/config/loader-yaml" },
      ],
    },
    {
      text: "Database",
      collapsed: false,
      items: [
        { text: "SQLC", link: "/guide/sqlc" },
        { text: "Queries", link: "/guide/sqlc/queries" },
        { text: "Migrations", link: "/guide/sqlc/migrations" },
      ],
    },
    {
      text: "Authentication",
      collapsed: false,
      items: [
        { text: "JWT", link: "/guide/auth/jwt" },
        { text: "API Keys", link: "/guide/auth/api-keys" },
        { text: "OAuth", link: "/guide/auth/oauth" },
      ],
    },
    {
      text: "Misc",
      collapsed: false,
      items: [
        { text: "Default values", link: "/guide/tools/defaults" },
        { text: "Cron jobs", link: "/guide/tools/cron" },
        { text: "Custom errors", link: "/guide/tools/custom-errors" },
        { text: "Custom validators", link: "/guide/tools/custom-validators" },
        { text: "Generators", link: "/guide/tools/generators" },
        { text: "Logger", link: "/guide/tools/logger" },
      ],
    },
  ],
  "/deployment/": [
    { text: "Deployment", link: "/deployment" },
    { text: "Docker", link: "/deployment/docker" },
    { text: "Kubernetes", link: "/deployment/kubernetes" },
    { text: "Cloudflare", link: "/deployment/cloudflare" },
    { text: "Digital Ocean", link: "/deployment/digital-ocean" },
  ],
} satisfies DefaultTheme.SidebarMulti;

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Gon",
  description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
  lastUpdated: true,
  themeConfig: {
    nav,
    sidebar,
    footer: {
      message: "Released under the MIT License.",
      copyright: "Copyright © 2024-present Juan Mesa",
    },
    search: {
      provider: "local",
    },
    editLink: {
      pattern: "https://github.com/jm2097/gon/edit/master/docs/:path",
      text: "Suggest changes to this page",
    },
    socialLinks: [{ icon: "github", link: "https://github.com/jm2097/gon" }],
  },
});

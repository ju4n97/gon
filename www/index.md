---
layout: home

hero:
  name: "Gon"
  text: Lorem ipsum dolor sit amet, consectetur.
  tagline: Lorem ipsum dolor sit amet, consectetur adipiscing.
  actions:
    - theme: brand
      text: What is Gon?
      link: /guide/what-is-gon
    - theme: alt
      text: Contribute
      link: /guide/quickstart
  image:
    src: https://raw.githubusercontent.com/vuejs/vitepress/af4717d6820233a011200d44abba53d0f66bfad3/art/vitepress-logo.svg
    alt: VitePress

features:
  - title: Composable
    icon: <span class="i-solar:layers-bold-duotone"></span>
    details: Lorem ipsum dolor sit amet, consectetur adipiscing elit
  - title: Blazingly fast
    icon: <span class="i-solar:rocket-bold-duotone"></span>
    details: Lorem ipsum dolor sit amet, consectetur adipiscing elit
  - title: Database documentation
    icon: <span class="i-solar:database-bold-duotone"></span>
    details: Lorem ipsum dolor sit amet, consectetur adipiscing elit
    link: /
---

<style>
:root {
  --vp-home-hero-name-color: transparent;
  --vp-home-hero-name-background: -webkit-linear-gradient(120deg, #bd34fe 30%, #41d1ff);

  --vp-home-hero-image-background-image: linear-gradient(-45deg, #bd34fe 50%, #47caff 50%);
  --vp-home-hero-image-filter: blur(44px);
}

@media (min-width: 640px) {
  :root {
    --vp-home-hero-image-filter: blur(56px);
  }
}

@media (min-width: 960px) {
  :root {
    --vp-home-hero-image-filter: blur(68px);
  }
}
</style>

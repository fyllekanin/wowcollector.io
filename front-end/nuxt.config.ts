// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  runtimeConfig: {},
  $development: {
    runtimeConfig: {
      baseURL: 'http://127.0.0.1:8888',
    },
  },
  $production: {
    runtimeConfig: {
      baseURL: '',
    },
  },
  ogImage: {
    enabled: false,
  },
  site: {
    url: 'https://wowcollector.io',
    name: 'WoW Collector',
    description:
      'The only World of Warcraft collection tracker you will ever need.',
    defaultLocale: 'en', // not needed if you have @nuxtjs/i18n installed
    indexable: true,
  },
  seo: {
    fallbackTitle: true,
    enabled: true,
  },
  robots: {
    allow: ['/'],
    disallow: ['/admin'],
  },
  srcDir: 'src',
  serverDir: 'server',
  extends: ['@nuxt/ui-pro'],
  modules: [
    '@nuxt/ui',
    '@nuxt/content',
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt',
    '@nuxt/icon',
    '@nuxt/image',
    '@vueuse/nuxt',
    'nuxt-gtag',
    '@nuxtjs/seo',
  ],
  icon: {
    provider: 'server',
  },
  components: [
    {
      path: '~/components',
      pathPrefix: false,
    },
    {
      path: '~/containers',
      pathPrefix: false,
    },
  ],
  colorMode: {
    preference: 'dark',
  },
  plugins: ['~/plugins/draggable.ts'],
  gtag: {
    id: 'G-XXXXXXXXXX',
  },
});

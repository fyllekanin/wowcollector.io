// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  runtimeConfig: {},
  $development: {
    runtimeConfig: {
      baseURL: 'http://127.0.0.1:8888',
      public: {
        BNET_CLIENT_ID: process.env.BNET_CLIENT_ID,
        BNET_REDIRECT_URI: process.env.BNET_REDIRECT_URI,
        BNET_SCOPE: 'openid wow.profile',
        DISCORD_CLIENT_ID: process.env.DISCORD_CLIENT_ID,
        DISCORD_REDIRECT_URI: process.env.DISCORD_REDIRECT_URI,
        DISCORD_SCOPE: 'identify+email+openid',
      },
    },
  },
  $production: {
    runtimeConfig: {
      baseURL: '',
      public: {
        BNET_CLIENT_ID: process.env.BNET_CLIENT_ID,
        BNET_REDIRECT_URI: process.env.BNET_REDIRECT_URI,
        BNET_SCOPE: 'openid wow.profile',
        DISCORD_CLIENT_ID: process.env.DISCORD_CLIENT_ID,
        DISCORD_REDIRECT_URI: process.env.DISCORD_REDIRECT_URI,
        DISCORD_SCOPE: 'identify+email+openid',
      },
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
    '@vueuse/motion/nuxt',
    '@nuxtjs/ngrok',
  ],
  icon: {
    provider: 'server',
    customCollections: [
      {
        prefix: 'custom',
        dir: './src/assets/custom',
      },
    ],
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
    id: process.env.GOOGLE_TAG_ID,
  },
  ngrok: {
    authtoken_from_env: true,
  },
});

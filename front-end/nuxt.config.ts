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
    runtimeConfig: {},
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
});

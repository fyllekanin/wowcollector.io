// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  runtimeConfig: {},
  $development: {},
  $production: {},
  srcDir: 'src',
  serverDir: 'server',
  extends: ['@nuxt/ui-pro'],
  modules: ['@nuxt/ui', '@nuxt/content'],
  components: [
    {
      path: '~/components',
      pathPrefix: false,
    },
  ],
});

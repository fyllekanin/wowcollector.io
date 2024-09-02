<script lang="ts" setup>
useHead({
  script: [
    {
      src: 'https://wow.zamimg.com/js/tooltips.js',
      async: true,
    },
  ],
  meta: [
    {
      name: 'viewport',
      content:
        'width=device-width, initial-scale=1, initial-scale=1, maximum-scale=1',
    },
  ],
});
</script>

<template>
  <div>
    <NuxtLoadingIndicator />

    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>

    <UNotifications />
    <UModals />
    <CookieBar />
  </div>
</template>

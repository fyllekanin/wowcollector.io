<script lang="ts" setup>
useHead({
  script: [
    {
      src: 'https://wow.zamimg.com/js/tooltips.js',
      async: true,
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

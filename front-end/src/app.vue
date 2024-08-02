<script lang="ts" setup>
useHead({
  script: [
    {
      src: 'https://wow.zamimg.com/js/tooltips.js',
      async: true,
    },
    {
      innerHTML:
        'const whTooltips = {colorLinks: false, iconizeLinks: true, renameLinks: false};',
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
  </div>
</template>

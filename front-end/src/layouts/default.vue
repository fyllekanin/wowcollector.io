<script setup lang="ts">
const route = useRoute();

const currentRoute = computed(() => route.path);
</script>

<template>
  <div>
    <AppHeader />

    <UMain>
      <slot />
    </UMain>

    <AppFooter />
  </div>
</template>

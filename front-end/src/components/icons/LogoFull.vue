<script lang="ts" setup>
const colorMode = useColorMode();
const isDark = computed(() => colorMode.value === 'dark');
</script>

<template>
  <img
    v-if="isDark"
    src="../../assets/img/logo.webp"
    alt="Logo"
    style="height: 32px; width: auto"
  />
  <img
    v-else
    src="../../assets/img/logo_black.webp"
    alt="Logo"
    style="height: 22px; width: auto"
  />
</template>

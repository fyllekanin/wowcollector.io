<script lang="ts" setup>
const slideoverOpen = ref<boolean>(false);

const mountsStore = useMountsStore();
</script>

<template>
  <UContainer>
    <MountFilter v-model="slideoverOpen" />
  </UContainer>
</template>

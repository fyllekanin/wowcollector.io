<script lang="ts" setup>
import { ref, onMounted, onUnmounted } from 'vue';

const isIntersecting = ref(false);
const elementRef = ref<HTMLElement | null>(null);

const observer = new IntersectionObserver((entries) => {
  entries.forEach((entry) => {
    isIntersecting.value = entry.isIntersecting;
  });
});

onMounted(() => {
  if (elementRef.value) {
    observer.observe(elementRef.value);
  }
});

onUnmounted(() => {
  observer.disconnect();
});
</script>

<template>
  <div ref="elementRef">
    <slot v-if="isIntersecting" />
  </div>
</template>

<script lang="ts" setup>
const isIntersecting = ref(false);
const elementRef = ref();

const observer = new IntersectionObserver(
  (entries) => {
    entries.forEach((entry) => {
      isIntersecting.value = entry.isIntersecting;
    });
  },
  {
    root: null,
    rootMargin: '0px',
    threshold: 0.1,
  }
);

onMounted(() => {
  if (!elementRef.value) return;
  observer.observe(elementRef.value);
});

onUnmounted(() => {
  observer.disconnect();
});
</script>

<template>
  <div ref="elementRef" v-if="isIntersecting">
    <slot />
  </div>
</template>

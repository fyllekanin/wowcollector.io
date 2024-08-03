<script lang="ts" setup>
const isIntersecting = ref(false);
const elementRef = ref<HTMLElement | null>(null);

const observer = new IntersectionObserver(
  (entries) => {
    entries.forEach((entry) => {
      isIntersecting.value = entry.isIntersecting;
    });
  },
  {
    threshold: 0.5,
  }
);

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
  <!-- <div ref="elementRef"> -->
  <slot ref="elementRef" v-if="isIntersecting" />
  <!-- </div> -->
</template>

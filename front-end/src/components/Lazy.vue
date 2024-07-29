<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue';

const isVisible = ref(false);
const el = ref<HTMLElement | null>(null);

const handleIntersection = (entries: IntersectionObserverEntry[]) => {
  const entry = entries[0];
  isVisible.value = entry.isIntersecting;
  if (entry.isIntersecting) {
    console.log('Component is in view');
  } else {
    console.log('Component is out of view');
  }
};

const intersectionObserver = new IntersectionObserver(handleIntersection);

onMounted(() => {
  nextTick(() => {
    if (el.value) {
      intersectionObserver.observe(el.value);
    }
  });
});

onBeforeUnmount(() => {
  intersectionObserver.disconnect();
});
</script>

<template>
  <div ref="el" class="wrapper">
    <div v-if="isVisible" class="list-item">
      <slot />
    </div>
  </div>
</template>

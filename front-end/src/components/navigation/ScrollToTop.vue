<script lang="ts" setup>
import { ref, onMounted, onUnmounted } from 'vue';

const showButton = ref(false);

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' });
}

function handleScroll() {
  showButton.value = window.scrollY > 100;
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll);
});
</script>

<template>
  <Transition name="slide-fade">
    <div v-if="showButton" class="fixed bottom-20 right-14 z-50">
      <UTooltip text="Scroll to top">
        <UButton
          color="primary"
          variant="solid"
          size="xl"
          :ui="{
            rounded: 'rounded-full',
          }"
          icon="i-heroicons-arrow-up-20-solid"
          @click="scrollToTop"
        />
      </UTooltip>
    </div>
  </Transition>
</template>

<style>
.slide-fade-enter-active {
  transition: all 0.2s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.2s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>

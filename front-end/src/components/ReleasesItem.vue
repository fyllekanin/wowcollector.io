<script setup lang="ts">
import { useIntersectionObserver } from '@vueuse/core';

import type { Changelog } from '~/types';

defineProps<{ date: Changelog }>();

const target = ref(null);
const targetIsVisible = ref(false);

useIntersectionObserver(
  target,
  ([{ isIntersecting }]) => {
    targetIsVisible.value = isIntersecting;
  },
  {
    rootMargin: '-68px 0px -68px 0px',
  }
);
</script>

<template>
  <div
    ref="target"
    :id="date.version"
    class="flex flex-col transition-opacity duration-500"
    :class="targetIsVisible ? 'opacity-100' : 'opacity-25'"
  >
    <time
      :datetime="date.timestamp"
      class="flex-shrink-0 text-sm/6 font-semibold text-gray-500 dark:text-gray-400"
      >{{
        new Date(Number(date.timestamp) * 1000).toLocaleString('en-GB', {
          year: 'numeric',
          month: 'short',
          day: 'numeric',
        })
      }}</time
    >

    <p
      v-if="date.version"
      class="text-gray-900 dark:text-white font-bold text-3xl mt-2 group hover:text-primary-500 dark:hover:text-primary-400 transition-[color]"
    >
      {{ date.version }}
    </p>
    <ul
      v-if="date.changes?.length"
      class="mt-2 space-y-1 text-gray-600 dark:text-gray-300"
    >
      <li
        v-for="(change, i) in date.changes"
        :key="i"
        class="text-sm/6 break-all"
      >
        {{ change.content }}
      </li>
    </ul>
  </div>
</template>

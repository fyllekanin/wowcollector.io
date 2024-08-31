<script setup lang="ts">
import { useIntersectionObserver } from '@vueuse/core';

import { ChangelogType, type Changelog } from '~/types';

const props = defineProps({
  date: {
    type: Object as PropType<Changelog>,
    required: true,
  },
});

const target = ref(null);
const targetIsVisible = ref(false);

const majorChanges = computed(() =>
  props.date.changes?.filter((change) => change.type === ChangelogType.BREAKING)
);
const minorChanges = computed(() =>
  props.date.changes?.filter((change) => change.type === ChangelogType.FEATURE)
);
const patchChanges = computed(() =>
  props.date.changes?.filter((change) => change.type === ChangelogType.BUG)
);

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
      class="text-gray-900 dark:text-white font-bold text-3xl mt-2 group hover:text-primary-500 dark:hover:text-primary-400 transition-[color] cursor-default"
    >
      {{ date.version }}
    </p>
    <div v-if="majorChanges.length" class="mt-2">
      <h3 class="text-gray-900 dark:text-white font-bold text-xl">
        Breaking Changes
      </h3>
      <ul
        class="mt-2 space-y-1 text-gray-600 dark:text-gray-300 list-disc list-inside"
      >
        <li
          v-for="(change, i) in majorChanges"
          :key="i"
          class="text-sm/6 break-all"
        >
          {{ change.content }}
        </li>
      </ul>
    </div>

    <div v-if="minorChanges.length" class="mt-2">
      <h3 class="text-gray-900 dark:text-white font-bold text-lg">Features</h3>
      <ul
        class="mt-2 space-y-1 text-gray-600 dark:text-gray-300 list-disc list-inside"
      >
        <li
          v-for="(change, i) in minorChanges"
          :key="i"
          class="text-sm/6 break-all"
        >
          {{ change.content }}
        </li>
      </ul>
    </div>

    <div v-if="patchChanges.length" class="mt-2">
      <h3 class="text-gray-900 dark:text-white font-bold text-xl">Bug Fixes</h3>
      <ul
        class="mt-2 space y-1 text-gray-600 dark:text-gray-300 list-disc list-inside"
      >
        <li
          v-for="(change, i) in patchChanges"
          :key="i"
          class="text-sm/6 break-all"
        >
          {{ change.content }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  progress: {
    type: Number,
    required: true,
  },
  collected: {
    type: Array,
    required: true,
  },
  available: {
    type: Array,
    required: true,
  },
  collection: {
    type: String,
    required: true,
    enum: ['mounts', 'achievements', 'toys', 'pets'],
  },
});

const phrase = computed(
  () =>
    ({
      mounts: 'collected',
      achievements: 'completed',
      toys: 'collected',
      pets: 'collected',
    }[props.collection])
);
</script>

<template>
  <div class="flex flex-col gap-4 items-center">
    <div class="flex grow gap-2 items-center w-full md:flex-col">
      <div class="flex flex-col w-full">
        <UProgress :value="progress" />
        <p
          class="text-center text-xs sm:text-sm text-nowrap text-gray-500 self-end pt-1"
        >
          {{ collected.length }} out of {{ available.length }} {{ collection }}
          {{ phrase }} ({{ progress }}%)
        </p>
      </div>
      <!-- <MountFilters v-if="collection === 'mounts'" />
      <AchievementFilters v-else-if="collection === 'achievements'" /> -->
      <slot />
    </div>
  </div>
</template>

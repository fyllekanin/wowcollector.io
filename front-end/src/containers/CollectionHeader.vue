<script lang="ts" setup>
const props = defineProps({
  progress: {
    type: Number,
    required: true,
  },
  collected: {
    type: Number,
    required: true,
  },
  total: {
    type: Number,
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
          {{ collected }} out of {{ total }} {{ collection }} {{ phrase }} ({{
            progress
          }}%)
        </p>
      </div>
      <slot />
    </div>
  </div>
</template>

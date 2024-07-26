<script lang="ts" setup>
const slideoverOpen = defineModel<boolean>();

// const { isCustomXL } = useScreenSize();
// console.log(isCustomXL.value);
</script>

<template>
  <UCard class="max-w-xs">
    <template #header>
      <div class="flex gap-2 items-center">
        <UIcon name="material-symbols:filter-alt" class="scale-125" />
        <span>Filters</span>
      </div>
    </template>

    <slot />
  </UCard>
  <!-- <USlideover v-else v-model="slideoverOpen">
    <UCard
      class="flex flex-col flex-1"
      :ui="{
        body: { base: 'flex-1' },
        ring: '',
        divide: 'divide-y divide-gray-100 dark:divide-gray-800',
      }"
    >
      <template #header>
        <UIcon name="filter" />
        <span>Filters</span>
      </template>

      <slot />
    </UCard>
  </USlideover> -->
</template>

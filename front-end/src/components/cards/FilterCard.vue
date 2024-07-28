<script lang="ts" setup></script>

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
</template>

<script lang="ts" setup>
import { SORT_TYPES } from '~/constants';

const slideoverOpen = ref<boolean>(false);

const mountsStore = useMountsStore();
const { filters, rootCategoryNames, subCategoryNames } =
  storeToRefs(mountsStore);

const misc = computed(() => filters.value.miscFilters ?? []) as ComputedRef<
  string[]
>;
</script>

<template>
  <UButton
    class="md:hidden"
    variant="ghost"
    color="gray"
    icon="material-symbols:filter-alt"
    @click="slideoverOpen = true"
    >Filters</UButton
  >
  <USlideover v-model="slideoverOpen" side="left">
    <UCard
      class="flex flex-col flex-1"
      :ui="{
        body: { base: 'flex-1' },
        ring: '',
        divide: 'divide-y divide-gray-100 dark:divide-gray-800',
      }"
    >
      <template #header>
        <div class="flex justify-between">
          <header class="flex items-center gap-2">
            <UIcon class="scale-125" name="material-symbols:filter-alt" />
            <span>Filters</span>
          </header>
          <UButton
            variant="ghost"
            icon="heroicons-solid:x"
            @click="slideoverOpen = false"
            color="gray"
            aria-label="Close"
          />
        </div>
      </template>

      <MountFilterForm class="md:hidden" />
    </UCard>
  </USlideover>
  <MountFilterForm class="hidden md:flex" />
</template>

<script lang="ts" setup>
import { RENDER_TYPES, SORT_TYPES } from '~/constants';

const mountsStore = useMountsStore();

const { mountFilters } = storeToRefs(mountsStore);

const baseCategories = computed(
  () => mountFilters.value.baseCategories ?? []
) as ComputedRef<string[]>;
const subCategories = computed(
  () => mountFilters.value.subCategories ?? []
) as ComputedRef<string[]>;
const misc = computed(() => mountFilters.value.misc ?? []) as ComputedRef<
  string[]
>;
</script>

<template>
  <div class="flex grow flex-wrap gap-4 items-end">
    <UInput
      v-model="mountFilters.search as string"
      class="self-end"
      placeholder="Search for a mount"
      icon="i-heroicons-magnifying-glass-20-solid"
      variant="none"
      :ui="{
        variant: {
          none: 'bg-transparent focus:ring-0 focus:shadow-none border-b border-gray-200 dark:border-gray-800 rounded-none focus:border-primary dark:focus:border-primary',
        },
      }"
    >
      <template #trailing>
        <UButton
          v-show="mountFilters.search !== ''"
          color="gray"
          variant="link"
          icon="i-heroicons-x-mark-20-solid"
          :padded="false"
          @click="mountsStore.setMountFilters({ search: '' })"
        />
      </template>
    </UInput>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Base Categories</span>
      <UInputMenu
        v-model="mountFilters.baseCategories"
        class="w-[150px]"
        :options="baseCategories"
      ></UInputMenu>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Sub Categories</span>
      <UInputMenu
        v-model="mountFilters.subCategories"
        class="w-[150px]"
        :options="subCategories"
      ></UInputMenu>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Misc</span>
      <UInputMenu
        v-model="mountFilters.misc"
        class="w-[150px]"
        :options="misc"
      ></UInputMenu>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Sort</span>
      <USelect
        class="self-end"
        v-model="mountFilters.sortType"
        :options="SORT_TYPES"
        icon="material-symbols:sort"
        placeholder="Sort by"
      />
    </div>
    <UButton
      class="h-min self-end"
      variant="ghost"
      color="gray"
      icon="codicon:collapse-all"
      >Collapse all</UButton
    >
    <UButtonGroup class="h-min lg:self-end" size="sm">
      <UTooltip
        v-for="({ label, value, icon }, index) in RENDER_TYPES"
        :text="label"
        :key="index"
      >
        <UButton
          v-model="mountFilters.renderType"
          :value="value"
          :icon="icon"
          :color="mountFilters.renderType === value ? 'primary' : 'light'"
          variant="outline"
          @click="mountsStore.setMountFilters({ renderType: value })"
        />
      </UTooltip>
    </UButtonGroup>
  </div>
</template>

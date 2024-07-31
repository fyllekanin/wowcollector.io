<script lang="ts" setup>
import { SORT_TYPES } from '~/constants';

const mountsStore = useMountsStore();
const { filters, rootCategoryNames, subCategoryNames } =
  storeToRefs(mountsStore);

const misc = computed(() => filters.value.miscFilters ?? []) as ComputedRef<
  string[]
>;
</script>

<template>
  <div class="flex grow flex-wrap gap-4 items-end">
    <UInput
      v-model="filters.search"
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
          v-show="filters.search !== ''"
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
      <UButtonGroup>
        <USelectMenu
          v-model="filters.rootCategories"
          class="w-[150px]"
          searchable
          clear-search-on-close
          multiple
          :options="rootCategoryNames"
        >
        </USelectMenu>
        <!--  clear all button -->
        <UTooltip
          text="Clear all"
          :prevent="filters.rootCategories.length === 0"
        >
          <UButton
            color="gray"
            icon="i-heroicons-x-mark-20-solid"
            :disabled="filters.rootCategories.length === 0"
            :padded="false"
            @click="mountsStore.setMountFilters({ rootCategories: [] })"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Sub Categories</span>
      <UButtonGroup>
        <USelectMenu
          v-model="filters.subCategories"
          class="w-[150px]"
          searchable
          clear-search-on-close
          multiple
          :options="subCategoryNames"
        ></USelectMenu>
        <UTooltip
          text="Clear all"
          :prevent="filters.subCategories.length === 0"
        >
          <UButton
            color="gray"
            icon="i-heroicons-x-mark-20-solid"
            :disabled="filters.subCategories.length === 0"
            :padded="false"
            @click="mountsStore.setMountFilters({ subCategories: [] })"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Misc</span>
      <UButtonGroup>
        <USelectMenu
          v-model="filters.miscFilters"
          class="w-[150px]"
          searchable
          clear-search-on-close
          multiple
          :options="misc"
        ></USelectMenu>
        <UTooltip text="Clear all" :prevent="filters.miscFilters.length === 0">
          <UButton
            color="gray"
            icon="i-heroicons-x-mark-20-solid"
            :disabled="filters.miscFilters.length === 0"
            :padded="false"
            @click="mountsStore.setMountFilters({ miscFilters: [] })"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1">
      <span class="text-sm">Sort</span>
      <USelect
        class="self-end"
        v-model="filters.sort"
        :options="SORT_TYPES"
        icon="material-symbols:sort"
        placeholder="Sort by"
      />
    </div>

    <!-- Upcoming Feature -->

    <!-- <UButton
      class="h-min self-end"
      variant="ghost"
      color="gray"
      icon="codicon:collapse-all"
      :disabled="filters.viewStyle !== 'list'"
      >Collapse all</UButton
    >
    <UButtonGroup class="h-min lg:self-end" size="sm">
      <UTooltip
        v-for="({ label, value, icon }, index) in RENDER_TYPES"
        :text="label"
        :key="index"
      >
        <UButton
          v-model="filters.viewStyle"
          :value="value"
          :icon="icon"
          :color="filters.viewStyle === value ? 'primary' : 'light'"
          variant="outline"
          @click="mountsStore.setMountFilters({ viewStyle: value })"
        />
      </UTooltip>
    </UButtonGroup> -->
  </div>
</template>

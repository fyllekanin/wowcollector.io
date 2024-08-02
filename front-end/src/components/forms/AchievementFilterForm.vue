<script lang="ts" setup>
import { SORT_TYPES } from '~/constants';

const achievementsStore = useAchievementsStore();
const { filters, rootCategoryNames, subCategoryNames } =
  storeToRefs(achievementsStore);

const misc = computed(() => filters.value.miscFilters ?? []) as ComputedRef<
  string[]
>;
</script>

<template>
  <UContainer
    class="w-full h-fit flex flex-col sm:flex-row sm:flex-wrap gap-5 items-center"
  >
    <UInput
      v-model="filters.search"
      class="w-full sm:w-[250px] sm:self-end"
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
          @click="achievementsStore.setAchievementFilters({ search: '' })"
        />
      </template>
    </UInput>
    <div class="flex flex-col gap-1 w-full sm:w-[150px]">
      <span class="text-sm">Base Categories</span>
      <UButtonGroup>
        <USelectMenu
          v-model="filters.rootCategories"
          class="w-full"
          searchable
          clear-search-on-close
          multiple
          :options="rootCategoryNames"
        >
        </USelectMenu>
        <UTooltip
          text="Clear all"
          :prevent="filters.rootCategories.length === 0"
        >
          <UButton
            color="gray"
            icon="i-heroicons-x-mark-20-solid"
            :disabled="filters.rootCategories.length === 0"
            :padded="false"
            @click="
              achievementsStore.setAchievementFilters({ rootCategories: [] })
            "
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1 w-full sm:w-[150px]">
      <span class="text-sm">Sub Categories</span>
      <UButtonGroup>
        <USelectMenu
          v-model="filters.subCategories"
          class="w-full"
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
            @click="
              achievementsStore.setAchievementFilters({ subCategories: [] })
            "
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1 w-full sm:w-[150px]">
      <span class="text-sm">Misc</span>
      <UButtonGroup>
        <USelectMenu
          class="w-full"
          v-model="filters.miscFilters"
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
            @click="
              achievementsStore.setAchievementFilters({ miscFilters: [] })
            "
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex flex-col gap-1 w-full sm:w-[150px]">
      <span class="text-sm">Sort</span>
      <USelect
        v-model="filters.sort"
        :options="SORT_TYPES"
        icon="material-symbols:sort"
        placeholder="Sort by"
      />
    </div>
    <div class="flex flex-col w-full sm:w-fit sm:self-end">
      <UButton
        class="self-end"
        color="gray"
        icon="material-symbols:filter-alt-off"
        :disabled="
          Object.values(filters).every(
            (filter) => filter?.length === 0 || filter === ''
          )
        "
        @click="achievementsStore.clearAchievementFilters"
        >Clear all filters</UButton
      >
    </div>
  </UContainer>
</template>

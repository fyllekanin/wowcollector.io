<script lang="ts" setup>
import { ACHIEVEMENT_SORT_TYPES } from '~/constants';

const { debounce } = useDebounce();
const achievementsStore = useAchievementsStore();
const { filters, rootCategoryNames, subCategoryNames } =
  storeToRefs(achievementsStore);

const misc = computed(() => filters.value.miscFilters ?? []) as ComputedRef<
  string[]
>;

const debouncableSearch = ref('');

watch(
  () => debouncableSearch.value,
  debounce((value) => {
    achievementsStore.setAchievementFilters({ search: value });
  }, 300),
  { immediate: true }
);
</script>

<template>
  <UContainer
    class="w-full h-fit flex flex-col sm:flex-row sm:flex-wrap gap-5 items-center"
  >
    <UButtonGroup class="w-full md:w-auto sm:self-end">
      <UInput
        v-model="debouncableSearch"
        class="w-full sm:w-[250px]"
        placeholder="Search for an achievement"
        icon="i-heroicons-magnifying-glass-20-solid"
      >
      </UInput>
      <UTooltip text="Clear search" :prevent="debouncableSearch === ''">
        <UButton
          color="gray"
          icon="i-heroicons-x-mark-20-solid"
          :disabled="debouncableSearch === ''"
          :padded="false"
          @click="debouncableSearch = ''"
        />
      </UTooltip>
    </UButtonGroup>
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
        :options="ACHIEVEMENT_SORT_TYPES"
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

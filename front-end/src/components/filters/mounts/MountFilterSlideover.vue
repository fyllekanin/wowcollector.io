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
    class="sm:hidden"
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

      <div class="flex flex-col gap-5 items-center">
        <UInput
          v-model="filters.search"
          class="w-full"
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
        <div class="flex flex-col gap-1 w-full">
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
        <div class="flex flex-col gap-1 w-full">
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
                @click="mountsStore.setMountFilters({ subCategories: [] })"
              />
            </UTooltip>
          </UButtonGroup>
        </div>
        <div class="flex flex-col gap-1 w-full">
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
            <UTooltip
              text="Clear all"
              :prevent="filters.miscFilters.length === 0"
            >
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
        <div class="flex flex-col gap-1 w-full">
          <span class="text-sm">Sort</span>
          <USelect
            v-model="filters.sort"
            :options="SORT_TYPES"
            icon="material-symbols:sort"
            placeholder="Sort by"
          />
        </div>
        <div class="flex flex-col w-full">
          <UButton
            class="self-end"
            color="gray"
            icon="material-symbols:filter-alt-off"
            :disabled="
              Object.values(filters).every(
                (filter) => filter?.length === 0 || filter === ''
              )
            "
            @click="mountsStore.clearMountFilters"
            >Clear all filters</UButton
          >
        </div>

        <!-- Upcoming Feature -->

        <!-- <div class="flex justify-evenly">
          <UButton
            variant="ghost"
            color="gray"
            icon="codicon:collapse-all"
            :disabled="filters.viewStyle !== 'grid'"
            >Collapse all</UButton
          >
          <UButtonGroup size="sm">
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
          </UButtonGroup>
        </div> -->
      </div>
    </UCard>
  </USlideover>
</template>

<script lang="ts" setup>
import { RENDER_TYPES, SORT_TYPES } from '~/constants';

const slideoverOpen = ref<boolean>(false);

const mountsStore = useMountsStore();
const { filters } = storeToRefs(mountsStore);

const rootCategories = computed(
  () => filters.value.rootCategories ?? []
) as ComputedRef<string[]>;
const subCategories = computed(
  () => filters.value.subCategories ?? []
) as ComputedRef<string[]>;
const misc = computed(() => filters.value.miscFilters ?? []) as ComputedRef<
  string[]
>;
</script>

<template>
  <UButton
    class="block sm:hidden"
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

      <div class="flex flex-col gap-5">
        <UInput
          v-model="filters.search as string"
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
          <UInputMenu
            v-model="filters.rootCategories"
            :options="rootCategories"
          ></UInputMenu>
        </div>
        <div class="flex flex-col gap-1">
          <span class="text-sm">Sub Categories</span>
          <UInputMenu
            v-model="filters.subCategories"
            :options="subCategories"
          ></UInputMenu>
        </div>
        <div class="flex flex-col gap-1">
          <span class="text-sm">Misc</span>
          <UInputMenu
            v-model="filters.miscFilters"
            :options="misc"
          ></UInputMenu>
        </div>
        <div class="flex flex-col gap-1">
          <span class="text-sm">Sort By</span>
          <USelect
            v-model="filters.sort"
            :options="SORT_TYPES"
            icon="material-symbols:sort"
            placeholder="Sort by"
          />
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

<script lang="ts" setup>
const mountsStore = useMountsStore();

const { filters } = storeToRefs(mountsStore);

const slideoverOpen = defineModel<boolean>();

const baseCategories = computed(
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
  <FilterCard v-model="slideoverOpen">
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
          :options="baseCategories"
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
        <UInputMenu v-model="filters.miscFilters" :options="misc"></UInputMenu>
      </div>
    </div>
  </FilterCard>
</template>

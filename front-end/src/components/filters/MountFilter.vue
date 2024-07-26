<script lang="ts" setup>
const mountsStore = useMountsStore();

const { mountFilters } = storeToRefs(mountsStore);

const slideoverOpen = defineModel<boolean>();

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
  <FilterCard v-model="slideoverOpen">
    <div class="flex flex-col gap-5">
      <UInput
        v-model="mountFilters.search"
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
          :options="baseCategories"
        ></UInputMenu>
      </div>
      <div class="flex flex-col gap-1">
        <span class="text-sm">Sub Categories</span>
        <UInputMenu
          v-model="mountFilters.subCategories"
          :options="subCategories"
        ></UInputMenu>
      </div>
      <div class="flex flex-col gap-1">
        <span class="text-sm">Misc</span>
        <UInputMenu v-model="mountFilters.misc" :options="misc"></UInputMenu>
      </div>
    </div>
  </FilterCard>
</template>

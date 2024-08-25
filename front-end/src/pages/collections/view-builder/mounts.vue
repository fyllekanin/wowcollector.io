<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
  middleware: 'view-builder-mounts',
});

const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/mounts/view-builder').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No mounts create view page found in the content.',
  });
}

const { debounce } = useDebounce();

const mountViewBuilderStore = useMountViewBuilderStore();
const {
  _cloneableCategory,
  _mounts,
  _searchFilter,
  _settings,
  successfulCreation,
} = storeToRefs(mountViewBuilderStore);

const createdViewId = ref('');

const debouncableSearch = ref('');
watch(
  () => debouncableSearch.value,
  debounce((value) => {
    mountViewBuilderStore.setSearchFilter(value);
  }, 300),
  { immediate: true }
);

onMounted(() => {
  window.addEventListener('mousedown', removeWowheadTooltips);
});
onUnmounted(() => {
  window.removeEventListener('mousedown', removeWowheadTooltips);
});

function removeWowheadTooltips() {
  [...document.getElementsByClassName('wowhead-tooltip')].forEach((item) => {
    item.remove();
  });
}

function onSuccess(value: string) {
  createdViewId.value = value;
  mountViewBuilderStore.setSuccessfulCreation(true);
}
</script>

<template>
  <SuccessfulViewCreation
    v-if="successfulCreation"
    v-model="createdViewId"
    to="mounts"
  />
  <div v-else class="w-full h-full">
    <ScreenTooSmall class="lg:hidden" />
    <MountViewBuilderContainer class="hidden lg:flex" @success="onSuccess">
      <template #sidebar-content>
        <div class="flex flex-col mt-6 gap-5">
          <draggable
            :list="_cloneableCategory"
            :group="{ name: 'category', pull: 'clone', put: false }"
            @start="mountViewBuilderStore.setNewIdForCloneableCategory"
          >
            <template #item="{ element: category }">
              <UCard
                class="select-none cursor-move"
                :ui="{
                  ring: 'ring-0 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out',
                  rounded: 'rounded-none',
                  body: {
                    padding: 'px-2 py-3 sm:p-3',
                  },
                }"
              >
                {{ category.name }}
              </UCard>
            </template>
          </draggable>
          <UDivider />
          <div class="flex flex-col gap-4 w-full">
            <h2 class="text-md font-semibold self-center">Settings</h2>
            <UDivider />
            <div class="flex gap-2 items-center justify-between">
              <span class="text-sm">Show borders</span>
              <UToggle v-model="_settings.showBorders" />
            </div>
            <div class="flex gap-2 items-center justify-between">
              <span class="text-sm">Show mount tooltips</span>
              <UToggle v-model="_settings.showMountTooltips" />
            </div>
            <UDivider />
          </div>
          <UInput
            v-model="debouncableSearch"
            placeholder="Search for a mount"
            icon="heroicons-outline:search"
          />
          <draggable
            class="flex grow flex-wrap gap-4 justify-center"
            :list="_mounts"
            :group="{ name: 'mount' }"
          >
            <template #item="{ element: mount }">
              <MountIcon
                v-if="
                  mount.name.toLowerCase().includes(_searchFilter.toLowerCase())
                "
                class="select-none cursor-move"
                :mount="mount"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showMountTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
        </div>
      </template>
      <template #main-content>
        <MountNestedDraggable />
      </template>
    </MountViewBuilderContainer>
  </div>
</template>

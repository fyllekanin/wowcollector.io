<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
  middleware: 'view-builder-toys',
});

const { data: page } = await useAsyncData('toys', () =>
  queryContent('/collections/view-builder/toys').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No toys create view page found in the content.',
  });
}

useSeoMeta({
  title: page.value.title,
  description: page.value.description,
  ogTitle: page.value.og.title,
  ogDescription: page.value.og.description,
  ogImage: page.value.og.image,
});

const { debounce } = useDebounce();

const toyViewBuilderStore = useToyViewBuilderStore();
const {
  _cloneableCategory,
  _toys,
  _searchFilter,
  _settings,
  successfulCreation,
  highlightCategoryDropzones,
  highlightToyDropzones,
} = storeToRefs(toyViewBuilderStore);

const createdViewId = ref('');

const debouncableSearch = ref('');
watch(
  () => debouncableSearch.value,
  debounce((value) => {
    toyViewBuilderStore.setSearchFilter(value);
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
  toyViewBuilderStore.setSuccessfulCreation(true);
}

function onLeave() {
  toyViewBuilderStore.resetStore();
  toyViewBuilderStore.setSuccessfulCreation(false);
  createdViewId.value = '';
}
</script>

<template>
  <SuccessModal
    v-if="successfulCreation"
    v-model="createdViewId"
    to="pets"
    @leave="onLeave"
  />
  <div v-else class="w-full h-full">
    <ScreenTooSmall class="lg:hidden" />
    <ToyViewBuilderContainer class="hidden lg:flex" @success="onSuccess">
      <template #sidebar-content>
        <div class="flex flex-col mt-6 gap-5">
          <draggable
            :list="_cloneableCategory"
            :group="{ name: 'category', pull: 'clone', put: false }"
            @start="toyViewBuilderStore.setNewIdForCloneableCategory"
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
              <span class="text-sm">Show toy tooltips</span>
              <UToggle v-model="_settings.showToyTooltips" />
            </div>
            <UDivider />
          </div>
          <UInput
            v-model="debouncableSearch"
            placeholder="Search for a toy"
            icon="heroicons-outline:search"
          />
          <draggable
            :class="[
              'flex grow flex-wrap gap-4 justify-center',
              highlightToyDropzones ? 'bg-green-900 bg-opacity-45' : '',
            ]"
            :list="_toys"
            :group="{ name: 'toy' }"
            @start="_settings.showToyTooltips = false"
            @end="_settings.showToyTooltips = true"
          >
            <template #item="{ element: toy }">
              <ToyIcon
                v-if="
                  toy.name.toLowerCase().includes(_searchFilter.toLowerCase())
                "
                class="select-none cursor-move"
                :toy="toy"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showToyTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
        </div>
      </template>
      <template #main-content>
        <ToyNestedDraggable />
      </template>
    </ToyViewBuilderContainer>
  </div>
</template>

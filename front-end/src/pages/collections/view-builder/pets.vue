<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
  middleware: 'view-builder-pets',
});

const { data: page } = await useAsyncData('pets', () =>
  queryContent('/collections/view-builder/pets').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No pets create view page found in the content.',
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

const petViewBuilderStore = usePetViewBuilderStore();
const {
  _cloneableCategory,
  _pets,
  _searchFilter,
  _settings,
  successfulCreation,
} = storeToRefs(petViewBuilderStore);

const createdViewId = ref('');

const debouncableSearch = ref('');
watch(
  () => debouncableSearch.value,
  debounce((value) => {
    petViewBuilderStore.setSearchFilter(value);
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
  petViewBuilderStore.setSuccessfulCreation(true);
}

function onLeave() {
  petViewBuilderStore.resetStore();
  petViewBuilderStore.setSuccessfulCreation(false);
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
    <PetViewBuilderContainer class="hidden lg:flex" @success="onSuccess">
      <template #sidebar-content>
        <div class="flex flex-col mt-6 gap-5">
          <draggable
            :list="_cloneableCategory"
            :group="{ name: 'category', pull: 'clone', put: false }"
            @start="petViewBuilderStore.setNewIdForCloneableCategory"
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
              <span class="text-sm">Show pet tooltips</span>
              <UToggle v-model="_settings.showPetTooltips" />
            </div>
            <UDivider />
          </div>
          <UInput
            v-model="debouncableSearch"
            placeholder="Search for a pet"
            icon="heroicons-outline:search"
          />
          <draggable
            class="flex grow flex-wrap gap-4 justify-center"
            :list="_pets"
            :group="{ name: 'pet' }"
            @start="_settings.showPetTooltips = false"
            @end="_settings.showPetTooltips = true"
          >
            <template #item="{ element: pet }">
              <PetIcon
                v-if="
                  pet.name.toLowerCase().includes(_searchFilter.toLowerCase())
                "
                class="select-none cursor-move"
                :pet="pet"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showPetTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
        </div>
      </template>
      <template #main-content>
        <PetNestedDraggable />
      </template>
    </PetViewBuilderContainer>
  </div>
</template>

<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
  middleware: 'achievement-view-builder',
});

const { data: page } = await useAsyncData('achievements', () =>
  queryContent('/collections/achievements/view-builder').findOne()
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

const achievementViewBuilderStore = useAchievementViewBuilderStore();
const { _cloneableCategory, _achievements, _searchFilter, _settings } =
  storeToRefs(achievementViewBuilderStore);

const debouncableSearch = ref('');

watch(
  () => debouncableSearch.value,
  debounce((value) => {
    achievementViewBuilderStore.setSearchFilter(value);
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
</script>

<template>
  <ScreenTooSmall class="lg:hidden" />
  <CreateViewContainer class="hidden lg:flex">
    <template #sidebar-content>
      <div class="flex flex-col mt-6 gap-5">
        <draggable
          :list="_cloneableCategory"
          :group="{ name: 'category', pull: 'clone', put: false }"
          @start="achievementViewBuilderStore.setNewIdForCloneableCategory"
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
            <span class="text-sm">Show achievement tooltips</span>
            <UToggle v-model="_settings.showAchievementTooltips" />
          </div>
          <UDivider />
        </div>
        <UInput
          v-model="debouncableSearch"
          placeholder="Search for an achievement"
          icon="heroicons-outline:search"
        />
        <draggable
          class="flex grow flex-wrap gap-4 justify-center"
          :list="_achievements"
          :group="{ name: 'achievement' }"
        >
          <template #item="{ element: achievement }">
            <AchievementIcon
              v-if="
                achievement.name
                  .toLowerCase()
                  .includes(_searchFilter.toLowerCase())
              "
              class="select-none cursor-move"
              :achievement="achievement"
              :clickable="false"
              build-mode
              :show-tooltip="_settings.showAchievementTooltips"
              :use-intersection-observer="false"
            />
          </template>
        </draggable>
      </div>
    </template>
    <template #main-content>
      <AchievementNestedDraggable />
    </template>
  </CreateViewContainer>
</template>

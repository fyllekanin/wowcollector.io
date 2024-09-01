<script lang="ts" setup>
import draggable from 'vuedraggable';

definePageMeta({
  layout: 'empty',
  middleware: 'view-builder-mounts',
});

const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/view-builder/mounts').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No mounts create view page found in the content.',
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

const mountViewBuilderStore = useMountViewBuilderStore();
const {
  _cloneableCategory,
  _mounts,
  _searchFilter,
  _settings,
  successfulCreation,
  highlightCategoryDropzones,
  highlightMountDropzones,
} = storeToRefs(mountViewBuilderStore);

const factions = ref([
  {
    label: 'Both',
    value: 'both',
    avatar: {
      src: 'https://cdn.discordapp.com/attachments/1161263238554599464/1279760555237838899/12d4f5a73e9c1b830c95229ac396a449.png?ex=66d59d65&is=66d44be5&hm=84d99f37f7c08df863528175f78e6853a19b5c6a48f31b1e7e0684708856139a&',
    },
  },
  {
    label: 'Alliance',
    value: 'alliance',
    avatar: {
      src: 'https://cdn.discordapp.com/attachments/1161263238554599464/1279760408932253727/alliance-logo-BDD77C0478-seeklogo.png?ex=66d59d42&is=66d44bc2&hm=0a52a366e8fe97841ed8b685763493d88b506b9b5928402acf6dce1a55174899&',
    },
  },
  {
    label: 'Horde',
    value: 'horde',
    avatar: {
      src: 'https://cdn.discordapp.com/attachments/1161263238554599464/1279760048746139742/horde_logo_by_ammeg88_d5sggp9-fullview.png?ex=66d59cec&is=66d44b6c&hm=57e20063b236f3f65801d407fadf2636c1f83210291f4c8ac6b85301ace3d37e&',
    },
  },
]);

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

function onLeave() {
  mountViewBuilderStore.resetStore();
  mountViewBuilderStore.setSuccessfulCreation(false);
  createdViewId.value = '';
}
</script>

<template>
  <SuccessfulViewCreation
    v-if="successfulCreation"
    v-model="createdViewId"
    to="mounts"
    @leave="onLeave"
  />
  <div v-else class="w-full h-full">
    <ScreenTooSmall class="lg:hidden" />
    <MountViewBuilderContainer class="hidden lg:flex" @success="onSuccess">
      <template #sidebar-content>
        <div class="flex flex-col mt-6 gap-5">
          <draggable
            :list="_cloneableCategory"
            :group="{ name: 'category', pull: 'clone', put: false }"
            @start="
              () => {
                mountViewBuilderStore.setNewIdForCloneableCategory();
                mountViewBuilderStore.setDragState(true, 'category');
              }
            "
            @end="
              () => {
                mountViewBuilderStore.clearDragState();
              }
            "
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
            <div class="flex gap-2 items-center justify-between">
              <span class="text-sm">Show faction</span>
              <USelectMenu
                class="w-1/2"
                v-model="_settings.showFaction"
                :options="factions"
              >
                <template #leading>
                  <UAvatar v-bind="_settings.showFaction.avatar" size="2xs" />
                </template>
              </USelectMenu>
            </div>
            <UDivider />
          </div>
          <UInput
            v-model="debouncableSearch"
            placeholder="Search for a mount"
            icon="heroicons-outline:search"
          />
          <draggable
            :class="[
              'flex grow flex-wrap gap-4 justify-center',
              highlightMountDropzones ? 'bg-green-900 bg-opacity-45' : '',
            ]"
            :list="_mounts"
            :group="{ name: 'mount' }"
            @start="
              () => {
                _settings.showMountTooltips = false;
                mountViewBuilderStore.setDragState(true, 'mount');
              }
            "
            @end="
              () => {
                _settings.showMountTooltips = true;
                mountViewBuilderStore.clearDragState();
              }
            "
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

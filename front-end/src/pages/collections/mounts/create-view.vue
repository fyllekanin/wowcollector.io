<script lang="ts" setup>
import draggable from 'vuedraggable';
import type { MountCategory } from '~/types';

definePageMeta({
  layout: 'empty',
  middleware: 'create-view',
});

const { data: page } = await useAsyncData('mounts', () =>
  queryContent('/collections/mounts/create-view').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No mounts create view page found in the content.',
  });
}

const mountsStore = useMountsStore();
const { allMounts } = storeToRefs(mountsStore);

const search = ref('');
const sidebarMounts = ref([...allMounts.value]);

const filteredMounts = computed(() => {
  if (!search.value) return sidebarMounts.value;
  return sidebarMounts.value.filter((mount) =>
    mount.name.toLowerCase().includes(search.value.toLowerCase())
  );
});

const cloneableCategory = ref<MountCategory>({
  id: useId(),
  name: 'New Category',
  mounts: [],
  categories: [],
  order: 0,
});
const categories = ref<MountCategory[]>([cloneableCategory.value]);

function onMove(event: {
  moved: { element: any; oldIndex: number; newIndex: number };
}) {
  if (event.moved.element) {
    const { element, oldIndex, newIndex } = event.moved;
    allMounts.value.splice(oldIndex, 1);
    allMounts.value.splice(newIndex, 0, element);
  }
}

function validate(event: Event) {
  const target = event.target as HTMLInputElement;
  target.blur();
  const { isRootCategory } = target.dataset;

  const categoryId = target.dataset.categoryId;
  const newName = target.innerText;

  if (isRootCategory) {
    const category = categories.value.find((cat) => cat.id === categoryId);
    if (category) {
      category.name = newName;
    }
  } else {
    const category = categories.value.find((cat) =>
      cat.categories?.find((subCat) => subCat.id === categoryId)
    );
    if (category) {
      const subCategory = category.categories?.find(
        (subCat) => subCat.id === categoryId
      );
      if (subCategory) {
        subCategory.name = newName;
      }
    }
  }
}

watch(
  () => categories,
  (newVal, oldVal) => {
    console.log('categories changed', newVal.value, oldVal.value);
  },
  { deep: true }
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
  <CreateViewContainer>
    <template #sidebar-content>
      <div class="flex flex-col mt-6 gap-10">
        <draggable
          :list="[cloneableCategory]"
          item-key="id"
          :group="{ name: 'mounts', pull: 'clone', put: false }"
        >
          <template #item>
            <UCard
              class="cursor-grab"
              :ui="{
                ring: 'ring-0 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out',
                rounded: 'rounded-none',
              }"
            >
              New category
            </UCard>
          </template>
        </draggable>
        <UInput
          v-model="search"
          placeholder="Search for a mount"
          icon="heroicons-outline:search"
        />
        <draggable
          class="flex grow flex-wrap gap-4 h-fit w-full p-2 border-[1px] border-gray-400 dark:border-gray-600"
          v-model="filteredMounts"
          :group="{ name: 'mounts', pull: true, put: true }"
          item-key="id"
          @change="onMove"
        >
          <template #item="{ element: mount }">
            <MountIcon
              class="select-none cursor-grab"
              :mount="mount"
              :clickable="false"
              build-mode
            />
          </template>
        </draggable>
      </div>
    </template>
    <template #main-content>
      <div class="flex flex-col w-full h-full p-10 gap-2 overflow-y-auto">
        <draggable
          class="flex flex-col gap-4 h-full w-full"
          v-model="categories"
          :group="{ name: 'mounts', pull: true, put: true }"
          item-key="name"
          @change="onMove"
        >
          <template #item="{ element: category }">
            <UCard
              class="flex flex-wrap gap-4 h-min w-full pb-10 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-grab"
              :ui="{ rounded: 'rounded-none' }"
            >
              <h2
                class="text-lg h-min w-fit px-2 border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text"
                contenteditable
                :spellcheck="false"
                :data-is-root-category="true"
                :data-category-id="category.id"
                @keydown.enter="validate"
                @blur="validate"
              >
                {{ category.name }}
              </h2>
              <draggable
                class="flex flex-wrap gap-4 w-full"
                v-model="category.mounts"
                :group="{ name: 'mounts', pull: true, put: true }"
                item-key="name"
                @change="onMove"
              >
                <template #item="{ element: mount }">
                  <MountIcon
                    class="select-none cursor-grab"
                    :mount="mount"
                    :clickable="false"
                    build-mode
                  />
                </template>
              </draggable>
            </UCard>
          </template>
        </draggable>
      </div>
    </template>
  </CreateViewContainer>
</template>

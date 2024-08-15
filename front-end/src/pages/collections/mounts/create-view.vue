<script lang="ts" setup>
import type { BuilderMountCategory, BuilderMountInformation } from '~/types';

// Set all mounts as one big array with two additional properties `category` to identify the category it belongs to and `level` to identify if it's root or sub.
// `level` is enum with values `root` and `sub`.
// The state in the store should have these properties and the getter should map the mounts to a regular mount category structure.

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

const toast = useToast();
const modal = useModal();

const search = ref('');
const viewBuilderStore = useViewBuilderStore();
const { mounts, flatMounts, _mountCategories } = storeToRefs(viewBuilderStore);

const mountCategories = computed(() => {
  return _mountCategories.value;
});

console.log(mounts.value);

const filteredMounts = computed(() => {
  if (!search.value) return flatMounts.value;
  return flatMounts.value?.filter((mount) =>
    mount.name.toLowerCase().includes(search.value.toLowerCase())
  );
});

function startDragClone(event: DragEvent, item: BuilderMountCategory) {
  console.log('startDragClone', event, item);
  if (!event.dataTransfer) return;

  event.dataTransfer.dropEffect = 'move';
  event.dataTransfer.effectAllowed = 'move';
  event.dataTransfer.setData('text/plain', JSON.stringify(item));
}
function startDrag(
  event: DragEvent,
  item: BuilderMountInformation | BuilderMountCategory
) {
  console.log('startDrag', event, item);
  if (!event.dataTransfer) return;

  event.dataTransfer.dropEffect = 'move';
  event.dataTransfer.effectAllowed = 'move';
  event.dataTransfer.setData('itemId', `${item.id}`);
}

// If null is passed as categoryId and level, it means the item is being dropped back to the sidebar
function onDrop(
  event: DragEvent,
  id: string | null,
  level: 'root' | 'sub' | null
) {
  const itemId = event.dataTransfer?.getData('itemId');
  const jsonItem = JSON.parse(
    event.dataTransfer?.getData('text/plain') || '{}'
  );
  const itemIsClonedCategory = jsonItem.hasOwnProperty('categories');

  console.log('1', { itemId, id, level });
  if (!itemId && !itemIsClonedCategory) return;

  if (itemIsClonedCategory) {
    if (level === null) return;

    if (level === 'root') {
      viewBuilderStore.addRootCategory(jsonItem as BuilderMountCategory);
    } else if (level === 'sub' && id) {
      viewBuilderStore.addSubCategory(jsonItem as BuilderMountCategory, id);
    }

    return;
  }

  const item =
    _mountCategories.value?.find((cat) => cat.id === itemId) ||
    mounts.value?.find((mount) => mount.id == itemId);

  console.log('2', { item });
  if (!item) return;

  const itemIsCategory = item.hasOwnProperty('categories');
  console.log('3', { itemIsCategory });

  const itemIsMount = item.hasOwnProperty('icon');
  console.log('4', { itemIsMount });

  if (itemIsCategory) {
    if (level === null) return;
    if (level === 'root') {
      const category = _mountCategories.value.find((cat) => cat.id === id);
      if (category) {
        category.categories?.push(item as BuilderMountCategory);
      }
    } else {
      const category = _mountCategories.value.find((cat) =>
        cat.categories?.find((subCat) => subCat.id === id)
      );
      if (category) {
        const subCategory = category.categories?.find(
          (subCat) => subCat.id === id
        );
        if (subCategory) {
          subCategory.categories?.push(item as BuilderMountCategory);
        }
      }
    }
  }

  // if (itemIsMount) {
  //   if (level === null && id === null) {
  //     viewBuilderStore.addMount(item as BuilderMountInformation);
  //   } else {
  //     const category = mountCategories.value.find((cat) => cat.id === id);
  //     if (category) {
  //       category.mounts?.push(item as BuilderMountInformation);
  //     } else {
  //       const category = mountCategories.value.find((cat) =>
  //         cat.categories?.find((subCat) => subCat.id === id)
  //       );
  //       if (category) {
  //         const subCategory = category.categories?.find(
  //           (subCat) => subCat.id === id
  //         );
  //         if (subCategory) {
  //           subCategory.mounts?.push(item as BuilderMountInformation);
  //         }
  //       }
  //     }
  //   }
  // }
}

watch(
  () => _mountCategories,
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

function validate(event: Event) {
  const target = event.target as HTMLInputElement;
  target.blur();
  const { isRootCategory } = target.dataset;

  const categoryId = target.dataset.categoryId;
  const newName = target.innerText;

  if (isRootCategory) {
    const category = _mountCategories.value.find(
      (cat) => cat.id === categoryId
    );
    if (category) {
      category.name = newName;
    }
  } else {
    const category = _mountCategories.value.find((cat) =>
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
</script>

<template>
  <CreateViewContainer>
    <template #sidebar-content>
      <div class="flex flex-col mt-6 gap-10">
        <UCard
          class="select-none cursor-grab"
          :draggable="true"
          :ui="{
            ring: 'ring-0 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out',
            rounded: 'rounded-none',
            body: {
              padding: 'px-2 py-2 sm:p-3',
            },
          }"
          @dragstart="
            startDragClone($event, {
              id: newId(10),
              name: 'New category',
              categories: [],
              mounts: [],
              order: 0,
            })
          "
        >
          New category
        </UCard>
        <UInput
          v-model="search"
          placeholder="Search for a mount"
          icon="heroicons-outline:search"
        />
        <div
          class="flex grow flex-wrap gap-4 justify-center"
          @drop="onDrop($event, null, null)"
          @dragenter.prevent
          @dragover.prevent
        >
          <MountIcon
            v-for="(mount, i) in filteredMounts"
            :key="i"
            :mount="mount"
            :clickable="false"
            build-mode
            class="select-none cursor-grab"
            @dragstart="startDrag($event, mount)"
          />
        </div>
      </div>
    </template>
    <template #main-content>
      <div
        class="flex flex-col w-full h-full p-10 gap-2 border-[1px] overflow-y-auto"
        @drop="onDrop($event, null, 'root')"
        @dragenter.prevent
        @dragover.prevent
      >
        <div
          class="flex flex-col gap-2 h-fit w-full p-2 pb-10 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-grab"
          v-for="category in mountCategories"
          :key="category.id"
          :draggable="true"
          @drop="onDrop($event, category.id as string, 'root')"
          @dragenter.prevent
          @dragover.prevent
        >
          <h2
            class="text-lg h-min w-fit px-2 border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text"
            contenteditable
            :spellcheck="false"
            @keydown.enter="validate"
            @blur="validate"
          >
            {{ category.name }}
          </h2>
        </div>
      </div>
    </template>
  </CreateViewContainer>
</template>

<!-- <script lang="ts" setup>
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

const viewBuilderStore = useViewBuilderStore();
const { _mounts: sidebarMounts, _mountCategories: categories } =
  storeToRefs(viewBuilderStore);

const search = ref('');

const filteredMounts = computed(() => {
  if (!search.value) return sidebarMounts.value;
  return sidebarMounts.value.filter((mount) =>
    mount.name.toLowerCase().includes(search.value.toLowerCase())
  );
});

const cloneableCategory = ref<MountCategory>({
  id: useId(),
  name: 'New category',
  mounts: [],
  categories: [],
  order: 0,
});

function onMove(event: {
  moved: { element: any; oldIndex: number; newIndex: number };
}) {
  if (event.moved.element) {
    const { element, oldIndex, newIndex } = event.moved;
    sidebarMounts.value?.splice(oldIndex, 1);
    sidebarMounts.value?.splice(newIndex, 0, element);
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
                body: {
                  padding: 'px-2 py-2 sm:p-3',
                },
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
          class="flex grow flex-wrap gap-4 justify-center h-fit w-full"
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
</template> -->

<script lang="ts" setup>
import type { BuilderMountCategory, BuilderMountInformation } from '~/types';

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
  if (!event.dataTransfer) return;

  event.dataTransfer.dropEffect = 'move';
  event.dataTransfer.effectAllowed = 'move';
  event.dataTransfer.setData('item', JSON.stringify(item));
  event.dataTransfer.setData('isNew', 'true');
}
function startDrag(
  event: DragEvent,
  item: BuilderMountInformation | BuilderMountCategory,
  from?: BuilderMountCategory
) {
  if (!event.dataTransfer) return;

  event.dataTransfer.dropEffect = 'move';
  event.dataTransfer.effectAllowed = 'move';
  event.dataTransfer.setData('item', JSON.stringify(item));
  if (from) event.dataTransfer.setData('from', JSON.stringify(from));
}

function getParsedJsonOr(value: string) {
  try {
    return JSON.parse(value);
  } catch (e) {
    return {};
  }
}

function onDrop(
  event: DragEvent,
  options?: {
    dropZoneParentId?: string;
    dropZoneKind?: 'root' | 'sub' | 'sub-mount' | 'side';
  }
) {
  event.stopPropagation();
  const item = event.dataTransfer?.getData('item');
  const parsedItem = getParsedJsonOr(item || '');

  const isNew = event.dataTransfer?.getData('isNew');

  const itemIsMount = parsedItem.hasOwnProperty('assets');
  const itemIsCategory = parsedItem.hasOwnProperty('categories');

  if (!itemIsMount && options?.dropZoneKind === 'sub-mount') return;

  if (isNew) {
    if (options?.dropZoneKind === 'root') {
      viewBuilderStore.addRootCategory(parsedItem as BuilderMountCategory);
    } else if (options?.dropZoneKind === 'sub') {
      viewBuilderStore.addSubCategory(
        parsedItem as BuilderMountCategory,
        options.dropZoneParentId!
      );
    }

    return;
  }

  if (itemIsCategory) {
    if (options?.dropZoneKind === 'root') {
      const alreadyExists = _mountCategories.value.find(
        (cat) => cat.id === parsedItem.id
      );
      if (alreadyExists) return;

      viewBuilderStore.addRootCategory(parsedItem as BuilderMountCategory);
    } else if (options?.dropZoneKind === 'sub') {
      const category = _mountCategories.value.find(
        (cat) => cat.id === options.dropZoneParentId
      );

      if (category) {
        const alreadyExists = category.categories?.find(
          (subCat) => subCat.id === parsedItem.id
        );
        if (alreadyExists) return;

        viewBuilderStore.addSubCategory(
          parsedItem as BuilderMountCategory,
          options.dropZoneParentId!
        );
      }
    }

    return;
  }
}

// function onDrop(
//   event: DragEvent,
//   id: string | null,
//   level: 'root' | 'sub' | null
// ) {
//   event.stopPropagation();
//   const itemId = event.dataTransfer?.getData('itemId');
//   const jsonItem = getParsedJsonOr(
//     event.dataTransfer?.getData('text/plain') || '',
//     {}
//   );

//   console.log(
//     // 'event',
//     // event,
//     'id',
//     id,
//     'level',
//     level,
//     'itemId',
//     itemId,
//     'jsonItem',
//     jsonItem
//   );

//   const itemIsClonedCategory =
//     jsonItem.hasOwnProperty('categories') && id === null;

//   if (!itemId && !itemIsClonedCategory) return;

//   if (itemIsClonedCategory) {
//     console.log('cloned category');
//     if (level === null) return;

//     if (level === 'root') {
//       viewBuilderStore.addRootCategory(jsonItem as BuilderMountCategory);
//     } else if (level === 'sub' && id) {
//       viewBuilderStore.addSubCategory(jsonItem as BuilderMountCategory, id);
//     }

//     return;
//   }

//   const item =
//     _mountCategories.value?.find((cat) => cat.id === itemId) ||
//     flatMounts.value?.find((mount) => mount.id == Number(itemId));

//   console.log('item', item);

//   if (!item) return;

//   const itemIsCategory = item.hasOwnProperty('categories');
//   const itemIsMount = item.hasOwnProperty('assets');

//   if (itemIsCategory) {
//     if (level === null) return;

//     if (level === 'root') {
//       const category = _mountCategories.value.find((cat) => cat.id === id);
//       if (category) {
//         category.categories?.push(item as BuilderMountCategory);
//       }
//     } else {
//       const category = _mountCategories.value.find((cat) => cat.id === id);
//       console.log('category', category);
//       if (category) {
//         category.categories?.push(item as BuilderMountCategory);
//         _mountCategories.value = _mountCategories.value.filter(
//           (cat) => cat.id !== itemId
//         );
//       }
//     }
//   }
//   if (itemIsMount) {
//     const category = _mountCategories.value?.find((cat) => {
//       return cat.id === id;
//     });
//     category?.mounts.push(item as any);
//   }
// }

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
          @drop="onDrop($event, { dropZoneKind: 'side' })"
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
        @drop="onDrop($event, { dropZoneKind: 'root' })"
        @dragenter.prevent
        @dragover.prevent
      >
        <div
          class="flex flex-col gap-2 h-fit w-full p-2 pb-10 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-grab"
          v-for="category in mountCategories"
          :key="category.id"
          :draggable="true"
          @drop="
            onDrop($event, {
              dropZoneKind: 'sub',
              dropZoneParentId: category.id,
            })
          "
          @dragstart="startDrag($event, category)"
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

            <div v-for="mount in category.mounts">
              {{ mount.name }}
            </div>
          </h2>

          <div
            class="flex flex-col gap-2 h-fit w-full p-2 pb-10 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-grab"
            v-for="subCategory in category.categories"
            :key="subCategory.id"
            :draggable="true"
            @drop="
              onDrop($event, {
                dropZoneKind: 'sub-mount',
                dropZoneParentId: subCategory.id,
              })
            "
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
              {{ subCategory.name }}
            </h2>

            <div v-for="mount in subCategory.mounts">
              {{ mount.name }}
            </div>
          </div>
        </div>
      </div>
    </template>
  </CreateViewContainer>
</template>

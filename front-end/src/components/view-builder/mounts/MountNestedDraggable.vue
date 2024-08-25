<script lang="ts" setup>
import draggable from 'vuedraggable';
import type { MountCategory, MountInformation } from '~/types';

const mountViewBuilderStore = useMountViewBuilderStore();
const { _mountCategories, _settings } = storeToRefs(mountViewBuilderStore);

const dragging = ref(false);
const dragItem = ref<MountCategory | MountInformation>();

const categoryCanBeDroppedInCategory = computed(() => {
  const itemIsCategory = dragItem.value?.hasOwnProperty('categories');
  if (!itemIsCategory) return false;

  const item = dragItem.value as MountCategory;
  return !item.categories?.length;
});

function removeCategory(categoryId: string, parentId?: string) {
  if (!parentId) {
    const category = _mountCategories.value.find(
      (cat) => cat.id === categoryId
    );
    if (!category) return;

    if (category.mounts?.length)
      mountViewBuilderStore.addMounts(category.mounts);

    if (category.categories?.length) {
      category.categories.forEach((subCat) => {
        if (subCat.mounts?.length)
          mountViewBuilderStore.addMounts(subCat.mounts);
      });
    }

    mountViewBuilderStore.removeRootCategory(categoryId);
    return;
  }

  const category = _mountCategories.value.find((cat) => cat.id === parentId);
  if (!category) return;

  const subCategory = category.categories?.find(
    (subCat) => subCat.id === categoryId
  );
  if (!subCategory) return;

  if (subCategory.mounts?.length)
    mountViewBuilderStore.addMounts(subCategory.mounts);
  mountViewBuilderStore.removeSubCategory(categoryId, parentId);
}

function dragRootStart(event: any) {
  mountViewBuilderStore.setDragState(true, 'category');
  const { category } = event.item.dataset;
  if (!category) return;
  dragItem.value = JSON.parse(category);
  dragging.value = true;
}

function dragRootEnd() {
  mountViewBuilderStore.clearDragState();
  dragItem.value = undefined;
  dragging.value = false;
}

function unFocus(event: Event) {
  const target = event.target as HTMLInputElement;
  target.blur();
}
</script>

<template>
  <draggable
    class="flex flex-col w-full h-full p-10 gap-2 overflow-y-auto"
    :list="_mountCategories"
    :group="{
      name: !categoryCanBeDroppedInCategory && dragging ? 'deny' : 'category',
    }"
    @start="dragRootStart"
    @end="dragRootEnd"
  >
    <template #item="{ element: category }">
      <div
        class="flex flex-col gap-5"
        :data-category="JSON.stringify(category)"
      >
        <UCard
          :ui="{
            ring: `ring-0 ${
              _settings.showBorders
                ? 'border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-move'
                : ''
            }`,
            rounded: 'rounded-none',
            body: {
              padding: 'p-0 sm:p-0',
            },
            shadow: _settings.showBorders ? '' : 'shadow-none',
          }"
        >
          <div class="flex justify-between">
            <UInput
              v-model="category.name"
              :ui="{
                wrapper: 'p-2',
                base: 'border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text w-min text-nowrap shadow-sm text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none',
                rounded: 'rounded-none',
              }"
              :data-is-root-category="false"
              :data-category-id="category.id"
              @keydown.enter="unFocus"
            />
            <UTooltip class="h-min mr-4 mt-4" text="Remove category">
              <UButton
                variant="ghost"
                color="red"
                icon="material-symbols:delete-rounded"
                @click="() => removeCategory(category.id)"
              />
            </UTooltip>
          </div>
          <draggable
            class="flex flex-wrap gap-2 min-h-10 p-4"
            :list="category.mounts"
            :group="{ name: 'mount' }"
            @start="_settings.showMountTooltips = false"
            @end="_settings.showMountTooltips = true"
          >
            <template #item="{ element: mount }">
              <MountIcon
                class="select-none cursor-grab"
                :mount="mount"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showMountTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
          <draggable
            v-if="category.categories"
            class="flex grow flex-wrap justify-start gap-2 min-h-10"
            :list="category.categories"
            :group="{ name: 'category' }"
            @start="mountViewBuilderStore.setDragState(true, 'category')"
            @end="mountViewBuilderStore.clearDragState"
          >
            <template #item="{ element: subCategory }">
              <div class="flex flex-col gap-5 p-4">
                <UCard
                  :ui="{
                    ring: _settings.showBorders
                      ? 'ring-0 border-[1px] border-dashed border-gray-400 dark:border-gray-600 hover:border-primary dark:hover:border-primary transition ease-in-out cursor-move'
                      : 'ring-0',
                    rounded: 'rounded-none',
                    body: {
                      padding: 'p-3 sm:p-3',
                    },
                    shadow: _settings.showBorders ? '' : 'shadow-none',
                  }"
                >
                  <div class="flex justify-between items-center gap-8">
                    <UInput
                      v-model="subCategory.name"
                      :ui="{
                        wrapper: 'p-2',
                        base: 'border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text w-min text-nowrap shadow-sm text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none',
                        rounded: 'rounded-none',
                      }"
                      :data-is-root-category="false"
                      :data-category-id="subCategory.id"
                      @keydown.enter="unFocus"
                    />
                    <UTooltip class="h-min" text="Remove category">
                      <UButton
                        variant="ghost"
                        color="red"
                        icon="material-symbols:delete-rounded"
                        @click="
                          () => removeCategory(subCategory.id, category.id)
                        "
                      />
                    </UTooltip>
                  </div>
                  <draggable
                    class="flex flex-wrap gap-2 py-5"
                    :list="subCategory.mounts"
                    :group="{ name: 'mount' }"
                    @start="_settings.showMountTooltips = false"
                    @end="_settings.showMountTooltips = true"
                  >
                    <template #item="{ element: mount }">
                      <MountIcon
                        class="select-none cursor-move"
                        :mount="mount"
                        :clickable="false"
                        build-mode
                        :show-tooltip="_settings.showMountTooltips"
                        :use-intersection-observer="false"
                      />
                    </template>
                  </draggable>
                </UCard>
              </div>
            </template>
          </draggable>
        </UCard>
      </div>
    </template>
  </draggable>
</template>

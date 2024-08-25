<script lang="ts" setup>
import draggable from 'vuedraggable';
import type { ToyCategory, ToyInformation } from '~/types';

const toyViewBuilderStore = useToyViewBuilderStore();
const { _toyCategories, _settings } = storeToRefs(toyViewBuilderStore);

const dragging = ref(false);
const dragItem = ref<ToyCategory | ToyInformation>();

const categoryCanBeDroppedInCategory = computed(() => {
  const itemIsCategory = dragItem.value?.hasOwnProperty('categories');
  if (!itemIsCategory) return false;

  const item = dragItem.value as ToyCategory;
  return !item.categories?.length;
});

function removeCategory(categoryId: string, parentId?: string) {
  if (!parentId) {
    const category = _toyCategories.value.find((cat) => cat.id === categoryId);
    if (!category) return;

    if (category.toys?.length) toyViewBuilderStore.addToys(category.toys);

    if (category.categories?.length) {
      category.categories.forEach((subCat) => {
        if (subCat.toys?.length) toyViewBuilderStore.addToys(subCat.toys);
      });
    }

    toyViewBuilderStore.removeRootCategory(categoryId);
    return;
  }

  const category = _toyCategories.value.find((cat) => cat.id === parentId);
  if (!category) return;

  const subCategory = category.categories?.find(
    (subCat) => subCat.id === categoryId
  );
  if (!subCategory) return;

  if (subCategory.toys?.length) toyViewBuilderStore.addToys(subCategory.toys);
  toyViewBuilderStore.removeSubCategory(categoryId, parentId);
}

function dragRootStart(event: any) {
  toyViewBuilderStore.setDragState(true, 'category');
  const { category } = event.item.dataset;
  if (!category) return;
  dragItem.value = JSON.parse(category);
  dragging.value = true;
}

function dragRootEnd() {
  toyViewBuilderStore.clearDragState();
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
    :list="_toyCategories"
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
              :data-is-root-category="true"
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
            :list="category.toys"
            :group="{ name: 'toy' }"
            @start="_settings.showToyTooltips = false"
            @end="_settings.showToyTooltips = true"
          >
            <template #item="{ element: toy }">
              <ToyIcon
                class="select-none cursor-grab"
                :toy="toy"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showToyTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
          <draggable
            v-if="category.categories"
            class="flex grow flex-wrap justify-start gap-2 min-h-10"
            :list="category.categories"
            :group="{ name: 'category' }"
            @start="toyViewBuilderStore.setDragState(true, 'category')"
            @end="toyViewBuilderStore.clearDragState"
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
                    :list="subCategory.toys"
                    :group="{ name: 'toy' }"
                    @start="_settings.showToyTooltips = false"
                    @end="_settings.showToyTooltips = true"
                  >
                    <template #item="{ element: toy }">
                      <ToyIcon
                        class="select-none cursor-move"
                        :toy="toy"
                        :clickable="false"
                        build-mode
                        :show-tooltip="_settings.showToyTooltips"
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

<script lang="ts" setup>
import draggable from 'vuedraggable';
import type { PetCategory, PetInformation } from '~/types';

const petViewBuilderStore = usePetViewBuilderStore();
const { _petCategories, _settings } = storeToRefs(petViewBuilderStore);

const dragging = ref(false);
const dragItem = ref<PetCategory | PetInformation>();

const categoryCanBeDroppedInCategory = computed(() => {
  const itemIsCategory = dragItem.value?.hasOwnProperty('categories');
  if (!itemIsCategory) return false;

  const item = dragItem.value as PetCategory;
  return !item.categories?.length;
});

function validate(event: Event) {
  const target = event.target as HTMLInputElement;
  target.blur();
  const { isRootCategory } = target.dataset;

  const categoryId = target.dataset.categoryId;
  const newName = target.innerText;

  if (isRootCategory) {
    const category = _petCategories.value.find((cat) => cat.id === categoryId);
    if (category) {
      category.name = newName;
    }
  } else {
    const category = _petCategories.value.find((cat) =>
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

function removeCategory(categoryId: string, parentId?: string) {
  if (!parentId) {
    const category = _petCategories.value.find((cat) => cat.id === categoryId);
    if (!category) return;

    if (category.pets?.length) petViewBuilderStore.addPets(category.pets);

    if (category.categories?.length) {
      category.categories.forEach((subCat) => {
        if (subCat.pets?.length) petViewBuilderStore.addPets(subCat.pets);
      });
    }

    petViewBuilderStore.removeRootCategory(categoryId);
    return;
  }

  const category = _petCategories.value.find((cat) => cat.id === parentId);
  if (!category) return;

  const subCategory = category.categories?.find(
    (subCat) => subCat.id === categoryId
  );
  if (!subCategory) return;

  if (subCategory.pets?.length) petViewBuilderStore.addPets(subCategory.pets);
  petViewBuilderStore.removeSubCategory(categoryId, parentId);
}

function dragRootStart(event: any) {
  petViewBuilderStore.setDragState(true, 'category');
  const { category } = event.item.dataset;
  if (!category) return;
  dragItem.value = JSON.parse(category);
  dragging.value = true;
}

function dragRootEnd() {
  petViewBuilderStore.clearDragState();
  dragItem.value = undefined;
  dragging.value = false;
}
</script>

<template>
  <draggable
    class="flex flex-col w-full h-full p-10 gap-2 overflow-y-auto"
    :list="_petCategories"
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
            <h2
              :class="[
                _settings.showBorders
                  ? 'p-2 ml-4 mt-4 border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text w-min text-nowrap shadow-sm text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none'
                  : 'p-2 ml-4 mt-4 cursor-text w-min text-nowrap text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none',
              ]"
              contenteditable
              :spellcheck="false"
              @keydown.enter="validate"
              @blur="validate"
              :data-is-root-category="true"
              :data-category-id="category.id"
            >
              {{ category.name }}
            </h2>
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
            :list="category.pets"
            :group="{ name: 'pet' }"
          >
            <template #item="{ element: pet }">
              <PetIcon
                class="select-none cursor-grab"
                :pet="pet"
                :clickable="false"
                build-mode
                :show-tooltip="_settings.showPetTooltips"
                :use-intersection-observer="false"
              />
            </template>
          </draggable>
          <draggable
            v-if="category.categories"
            class="flex grow flex-wrap justify-start gap-2 min-h-10"
            :list="category.categories"
            :group="{ name: 'category' }"
            @start="petViewBuilderStore.setDragState(true, 'category')"
            @end="petViewBuilderStore.clearDragState"
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
                    <h2
                      :class="[
                        _settings.showBorders
                          ? 'p-2 border-[1px] border-dashed border-gray-400 dark:border-gray-600 cursor-text w-min text-nowrap shadow-sm text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none'
                          : 'p-2 cursor-text w-min text-nowrap text-gray-900 dark:text-white focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400 focus:outline-none',
                      ]"
                      contenteditable
                      :spellcheck="false"
                      @keydown.enter="validate"
                      @blur="validate"
                      :data-is-root-category="false"
                      :data-category-id="subCategory.id"
                    >
                      {{ subCategory.name }}
                    </h2>
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
                    :list="subCategory.pets"
                    :group="{ name: 'pet' }"
                  >
                    <template #item="{ element: pet }">
                      <PetIcon
                        class="select-none cursor-move"
                        :pet="pet"
                        :clickable="false"
                        build-mode
                        :show-tooltip="_settings.showPetTooltips"
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

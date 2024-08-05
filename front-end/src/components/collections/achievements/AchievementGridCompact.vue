<script lang="ts" setup>
import type { AchievementCategory } from '~/types';

const loading = ref(false);
const loadingCategoryId = ref<number | null>(null);

const toast = useToast();
const achievementsStore = useAchievementsStore();
const characterStore = useCharacterStore();

const { achievements } = storeToRefs(achievementsStore);
const { character } = storeToRefs(characterStore);

const accordionItems = computed(() =>
  achievements.value.map((category, i) => ({
    ...category,
    ...((i === 0 || i === 1) && { defaultOpen: true }),
  }))
);

const fetchAchievementCategory = async (
  category: AchievementCategory,
  open: boolean
) => {
  if ((category.achievements && category.categories) || open) return;

  try {
    loading.value = true;
    loadingCategoryId.value = category.id;
    const data = await $fetch(
      `/api/character/${character.value?.region}/${character.value?.realm}/${character.value?.name}/achievements`,
      {
        query: { rootCategoryId: category.id },
      }
    );

    if (data?.category)
      achievementsStore.mergeAchievementCategory({
        ...data.category,
        id: category.id,
      });
  } catch (error) {
    console.error(error);
    toast.add({
      title: 'Error',
      description: (error as Error).message || 'Failed to fetch achievements',
      color: 'red',
    });
  } finally {
    loading.value = false;
    loadingCategoryId.value = null;
  }
};
</script>

<template>
  <UAccordion
    :items="accordionItems"
    :ui="{ wrapper: 'flex flex-col w-full' }"
    multiple
  >
    <template #default="{ item, open }">
      <UButton
        color="gray"
        variant="ghost"
        class="border-b border-gray-200 dark:border-gray-700"
        :ui="{ rounded: 'rounded-none', padding: { sm: 'p-3' } }"
        @click="() => fetchAchievementCategory(item, open)"
      >
        <span>{{ item.name }}</span>

        <template #trailing>
          <UIcon
            name="i-heroicons-chevron-right-20-solid"
            class="w-5 h-5 ms-auto transform transition-transform duration-200"
            :class="[open && 'rotate-90']"
          />
        </template>
      </UButton>
    </template>
    <template #item="{ item }">
      <!-- <UContainer
        class="flex flex-wrap w-full justify-center md:justify-start px-3 lg:px-3 sm:px-0 py-2 lg:py-2 sm:py-2 mx-0 gap-4"
      > -->

      <!-- <UProgress
        v-if="loading && loadingCategoryId === item.id"
        class="w-full"
        animation="carousel"
      /> -->

      <!-- <USkeleton
          v-for="i in 300"
          :key="i"
          class="h-8 w-8"
          :ui="{
            rounded: 'none',
          }"
        /> -->
      <!-- </UContainer> -->
      <UContainer
        v-if="item.achievements?.length"
        class="flex flex-wrap w-full justify-center md:justify-start px-3 lg:px-3 sm:px-0 py-2 lg:py-2 sm:py-2 mx-0 gap-4"
      >
        <div v-for="(achievement, j) in item.achievements" :key="j">
          <AchievementIcon :achievement="achievement" />
        </div>
      </UContainer>
      <UContainer
        class="flex flex-wrap w-full self-start px-3 lg:px-3 sm:px-0 py-2 lg:py-2 sm:py-2 mx-0 gap-4"
      >
        <UContainer
          v-for="(subCategory, j) in item.categories?.filter(
          (c: AchievementCategory) => c.achievements?.length
        )"
          :key="j"
          class="flex grow flex-wrap px-0 lg:px-0 sm:px-0 mx-0"
        >
          <div class="flex flex-col gap-4">
            <h3 class="text-xs">{{ subCategory.name }}</h3>
            <UContainer
              class="flex grow flex-wrap gap-2 px-0 lg:px-0 sm:px-0 mx-0"
            >
              <div
                v-for="(achievement, k) in subCategory.achievements"
                :key="k"
              >
                <AchievementIcon :achievement="achievement" />
              </div>
            </UContainer>
          </div>
        </UContainer>
      </UContainer>
    </template>
  </UAccordion>
</template>

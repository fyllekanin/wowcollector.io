<script lang="ts" setup>
import type { AchievementCategory } from '~/types';

const achievementsStore = useAchievementsStore();

const { achievements } = storeToRefs(achievementsStore);

onUnmounted(() => {
  [...document.getElementsByClassName('wowhead-tooltip')].forEach((item) => {
    item.remove();
  });
});
</script>

<template>
  <UAccordion
    :items="achievements"
    :ui="{ wrapper: 'flex flex-col w-full' }"
    multiple
  >
    <template #default="{ item, open }">
      <UButton
        color="gray"
        variant="ghost"
        class="border-b border-gray-200 dark:border-gray-700"
        :ui="{ rounded: 'rounded-none', padding: { sm: 'p-3' } }"
        @click="() => achievementsStore.toggleCategoryOpen(item.id)"
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
    <template #item="{ item, open }">
      <UContainer
        v-if="item.achievements?.length && open"
        class="flex flex-wrap w-full justify-center md:justify-start px-3 lg:px-3 sm:px-0 py-2 lg:py-2 sm:py-2 mx-0 gap-4"
      >
        <div v-for="(achievement, j) in item.achievements" :key="j">
          <AchievementIcon :achievement="achievement" />
        </div>
      </UContainer>
      <UContainer
        v-if="open"
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

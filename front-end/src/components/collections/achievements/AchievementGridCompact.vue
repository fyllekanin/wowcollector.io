<script lang="ts" setup>
const achievementsStore = useAchievementsStore();

const { achievements } = storeToRefs(achievementsStore);
</script>

<template>
  <UContainer
    v-for="(category, i) in achievements"
    :key="i"
    class="flex flex-col w-full justify-center md:justify-start px-0 lg:px-0 sm:px-0 mx-0 gap-2"
  >
    <h2 class="text-lg">{{ category.name }}</h2>
    <UContainer
      v-if="category.achievements?.length"
      class="flex flex-wrap w-full justify-center md:justify-start px-0 lg:px-0 sm:px-0 mx-0 gap-4"
    >
      <div v-for="(achievement, j) in category.achievements" :key="j">
        <AchievementIcon :achievement="achievement" />
      </div>
    </UContainer>
    <UContainer
      class="flex flex-wrap w-full self-start px-0 lg:px-0 sm:px-0 mx-0 gap-4"
    >
      <UContainer
        v-for="(subCategory, j) in category.categories?.filter(
          (c) => c.achievements?.length
        )"
        :key="j"
        class="flex grow flex-wrap px-0 lg:px-0 sm:px-0 mx-0"
      >
        <div class="flex flex-col gap-4">
          <h3 class="text-xs">{{ subCategory.name }}</h3>
          <UContainer
            class="flex grow flex-wrap gap-2 px-0 lg:px-0 sm:px-0 mx-0"
          >
            <div v-for="(achievement, k) in subCategory.achievements" :key="k">
              <AchievementIcon :achievement="achievement" />
            </div>
          </UContainer>
        </div>
      </UContainer>
    </UContainer>
  </UContainer>
</template>

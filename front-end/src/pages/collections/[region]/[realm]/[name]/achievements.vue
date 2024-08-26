<script lang="ts" setup>
const { data: page } = await useAsyncData('achievements', () =>
  queryContent('/collections/achievements').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No achievements page found in the content.',
  });
}

definePageMeta({
  middleware: 'collection-achievements',
});
useHead({
  title: 'WoW Collector - Achievements',
  meta: [
    {
      name: 'description',
      content: 'Achievements collection page',
    },
  ],
});
useSeoMeta({
  title: page.value.title,
  description: page.value.description,
  ogTitle: page.value.og.title,
  ogDescription: page.value.og.description,
  ogImage: page.value.og.image,
});

const achievementsStore = useAchievementsStore();
const { achievements } = storeToRefs(achievementsStore);

const total = computed(() => {
  if (!achievements.value) return 0;
  return flatMapAchievements(achievements.value).length;
});
const completed = computed(() => {
  if (!achievements.value) return 0;
  return flatMapAchievements(achievements.value).filter(
    (achievement) => achievement.isCompleted
  ).length;
});
const percentageAchievementsCompleted = computed(() => {
  if (!achievements.value) return 0;
  return Math.round((completed.value / total.value) * 100);
});
</script>

<template>
  <UContainer class="flex flex-col gap-4 pb-6">
    <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
    <CollectionHeader
      :progress="percentageAchievementsCompleted"
      :collected="completed"
      :total="total"
      collection="achievements"
    >
      <AchievementFilters />
    </CollectionHeader>

    <AchievementGridCompact />
  </UContainer>
  <ScrollToTop />
</template>

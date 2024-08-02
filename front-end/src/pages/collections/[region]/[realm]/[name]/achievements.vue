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
  middleware: 'achievements',
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
});

const achievementsStore = useAchievementsStore();
const { achievements } = storeToRefs(achievementsStore);

const availableAchievements = computed(() => {
  if (!achievements.value) return [];
  return flatMapAchievements(achievements.value);
});
const completedAchievements = computed(() => {
  if (!achievements.value) return [];
  return flatMapAchievements(achievements.value).filter(
    (mount) => mount.isCompleted
  );
});
const percentageAchievementsCompleted = computed(() => {
  if (!achievements.value) return 0;
  const flattenedAchievements = flatMapAchievements(achievements.value);
  console.log(
    flattenedAchievements.filter((achievement) => achievement.id === 8891)
  );
  const completedAchievements = flattenedAchievements.filter(
    (mount) => mount.isCompleted
  );
  return Math.round(
    (completedAchievements.length / flattenedAchievements.length) * 100
  );
});
</script>

<template>
  <UContainer class="flex flex-col gap-4">
    <CollectionHeader
      :progress="percentageAchievementsCompleted"
      :collected="completedAchievements"
      :available="availableAchievements"
      collection="achievements"
    >
      <AchievementFilters />
    </CollectionHeader>

    <AchievementGridCompact />
  </UContainer>
  <ScrollToTop />
</template>

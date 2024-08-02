<script lang="ts" setup>
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
  title: 'WoW Collector - Achievements',
  description: 'Achievements collection page',
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
      collection="mounts"
    >
      <AchievementFilters />
    </CollectionHeader>

    <AchievementGridCompact />
  </UContainer>
  <ScrollToTop />
</template>

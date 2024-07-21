<script lang="ts" setup>
const { data: page } = await useAsyncData('landing', () =>
  queryContent('/landing').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
  });
}

useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});
</script>

<template>
  <ULandingGrid class="max-w-[1200px] mx-4 lg:mx-auto mt-24">
    <ULandingCard
      v-for="(card, index) in page?.collectables"
      :key="index"
      :class="card.class"
      :icon="card.icon"
      :title="card.title"
      :description="card.description"
      :to="card.to"
    />
  </ULandingGrid>
</template>

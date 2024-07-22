<script lang="ts" setup>
const { data: page } = await useAsyncData('index', () =>
  queryContent('/').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No landing page found in the content.',
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
      v-for="(card, index) in page?.collections"
      :key="index"
      :class="card.class"
      :icon="card.icon"
      :title="card.title"
      :description="card.description"
      :to="card.to"
    />
  </ULandingGrid>
</template>

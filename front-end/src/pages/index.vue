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
  <UContainer class="flex flex-col gap-2 max-w-[1200px] max-4 lg:mx-auto pt-12">
    <ULandingGrid>
      <ULandingCard
        v-for="(card, index) in page?.services"
        :key="index"
        :class="card.class"
        :icon="card.icon"
        :title="card.title"
        :description="card.description"
        :to="card.to"
        :color="card.color"
      />
    </ULandingGrid>
  </UContainer>
</template>

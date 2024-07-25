<script lang="ts" setup>
const { data: page } = await useAsyncData('search', () =>
  queryContent('/search').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No search page found in the content.',
  });
}

useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

defineEmits(['success']);

const { isMobile } = useScreenSize();
</script>

<template>
  <UContainer class="w-full flex items-center justify-center">
    <UCard
      class="w-96"
      :ui="{
        ring: isMobile ? '' : 'ring-1 ring-gray-200 dark:ring-gray-800',
        shadow: isMobile ? 'shadow-none' : 'shadow',
      }"
    >
      <CharacterSearchForm @success="() => $emit('success')" />
    </UCard>
  </UContainer>
</template>

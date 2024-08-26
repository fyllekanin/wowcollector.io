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
  ogTitle: page.value.og.title,
  ogDescription: page.value.og.description,
  ogImage: page.value.og.image,
});

const router = useRouter();

function onSuccess() {
  router.push({
    path: '/',
  });
}
</script>

<template>
  <UContainer class="flex flex-col gap-4 max-w-[1200px] max-4 lg:mx-auto">
    <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
    <CharacterSearchCard @success="onSuccess" />
  </UContainer>
</template>

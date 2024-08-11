<script lang="ts" setup>
// Temporary disables the page until the content is ready
definePageMeta({
  middleware: () => navigateTo('/'),
});

const { data: page } = await useAsyncData('about', () =>
  queryContent('/footer/about').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No about page found in the content.',
  });
}
</script>

<template>
  <UContainer class="flex flex-col gap-5 max-w-[1200px] mx-auto mb-10">
    <h1 class="text-2xl font-bold">{{ page?.title }}</h1>
    <p class="mt-2 text-md">{{ page?.introduction }}</p>

    <div
      class="flex flex-col gap-5"
      v-for="(section, i) in page?.sections"
      :key="i"
    >
      <h2 class="text-xl font-bold">{{ section.title }}</h2>
      <p class="mt-2 text-md">{{ section.content }}</p>

      <div v-for="(subsection, j) in section.subsections" :key="j">
        <h3 class="text-lg font-bold">{{ subsection.title }}</h3>
        <p class="mt-2 text-md">{{ subsection.content }}</p>
      </div>
    </div>

    <p>{{ page?.conclusion }}</p>
  </UContainer>
</template>

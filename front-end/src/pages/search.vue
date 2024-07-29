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

const router = useRouter();
const characterStore = useCharacterStore();

const { character } = storeToRefs(characterStore);

function onSuccess() {
  router.push({
    path: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/mounts`,
  });
}
</script>

<template>
  <UContainer class="w-full flex items-center justify-center">
    <CharacterSearchCard @success="onSuccess" />
  </UContainer>
</template>

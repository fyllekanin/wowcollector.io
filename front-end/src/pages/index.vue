<script lang="ts" setup>
import CharacterSearchModal from '~/components/modals/CharacterSearchModal.vue';

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

const { character } = useCharacterStore();
const modal = useModal();
const router = useRouter();

function onSearch(to: string) {
  modal.open(CharacterSearchModal, {
    onSuccess: () => {
      router.push(to);
      modal.close();
    },
  });
}
</script>

<template>
  <ULandingGrid class="max-w-[1200px] mx-4 lg:mx-auto">
    <ULandingCard
      v-for="(card, index) in page?.collections"
      class="cursor-pointer"
      :key="index"
      :class="card.class"
      :icon="card.icon"
      :title="card.title"
      :description="card.description"
      :to="character ? card.to : undefined"
      :color="card.color"
      @click="character ? undefined : () => onSearch(card.to)"
    />
  </ULandingGrid>
</template>

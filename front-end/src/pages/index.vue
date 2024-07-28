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

const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);
const modal = useModal();
const router = useRouter();

function onSearch(to: string) {
  modal.open(CharacterSearchModal, {
    onSuccess: () => {
      router.push(
        `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/${to}`
      );
      modal.close();
    },
  });
}
</script>

<template>
  <ULandingGrid class="max-w-[1200px] mx-4 lg:mx-auto pt-12">
    <ULandingCard
      v-for="(card, index) in page?.collections"
      class="cursor-pointer"
      :key="index"
      :class="card.class"
      :icon="card.icon"
      :title="card.title"
      :description="card.description"
      :to="
        character
          ? `/collections/${character.region}/${character.realm}/${character.name}/${card.to}`
          : undefined
      "
      :color="card.color"
      :ui="{
        // hack to fix the hover effect on the cards until it's fixed in Nuxt UI.
        to: character
          ? `hover:ring-${card.color}-500 dark:hover:ring-${card.color}-400 transition-shadow duration-200`
          : `hover:ring-${card.color}-500 dark:hover:ring-${card.color}-400 transition-shadow duration-200`,
      }"
      @click="character ? undefined : onSearch(card.to)"
    />
  </ULandingGrid>
</template>

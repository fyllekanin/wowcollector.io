<script lang="ts" setup>
import CharacterSearchModal from '~/components/modals/CharacterSearchModal.vue';

const { data: page } = await useAsyncData('collections', () =>
  queryContent('/collections').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No collections page found in the content.',
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
  <UContainer class="flex flex-col gap-6 max-w-[1200px] max-4 lg:mx-auto pb-12">
    <UBreadcrumb :links="mapContentNavigation(page?.breadcrumbs)" />
    <ULandingGrid>
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
        @click="character ? undefined : onSearch(card.to)"
      />
    </ULandingGrid>
  </UContainer>
</template>

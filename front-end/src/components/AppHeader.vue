<script lang="ts" setup>
import CharacterSearchModal from './modals/CharacterSearchModal.vue';

import type { HeaderLink } from '@nuxt/ui-pro/types';

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

const links = [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collections',
    to: '/collections',
    children: [
      {
        label: 'Mounts',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/mounts`,
        disabled: !character.value,
        click: () => {
          if (!character.value) {
            onSearch('mounts');
          }
        },
      },
      {
        label: 'Achievements',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/achievements`,
        disabled: !character.value,
        click: () => {
          if (!character.value) {
            onSearch('achievements');
          }
        },
      },
      {
        label: 'Pets',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/pets`,
        disabled: !character.value,
        click: () => {
          if (!character.value) {
            onSearch('pets');
          }
        },
      },
      {
        label: 'Toys',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/toys`,
        disabled: !character.value,
        click: () => {
          if (!character.value) {
            onSearch('toys');
          }
        },
      },
    ],
  },
  {
    label: 'Leaderboards',
    to: '/leaderboards',
  },
  {
    label: 'Character Search',
    to: '/search',
  },
] as HeaderLink[];
</script>

<template>
  <UHeader :links="links">
    <template #logo>
      <Logo width="48px" height="48px" />
      <p class="hidden min-[375px]:block self-center mt-1">WOW Collector</p>
    </template>

    <template #right>
      <UButton icon="simple-icons:battledotnet" to="/login" color="gray"
        >Sign In</UButton
      >
      <UColorModeButton class="hidden lg:flex" />
    </template>

    <template #panel>
      <UNavigationTree :links="links" default-open />
      <UColorModeSelect class="lg:hidden pt-5" />
    </template>
  </UHeader>
</template>

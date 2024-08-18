<script lang="ts" setup>
import CharacterSearchModal from './modals/CharacterSearchModal.vue';

import type { HeaderLink } from '@nuxt/ui-pro/types';

const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);

const characterExists = computed(() => !!character.value);

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

const links = computed(() => [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collections',
    children: [
      {
        label: 'Mounts',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/mounts`,
        disabled: !characterExists.value,
        click: () => {
          if (!characterExists.value) {
            onSearch('mounts');
          }
        },
      },
      {
        label: 'Achievements',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/achievements`,
        disabled: !characterExists.value,
        click: () => {
          if (!characterExists.value) {
            onSearch('achievements');
          }
        },
      },
      {
        label: 'Pets',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/pets`,
        disabled: !characterExists.value,
        click: () => {
          if (!characterExists.value) {
            onSearch('pets');
          }
        },
      },
      {
        label: 'Toys',
        to: `/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/toys`,
        disabled: !characterExists.value,
        click: () => {
          if (!characterExists.value) {
            onSearch('toys');
          }
        },
      },
    ],
  },
  {
    label: 'Character Search',
    to: '/search',
  },
  {
    label: 'View Builder',
    children: [
      {
        label: 'Mounts',
        to: '/collections/mounts/view-builder',
      },
      {
        label: 'Achievements',
        to: '/collections/achievements/view-builder',
      },
      {
        label: 'Pets',
        to: '/collections/pets/view-builder',
      },
      {
        label: 'Toys',
        to: '/collections/toys/view-builder',
      },
    ],
  },
]) as ComputedRef<HeaderLink[]>;
</script>

<template>
  <UHeader :links="links">
    <template #logo>
      <Logo width="48px" height="48px" />
      <p class="hidden min-[375px]:block self-center mt-1">WoW Collector</p>
    </template>

    <template #right>
      <UColorModeButton class="hidden lg:flex" />
      <ActiveCharacterDropdown v-if="character" class="hidden lg:flex" />
    </template>

    <template #panel>
      <div class="flex flex-col gap-8">
        <UNavigationTree :links="links" default-open />
        <UDivider />
        <div class="flex grow flex-wrap items-center justify-between gap-2">
          <div class="flex items-center gap-3">
            <UAvatar :src="character?.assets?.avatar" />
            <div class="flex gap-1">
              <span class="text-sm"
                >{{ character?.name }} - {{ character?.realm }}</span
              >
            </div>
          </div>
          <UButton
            icon="material-symbols:logout"
            to="/search"
            variant="ghost"
            color="red"
            @click="characterStore.clearCharacter"
            >Change character</UButton
          >
        </div>
        <UColorModeSelect class="lg:hidden" />
      </div>
    </template>
  </UHeader>
</template>

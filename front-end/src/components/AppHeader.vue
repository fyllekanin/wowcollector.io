<script lang="ts" setup>
import CharacterSearchModal from './modals/CharacterSearchModal.vue';

import { Icons } from '~/constants';

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
        description: 'View your mount collection.',
        icon: Icons.MOUNTS,
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
        description: 'View your achievement collection.',
        icon: Icons.ACHIEVEMENTS,
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
        description: 'View your pet collection.',
        icon: Icons.PETS,
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
        description: 'View your toy collection.',
        icon: Icons.TOYS,
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
        label: 'Discovery',
        description: 'Discover views created by the community.',
        to: '/collections/view-builder/discovery',
        icon: Icons.COMPASS,
      },
      {
        label: 'Mounts',
        description: 'Create a custom view and share it with others!',
        to: '/collections/view-builder/mounts',
        icon: Icons.MOUNTS,
      },
      {
        label: 'Pets',
        description: 'Create a custom view and share it with others!',
        to: '/collections/view-builder/pets',
        icon: Icons.PETS,
      },
      {
        label: 'Toys',
        description: 'Create a custom view and share it with others!',
        to: '/collections/view-builder/toys',
        icon: Icons.TOYS,
      },
    ],
  },
]) as ComputedRef<HeaderLink[]>;
</script>

<template>
  <UHeader :links="links">
    <template #logo>
      <LogoFull />
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

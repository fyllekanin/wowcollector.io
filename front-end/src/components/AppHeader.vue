<script lang="ts" setup>
import type { HeaderLink, NavigationTree } from '@nuxt/ui-pro/types';

const links = [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collectables',
    children: [
      {
        label: 'Mounts',
        to: '/collectables/mounts',
      },
      {
        label: 'Achievements',
        to: '/collectables/achievements',
        disabled: true,
      },
      {
        label: 'Pets',
        to: '/collectables/pets',
        disabled: true,
      },
      {
        label: 'Toys',
        to: '/collectables/toys',
        disabled: true,
      },
    ],
  },
] as HeaderLink[];

const navigationTreeLinks = [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collectables',
    children: [
      {
        label: 'Mounts',
        to: '/collectables/mounts',
      },
      {
        label: 'Achievements',
        to: '/collectables/achievements',
        disabled: true,
      },
      {
        label: 'Pets',
        to: '/collectables/pets',
        disabled: true,
      },
      {
        label: 'Toys',
        to: '/collectables/toys',
        disabled: true,
      },
    ],
  },
] as NavigationTree[];
</script>

<template>
  <UHeader :links="links">
    <template #logo>
      <Logo width="48px" height="48px" />
      <p class="self-center mt-1">WoW Collector</p>
    </template>

    <template #right>
      <UButton icon="simple-icons:battledotnet" to="/login" color="gray"
        >Sign in</UButton
      >
    </template>

    <template #panel>
      <UNavigationTree :links="navigationTreeLinks" default-open />
    </template>
  </UHeader>
</template>

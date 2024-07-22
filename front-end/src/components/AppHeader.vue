<script lang="ts" setup>
import type { HeaderLink, NavigationTree } from '@nuxt/ui-pro/types';

const links = [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collections',
    children: [
      {
        label: 'Mounts',
        to: '/collections/mounts',
      },
      {
        label: 'Achievements',
        to: '/collections/achievements',
        disabled: true,
      },
      {
        label: 'Pets',
        to: '/collections/pets',
        disabled: true,
      },
      {
        label: 'Toys',
        to: '/collections/toys',
        disabled: true,
      },
    ],
  },
  {
    label: 'Search',
    to: '/search',
  },
] as HeaderLink[];

const navigationTreeLinks = [
  {
    label: 'Home',
    to: '/',
  },
  {
    label: 'Collections',
    children: [
      {
        label: 'Mounts',
        to: '/collections/mounts',
      },
      {
        label: 'Achievements',
        to: '/collections/achievements',
        disabled: true,
      },
      {
        label: 'Pets',
        to: '/collections/pets',
        disabled: true,
      },
      {
        label: 'Toys',
        to: '/collections/toys',
        disabled: true,
      },
    ],
  },
  {
    label: 'Search',
    to: '/search',
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
      <UColorModeButton />
    </template>

    <template #panel>
      <UNavigationTree :links="navigationTreeLinks" default-open />
    </template>
  </UHeader>
</template>

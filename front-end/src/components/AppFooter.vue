<script setup lang="ts">
import type { FooterLink } from '@nuxt/ui-pro/types';

const links = [
  {
    label: 'About Us',
    to: '/about',
  },
  {
    label: 'Releases',
    to: '/releases',
  },
  {
    label: 'Suggestions',
    to: '/suggestions',
  },
] as FooterLink[];
</script>

<template>
  <UFooter :links="links" />
</template>

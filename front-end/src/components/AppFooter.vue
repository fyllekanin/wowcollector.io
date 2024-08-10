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
    label: 'Feedback',
    to: '/feedback',
  },
] as FooterLink[];
</script>

<template>
  <UDivider />
  <UFooter :links="links" />
</template>

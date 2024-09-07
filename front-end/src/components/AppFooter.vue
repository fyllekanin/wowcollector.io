<script setup lang="ts">
import CookieModal from '~/components/cookie/CookieModal.vue';

import type { FooterLink } from '@nuxt/ui-pro/types';

const modal = useModal();

const links = [
  // {
  //   label: 'About Us',
  //   to: '/about',
  // },
  {
    label: 'Releases',
    to: '/releases',
  },
  {
    label: 'Feedback',
    to: '/feedback',
  },
  {
    label: 'Terms of Service',
    to: '/terms',
  },
  {
    label: 'Privacy Policy',
    to: '/privacy',
  },
  {
    label: 'Cookies',
    click: () => {
      modal.open(CookieModal, {
        onClose: () => {
          modal.close();
        },
      });
    },
  },
] as FooterLink[];
</script>

<template>
  <UDivider />
  <UFooter :links="links" />
</template>

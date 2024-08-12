<script lang="ts" setup>
import type { PropType } from 'vue';
import type { PetInformation } from '~/types';

defineProps({
  pet: {
    type: Object as PropType<PetInformation>,
    required: true,
  },
});
</script>

<template>
  <a
    :href="`https://www.wowhead.com/battle-pet/${pet.id}`"
    target="_blank"
    :data-wowhead="`battle-pet=${pet.id}`"
  >
    <img
      :src="pet.assets.largeIcon"
      :class="[
        !pet.isCollected
          ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
          : '',
      ]"
      width="38"
      @error="
        pet.assets.largeIcon =
          'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
      "
    />
  </a>
</template>

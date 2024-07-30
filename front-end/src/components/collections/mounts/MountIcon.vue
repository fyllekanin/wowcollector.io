<script lang="ts" setup>
import type { PropType } from 'vue';
import type { MountInformation } from '~/types';

defineProps({
  mount: {
    type: Object as PropType<MountInformation>,
    required: true,
  },
});
</script>

<template>
  <Lazy>
    <a
      :href="`https://www.wowhead.com/mount/${mount.id}`"
      target="_blank"
      :data-wowhead="`mount=${mount.id}`"
    >
      <img
        :src="mount.assets.largeIcon"
        :class="[
          !mount.isCollected
            ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
            : '',
        ]"
        width="38"
        :on-error="(e: any) => (e.target.src = mount.creatureDisplay)"
        @error="
          mount.assets.smallIcon =
            'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
        "
      />
    </a>
  </Lazy>
</template>

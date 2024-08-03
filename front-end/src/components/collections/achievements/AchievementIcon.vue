<script lang="ts" setup>
import type { PropType } from 'vue';
import type { AchievementInformation } from '~/types';

defineProps({
  achievement: {
    type: Object as PropType<AchievementInformation>,
    required: true,
  },
});
</script>

<template>
  <Lazy>
    <a
      :href="`https://www.wowhead.com/achievement/${achievement.id}`"
      target="_blank"
      :data-wowhead="`achievement=${achievement.id}`"
    >
      <nuxt-img
        :src="achievement.icon"
        :class="[
          !achievement.isCompleted
            ? 'brightness-50 grayscale blur-[1px] transition ease-in-out hover:grayscale-0 hover:blur-[0px] hover:brightness-100 hover:ring-1 hover:ring-primary'
            : '',
        ]"
        width="38"
        loading="lazy"
        @error="
          achievement.icon =
            'https://wow.zamimg.com/images/wow/icons/large/inv_misc_questionmark.jpg'
        "
      />
    </a>
  </Lazy>
</template>

<script lang="ts" setup>
import ConfettiExplosion from 'vue-confetti-explosion';

const viewId = defineModel<string>({
  required: true,
});
defineProps({
  to: {
    type: String as PropType<'mounts' | 'pets' | 'toys' | 'achievements'>,
    required: true,
  },
});
const emit = defineEmits(['leave']);

const clipboard = useClipboard();
const toast = useToast();
const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);

const visible = ref(true);

const explode = async () => {
  visible.value = false;
  await nextTick();
  visible.value = true;
};

function copyLink() {
  clipboard.copy(viewId.value);
  toast.add({
    title: 'View ID copied',
    color: 'green',
    timeout: 2000,
  });
}

onMounted(() => {
  clipboard.copy(viewId.value);
  toast.add({
    title: 'View ID copied',
    color: 'green',
    timeout: 2000,
  });
});

onBeforeUnmount(() => {
  emit('leave');
});
onBeforeRouteLeave(() => {
  emit('leave');
});
</script>

<template>
  <div class="flex flex-col items-center justify-center h-full w-full">
    <ConfettiExplosion v-if="visible" :stageHeight="1300" />
    <div class="flex flex-col items-center justify-center mt-20 gap-5 w-full">
      <div class="flex gap-2 items-center">
        <h2 class="text-2xl font-bold text-center">Congratulations!</h2>
        <span class="cursor-pointer select-none" @click="explode"> 🎉 </span>
      </div>
      <h3 class="text-lg font-semibold text-center">
        You have successfully created a new view!
      </h3>
      <p class="text-sm text-center mt-2">
        Here's your unique view ID, keep it safe as this is the only way to
        access it currently. You can load the view by going to the collections
        page and entering the view ID.
      </p>
      <UButtonGroup>
        <UInput
          :style="`width: ${viewId.length + 5}ch; max-width: 80vw;`"
          v-model="viewId"
          readonly
          icon="heroicons-outline:link"
          placeholder="Link to view"
        />
        <UTooltip text="Go to view">
          <UButton
            @click="copyLink"
            icon="heroicons-outline:clipboard"
            color="gray"
          />
        </UTooltip>
      </UButtonGroup>
    </div>
    <div class="flex gap-2 items-center mt-8">
      <UButton to="/" icon="heroicons-outline:home" color="gray">
        Go back home
      </UButton>
      <div v-if="character" class="flex gap-2 items-center mt-8">
        or
        <UButton
          v-if="character"
          :to="`/collections/${character.region}/${character.realm}/${character.name}/${to}?viewId=${viewId}`"
          icon="heroicons-outline:eye"
          color="gray"
        >
          Go to view
        </UButton>
      </div>
    </div>
    <ConfettiExplosion />
  </div>
</template>

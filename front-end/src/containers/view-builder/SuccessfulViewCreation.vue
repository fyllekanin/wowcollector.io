<script lang="ts" setup>
import ConfettiExplosion from 'vue-confetti-explosion';

const viewId = defineModel<string>({
  required: true,
});
const props = defineProps({
  to: {
    type: String,
    required: true,
    enum: ['mounts', 'pets', 'toys', 'achievements'],
  },
});

const clipboard = useClipboard();
const toast = useToast();
const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);

const runtimeConfig = useRuntimeConfig();

const link = ref(
  character.value
    ? `${runtimeConfig.public.baseURLFront}/collections/${character.value?.region}/${character.value?.realm}/${character.value?.name}/${props.to}?viewId=${viewId.value}`
    : viewId.value
);

function copyLink() {
  clipboard.copy(link.value);
  toast.add({
    title: 'Link copied',
    color: 'green',
    timeout: 2000,
  });
}
</script>

<template>
  <div class="flex flex-col items-center justify-center h-full w-full">
    <ConfettiExplosion />
    <div class="flex flex-col items-center justify-center mt-20 gap-5 w-full">
      <h2 class="text-2xl font-bold text-center">Congratulations!</h2>
      <h3 class="text-lg font-semibold text-center">
        You have successfully created a new view!
      </h3>
      <p class="text-sm text-center mt-2">
        Here's your link to view it, keep it safe as this is the only way to
        access it currently.
      </p>
      <UButtonGroup>
        <UInput
          :style="`width: ${link.length}ch; max-width: 80vw;`"
          v-model="link"
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
    <div class="mt-8">
      <UButton v-if="character" :to="link"> Go to view </UButton>
    </div>
    <ConfettiExplosion />
  </div>
</template>

<script lang="ts" setup>
const characterStore = useCharacterStore();
const { character } = storeToRefs(characterStore);

const items = computed(() => [
  [
    {
      name: character.value?.name,
      realm: character.value?.realm,
      slot: 'account',
      disabled: true,
    },
  ],
  [
    {
      label: 'Change character',
      icon: 'material-symbols:logout',
      to: '/search',
    },
  ],
]);
</script>

<template>
  <UDropdown
    mode="hover"
    :close-delay="50"
    :items="items"
    :ui="{ item: { disabled: 'cursor-text select-text' } }"
    :popper="{ placement: 'bottom-start' }"
  >
    <UAvatar :src="character?.assets?.avatar" />

    <template #account="{ item }">
      <div class="text-left flex flex-col">
        <span class="font-bold">{{ item.name }}</span>
        <span>{{ item.realm }}</span>
      </div>
    </template>

    <template #item="{ item }">
      <UButton
        :icon="item.icon"
        :to="item.to"
        :disabled="item.disabled"
        variant="ghost"
        color="red"
        class="w-full"
        @click="characterStore.clearCharacter"
        >{{ item.label }}</UButton
      >
    </template>
  </UDropdown>
</template>

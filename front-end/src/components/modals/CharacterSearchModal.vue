<script lang="ts" setup>
defineEmits(['error', 'success']);
</script>

<template>
  <UModal>
    <UCard>
      <CharacterSearchForm @success="$emit('success')" />
    </UCard>
  </UModal>
</template>

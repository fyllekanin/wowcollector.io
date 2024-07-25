<script lang="ts" setup>
import { object, string, type InferType } from 'yup';
import type { FormError, FormSubmitEvent } from '#ui/types';
import type { RealmsRegions } from '~/types';

const { data: page } = await useAsyncData('search', () =>
  queryContent('/search').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No search page found in the content.',
  });
}

const { data: realmsRegions } = await useAsyncData<RealmsRegions>(
  'realms-regions',
  () => $fetch('/api/battle-net/realms-regions')
);

useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

const emit = defineEmits(['success', 'error']);

const schema = object({
  name: string().required(),
  realm: string().required(),
  region: string().required(),
});

type Schema = InferType<typeof schema>;
const state = reactive({
  name: undefined,
  realm: undefined,
  region: 'eu',
});
const loading = ref(false);

const realms = computed(() =>
  realmsRegions.value?.realms
    .filter(({ region }) => state.region === region)
    .map(({ name, slug }) => ({ label: name, value: slug }))
);
const regions = computed(() =>
  realmsRegions.value?.regions.map(({ name, value }) => ({
    label: name,
    value,
  }))
);

async function onSubmit(event: FormSubmitEvent<Schema>) {
  loading.value = true;

  try {
    emit('success');
  } catch (error) {
    console.error(error);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
    <p>Search for a character</p>

    <UDivider />

    <UFormGroup label="Character Name" name="name">
      <UInput v-model="state.name" />
    </UFormGroup>

    <UFormGroup label="Realm" name="realm">
      <USelectMenu
        v-model="state.realm"
        searchable
        value-attribute="value"
        :options="realms"
      />
    </UFormGroup>

    <UFormGroup label="Region" name="region">
      <URadioGroup
        class="[&>fieldset]:flex [&>fieldset]:gap-3"
        v-model="state.region"
        :options="regions"
      />
    </UFormGroup>

    <UDivider />

    <div class="w-full flex justify-end">
      <UButton type="submit" :loading="loading"> Search </UButton>
    </div>
  </UForm>
</template>

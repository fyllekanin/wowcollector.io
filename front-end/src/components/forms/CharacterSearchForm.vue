<script lang="ts" setup>
import { object, string, type InferType } from 'yup';
import type { FormSubmitEvent } from '#ui/types';

import type {
  RealmsRegions,
  CollectionInformationResponse,
  MountCategory,
} from '~/types';

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
useSeoMeta({
  title: page.value.title,
  description: page.value.description,
});

const toast = useToast();
const characterStore = useCharacterStore();
const mountsStore = useMountsStore();
const realmsRegionsStore = useRealmsRegionsStore();
const { realms, regions } = storeToRefs(realmsRegionsStore);

if (!realms.value || !regions.value) {
  const { data: realmsRegions, error } = await useAsyncData<RealmsRegions>(
    'realms-regions',
    () => $fetch('/api/battle-net/realms-regions')
  );

  if (!realmsRegions.value) {
    throw createError({
      statusCode: error.value?.statusCode || 500,
      statusMessage:
        error.value?.statusMessage ||
        error.value?.message ||
        'Internal Server Error',
      fatal: true,
      cause: error,
    });
  }

  realmsRegionsStore.setRealms(realmsRegions.value?.realms);
  realmsRegionsStore.setRegions(realmsRegions.value?.regions);
}

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

const realmsMapped = computed(() =>
  realms.value
    ?.filter(({ region }) => state.region === region)
    .map(({ name, slug }) => ({ label: name, value: slug }))
);
const regionsMapped = computed(() =>
  regions.value?.map(({ name, value }) => ({
    label: name,
    value,
  }))
);

async function onSubmit(event: FormSubmitEvent<Schema>) {
  loading.value = true;

  const { data: character } = event;

  try {
    const mounts = await $fetch<CollectionInformationResponse<MountCategory>>(
      `/api/character/${character.region}/${character.realm}/${character.name}/mounts`
    );

    if (!mounts) {
      throw new Error('Character not found');
    }

    characterStore.setCharacter(character);
    mountsStore.setMounts(mounts);

    emit('success');
  } catch (error) {
    console.error(error);
    loading.value = false;

    toast.add({
      title: 'Error',
      // @ts-expect-error - error is an instance of Error
      description: error.message,
      color: 'red',
    });

    emit('error', error);
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
      <UInputMenu
        v-model="state.realm"
        searchable
        value-attribute="value"
        :options="realmsMapped"
      />
    </UFormGroup>

    <UFormGroup label="Region" name="region">
      <URadioGroup
        class="[&>fieldset]:flex [&>fieldset]:gap-3"
        v-model="state.region"
        :options="regionsMapped"
      />
    </UFormGroup>

    <UDivider />

    <div class="w-full flex justify-end">
      <UButton type="submit" :loading="loading"> Search </UButton>
    </div>
  </UForm>
</template>

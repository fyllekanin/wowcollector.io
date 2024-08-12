<script lang="ts" setup>
import ConfirmationModal from '~/components/modals/ConfirmationModal.vue';

const { data: page } = await useAsyncData('feedback', () =>
  queryContent('/footer/feedback').findOne()
);
if (!page.value) {
  throw createError({
    statusCode: 404,
    statusMessage: 'Page not found',
    fatal: true,
    cause: 'No feedback page found in the content.',
  });
}

const state = reactive({
  description: '',
  attachments: null as FileList | null,
  currentUserExperience: {
    rating: 0,
    preferNotToSay: true,
  },
  email: '',
  battleTag: '',
});

const stateValid = computed(() => state.description.length > 0);

const modal = useModal();
const toast = useToast();

async function onConfirm(reportKind: 'feedback' | 'bug') {
  try {
    const status = await $fetch('/api/feedback', {
      method: 'POST',
      body: JSON.stringify({
        description: state.description,
        attachments: state.attachments,
        email: state.email,
        battleTag: state.battleTag,
        rating: state.currentUserExperience.rating,
        type: reportKind,
      }),
    });

    if (status && status >= 400) {
      return toast.add({
        title: 'Feedback not submitted',
        description: `An error occurred while submitting your ${reportKind}. Please try again later.`,
        color: 'red',
      });
    }

    toast.add({
      title: 'Feedback submitted',
      description: `Your ${reportKind} has been submitted successfully.`,
      color: 'green',
    });

    resetState();

    modal.close();
  } catch (error) {
    toast.add({
      title: 'Feedback not submitted',
      description: `An error occurred while submitting your ${reportKind}. Please try again later.`,
      color: 'red',
    });
  }
}

function openConfirmationModal(reportKind: 'feedback' | 'bug') {
  modal.open(ConfirmationModal, {
    title: `Submit ${reportKind}`,
    message: `Are you sure you want to submit your ${reportKind}?`,
    buttonText: 'Submit',
    onConfirm: () => onConfirm(reportKind),
    onCancel: () => {
      modal.close();
    },
  });
}

function resetState() {
  state.description = '';
  state.attachments = null;
  state.currentUserExperience.rating = 0;
  state.currentUserExperience.preferNotToSay = true;
  state.email = '';
  state.battleTag = '';
}

watch(
  () => state.currentUserExperience?.preferNotToSay,
  (value) => {
    if (value && state.currentUserExperience?.rating !== 0) {
      state.currentUserExperience.rating = 0;
    }
  }
);
</script>

<template>
  <UContainer class="flex max-w-[1000px] mx-auto mb-10">
    <UTabs class="w-full" :items="page?.tabs" orientation="horizontal">
      <template #item="{ item }">
        <UContainer
          class="flex flex-col gap-4"
          :ui="{
            base: 'mx-0',
            padding: 'px-0 sm:px-0 lg:px-0',
            constrained: 'max-w-full',
          }"
        >
          <UCard>
            <div class="flex flex-col gap-2">
              <p class="text-xl font-semibold leading-6">
                {{ item.label }}
              </p>
              <p class="mt-1 text-sm">
                {{ item.description }}
              </p>
              <p class="mt-1 text-xs italic">
                {{ item.additionalNotes }}
              </p>
            </div>
          </UCard>

          <div v-if="item.feedbackType === 'feedback'" class="space-y-6">
            <UCard>
              <template #header>
                <UFormGroup label="Description" name="feedback" required>
                  <UTextarea
                    v-model="state.description"
                    resize
                    autoresize
                    :rows="10"
                    placeholder="Enter your suggestions here, be as detailed as possible."
                  />
                </UFormGroup>
              </template>
            </UCard>

            <UCard>
              <UFormGroup label="Attachments" name="attachments">
                <UButtonGroup>
                  <UInput
                    type="file"
                    size="md"
                    icon="i-heroicons-folder"
                    accept="image/*,video/*"
                    multiple
                    @change="($event: Event) => {
                      const target = $event.target as HTMLInputElement;
                      state.attachments = target.files;
                    }"
                  />
                  <UButton
                    icon="i-heroicons-trash"
                    color="red"
                    :disabled="!state.attachments?.length"
                  />
                </UButtonGroup>
                <p class="text-xs italic mt-2">
                  Attachments are optional, however any screenshots or videos
                  that you have can greatly help us replicate and identify the
                  bug.
                </p>
              </UFormGroup>
            </UCard>

            <UCard>
              <UFormGroup
                class="flex flex-col gap-4"
                label="How would you rate the current user experience so far?"
                name="currentUserExperience"
              >
                <div class="flex items-center justify-between max-w-sm">
                  <p>Poor</p>
                  <URadio
                    v-model="state.currentUserExperience.rating"
                    class="flex flex-col-reverse gap-2 items-center"
                    :ui="{
                      inner: 'ms-0 flex flex-col gap-2',
                    }"
                    :disabled="state.currentUserExperience.preferNotToSay"
                    v-for="i in 5"
                    :key="i"
                  >
                    <template #label>
                      {{ i }}
                    </template>
                  </URadio>
                  <p>Excellent</p>
                </div>
                <UCheckbox
                  v-model="state.currentUserExperience.preferNotToSay"
                  class="mt-6"
                  label="I prefer not to say"
                />
              </UFormGroup>
            </UCard>

            <UCard>
              <div class="flex flex-col sm:flex-row gap-6">
                <div class="flex flex-col gap-4">
                  <UFormGroup label="Email Address" name="email">
                    <UInput
                      v-model="state.email"
                      class="max-w-[250px] min-w-[150px]"
                      icon="i-heroicons-envelope"
                      type="email"
                      placeholder="Email Address"
                    />
                  </UFormGroup>
                  <UFormGroup label="Battle tag" name="battleTag">
                    <UInput
                      v-model="state.battleTag"
                      class="max-w-[250px] min-w-[200px]"
                      icon="simple-icons:battledotnet"
                      type="text"
                      placeholder="Battle tag"
                    />
                  </UFormGroup>
                </div>
                <p class="text-sm italic self-end">
                  We will only use your contact information to follow up on your
                  feedback if necessary. (Completely optional)
                </p>
              </div>
            </UCard>
          </div>

          <div v-else-if="item.feedbackType === 'bug'" class="space-y-6">
            <UCard>
              <template #header>
                <UFormGroup label="Description" name="bugreport" required>
                  <UTextarea
                    v-model="state.description"
                    resize
                    autoresize
                    :rows="10"
                    placeholder="Enter the details of the bug you encountered here, be as detailed as possible."
                  />
                </UFormGroup>
              </template>
            </UCard>

            <UCard>
              <UFormGroup label="Attachments" name="attachments">
                <UButtonGroup>
                  <UInput
                    type="file"
                    size="md"
                    icon="i-heroicons-folder"
                    accept="image/*,video/*"
                    multiple
                    @change="($event: Event) => {
                      const target = $event.target as HTMLInputElement;
                      state.attachments = target.files;
                    }"
                  />
                  <UButton
                    icon="i-heroicons-trash"
                    color="red"
                    :disabled="!state.attachments?.length"
                  />
                </UButtonGroup>
                <p class="text-xs italic mt-2">
                  Attachments are optional, however any screenshots or videos
                  that you have can greatly help us replicate and identify the
                  bug.
                </p>
              </UFormGroup>
            </UCard>

            <UCard>
              <div class="flex flex-col sm:flex-row gap-6">
                <div class="flex flex-col gap-4">
                  <UFormGroup label="Email Address" name="email">
                    <UInput
                      v-model="state.email"
                      class="max-w-[250px] min-w-[150px]"
                      icon="i-heroicons-envelope"
                      type="email"
                      placeholder="Email Address"
                    />
                  </UFormGroup>
                  <UFormGroup label="Battle tag" name="battleTag">
                    <UInput
                      v-model="state.battleTag"
                      class="max-w-[250px] min-w-[200px]"
                      icon="simple-icons:battledotnet"
                      type="text"
                      placeholder="Battle tag"
                    />
                  </UFormGroup>
                </div>
                <p class="text-sm italic self-end">
                  We will only use your contact information to follow up on your
                  feedback if necessary. (Completely optional)
                </p>
              </div>
            </UCard>
          </div>
        </UContainer>

        <div class="flex w-full justify-end mt-2">
          <UButton
            color="primary"
            :disabled="!stateValid"
            @click="() => openConfirmationModal(item.feedbackType)"
          >
            Submit {{ item.label }}</UButton
          >
        </div>
      </template>
    </UTabs>
  </UContainer>
</template>

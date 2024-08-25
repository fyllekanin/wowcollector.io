<script setup lang="ts">
import ConfirmationModal from '~/components/modals/ConfirmationModal.vue';
import FAQModal from '~/components/modals/builder/FAQModal.vue';
import SubmitViewModal from '~/components/modals/builder/SubmitViewModal.vue';

import type { EmitState } from '~/types';

const mountViewBuilderStore = useMountViewBuilderStore();

const modal = useModal();
const toast = useToast();

const emit = defineEmits(['success']);

function openModal() {
  if (!mountViewBuilderStore.hasChanges) {
    const router = useRouter();
    router.push('/');
    return;
  }

  modal.open(ConfirmationModal, {
    title: 'Exit build mode',
    message:
      'You have made some changes, are you sure you want to exit build mode?',
    additionaInformation: 'All progress will be lost.',
    onConfirm: () => {
      mountViewBuilderStore.resetStore();
      const router = useRouter();
      router.push('/');
      modal.close();
    },
    onCancel: () => {
      modal.close();
    },
  });
}

function openHelpModal() {
  modal.open(FAQModal, {
    kind: 'mounts',
    onClose: () => {
      modal.close();
    },
  });
}

function openNamePromptModal() {
  modal.open(SubmitViewModal, {
    onConfirm: async (state: EmitState) => {
      try {
        const viewId = await $fetch('/api/item-view/mount', {
          method: 'POST',
          body: JSON.stringify({
            name: state.name,
            categories: mountViewBuilderStore.getFinalCategories,
            isUnknownIncluded: state.isUnknownIncluded,
          }),
        });

        if (!viewId) {
          throw new Error('Failed to create view');
        }

        emit('success', viewId);

        modal.close();
      } catch (error) {
        console.error(error);

        toast.add({
          title: 'Failed to create view',
          description: 'An error occurred while creating the view.',
          color: 'red',
        });
      }
    },
    onCancel: () => {
      modal.close();
    },
  });
}
</script>

<template>
  <UDashboardLayout>
    <UDashboardPanel
      :width="400"
      :resizable="{ min: 200, max: 500 }"
      collapsible
    >
      <UDashboardPanelContent>
        <div class="flex flex-col w-full h-full">
          <UDashboardNavbar
            class="!border-transparent"
            :ui="{ left: 'flex-1', wrapper: 'pb-4 px-0' }"
          >
            <template #left>
              <div class="flex grow flex-wrap items-center justify-center">
                <div class="flex cursor-pointer" @click="openModal">
                  <Logo class="mb-1" width="48px" height="48px" />
                  <h2
                    class="flex self-center flex-shrink-0 font-bold text-xl text-gray-900 dark:text-white items-end gap-1.5"
                  >
                    WoW Collector
                  </h2>
                </div>
              </div>
            </template>
          </UDashboardNavbar>

          <UDivider />

          <slot name="sidebar-content" />
        </div>
      </UDashboardPanelContent>
    </UDashboardPanel>

    <div class="flex flex-col w-full h-full">
      <UDashboardNavbar class="" :ui="{ wrapper: 'pb-' }">
        <template #title>
          <div class="flex gap-2 items-center">
            <UIcon class="scale-125" name="mdi:tools" />
            <h2 class="text-lg font-semibold">Mount View Builder</h2>
          </div>
        </template>

        <template #center>
          <UButton variant="ghost" color="gray" @click="openHelpModal"
            >What is this?</UButton
          >
        </template>

        <template #right>
          <UButton
            :disabled="
              !mountViewBuilderStore.hasChanges ||
              !mountViewBuilderStore.isValid
            "
            icon="mdi:plus"
            @click="openNamePromptModal"
            >Create view</UButton
          >
        </template>
      </UDashboardNavbar>
      <slot name="main-content" />
    </div>
  </UDashboardLayout>
</template>

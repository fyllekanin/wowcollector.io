<script setup lang="ts">
import ConfirmationModal from '~/components/modals/ConfirmationModal.vue';
import FAQModal from '~/components/modals/builder/FAQModal.vue';

const viewBuilderStore = useViewBuilderStore();

const modal = useModal();

function openModal() {
  if (!viewBuilderStore.hasChanges) {
    const router = useRouter();
    router.push('/');
    return;
  }

  modal.open(ConfirmationModal, {
    title: 'Exit build mode',
    message:
      'You have made some changes, are you sure you want to exit build mode? All progress will be lost.',
    onConfirm: () => {
      viewBuilderStore.resetStore();
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
            :ui="{ left: 'flex-1', wrapper: 'pb-4' }"
          >
            <template #left>
              <div class="flex grow flex-wrap items-center justify-between">
                <div class="flex cursor-pointer" @click="openModal">
                  <Logo class="mb-1" width="48px" height="48px" />
                  <h2
                    class="flex self-center flex-shrink-0 font-bold text-xl text-gray-900 dark:text-white items-end gap-1.5"
                  >
                    WoW Collector
                  </h2>
                </div>
                <div class="flex gap-2 items-center">
                  <h2 class="text-lg font-semibold">View Builder</h2>
                  <UIcon name="mdi:tools" />
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
      <UDashboardNavbar class="" :ui="{ left: 'flex-1', wrapper: 'pb-' }">
        <template #left>
          <UButton variant="ghost" color="gray" @click="openHelpModal"
            >What is this?</UButton
          >
        </template>
        <template #right>
          <UButton :disabled="!viewBuilderStore.hasChanges" icon="mdi:plus"
            >Create view</UButton
          >
        </template>
      </UDashboardNavbar>
      <slot name="main-content" />
    </div>
  </UDashboardLayout>
</template>

<script setup lang="ts">
import ConfirmationModal from '~/components/modals/ConfirmationModal.vue';

const modal = useModal();

function openModal() {
  modal.open(ConfirmationModal, {
    title: 'Exit build mode',
    message: 'Are you sure you want to exit build mode?',
    onConfirm: () => {
      const router = useRouter();
      router.push('/');
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
            :ui="{ left: 'flex-1 cursor-pointer' }"
          >
            <template #left>
              <div class="flex grow flex-wrap items-center justify-between">
                <div class="flex" @click="openModal">
                  <Logo width="48px" height="48px" />
                  <h2
                    class="flex self-center mt-1 flex-shrink-0 font-bold text-xl text-gray-900 dark:text-white items-end gap-1.5"
                  >
                    WoW Collector
                  </h2>
                </div>
                <!-- <UButton variant="ghost" color="gray">Need help?</UButton> -->
                <!-- <UTooltip
                  text="Grab an icon or category and drop it inside the content panel"
                >
                  <UIcon
                    name="material-symbols-light:help-outline"
                    class="scale-125"
                  />
                </UTooltip> -->
              </div>
            </template>
          </UDashboardNavbar>

          <UDivider />

          <slot name="sidebar-content" />
        </div>
      </UDashboardPanelContent>
    </UDashboardPanel>

    <slot name="main-content" />
  </UDashboardLayout>
</template>

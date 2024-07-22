import { ref, onMounted, onUnmounted } from 'vue';

export function useScreenSize() {
  const isMobileRef = ref(false);
  const isTabletRef = ref(false);
  const isDesktopRef = ref(false);
  const isCustomXLRef = ref(false);
  const sizeRef = ref();

  const isMobile = computed(() => isMobileRef.value);
  const isTablet = computed(() => isTabletRef.value);
  const isDesktop = computed(() => isDesktopRef.value);
  const isCustomXL = computed(() => isCustomXLRef.value);
  const size = computed(() => sizeRef.value);

  const updateScreenSize = () => {
    isMobileRef.value = window.innerWidth <= 480;
    isTabletRef.value = window.innerWidth <= 1024;
    isDesktopRef.value = window.innerWidth <= 1440;
    isCustomXLRef.value = window.innerWidth > 1440;
    sizeRef.value = window.innerWidth;
  };

  onMounted(() => {
    updateScreenSize();
    window.addEventListener('resize', updateScreenSize);
  });

  onUnmounted(() => {
    window.removeEventListener('resize', updateScreenSize);
  });

  return { isMobile, isTablet, isDesktop, isCustomXL, size };
}

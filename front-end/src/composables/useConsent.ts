import type { ConsentSettings } from '~/types';

const defaultConsent: ConsentSettings = {
  essential: true,
  analytics: false,
};

export function useConsent() {
  const consentCookie = useCookie<ConsentSettings>('_wc_consent_optin', {
    maxAge: 365 * 24 * 60 * 60,
  }); // Expires in 1 year
  const consent = ref<ConsentSettings>(consentCookie.value || defaultConsent);

  const isConsentGiven = computed(() => consentCookie.value);

  const setConsent = (newConsent: Partial<ConsentSettings>) => {
    consent.value = { ...consent.value, ...newConsent };
    consentCookie.value = consent.value;
  };

  const resetConsent = () => {
    consent.value = defaultConsent;
    consentCookie.value = defaultConsent;
  };

  return {
    consent,
    isConsentGiven,
    setConsent,
    resetConsent,
  };
}

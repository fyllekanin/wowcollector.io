import type { Auth } from '~/types';

export function useAuth() {
  const atCookie = useCookie('wc_at');
  const rtCookie = useCookie('wc_rt');

  const login = async (code: string, redirectUri: string, scope: string) => {
    const response = await $fetch(
      `/api/auth/battle-net?code=${code}&redirect_uri=${redirectUri}&scope=${scope}`
    );

    if (!response?.tokens?.accessToken || !response?.tokens?.refreshToken) {
      throw new Error('Failed to login');
    }

    atCookie.value = response.tokens.accessToken;
    rtCookie.value = response.tokens.refreshToken;

    return response;
  };

  const refresh = async (refreshToken: string) => {
    const response = await $fetch<Auth>(
      `/api/auth/refresh?refreshToken=${refreshToken}`
    );

    if (!response?.tokens?.accessToken || !response?.tokens?.refreshToken) {
      throw new Error('Failed to refresh');
    }

    atCookie.value = response.tokens.accessToken;
    rtCookie.value = response.tokens.refreshToken;

    return response;
  };

  return {
    atCookie,
    rtCookie,
    login,
    refresh,
  };
}

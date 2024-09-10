import { get } from './HttpService';

import type { Auth } from '~/types';

export default class AuthService {
  static async login(code: string, redirectUri: string, scope: string) {
    try {
      const response = await get<Auth>({
        url: `/api/v1/auth/battle-net?code=${code}&redirectUri=${redirectUri}&scope=${scope}`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

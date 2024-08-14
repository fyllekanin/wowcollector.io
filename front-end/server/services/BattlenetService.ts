import { get } from './HttpService';
import type { MountInformation, RealmsRegions } from '~/types';

export default class BattlenetService {
  static async getRealmsAndRegions() {
    try {
      const response = await get<RealmsRegions>({
        url: '/api/v1/battle-net/realms-regions',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getMounts() {
    try {
      const response = await get<MountInformation[]>({
        url: '/api/v1/battle-net/mounts',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

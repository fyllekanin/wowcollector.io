import { get } from './HttpService';
import type { RealmsRegions, AchievementRootCategory } from '~/types';

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

  static async getAchievementRootCategories() {
    try {
      const response = await get<AchievementRootCategory[]>({
        url: '/api/v1/battle-net/achievement-root-categories',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

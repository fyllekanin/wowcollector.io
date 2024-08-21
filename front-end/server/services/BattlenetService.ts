import { get } from './HttpService';
import type {
  AchievementInformation,
  MountInformation,
  PetInformation,
  RealmsRegions,
  ToyInformation,
} from '~/types';

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

  static async getAchievements() {
    try {
      const response = await get<AchievementInformation[]>({
        url: '/api/v1/battle-net/achievements',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getToys() {
    try {
      const response = await get<ToyInformation[]>({
        url: '/api/v1/battle-net/toys',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getPets() {
    try {
      const response = await get<PetInformation[]>({
        url: '/api/v1/battle-net/pets',
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

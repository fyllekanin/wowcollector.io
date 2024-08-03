import { get } from './HttpService';

import type { AchievementCategory, MountCategory } from '~/types';

export default class CharacterService {
  static async getCharacterMounts(name: string, realm: string, region: string) {
    try {
      const response = await get<MountCategory[]>({
        url: `/api/v1/character/${region}/${realm}/${name}/mounts`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getCharacterAchievements(
    name: string,
    realm: string,
    region: string,
    rootCategoryId?: number | string
  ) {
    try {
      const response = await get<AchievementCategory[]>({
        url: `/api/v1/character/${region}/${realm}/${name}/achievements`,
        config: {
          params: {
            rootCategoryId,
          },
        },
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

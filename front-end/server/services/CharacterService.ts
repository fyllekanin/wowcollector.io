import { get } from './HttpService';

import type { MountCategory } from '~/types/collections/mounts';

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
}

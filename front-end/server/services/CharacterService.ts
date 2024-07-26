import { get } from './HttpService';

export default class CharacterService {
  static async getCharacterMounts(name: string, realm: string, region: string) {
    try {
      const response = await get<unknown>({
        url: `/api/v1/character/${region}/${realm}/${name}/mounts`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

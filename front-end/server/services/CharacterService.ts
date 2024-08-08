import { get } from './HttpService';

import type {
  AchievementCategory,
  MountCategory,
  PetCategory,
  ToyCategory,
  Character,
} from '~/types';

export default class CharacterService {
  static async getCharacterInformation(
    name: string,
    realm: string,
    region: string
  ) {
    try {
      const response = await get<Character>({
        url: `/api/v1/character/${region}/${realm}/${name}`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

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
    region: string
  ) {
    try {
      const response = await get<AchievementCategory[]>({
        url: `/api/v1/character/${region}/${realm}/${name}/achievements`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getCharacterToys(name: string, realm: string, region: string) {
    try {
      const response = await get<ToyCategory[]>({
        url: `/api/v1/character/${region}/${realm}/${name}/toys`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async getCharacterPets(name: string, realm: string, region: string) {
    try {
      const response = await get<PetCategory[]>({
        url: `/api/v1/character/${region}/${realm}/${name}/pets`,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

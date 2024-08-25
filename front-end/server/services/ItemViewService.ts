import { post } from './HttpService';

import type { MountViewPayload, PetViewPayload, ToyViewPayload } from '~/types';

export default class ItemViewService {
  static async createMountView(
    postData: MountViewPayload
  ): Promise<string | null> {
    try {
      const response = await post<MountViewPayload, string>({
        url: '/api/v1/item-view/mount',
        body: postData,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async createPetView(postData: PetViewPayload): Promise<string | null> {
    try {
      const response = await post<PetViewPayload, string>({
        url: '/api/v1/item-view/pet',
        body: postData,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  static async createToyView(postData: ToyViewPayload): Promise<string | null> {
    try {
      const response = await post<ToyViewPayload, string>({
        url: '/api/v1/item-view/toy',
        body: postData,
      });

      return response.data;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

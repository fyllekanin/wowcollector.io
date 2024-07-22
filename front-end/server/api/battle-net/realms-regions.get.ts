import BattlenetService from '~~/server/services/BattlenetService';

import type { RealmsRegions } from '~/types';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const data = await BattlenetService.getRealmsAndRegions();

    if (data) return data;
  },
});

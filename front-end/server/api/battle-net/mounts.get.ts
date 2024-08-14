import BattlenetService from '~~/server/services/BattlenetService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const mounts = await BattlenetService.getMounts();
    if (mounts) return mounts;
  },
});

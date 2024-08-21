import BattlenetService from '~~/server/services/BattlenetService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const toys = await BattlenetService.getToys();
    if (toys) return toys;
  },
});

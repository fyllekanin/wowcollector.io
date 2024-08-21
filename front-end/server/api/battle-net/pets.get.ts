import BattlenetService from '~~/server/services/BattlenetService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const pets = await BattlenetService.getPets();
    if (pets) return pets;
  },
});

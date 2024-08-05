import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const toys = await CharacterService.getCharacterToys(name, realm, region);
    if (toys) return toys;
  },
});

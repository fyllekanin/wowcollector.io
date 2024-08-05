import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const pets = await CharacterService.getCharacterPets(name, realm, region);
    if (pets) return pets;
  },
});

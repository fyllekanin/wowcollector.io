import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const { viewId } = getQuery(event);

    const pets = await CharacterService.getCharacterPets(
      name.toLowerCase(),
      realm.toLowerCase(),
      region.toLowerCase(),
      viewId as string
    );
    if (pets) return pets;
  },
});

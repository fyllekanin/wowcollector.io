import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const { viewId } = getQuery(event);

    const toys = await CharacterService.getCharacterToys(
      name.toLowerCase(),
      realm.toLowerCase(),
      region.toLowerCase(),
      viewId as string
    );
    if (toys) return toys;
  },
});

import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const character = await CharacterService.getCharacterInformation(
      name.toLowerCase(),
      realm.toLowerCase(),
      region.toLowerCase()
    );
    if (character) return character;
  },
});

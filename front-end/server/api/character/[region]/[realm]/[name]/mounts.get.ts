import CharacterService from '~~/server/services/CharacterService';

export default defineEventHandler({
  onRequest: (_) => {},
  handler: async (event) => {
    const { region, realm, name } = getRouterParams(event);
    const mounts = await CharacterService.getCharacterMounts(
      name.toLowerCase(),
      realm.toLowerCase(),
      region.toLowerCase()
    );
    if (mounts) return mounts;
  },
});

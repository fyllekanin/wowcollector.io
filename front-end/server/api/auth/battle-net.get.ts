import AuthService from '~~/server/services/AuthService';

export default defineEventHandler({
  handler: async (event) => {
    const query = getQuery(event);

    const response = await AuthService.login(
      query.code as string,
      query.redirect_uri as string,
      query.scope as string
    );
    return response;
  },
});

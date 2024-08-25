import ItemViewService from '../../services/ItemViewService';

export default defineEventHandler({
  handler: async (event) => {
    const postData = await readBody(event);

    const response = ItemViewService.createToyView(postData);
    return response;
  },
});

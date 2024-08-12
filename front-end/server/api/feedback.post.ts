import FeedbackService from '../services/FeedbackService';

export default defineEventHandler({
  handler: async (event) => {
    const feedback = await readBody(event);

    const response = FeedbackService.submitFeedback(feedback);
    return response;
  },
});

import { postFormData } from './HttpService';

import type { Feedback } from '~/types';

export default class FeedbackService {
  static async submitFeedback(feedback: Feedback) {
    try {
      const response = await postFormData({
        url: '/api/v1/feedback',
        body: feedback,
      });

      return response.status;
    } catch (error) {
      console.error(error);
      return null;
    }
  }
}

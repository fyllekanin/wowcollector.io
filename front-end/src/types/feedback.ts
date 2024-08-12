export interface Feedback {
  description: string;
  attachments?: FormData[];
  email?: string;
  battleTag?: string;
  rating?: number;
  type: 'bug' | 'feedback';
}

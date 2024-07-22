import type { AxiosRequestConfig } from 'axios';

export interface HttpParamater<T> {
  url: string;
  body?: T;
  config?: AxiosRequestConfig;
}

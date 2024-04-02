// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import axios from 'axios';

export const riftClient = axios.create({
  baseURL: '/api/rift',
});
riftClient.interceptors.request.use(config => {
  config.timeout = 30000;
  return config;
});

export interface Response {
    error: boolean;
    message: string;
}

export interface Setting {
    name: string;
    value: boolean;
}
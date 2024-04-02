// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Setting, riftClient } from './client.gen';
import { Page } from './models.gen';

export interface PageIndexRequest {
  page: number;
  limit: number;
}
export interface PageIndexResponse extends Response {
  result: Page[];
  total: number;
}
export interface PageCreateRequest {
  subject: Page;
}
export interface PageCreateResponse extends Response {
  result: Page;
}
export interface PageShowRequest {
  id: string;
}
export interface PageShowResponse extends Response {
  result: Page;
}
export interface PageUpdateRequest {
  id: string;
  subject: Page;
}
export interface PageUpdateResponse extends Response {
  result: Page;
}
export interface PageSettingsRequest {
  id: string;
  setting: Setting;
}
export interface PageSettingsResponse extends Response {
  result: Page;
}
export interface PageDeleteRequest {
  id: string;
}
export interface PageDeleteResponse extends Response {
  result: Page;
}

export const PageIndex = async (params: PageIndexRequest) => {
  const response = await riftClient.get(`/page/?page=${params.page}&limit=${params.limit}`);
  return response.data as PageIndexResponse;
};
export const PageCreate = async (params: PageCreateRequest) => {
  const response = await riftClient.post(`/page/?`, params.subject);
  return response.data as PageCreateResponse;
};
export const PageShow = async (params: PageShowRequest) => {
  const response = await riftClient.get(`/page/${params.id}?`);
  return response.data as PageShowResponse;
};
export const PageUpdate = async (params: PageUpdateRequest) => {
  const response = await riftClient.put(`/page/${params.id}?`, params.subject);
  return response.data as PageUpdateResponse;
};
export const PageSettings = async (params: PageSettingsRequest) => {
  const response = await riftClient.patch(`/page/${params.id}?`, params.setting);
  return response.data as PageSettingsResponse;
};
export const PageDelete = async (params: PageDeleteRequest) => {
  const response = await riftClient.delete(`/page/${params.id}?`);
  return response.data as PageDeleteResponse;
};

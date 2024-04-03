// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Response, Setting, Visit, riftClient } from '.';

export interface VisitIndexRequest {
  page: number;
  limit: number;
}
export interface VisitIndexResponse extends Response {
  result: Visit[];
  total: number;
}
export const VisitIndex = async (params: VisitIndexRequest) => {
  const response = await riftClient.get(`/visit/?page=${params.page}&limit=${params.limit}`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitIndexResponse;
};

export interface VisitCreateRequest {
  subject: Visit;
}
export interface VisitCreateResponse extends Response {
  result: Visit;
}
export const VisitCreate = async (params: VisitCreateRequest) => {
  const response = await riftClient.post(`/visit/?`, params.subject);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitCreateResponse;
};

export interface VisitShowRequest {
  id: string;
}
export interface VisitShowResponse extends Response {
  result: Visit;
}
export const VisitShow = async (params: VisitShowRequest) => {
  const response = await riftClient.get(`/visit/${params.id}?`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitShowResponse;
};

export interface VisitUpdateRequest {
  id: string;
  subject: Visit;
}
export interface VisitUpdateResponse extends Response {
  result: Visit;
}
export const VisitUpdate = async (params: VisitUpdateRequest) => {
  const response = await riftClient.put(`/visit/${params.id}?`, params.subject);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitUpdateResponse;
};

export interface VisitSettingsRequest {
  id: string;
  setting: Setting;
}
export interface VisitSettingsResponse extends Response {
  result: Visit;
}
export const VisitSettings = async (params: VisitSettingsRequest) => {
  const response = await riftClient.patch(`/visit/${params.id}?`, params.setting);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitSettingsResponse;
};

export interface VisitDeleteRequest {
  id: string;
}
export interface VisitDeleteResponse extends Response {
  result: Visit;
}
export const VisitDelete = async (params: VisitDeleteRequest) => {
  const response = await riftClient.delete(`/visit/${params.id}?`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VisitDeleteResponse;
};

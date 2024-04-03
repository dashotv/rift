// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Response, Setting, Video, riftClient } from '.';

export interface VideoIndexRequest {
  page: number;
  limit: number;
}
export interface VideoIndexResponse extends Response {
  result: Video[];
  total: number;
}
export const VideoIndex = async (params: VideoIndexRequest) => {
  const response = await riftClient.get(`/video/?page=${params.page}&limit=${params.limit}`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoIndexResponse;
};

export interface VideoCreateRequest {
  subject: Video;
}
export interface VideoCreateResponse extends Response {
  result: Video;
}
export const VideoCreate = async (params: VideoCreateRequest) => {
  const response = await riftClient.post(`/video/?`, params.subject);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoCreateResponse;
};

export interface VideoShowRequest {
  id: string;
}
export interface VideoShowResponse extends Response {
  result: Video;
}
export const VideoShow = async (params: VideoShowRequest) => {
  const response = await riftClient.get(`/video/${params.id}?`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoShowResponse;
};

export interface VideoUpdateRequest {
  id: string;
  subject: Video;
}
export interface VideoUpdateResponse extends Response {
  result: Video;
}
export const VideoUpdate = async (params: VideoUpdateRequest) => {
  const response = await riftClient.put(`/video/${params.id}?`, params.subject);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoUpdateResponse;
};

export interface VideoSettingsRequest {
  id: string;
  setting: Setting;
}
export interface VideoSettingsResponse extends Response {
  result: Video;
}
export const VideoSettings = async (params: VideoSettingsRequest) => {
  const response = await riftClient.patch(`/video/${params.id}?`, params.setting);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoSettingsResponse;
};

export interface VideoDeleteRequest {
  id: string;
}
export interface VideoDeleteResponse extends Response {
  result: Video;
}
export const VideoDelete = async (params: VideoDeleteRequest) => {
  const response = await riftClient.delete(`/video/${params.id}?`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.Message) {
      throw new Error(response.data.Message);
    }
    throw new Error('unknown error');
  }

  return response.data as VideoDeleteResponse;
};

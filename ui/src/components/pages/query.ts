import { Page, PageCreate, PageDelete, PageIndex, PageRefresh, PageUpdate, PageVideos, PageVisits } from 'client';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

export const usePagesQuery = (page: number, limit = 50) =>
  useQuery({
    queryKey: ['pages', page],
    queryFn: () => PageIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageVisitsQuery = (id: string, page: number, limit = 10) =>
  useQuery({
    queryKey: ['pages', id, 'videos', page],
    queryFn: () => PageVisits({ id, page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageVideosQuery = (id: string, page: number, limit = 10) =>
  useQuery({
    queryKey: ['pages', id, 'videos', page],
    queryFn: () => PageVideos({ id, page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (page: Page) => {
      if (!page.id) {
        throw new Error('page id required');
      }
      return PageUpdate({ id: page.id, subject: page });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['pages'] });
    },
  });
};

export const usePageCreateMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (page: Page) => PageCreate({ subject: page }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['pages'] });
    },
  });
};

export const usePageDeleteMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: async (page: Page) => {
      if (!page.id) {
        throw new Error('page id required');
      }
      return PageDelete({ id: page.id });
    },
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ['pages'] });
    },
  });
};
export const usePageRefreshMutation = () => {
  return useMutation({
    mutationFn: async (id: string) => PageRefresh({ id }),
  });
};

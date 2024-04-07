import { PageIndex, PageRefresh, PageVideos, PageVisits } from 'client';

import { useMutation, useQuery } from '@tanstack/react-query';

export const usePagesQuery = (page: number, limit = 50) =>
  useQuery({
    queryKey: ['pages', page],
    queryFn: () => PageIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageVisitsQuery = (id: string, page: number, limit = 50) =>
  useQuery({
    queryKey: ['pages', id, 'videos', page],
    queryFn: () => PageVisits({ id, page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageVideosQuery = (id: string, page: number, limit = 50) =>
  useQuery({
    queryKey: ['pages', id, 'videos', page],
    queryFn: () => PageVideos({ id, page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const usePageRefreshMutation = () => {
  return useMutation({
    onMutate: async (id: string) => PageRefresh({ id }),
  });
};

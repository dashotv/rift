import { VideoIndex } from 'client';

import { useQuery } from '@tanstack/react-query';

export const useVideosQuery = (page: number, limit = 25) =>
  useQuery({
    queryKey: ['jobs', page, status],
    queryFn: () => VideoIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

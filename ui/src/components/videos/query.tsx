import { VideoDelete, VideoIndex } from 'client';

import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';

export const useVideosQuery = (page: number, limit = 25) =>
  useQuery({
    queryKey: ['videos', page, limit],
    queryFn: () => VideoIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useVideosDeleteMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    onMutate: async (id: string) => VideoDelete({ id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['videos'] });
    },
  });
};

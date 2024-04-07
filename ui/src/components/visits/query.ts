import { VisitDelete, VisitIndex } from 'client';

import { useMutation, useQuery } from '@tanstack/react-query';
import { useQueryClient } from '@tanstack/react-query';

export const useVisitsQuery = (page: number, limit = 25) =>
  useQuery({
    queryKey: ['visits', page, limit],
    queryFn: () => VisitIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

export const useVisitsDeleteMutation = () => {
  const queryClient = useQueryClient();
  return useMutation({
    onMutate: async (id: string) => VisitDelete({ id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['visits'] });
    },
  });
};

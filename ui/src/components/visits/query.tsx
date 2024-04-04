import { VisitIndex } from 'client';

import { useQuery } from '@tanstack/react-query';

export const useVisitsQuery = (page: number, limit = 25) =>
  useQuery({
    queryKey: ['visits', page, limit],
    queryFn: () => VisitIndex({ page, limit }),
    placeholderData: previousData => previousData,
    retry: false,
  });

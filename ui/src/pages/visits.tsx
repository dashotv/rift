import { useEffect, useState } from 'react';
import { Helmet } from 'react-helmet-async';

import { Grid, Pagination, Stack } from '@mui/material';

import { VisitsList, useVisitsQuery } from 'components/visits';

const pagesize = 50;
const Videos = () => {
  const [page, setPage] = useState(1);
  const [count, setCount] = useState(0);
  const { data } = useVisitsQuery(page);

  useEffect(() => {
    if (!data?.total) return;
    setCount(Math.ceil((data.total || 0) / pagesize)); // Math.ceil((data?.count || 0) / pagesize)
  }, [data?.total]);

  const handleChange = (_event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };

  return (
    <>
      <Helmet>
        <title>Rift - Videos</title>
        <meta name="description" content="runic" />
      </Helmet>

      <Grid container spacing={0} sx={{ mb: 2 }}>
        <Grid item xs={12} md={6}>
          <Stack direction="row" spacing={0} alignItems="center"></Stack>
        </Grid>
        <Grid item xs={12} md={6} justifyContent="end">
          {data && (
            <Pagination
              sx={{ display: 'flex', justifyContent: 'end', height: '48px' }}
              page={page}
              count={count}
              onChange={handleChange}
            />
          )}
        </Grid>
      </Grid>

      {data?.result ? <VisitsList data={data.result} /> : null}
    </>
  );
};

export default Videos;

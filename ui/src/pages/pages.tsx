import { useState } from 'react';
import { Helmet } from 'react-helmet-async';

import { Grid, Stack } from '@mui/material';

import { PageList, usePagesQuery } from 'components/pages';

const Videos = () => {
  const [page] = useState(1);
  const { data } = usePagesQuery(page);

  return (
    <>
      <Helmet>
        <title>Minion - Jobs</title>
        <meta name="description" content="runic" />
      </Helmet>

      <Grid container spacing={0} sx={{ mb: 2 }}>
        <Grid item xs={12} md={6}>
          <Stack direction="row" spacing={0} alignItems="center">
            <div>buttons</div>
          </Stack>
        </Grid>
        <Grid item xs={12} md={6} justifyContent="end">
          <div>pagination</div>
        </Grid>
      </Grid>

      {data?.result ? <PageList data={data.result} /> : null}
    </>
  );
};

export default Videos;

import { useEffect, useState } from 'react';
import { Helmet } from 'react-helmet-async';

import { Page } from 'client';

import QueueIcon from '@mui/icons-material/Queue';
import { Grid, IconButton, Pagination, Stack } from '@mui/material';

import { PageList, usePagesQuery } from 'components/pages';
import { PagesEditDialog } from 'components/pages/edit';

const pagesize = 50;
const Videos = () => {
  const [page, setPage] = useState(1);
  const [count, setCount] = useState(0);
  const [editing, setEditing] = useState<Page | null>(null);
  const { data } = usePagesQuery(page);

  const handleChange = (_event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };

  useEffect(() => {
    if (!data?.total) return;
    setCount(Math.ceil((data.total || 0) / pagesize)); // Math.ceil((data?.count || 0) / pagesize)
  }, [data?.total]);

  return (
    <>
      <Helmet>
        <title>Rift - Pages</title>
        <meta name="description" content="runic" />
      </Helmet>

      <Grid container spacing={0} sx={{ mb: 2 }}>
        <Grid item xs={12} md={6}>
          <Stack direction="row" spacing={0} alignItems="center">
            <IconButton aria-label="refresh" color="primary" onClick={() => setEditing({ name: '', url: '' })}>
              <QueueIcon />
            </IconButton>
          </Stack>
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

      {data?.result ? <PageList data={data.result} setEditing={setEditing} /> : null}
      {editing && <PagesEditDialog subject={editing} setEditing={setEditing} />}
    </>
  );
};

export default Videos;

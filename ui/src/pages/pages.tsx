import { useState } from 'react';
import { Helmet } from 'react-helmet-async';

import { Page } from 'client';

import QueueIcon from '@mui/icons-material/Queue';
import { Grid, IconButton, Stack } from '@mui/material';

import { PageList, usePagesQuery } from 'components/pages';
import { PagesEditDialog } from 'components/pages/edit';

const Videos = () => {
  const [page] = useState(1);
  const [editing, setEditing] = useState<Page | null>(null);
  const { data } = usePagesQuery(page);

  return (
    <>
      <Helmet>
        <title>Rift - Pages</title>
        <meta name="description" content="runic" />
      </Helmet>

      <Grid container spacing={0} sx={{ mb: 2 }}>
        <Grid item xs={12} md={6}>
          <Stack direction="row" spacing={0} alignItems="center">
            <IconButton
              aria-label="refresh"
              color="primary"
              onClick={() => setEditing({ name: '', url: '', scraper: 'myanime', downloader: 'metube' })}
            >
              <QueueIcon />
            </IconButton>
          </Stack>
        </Grid>
        <Grid item xs={12} md={6} justifyContent="end"></Grid>
      </Grid>

      {data?.result ? <PageList data={data.result} setEditing={setEditing} /> : null}
      {editing && <PagesEditDialog subject={editing} setEditing={setEditing} />}
    </>
  );
};

export default Videos;

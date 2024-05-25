import { useState } from 'react';

import { Page } from 'client';

import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
import ReplayCircleFilledIcon from '@mui/icons-material/ReplayCircleFilled';
import StarIcon from '@mui/icons-material/Star';
import Box from '@mui/material/Box';
import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { ButtonMap, ButtonMapButton, Chrono, Row } from '@dashotv/components';

import { PagesDialog, usePageDeleteMutation, usePageMutation, usePageRefreshMutation } from '.';

export const PageList = ({ data, setEditing }: { data: Page[]; setEditing: (p: Page) => void }) => {
  const [viewing, setViewing] = useState<Page | null>(null);
  // const [editing, setEditing] = useState<Page | null>(null);
  const pageUpdate = usePageMutation();
  const pageRefresh = usePageRefreshMutation();
  const pageDelete = usePageDeleteMutation();
  const view = (row: Page) => {
    console.log(row);
    setViewing(row);
  };
  const handleClose = () => {
    setViewing(null);
  };
  const refresh = (row: Page) => {
    if (!row.id) return;
    // console.log('refresh:', row);
    pageRefresh.mutate(row.id);
  };
  const edit = (row: Page) => {
    // console.log(row);
    setEditing(row);
  };
  const del = (row: Page) => {
    pageDelete.mutate(row);
  };
  const actions = (row: Page) => {
    const buttons: ButtonMapButton[] = [
      {
        title: 'Enabled',
        Icon: StarIcon,
        color: row.enabled ? 'secondary' : 'disabled',
        click: ev => {
          ev.preventDefault();
          pageUpdate.mutate({ ...row, enabled: !row.enabled });
        },
      },
      {
        title: 'Edit',
        Icon: EditIcon,
        color: 'primary',
        click: ev => {
          ev.preventDefault();
          edit(row);
        },
      },
      {
        title: 'Refresh',
        Icon: ReplayCircleFilledIcon,
        color: 'primary',
        click: ev => {
          ev.preventDefault();
          refresh(row);
        },
      },
      {
        title: 'Delete',
        Icon: DeleteForeverIcon,
        color: 'error',
        click: ev => {
          ev.preventDefault();
          del(row);
        },
      },
    ];
    return <ButtonMap {...{ buttons }} size="small" />;
  };

  return (
    <Paper elevation={0} sx={{ width: '100%' }}>
      {data.map((row: Page) => (
        <Row key={row.id}>
          <Stack direction={{ xs: 'column', md: 'row' }} spacing={{ xs: 0, md: 1 }} alignItems="center">
            <Stack
              direction="row"
              spacing={1}
              width="100%"
              maxWidth={{ xs: '100%', md: '800px' }}
              pr="3px"
              alignItems="center"
            >
              <Typography
                component="div"
                fontWeight="bolder"
                noWrap
                color="primary"
                sx={{ pr: 1, '& a': { color: 'primary.main' } }}
              >
                <Link href="#" onClick={() => view(row)}>
                  {row.name}
                </Link>
              </Typography>
              <Typography variant="subtitle2" color="primary.dark">
                {row.scraper || 'myanime'}
              </Typography>
              <Typography variant="subtitle2" color="secondary.dark">
                {row.downloader || 'metube'}
              </Typography>
              {/* <Group group={row.scraper} author="myanime" variant="default" /> */}
            </Stack>
            <Stack
              direction="row"
              spacing={1}
              alignItems="center"
              sx={{ width: '100%', justifyContent: { xs: 'start', md: 'end' } }}
            >
              <Stack width={{ xs: '100%', md: 'auto' }} direction="row" spacing={1} alignItems="center">
                <Typography noWrap variant="subtitle2" color="gray" pl="3px" width="100%">
                  {row.processed_at && <Chrono fromNow>{row.processed_at}</Chrono>}
                </Typography>
                <Box>{actions && actions(row)}</Box>
              </Stack>
            </Stack>
          </Stack>
        </Row>
      ))}
      {viewing && <PagesDialog {...{ open, close: handleClose }} page={viewing} />}
    </Paper>
  );
};

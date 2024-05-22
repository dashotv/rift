import { useState } from 'react';

import { Page } from 'client';

import CloseIcon from '@mui/icons-material/Close';
import { Dialog, DialogContent, DialogTitle, IconButton, Paper, Stack, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

import { Chrono } from '@dashotv/components';

import { VideosList } from 'components/videos';

import { usePageVideosQuery } from '.';

export const PagesDialog = ({
  close,
  page: { id, name, processed_at, url, scraper, downloader },
}: {
  close: () => void;
  page: Page;
}) => {
  const [open, setOpen] = useState(true);
  const handleClose = () => {
    setOpen(false);
    close();
  };

  if (!id) {
    throw new Error('Page ID is required');
    return null;
  }

  const { data } = usePageVideosQuery(id, 1);

  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="md" fullWidth fullScreen={fullScreen}>
      <DialogTitle>
        <Stack direction="row" spacing={3} alignItems="center" justifyContent="start" width="100%">
          <Typography fontWeight="bolder" color="primary">
            {name}
          </Typography>
        </Stack>
        <IconButton
          aria-label="close"
          onClick={handleClose}
          sx={{
            position: 'absolute',
            right: 8,
            top: 8,
            color: theme => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>
      </DialogTitle>
      <DialogContent>
        <Stack width="100%" direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
          <Typography minWidth="125px" variant="subtitle1" color="textSecondary">
            ID
          </Typography>
          <Typography variant="button">{id}</Typography>
        </Stack>
        <Stack width="100%" direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
          <Typography minWidth="125px" variant="subtitle1" color="textSecondary">
            Name
          </Typography>
          <Typography variant="subtitle1">{name}</Typography>
          <Typography variant="subtitle2" color="primary">
            {scraper || 'myanime'}
          </Typography>
          <Typography variant="subtitle2" color="secondary">
            {downloader || 'metube'}
          </Typography>
        </Stack>
        <Stack width="100%" direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
          <Typography minWidth="125px" variant="subtitle1" color="textSecondary">
            URL
          </Typography>
          <Typography variant="subtitle1">{url}</Typography>
        </Stack>
        <Stack width="100%" direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
          <Typography minWidth="125px" variant="subtitle1" color="textSecondary">
            Processed At
          </Typography>
          <Typography noWrap variant="subtitle2" color="gray" pl="3px" width="100%">
            {processed_at && <Chrono fromNow>{processed_at}</Chrono>}
          </Typography>
        </Stack>
        <Paper elevation={0} sx={{ p: 1, width: '100%' }}>
          {/* <Row>{data?.result?.map((row: Video) => <div>{row.title}</div>)}</Row> */}
          <VideosList data={data?.result || []} />
        </Paper>
      </DialogContent>
    </Dialog>
  );
};

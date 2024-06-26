import { useState } from 'react';

import { Page } from 'client';

import CloseIcon from '@mui/icons-material/Close';
import { Dialog, DialogContent, DialogTitle, IconButton, Stack, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

import { Chrono, MediumTabMap, MediumTabs } from '@dashotv/components';

import { VideosList } from 'components/videos';
import { VisitsList } from 'components/visits';

import { usePageVideosQuery, usePageVisitsQuery } from '.';

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
  const { data: visits } = usePageVisitsQuery(id, 1);

  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  const tabs: MediumTabMap = {
    Videos: <VideosList data={data?.result || []} />,
    Visits: <VisitsList data={visits?.result || []} />,
  };

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
          <Typography variant="subtitle2" color="primary.dark">
            {scraper || 'myanime'}
          </Typography>
          <Typography variant="subtitle2" color="secondary.dark">
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
        <MediumTabs data={tabs} />
      </DialogContent>
    </Dialog>
  );
};

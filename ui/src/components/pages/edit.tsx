import { useState } from 'react';
import { useForm } from 'react-hook-form';

import { Page } from 'client';

import CloseIcon from '@mui/icons-material/Close';
import { Box, Button, Dialog, DialogContent, DialogTitle, IconButton, Stack, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

import { Option, Select, Text } from '@dashotv/components';

import { usePageCreateMutation, usePageMutation } from '.';

const scraperOptions: Option[] = [
  { value: 'myanime', label: 'MyAnime.live' },
  { value: 'jhdanime', label: 'JHDAnime.live' },
  { value: 'animexin', label: 'AnimeXin.vip' },
];
const downloaderOptions: Option[] = [{ value: 'metube', label: 'MeTube' }];

export const PagesEditDialog = ({
  setEditing,
  subject,
}: {
  setEditing: (page: Page | null) => void;
  subject: Page;
}) => {
  const [open, setOpen] = useState(true);
  const { control, handleSubmit } = useForm<Page>({ values: subject });
  const update = usePageMutation();
  const create = usePageCreateMutation();

  const handleClose = () => {
    setEditing(null);
    setOpen(false);
  };

  const submit = (data: Page) => {
    console.log('page:', data);
    if (!data?.id) {
      create.mutate(data, { onSuccess: handleClose, onError: e => console.log('error', e) });
      return;
    }
    update.mutate(data, { onSuccess: handleClose });
  };

  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  return (
    <Dialog open={open} onClose={handleClose} maxWidth="md" fullWidth fullScreen={fullScreen}>
      <DialogTitle>
        <Stack direction="row" spacing={3} alignItems="center" justifyContent="start" width="100%">
          <Typography fontWeight="bolder" color="primary">
            Create Page
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
        <Box component="form" noValidate autoComplete="off" onSubmit={handleSubmit(submit)}>
          <Stack direction="column" spacing={1}>
            <Text name="name" control={control} />
            <Text name="url" control={control} />
            <Stack direction="row" spacing={1}>
              <Select name="scraper" control={control} options={scraperOptions} />
              <Select name="downloader" control={control} options={downloaderOptions} />
            </Stack>
          </Stack>
          <Stack direction="row" spacing={1} sx={{ mt: 3, width: '100%', justifyContent: 'end' }}>
            <Button variant="contained" onClick={() => handleClose()}>
              Cancel
            </Button>
            <Button variant="contained" type="submit">
              Ok
            </Button>
          </Stack>
        </Box>
      </DialogContent>
    </Dialog>
  );
};

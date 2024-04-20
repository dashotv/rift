import { useState } from 'react';
import { useForm } from 'react-hook-form';

import { Page } from 'client';

import CloseIcon from '@mui/icons-material/Close';
import QueueIcon from '@mui/icons-material/Queue';
import { Box, Button, Dialog, DialogContent, DialogTitle, IconButton, Stack, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

import { Text } from '@dashotv/components';

export const PagesCreateDialog = ({ subject }: { subject: Page }) => {
  const [open, setOpen] = useState(false);
  const { control, handleSubmit } = useForm<Page>({ values: subject });

  const handleClose = () => {
    setOpen(false);
    close();
  };

  // const close = () => {
  //   setOpen(false);
  //   handleClose();
  // };

  const submit = (data: Page) => {
    setOpen(false);
    console.log('page:', data);
  };

  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  return (
    <>
      <IconButton aria-label="refresh" color="primary" onClick={() => setOpen(true)}>
        <QueueIcon />
      </IconButton>
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
    </>
  );
};

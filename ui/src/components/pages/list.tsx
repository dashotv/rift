import { useState } from 'react';

import { Page } from 'client';

import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { Chrono, Row } from '@dashotv/components';

import { PagesDialog } from '.';

export const PageList = ({ data }: { data: Page[] }) => {
  const [viewing, setViewing] = useState<Page | null>(null);
  const view = (row: Page) => {
    console.log(row);
    setViewing(row);
  };
  const handleClose = () => {
    setViewing(null);
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
                {/* <Box>{actions && actions(row)}</Box> */}
              </Stack>
            </Stack>
          </Stack>
        </Row>
      ))}
      {viewing && <PagesDialog {...{ open, close: handleClose }} page={viewing} />}
    </Paper>
  );
};

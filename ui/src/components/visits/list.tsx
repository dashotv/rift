import { useEffect, useState } from 'react';

import { Visit } from 'client';

import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { Chrono, Row } from '@dashotv/components';

import { usePagesQuery } from 'components/pages';

export interface PagesMap {
  [key: string]: string;
}

export const VisitsList = ({ data }: { data: Visit[] }) => {
  const { data: pages } = usePagesQuery(1);
  const [pagesMap, setPagesMap] = useState<PagesMap>({});
  useEffect(() => {
    if (pages?.result) {
      const p: PagesMap = {};
      pages.result.forEach(page => {
        if (page.id && page.name) {
          p[page.id] = page.name;
        }
      });
      setPagesMap(p);
    }
  }, [pages]);
  const view = (row: Visit) => {
    console.log(row);
  };

  return (
    <Paper elevation={0} sx={{ width: '100%' }}>
      {data.map((row: Visit) => (
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
                  {row.url}
                </Link>
              </Typography>
            </Stack>
            <Stack
              direction="row"
              spacing={1}
              alignItems="center"
              sx={{ width: '100%', justifyContent: { xs: 'start', md: 'end' } }}
            >
              <Stack
                width={{ xs: '100%', md: 'auto' }}
                minWidth="300px"
                direction="row"
                spacing={1}
                alignItems="center"
                justifyContent="end"
                textAlign="right"
              >
                <Typography noWrap variant="subtitle2" color="gray" pl="3px" width="100%">
                  {row.page_id && pagesMap[row.page_id]}
                </Typography>
                <Typography noWrap variant="subtitle2" color="gray" pl="3px" width="100%">
                  {row.created_at && <Chrono fromNow>{row.created_at}</Chrono>}
                </Typography>
                {/* <Box>{actions && actions(row)}</Box> */}
              </Stack>
            </Stack>
          </Stack>
        </Row>
      ))}
      {/* {viewing && <ReleaseDialog {...{ open, handleClose }} release={viewing} actions={actions} />} */}
    </Paper>
  );
};

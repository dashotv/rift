import Truncate from 'react-truncate-inside';

import { Video } from 'client';

import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { Chrono, Group, Megabytes, Resolution, Row } from '@dashotv/components';

export const VideosList = ({ data }: { data: Video[] }) => {
  const view = (row: Video) => {
    console.log(row);
  };

  return (
    <Paper elevation={0} sx={{ width: '100%' }}>
      {data.map((row: Video) => (
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
                  <Truncate text={`${row.title} ${row.season}x${row.episode}`} ellipsis=" ... " />
                </Link>
              </Typography>
              <Stack
                display={{ xs: 'none', md: 'inherit' }}
                direction="row"
                spacing={1}
                alignItems="center"
                sx={{ pl: 1 }}
              >
                <Group group={row.source} author="" />
                <Resolution resolution={row.resolution} variant="default" />
              </Stack>
            </Stack>
            <Stack
              direction="row"
              spacing={1}
              alignItems="center"
              sx={{ width: '100%', justifyContent: { xs: 'start', md: 'end' } }}
            >
              <Stack
                display={{ xs: 'inherit', md: 'none' }}
                direction="row"
                spacing={1}
                alignItems="center"
                sx={{ pl: 1 }}
              >
                <Resolution resolution={row.resolution} variant="default" />
              </Stack>
              <Stack width={{ xs: '100%', md: 'auto' }} direction="row" spacing={1} alignItems="center">
                {row.size ? <Megabytes value={row.size} ord="bytes" /> : null}
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

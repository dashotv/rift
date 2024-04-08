import { Box } from '@mui/material';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { RoutingTabs, RoutingTabsRoute } from '@dashotv/components';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import Pages from 'pages/pages';
import Videos from 'pages/videos';
import Visits from 'pages/visits';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
  components: {
    MuiLink: {
      styleOverrides: {
        root: {
          textDecoration: 'none',
        },
      },
    },
  },
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 5,
      staleTime: 5 * 1000,
      throwOnError: true,
    },
  },
});

const App = ({ mount }: { mount: string }) => {
  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Pages',
      to: '',
      element: <Pages />,
    },
    {
      label: 'Visits',
      to: 'visits',
      element: <Visits />,
    },
    {
      label: 'Videos',
      to: 'videos',
      element: <Videos />,
    },
  ];
  return (
    <ThemeProvider theme={darkTheme}>
      <QueryClientProvider client={queryClient}>
        <CssBaseline />
        <Box sx={{ pr: 2, pl: 2 }}>
          <RoutingTabs data={tabsMap} mount={mount} />
        </Box>
      </QueryClientProvider>
    </ThemeProvider>
  );
};

export default App;

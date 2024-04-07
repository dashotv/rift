import { Container } from '@mui/material';
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

const App = () => {
  const tabsMap: RoutingTabsRoute[] = [
    {
      label: 'Pages',
      to: '/rift',
      element: <Pages />,
    },
    {
      label: 'Visits',
      to: '/rift/visits',
      element: <Visits />,
    },
    {
      label: 'Videos',
      to: '/rift/videos',
      element: <Videos />,
    },
  ];
  return (
    <ThemeProvider theme={darkTheme}>
      <QueryClientProvider client={queryClient}>
        <CssBaseline />
        <Container>
          <RoutingTabs data={tabsMap} route={'/'} />
        </Container>
      </QueryClientProvider>
    </ThemeProvider>
  );
};

export default App;

import { Container } from '@mui/material';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import { RoutingTabs, RoutingTabsRoute } from '@dashotv/components';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import Pages from 'pages/pages';
import Videos from 'pages/videos';

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
      to: '',
      element: <Pages />,
    },
    {
      label: 'Videos',
      to: '/videos',
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

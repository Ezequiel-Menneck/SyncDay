import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ErrorPage from './components/layout/error-page';
import Layout from './components/layout/layout';
import Financial from './components/pages/financial';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: '/financial',
        element: <Financial />,
      },
      {
        path: '/diet',
        element: <div>Hello dieta!</div>,
      },
      {
        path: '/notes',
        element: <div>Hello notas!</div>,
      },
    ],
  },
]);

const AppRoutes = () => {
  return <RouterProvider router={router} />;
};

export default AppRoutes;

import { Outlet } from 'react-router-dom';
import RoutesMenu from '../routes-menu';

const Layout = () => {
  return (
    <>
      <div className='flex flex-col sm:flex-row '>
        <div className='absolute sm:relative bottom-0 left-0 right-0'>
          <RoutesMenu />
        </div>
        <Outlet />
      </div>
    </>
  );
};

export default Layout;

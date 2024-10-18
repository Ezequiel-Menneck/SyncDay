import { useState } from 'react';
import { useLocation } from 'react-router-dom';
import NavigationLink from '../../common/navigation-link';

const RoutesMenu = () => {
  const location = useLocation();
  const [selectedLink, setSelectedLink] = useState<string>(
    location.pathname.slice(1, location.pathname.length)
  );

  const handleSelectedLink: (route: string) => void = (route: string) => {
    setSelectedLink(route);
  };

  const links: Array<string> = Array.of('financial', 'diet', 'notes');

  return (
    <>
      <section className='w-full h-[72px] sm:h-svh sm:w-20 bg-cyan-700'>
        <div className='flex sm:flex flex-row sm:flex-col items-center justify-center gap-3 sm:gap-1'>
          {links.map((link) => (
            <NavigationLink
              key={link}
              link={link}
              selected={
                link === selectedLink ||
                location.pathname.slice(0, link.length) === link
              }
              onClick={() => handleSelectedLink(link)}
            />
          ))}
        </div>
      </section>
    </>
  );
};

export default RoutesMenu;

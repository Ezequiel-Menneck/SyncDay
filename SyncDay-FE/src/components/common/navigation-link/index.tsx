import AttachMoneyIcon from '@mui/icons-material/AttachMoney';
import ChecklistIcon from '@mui/icons-material/Checklist';
import RestaurantIcon from '@mui/icons-material/Restaurant';
import { Link } from 'react-router-dom';

type IconsMapping = {
  financial: JSX.Element;
  diet: JSX.Element;
  notes: JSX.Element;
  [key: string]: JSX.Element;
};

const iconsMapping: IconsMapping = {
  financial: <AttachMoneyIcon />,
  diet: <RestaurantIcon />,
  notes: <ChecklistIcon />,
} as const;

type IconKey = keyof typeof iconsMapping;

const NavigationLink: React.FC<{
  link: string;
  selected: boolean;
  onClick: () => void;
}> = ({ link, selected, onClick }) => {
  return (
    <>
      <Link
        onClick={onClick}
        className={`flex items-center h-5 sm:h-11 justify-center p-7 sm:p-7 rounded-lg m-2 sm:m-3 font-bold duration-200 w-auto sm:w-12 sm:hover:bg-cyan-500 ${
          selected
            ? 'bg-slate-50 text-black hover:bg-slate-50'
            : 'bg-cyan-600 text-white'
        }`}
        to={link}
      >
        {iconsMapping[link as IconKey]}
      </Link>
    </>
  );
};

export default NavigationLink;

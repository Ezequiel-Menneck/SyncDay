import FinancialCard from '../../common/financial-card';
import Transaction from '../../common/transaction-card/transaction';

const Financial = () => {
  return (
    <>
      <div className='h-[calc(100vh-72px)] sm:h-full w-screen'>
        <div className='flex flex-col h-full w-5/6 items-center m-auto'>
          <FinancialCard />
          <Transaction />
          <Transaction />
          <Transaction />
          <Transaction />
          <Transaction />
        </div>
      </div>
    </>
  );
};

export default Financial;

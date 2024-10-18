const FinancialCard = () => {
  return (
    <>
      <section className='flex flex-col h-1/3 w-full items-center justify-center mt-12 mb-12'>
        <div className='flex flex-col items-center justify-center border-2 border-red-400 w-full h-3/4 rounded-2xl bg-gray-300  '>
          <div className='flex flex-col w-3/4 h-3/4 items-center justify-center'>
            <span className='font-extrabold text-4xl'>R$: 10000</span>
          </div>
          <div className='flex flex-col h-96 align-bottom items-center justify-center '>
            <span className='text-green-600 text-lg font-semibold'>
              Income: R$: 20000
            </span>
            <span className='text-red-600 text-lg font-semibold'>
              Expense: R$: 10000
            </span>
          </div>
        </div>
      </section>
    </>
  );
};

export default FinancialCard;

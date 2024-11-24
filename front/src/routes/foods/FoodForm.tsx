type Props = {
  closeModal: Function;
};

const FoodForm = (p: Props) => {
  const onSubmit = (data: any) => {
    p.closeModal();
  };

  return (
    <form id='foodForm' onSubmit={onSubmit}>
      <input
        type='text'
        placeholder='Название'
        className='input input-bordered w-full max-w-xs'
      />
      <button type='submit' className='btn btn-primary'>
        Save
      </button>
      <div id='btn' className='btn' onClick={() => p.closeModal()}>
        Close
      </div>
    </form>
  );
};

export default FoodForm;

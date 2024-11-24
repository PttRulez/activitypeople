import Modal from "@/components/Modal";
import { modalId } from "@/types/constants";
import FoodForm from "./FoodForm";

const Foods = () => {
  return (
    <div className='p-4 text-3xl flex justify-between'>
      <p>Foods bleat'</p>
      <button
        className='btn btn-secondary'
        onClick={() =>
          (document.getElementById(modalId) as HTMLDialogElement).showModal()
        }
      >
        Add Food
      </button>
      <Modal modalId={modalId} title='Добавь еду'>
        <FoodForm
          closeModal={() =>
            (document.getElementById(modalId) as HTMLDialogElement).close()
          }
        />
      </Modal>
    </div>
  );
};

export default Foods;

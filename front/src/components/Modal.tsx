import React from "react";

type Props = {
  children: React.ReactElement;
  modalId: string;
  title: string;
};
const Modal = (p: Props) => {
  return (
    <dialog id={p.modalId} className='modal' hx-ext='response-targets'>
      <div className='modal-box'>
        <h3 className='text-lg font-bold'>{p.title}</h3>
        <div className='modal-action'>{p.children}</div>
      </div>
      <form method='dialog' className='modal-backdrop'>
        <button>close</button>
      </form>
    </dialog>
  );
};
export default Modal;

import classNames from "classnames";
import React, { ChangeEvent, useEffect, useRef, useState } from "react";

interface Props<D> {
  getOptionLabel(i: D): string;
  items: D[];
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  inputProps: {
    placeholder?: string;
    [key: string]: any;
  };
  onChoose(val: D): void;
}

const Autocomplete = <D extends object>(props: Props<D>) => {
  const { getOptionLabel, items, inputProps, onChange, onChoose } = props;
  const ref = useRef<HTMLDivElement>(null);
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (items.length > 0) {
      setOpen(true);
    } else {
      setOpen(false);
    }
  }, [items]);

  return (
    <div
      className={classNames({
        "dropdown w-full": true,
        "dropdown-open": open,
      })}
      ref={ref}
    >
      <input
        {...inputProps}
        type='text'
        className='input input-bordered w-full'
        tabIndex={0}
        onChange={(e) => {
          onChange(e);
        }}
        autoComplete='off'
      />

      {items.length > 0 && open && (
        <div className='dropdown-content bg-base-200 top-14 max-h-96 overflow-auto flex-col rounded-md z-10'>
          <ul
            className='menu menu-compact '
            style={{ width: ref.current?.clientWidth }}
          >
            {items.map((item, index) => {
              return (
                <li
                  key={index}
                  tabIndex={index + 1}
                  onClick={async () => {
                    onChoose(item);
                    setOpen(false);
                  }}
                  className='border-b border-b-base-content/10 w-full'
                >
                  <span>{getOptionLabel(item)}</span>
                </li>
              );
            })}
          </ul>
        </div>
      )}
    </div>
  );
};

export default Autocomplete;

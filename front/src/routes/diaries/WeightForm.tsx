import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import dayjs from "dayjs";
import { Controller, useForm } from "react-hook-form";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { CreateWeightData, CreateWeightSchema } from "src/validation/weight";

type Props = {
  onSuccess: Function;
};

const WeightForm = (p: Props) => {
  const {
    control,
    formState: { errors },
    handleSubmit,
    register,
    reset,
    setValue,
    watch,
  } = useForm<CreateWeightData>({
    defaultValues: {
      date: dayjs().toISOString(),
    },
    resolver: zodResolver(CreateWeightSchema),
  });

  const axios = useAxiosPrivate();

  const createWeight = useMutation({
    mutationFn: async (data: CreateWeightData) => {
      await axios.post("/weight", data);
    },
    onSuccess: (data: any) => {
      reset();
      p.onSuccess();
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: CreateWeightData) => {
    createWeight.mutate(data);
  };

  return (
    <form id='foodForm' onSubmit={handleSubmit(onSubmit)} className='pt-3'>
      <div className='flex flex-col gap-2'>
        <section>
          <Controller
            control={control}
            name='date'
            render={({ field }) => (
              <input
                type='date'
                className='input input-bordered w-full'
                autoComplete='off'
                onChange={(e) => {
                  setValue(
                    "date",
                    dayjs(new Date(e.target.value)).toISOString()
                  );
                }}
                value={dayjs(new Date(watch("date"))).format("YYYY-MM-DD")}
              />
            )}
          />
        </section>

        <section>
          <div className='label'>
            <span className='label-text'>Вес, кг</span>
          </div>
          <input
            type='number'
            step='.1'
            className='input input-bordered w-full'
            {...register("weight", {
              valueAsNumber: true,
            })}
          />
          <div className='label'>
            <span className='label-text-alt text-error'>
              {errors.weight?.message}
            </span>
          </div>
        </section>

        <div className='flex justify-end gap-3 pt-5'>
          <button type='submit' className='btn btn-primary'>
            Сохранить
          </button>
          <div id='btn' className='btn' onClick={() => p.onSuccess()}>
            Отмена
          </div>
        </div>
      </div>
    </form>
  );
};

export default WeightForm;

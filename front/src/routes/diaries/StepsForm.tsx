import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import dayjs from "dayjs";
import { Controller, useForm } from "react-hook-form";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { CreateStepsData, CreateStepsSchema } from "src/validation/steps";

type Props = {
  onSuccess: Function;
};

const StepsForm = (p: Props) => {
  const {
    control,
    formState: { errors },
    handleSubmit,
    register,
    reset,
    setValue,
    watch,
  } = useForm<CreateStepsData>({
    defaultValues: {
      date: dayjs().toISOString(),
    },
    resolver: zodResolver(CreateStepsSchema),
  });

  const axios = useAxiosPrivate();
  const queryClient = useQueryClient();
  const createSteps = useMutation({
    mutationFn: async (data: CreateStepsData) => {
      await axios.post("/steps", data);
    },
    onSuccess: (data: any) => {
      queryClient.invalidateQueries({ queryKey: ["diaries"] });
      reset();
      p.onSuccess();
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: CreateStepsData) => {
    createSteps.mutate(data);
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
                  if (e.target.value === "") return;
                  console.log(e.target.value);
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
            <span className='label-text'>Шаги</span>
          </div>
          <input
            type='number'
            step='.1'
            className='input input-bordered w-full'
            {...register("steps", {
              valueAsNumber: true,
            })}
          />
          <div className='label'>
            <span className='label-text-alt text-error'>
              {errors.steps?.message}
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

export default StepsForm;

import { CreateFoodData, CreateFoodSchema } from "src/validation/food";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";

type Props = {
  onSuccess: Function;
};

const FoodForm = (p: Props) => {
  const {
    formState: { errors },
    handleSubmit,
    register,
    reset,
  } = useForm<CreateFoodData>({
    resolver: zodResolver(CreateFoodSchema),
  });

  const axios = useAxiosPrivate();

  const createFood = useMutation({
    mutationFn: async (data: CreateFoodData) => {
      await axios.post("/food", data);
    },
    onSuccess: (data: any) => {
      reset();
      p.onSuccess();
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: CreateFoodData) => {
    createFood.mutate(data);
  };

  return (
    <form id='foodForm' onSubmit={handleSubmit(onSubmit)} className='pt-3'>
      <div className='flex flex-col gap-2'>
        <section>
          <div className='label'>
            <span className='label-text'>Название</span>
          </div>
          <input
            type='text'
            className='input input-bordered w-full'
            autoComplete='off'
            {...register("name")}
          />
          <div className='label'>
            <span className='label-text-alt text-error'>
              {errors.name?.message}
            </span>
          </div>
        </section>

        <section className='flex justify-between'>
          <div>
            <div className='label'>
              <span className='label-text'>Калорийность на 100 г</span>
            </div>
            <input
              type='number'
              className='input input-bordered w-full'
              {...register("caloriesPer100", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.caloriesPer100?.message}
              </span>
            </div>
          </div>

          <div>
            <div className='label'>
              <span className='label-text'>Углеводов на 100 г</span>
            </div>
            <input
              type='number'
              className='input input-bordered w-full'
              {...register("carbs", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.carbs?.message}
              </span>
            </div>
          </div>
        </section>

        <section className='flex justify-between'>
          <div>
            <div className='label'>
              <span className='label-text'>Жир на 100 г</span>
            </div>
            <input
              type='number'
              className='input input-bordered w-full'
              {...register("fat", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.fat?.message}
              </span>
            </div>
          </div>

          <div>
            <div className='label'>
              <span className='label-text'>Белки на 100 г</span>
            </div>
            <input
              type='number'
              step='.1'
              className='input input-bordered w-full'
              {...register("protein", {
                valueAsNumber: true,
              })}
            />
            <div className='label'>
              <span className='label-text-alt text-error'>
                {errors.protein?.message}
              </span>
            </div>
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

export default FoodForm;

import {
  CreateMealData,
  CreateMealSchema,
  FoodInMealData,
} from "@/validation/meal";
import { zodResolver } from "@hookform/resolvers/zod";
import dayjs from "dayjs";
import { useEffect } from "react";
import { useFieldArray, useForm, useWatch } from "react-hook-form";
import FoodSearch from "./FoodSearch";
import { useMutation } from "@tanstack/react-query";
import useAxiosPrivate from "@/hooks/useAxiosPrivate";

type Props = {
  onSuccess: Function;
};
const emptyFood: FoodInMealData = {
  calories: 0,
  caloriesPer100: 0,
  id: 0,
  name: "",
  weight: 0,
};

const MealForm = (p: Props) => {
  const {
    control,
    formState: { errors },
    handleSubmit,
    register,
    reset,
    setValue,
    watch,
  } = useForm<CreateMealData>({
    defaultValues: {
      calories: 0,
      date: dayjs().format("YYYY-MM-DD"),
      foods: [emptyFood],
      name: "Обед",
    },
    resolver: zodResolver(CreateMealSchema),
  });

  const {
    append,
    fields: foodFields,
    update,
  } = useFieldArray<CreateMealData, "foods">({
    control,
    name: "foods",
  });

  const axios = useAxiosPrivate();

  const createMeal = useMutation({
    mutationFn: async (data: CreateMealData) => {
      await axios.post("/meal", data);
    },
    onSuccess: (data: any) => {
      reset();
      p.onSuccess();
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: CreateMealData) => {
    createMeal.mutate(data);
  };

  const watchedWeight = useWatch({
    control,
    name: "foods",
  });
  // const watchedWeight = watch("foods")?.map((item) => item?.weight);

  useEffect(() => {
    const c = watchedWeight.reduce((acc, curr) => acc + curr.calories, 0);
    setValue("calories", c);
  }, [watchedWeight]);

  return (
    <form id='mealForm' onSubmit={handleSubmit(onSubmit)} className='pt-3'>
      <section className='flex justify-between gap-3'>
        <div>
          <input
            placeholder='Название'
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
        </div>
        <div>{watch("calories")} калорий</div>
      </section>
      <section>
        <input
          type='date'
          className='input input-bordered w-full'
          autoComplete='off'
          {...register("date")}
        />
      </section>
      <div className='label'>
        <span className='label-text'>Что ели?</span>
      </div>
      <div className='grid grid-cols-[4fr_1fr_1fr] gap-2'>
        {foodFields.map((field, index) => (
          <>
            <FoodSearch
              key={field.name + index}
              inputProps={{ ...register(`foods.${index}.name`) }}
              onChoose={(f) => {
                update(index, {
                  calories: 0,
                  caloriesPer100: f.calories,
                  id: f.id,
                  name: f.name,
                  weight: 0,
                });
              }}
            />
            <input
              type='number'
              placeholder='0 г'
              className='input input-bordered w-full'
              value={watch(`foods.${index}.weight`).toString()}
              onChange={(e) => {
                setValue(`foods.${index}.weight`, Number(e.target.value));
                const c = Math.floor(
                  (Number(e.target.value) / 100) * field.caloriesPer100
                );
                setValue(`foods.${index}.calories`, c);
              }}
            />
            <input
              type='number'
              className='input input-bordered w-full'
              value={watch(`foods.${index}.calories`)}
              disabled
            />
          </>
        ))}
      </div>
      <div className='flex justify-between gap-3 pt-5'>
        <div className='ml-1 cursor-pointer' onClick={() => append(emptyFood)}>
          +
        </div>
        <div>
          <button type='submit' className='btn btn-primary'>
            Сохранить
          </button>
          <div
            id='btn'
            className='btn'
            onClick={() => {
              p.onSuccess();
              reset();
            }}
          >
            Отмена
          </div>
        </div>
      </div>
    </form>
  );
};

export default MealForm;

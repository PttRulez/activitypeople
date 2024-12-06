import {
  CreateMealData,
  CreateMealSchema,
  FoodInMealData,
} from "src/validation/meal";
import { zodResolver } from "@hookform/resolvers/zod";
import dayjs from "dayjs";
import { useEffect } from "react";
import { Controller, useFieldArray, useForm, useWatch } from "react-hook-form";
import FoodSearch from "./FoodSearch";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";

type Props = {
  onSuccess: Function;
};
const emptyFood: FoodInMealData = {
  calories: 0,
  caloriesPer100: 0,
  name: "",
  weight: 0,
};

const MealForm = (p: Props) => {
  const queryClient = useQueryClient();
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
      date: dayjs().toISOString(),
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
      queryClient.invalidateQueries({ queryKey: ["meals"] });
    },
    onError: (error: any) => {
      console.error("onError", error);
    },
  });

  const onSubmit = (data: CreateMealData) => {
    console.log("onSubmit", data);
    createMeal.mutate(data);
  };

  const watchedFoods = useWatch({
    control,
    name: "foods",
  });

  useEffect(() => {
    const c = watchedFoods.reduce((acc, curr) => acc + curr.calories, 0);
    setValue("calories", c);
  }, [watchedFoods]);

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
        <Controller
          control={control}
          name='date'
          render={({ field }) => (
            <input
              type='date'
              className='input input-bordered w-full'
              autoComplete='off'
              onChange={(e) => {
                setValue("date", dayjs(new Date(e.target.value)).toISOString());
              }}
              value={dayjs(new Date(watch("date"))).format("YYYY-MM-DD")}
            />
          )}
        />
      </section>

      <div className='label grid grid-cols-[3fr_1fr_1fr_1fr] gap-2 text-xs text-center'>
        <span className='label-text'>Что ели?</span>
        <span className='label-text'>вес</span>
        <span className='label-text'>ккал/100</span>
        <span className='label-text'>ккал</span>
      </div>
      <div>
        {foodFields.map((field, index) => (
          <div className='grid grid-cols-[3fr_1fr_1fr_1fr] gap-2 mb-2'>
            <FoodSearch
              reactKey={field.name + index}
              inputProps={{
                ...register(`foods.${index}.name`),
                value: watch(`foods.${index}.name`),
                onChange: (e: React.ChangeEvent<HTMLInputElement>) => {
                  setValue(`foods.${index}.name`, e.target.value);
                },
              }}
              onChoose={(f) => {
                update(index, {
                  calories: 0,
                  caloriesPer100: f.caloriesPer100,
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
                  (Number(e.target.value) / 100) *
                    watch(`foods.${index}.caloriesPer100`)
                );
                setValue(`foods.${index}.calories`, c);
              }}
            />
            <input
              type='number'
              className='input input-bordered w-full'
              {...register(`foods.${index}.caloriesPer100`, {
                valueAsNumber: true,
              })}
              onChange={(e) => {
                setValue(
                  `foods.${index}.caloriesPer100`,
                  Number(e.target.value)
                );
                const c = Math.floor(
                  (watch(`foods.${index}.weight`) / 100) *
                    Number(e.target.value)
                );
                setValue(`foods.${index}.calories`, c);
              }}
              value={watch(`foods.${index}.caloriesPer100`)}
            />
            <input
              type='number'
              className='input input-bordered w-full'
              value={watch(`foods.${index}.calories`)}
              disabled
            />
          </div>
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

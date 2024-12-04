import Modal from "src/components/Modal";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { MealResponse } from "src/types/food";
import { useQuery } from "@tanstack/react-query";
import dayjs from "dayjs";
import React, { useMemo } from "react";
import FoodForm from "./FoodForm";
import MealForm from "./MealForm";

type Day = { date: Date; meals: MealResponse[] };

const Foods = () => {
  const axios = useAxiosPrivate();

  const { data: meals } = useQuery({
    queryKey: ["meals"],
    queryFn: async () => {
      const res = await axios.get<MealResponse[]>("/meal");
      return res.data;
    },
  });

  const days = useMemo<Day[]>(() => {
    if (!meals || meals.length === 0) return [];

    let sorterMeals = meals.sort((a, b) => (a.date > b.date ? -1 : 1));
    let curDate = sorterMeals[0].date;
    let curIndex = 0;
    let days: Day[] = [{ date: curDate, meals: [] }];

    for (const m of sorterMeals) {
      if (dayjs(m.date).isSame(curDate, "day")) {
        days[curIndex].meals.push(m);
      } else {
        curDate = m.date;
        days.push({ date: curDate, meals: [m] });
        curIndex++;
      }
    }

    return days;
  }, [meals]);
  console.log("DAYS:", days);
  return (
    <div className='p-4'>
      <section className='text-3xl flex justify-end gap-4'>
        <button
          className='btn btn-secondary'
          onClick={() =>
            (
              document.getElementById("mealModal") as HTMLDialogElement
            ).showModal()
          }
        >
          Add Meal
        </button>
        <button
          className='btn btn-secondary'
          onClick={() =>
            (
              document.getElementById("foodModal") as HTMLDialogElement
            ).showModal()
          }
        >
          Add Food
        </button>
      </section>

      <section>
        {days?.map((d) => (
          <div className='mb-8' key={d.date.toString()}>
            <p>
              {new Date(d.date).toLocaleString("Ru-ru", {
                month: "numeric",
                day: "numeric",
              })}{" "}
              -{" "}
              <span>{d.meals.reduce((acc, cur) => cur.calories + acc, 0)}</span>
            </p>

            <div className='divider'></div>

            {d.meals.map((m) => (
              <div className='collapse bg-base-200 mb-4' key={m.id + m.name}>
                <input type='checkbox' />
                <div className='collapse-title  font-medium'>
                  {`${m.name} - ${m.calories}`}
                </div>
                <div className='collapse-content'>
                  {m.foods.map((f) => (
                    <div className='text-sm' key={f.name + m.id}>
                      {`${f.name} - ${f.weight} г - ${f.calories} калорий`}
                    </div>
                  ))}
                </div>
              </div>
            ))}
          </div>
        ))}
      </section>

      <Modal modalId='mealModal' title='Новый прием пищи'>
        <MealForm
          onSuccess={() =>
            (document.getElementById("mealModal") as HTMLDialogElement).close()
          }
        />
      </Modal>
      <Modal modalId='foodModal' title='Добавь еду'>
        <FoodForm
          onSuccess={() =>
            (document.getElementById("foodModal") as HTMLDialogElement).close()
          }
        />
      </Modal>
    </div>
  );
};

export default Foods;

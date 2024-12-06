import { useQuery } from "@tanstack/react-query";
import classNames from "classnames";
import dayjs from "dayjs";
import { useMemo, useState } from "react";
import Modal from "src/components/Modal";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { DiariesResponse, DiaryResponse } from "src/types/diary";
import StepsForm from "./StepsForm";
import WeightForm from "./WeightForm";
import useAuth from "src/hooks/useAuth";

type DiaryDay = {
  date: string;
  diary: DiaryResponse | undefined;
};

const Diaries = () => {
  const axios = useAxiosPrivate();
  const { auth } = useAuth();
  const [until, setUntil] = useState(dayjs().endOf("week").add(1, "day"));
  const [from, setFrom] = useState(
    dayjs().endOf("week").add(1, "day").subtract(27, "day")
  );

  const diaryQuery = useQuery({
    queryKey: ["diaries"],
    queryFn: async () => {
      const res = await axios.get<DiariesResponse>("/diary", {
        params: {
          from: from.format("YYYY-MM-DD"),
          until: until.format("YYYY-MM-DD"),
        },
      });
      return res.data;
    },
  });

  let diaries = useMemo<{ weight: number; diaries: DiaryDay[] }[]>(() => {
    if (!diaryQuery.data) return [];

    let res: { weight: number; diaries: DiaryDay[] }[] = [];
    let curDate = until;
    let curIndex = 0;

    while (curDate >= from) {
      res[curIndex] = {
        weight: 0,
        diaries: [],
      };
      for (let i = 0; i < 7; i++) {
        res[curIndex].diaries.unshift({
          date: curDate.format("DD.MM"),
          diary: diaryQuery.data.diaries[curDate.format("YYYY-MM-DD")],
        });
        curDate = curDate.subtract(1, "day");
      }
      res[curIndex].weight = res[curIndex].diaries.reduce((acc, cur) => {
        if (!cur.diary) return acc;
        return cur.diary?.weight > 0 ? cur.diary?.weight : acc;
      }, 0);
      curIndex++;
    }

    return res;
  }, [diaryQuery]);

  return (
    <div className='p-4'>
      <section className='text-3xl flex justify-end gap-4'>
        <button
          className='btn btn-secondary'
          onClick={() =>
            (
              document.getElementById("stepsModal") as HTMLDialogElement
            ).showModal()
          }
        >
          Add Steps
        </button>

        <button
          className='btn btn-secondary'
          onClick={() =>
            (
              document.getElementById("weightModal") as HTMLDialogElement
            ).showModal()
          }
        >
          Add Weight
        </button>
      </section>
      <section className='mt-14'>
        {diaries.map((week, index) => (
          <div key={index} className='flex gap-3'>
            <ul className='timeline flex justify-between w-full mb-20 flex-auto'>
              {week.diaries.map(({ date, diary: d }, index) => (
                <li className='grow' key={date}>
                  {index > 0 && <hr className='grow' />}
                  {d && (
                    <div
                      className={classNames({
                        "tooltip timeline-start border-2  rounded p-2": true,
                        "border-green-500": d.calories <= 0,
                        "border-red-500": d.calories > 0,
                      })}
                      data-tip={`\n ${d.caloriesConsumed} - ${auth.user.bmr} - ${d.caloriesBurned} (${d.steps})`}
                    >
                      {d?.calories}
                    </div>
                  )}
                  <div className='timeline-middle tooltip' data-tip={date}>
                    <svg
                      xmlns='http://www.w3.org/2000/svg'
                      viewBox='0 0 20 20'
                      fill='currentColor'
                      className='h-5 w-5 '
                    >
                      <path
                        fillRule='evenodd'
                        d='M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z'
                        clipRule='evenodd'
                      />
                    </svg>
                  </div>
                  {d?.weight && d.weight > 0 ? (
                    <div className='timeline-end timeline-box'>{d.weight}</div>
                  ) : (
                    ""
                  )}
                  {index < 6 && <hr />}
                </li>
              ))}
            </ul>
            <div className='text-center'>
              <div className='stat-title'>Результат</div>
              <div
                className={classNames({
                  "stat-value": true,
                  "text-green-500":
                    week.weight - diaries[index + 1]?.weight <= 0.2,
                  "text-red-500":
                    week.weight - diaries[index + 1]?.weight > 0.2,
                })}
              >
                {week.weight}
              </div>
            </div>
          </div>
        ))}
      </section>
      <Modal modalId='weightModal' title='Новый вес'>
        <WeightForm
          onSuccess={() =>
            (
              document.getElementById("weightModal") as HTMLDialogElement
            ).close()
          }
        />
      </Modal>
      <Modal modalId='stepsModal' title='Инфа по шагам'>
        <StepsForm
          onSuccess={() =>
            (document.getElementById("stepsModal") as HTMLDialogElement).close()
          }
        />
      </Modal>
    </div>
  );
};

export default Diaries;

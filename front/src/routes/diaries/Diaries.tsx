import Modal from "src/components/Modal";
import WeightForm from "./WeightForm";
import { useQuery } from "@tanstack/react-query";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { DiaryResponse } from "src/types/diary";
import { useEffect, useMemo, useState } from "react";
import dayjs from "dayjs";

const Diaries = () => {
  const axios = useAxiosPrivate();
  const [until, setUntil] = useState(dayjs());
  const [from, setFrom] = useState(dayjs().subtract(28, "day"));

  const diaryQuery = useQuery({
    queryKey: ["diaries"],
    queryFn: async () => {
      const res = await axios.get<Record<string, DiaryResponse>>("/diary", {
        params: {
          from: from.format("YYYY-MM-DD"),
          until: until.format("YYYY-MM-DD"),
        },
      });
      return res.data;
    },
  });

  let diaries = useMemo(() => {
    if (!diaryQuery.data) return [];

    let curDate = from;

    while (curDate <= until) {
      curDate = dayjs(curDate).add(1, "day");
    }
  }, [diaryQuery]);

  return (
    <div className='p-4'>
      <section className='text-3xl flex justify-end gap-4'>
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

      <section></section>

      <Modal modalId='weightModal' title='Новый вес'>
        <WeightForm
          onSuccess={() =>
            (
              document.getElementById("weightModal") as HTMLDialogElement
            ).close()
          }
        />
      </Modal>
    </div>
  );
};

export default Diaries;

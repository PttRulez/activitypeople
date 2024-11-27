import useAxiosPrivate from "@/hooks/useAxiosPrivate";
import { ActivityResponse } from "@/types/activity";
import { monthNameMap } from "@/types/enums";
import { getDay } from "@/utils/utils";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import dayjs from "dayjs";
import { useMemo } from "react";
import { Link, useSearchParams } from "react-router-dom";
import ActivityCard from "./ActivityCard";

type ActivitiesState = {
  monthName: string;
  monthNumber: number;
  nextButtonLink: string | null;
  prevButtonLink: string;
  year: number;
};

type Day = ActivityResponse | null;

const Activities = () => {
  const queryClient = useQueryClient();
  const [searchParams, _] = useSearchParams();
  const now = dayjs();
  let date = dayjs();
  let year = Number(searchParams.get("year"));
  let month = Number(searchParams.get("month"));
  if (year) {
    date = date.year(year);
  } else {
    year = now.year();
  }
  if (month) {
    date = date.month(month - 1);
  } else {
    month = now.month() + 1;
  }

  const calendar = useMemo<ActivitiesState>(() => {
    let nextButtonLink: string | null = null;
    if (date.isBefore(now)) {
      nextButtonLink = `/activities?year=${date.add(1, "month").year()}&month=${
        date.add(1, "month").month() + 1
      }`;
    }

    return {
      monthName: monthNameMap[date.month()],
      monthNumber: date.month() + 1,
      nextButtonLink: nextButtonLink,
      prevButtonLink: `/activities?year=${date
        .subtract(1, "month")
        .year()}&month=${date.subtract(1, "month").month() + 1}`,
      year,
    };
  }, [year, month]);

  const axios = useAxiosPrivate();

  const query = useQuery({
    queryKey: ["activities", year, month],
    queryFn: async () => {
      let params: any = {};
      if (calendar.monthNumber) {
        params.month = calendar.monthNumber;
      }
      if (calendar.monthNumber) {
        params.year = calendar.year;
      }
      const data = await axios.get<ActivityResponse[]>("/activities", {
        params,
      });
      return data;
    },
  });

  const syncStrava = useMutation({
    mutationKey: ["sync-strava"],
    mutationFn: async () => {
      await axios.get("/sync-strava");
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["activities"] });
      console.log("Synced with Strava!");
    },
  });

  const days: Day[] = useMemo(() => {
    if (!query.data) return [];

    const date = dayjs(`${calendar.year}-${calendar.monthNumber}-01`);
    const daysBeforeFirst = getDay(date) - 1;
    const daysAfterLast = 7 - getDay(date.endOf("month"));
    const totalDays = daysBeforeFirst + daysAfterLast + date.daysInMonth();
    console.log(totalDays);
    const days = new Array(totalDays).fill(null);
    console.log("daysBeforeFirst", daysBeforeFirst);
    for (const a of query.data.data) {
      days[dayjs(a.date).date() - 1 + daysBeforeFirst] = a;
    }

    return days;
  }, [query.data, calendar]);

  return (
    calendar && (
      <>
        <div className='p-4 text-3xl flex justify-between'>
          <div className='min-w-56 flex justify-between'>
            <Link to={calendar.prevButtonLink}>
              <i className='fa-solid fa-arrow-left  cursor-pointer'></i>
            </Link>
            <span>{calendar.monthName}</span>
            {calendar.nextButtonLink && (
              <Link to={calendar.nextButtonLink}>
                <i className='fa-solid fa-arrow-right cursor-pointer'></i>
              </Link>
            )}
            <span className='invisible'>
              <i className='fa-solid fa-arrow-right cursor-pointer'></i>
            </span>
          </div>
          <div>
            <button
              className='btn btn-primary'
              onClick={() => syncStrava.mutate()}
            >
              Sync Strava
              {syncStrava.status === "pending" && (
                <i className='fa-solid fa-spinner animate-spin htmx-indicator'></i>
              )}
            </button>
          </div>
        </div>
        <div className='grid grid-cols-7 gap-4'>
          {days.map((a) => (
            <ActivityCard a={a} />
          ))}
        </div>
      </>
    )
  );
};

export default Activities;

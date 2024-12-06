import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { ActivityResponse } from "src/types/activity";
import { monthNameMap } from "src/types/enums";
import { getDay } from "src/utils/utils";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import dayjs from "dayjs";
import { useMemo } from "react";
import { Link, useSearchParams } from "react-router-dom";
import ActivityDayCard from "./ActivityDayCard";

type ActivitiesState = {
  daysBeforeFirst: number;
  firstDate: dayjs.Dayjs;
  lastDate: dayjs.Dayjs;
  monthName: string;
  nextButtonLink: string | null;
  prevButtonLink: string;
  totalDays: number;
};

export type ActivityDay = {
  activities: ActivityResponse[];
  date: Date;
};

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

    const daysBeforeFirst = getDay(date.startOf("month")) - 1;
    const daysAfterLast = 7 - getDay(date.endOf("month"));
    const totalDays = daysBeforeFirst + daysAfterLast + date.daysInMonth();

    return {
      daysBeforeFirst,
      firstDate: date.startOf("month").subtract(daysBeforeFirst, "days"),
      lastDate: date.endOf("month").add(daysAfterLast, "days"),
      monthName: monthNameMap[date.month()],
      nextButtonLink: nextButtonLink,
      prevButtonLink: `/activities?year=${date
        .subtract(1, "month")
        .year()}&month=${date.subtract(1, "month").month() + 1}`,
      year,
      totalDays,
    };
  }, [year, month]);

  const axios = useAxiosPrivate();

  const { data: activitiesResponse } = useQuery({
    queryKey: ["activities", year, month],
    queryFn: async () => {
      const data = await axios.get<Record<string, ActivityResponse[]>>(
        "/activities",
        {
          params: {
            from: calendar.firstDate.format("YYYY-MM-DD"),
            until: calendar.lastDate.format("YYYY-MM-DD"),
          },
        }
      );
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
    },
  });

  const days: ActivityDay[] = useMemo(() => {
    if (!activitiesResponse) return [];

    const days = [];
    let curDate = calendar.firstDate;
    while (!curDate.isAfter(calendar.lastDate)) {
      let day = {
        activities: [] as ActivityResponse[],
        date: curDate.toDate(),
      };
      day.activities =
        activitiesResponse.data[curDate.format("YYYY-MM-DD")] ?? [];
      days.push(day);

      curDate = curDate.add(1, "day");
    }

    return days;
  }, [activitiesResponse, calendar]);

  return (
    calendar && (
      <div>
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
          {days.map((d, i) => (
            <ActivityDayCard
              key={d.date.toLocaleString("ru-RU", {
                year: "numeric",
                month: "2-digit",
                day: "2-digit",
              })}
              day={d}
            />
          ))}
        </div>
      </div>
    )
  );
};

export default Activities;

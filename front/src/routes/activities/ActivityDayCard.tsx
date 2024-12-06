import { useMutation, useQueryClient } from "@tanstack/react-query";
import classNames from "classnames";
import dayjs from "dayjs";
import React from "react";
import useAxiosPrivate from "src/hooks/useAxiosPrivate";
import { SportType } from "src/types/activity";
import { ActivityDay } from "./Activities";

type Props = {
  day: ActivityDay;
};

const ActivityDayCard = ({ day: d }: Props) => {
  const axios = useAxiosPrivate();
  const queryClient = useQueryClient();
  const hydrateActivity = useMutation({
    mutationKey: ["hydrate"],
    mutationFn: async ({ id, sourceId }: { id: number; sourceId: number }) => {
      await axios.get(`/hydrate/${sourceId}`);
      return id;
    },
    onSuccess: (id: number) => {
      queryClient.invalidateQueries({ queryKey: ["activities"] });
    },
  });

  if (d.activities.length === 0) {
    return (
      <Wrapper>
        <div className='card-body'>
          <div>
            <div className='flex'>
              <p>{dayjs(d.date).format("DD.MM")}</p>
            </div>
          </div>
        </div>
      </Wrapper>
    );
  }

  if (d.activities.length === 1) {
    const a = d.activities[0];
    return (
      <Wrapper>
        <div className='card-body'>
          <div>
            <div className='flex'>
              <p>{dayjs(d.date).format("DD.MM")}</p>
              <SportIcon
                activityId={a.id}
                calories={a.calories}
                sportType={a.sportType}
                onClick={() => {
                  hydrateActivity.mutate({
                    id: a.id,
                    sourceId: a.sourceId,
                  });
                }}
              />
            </div>
            <h2
              className='card-title text-ellipsis whitespace-nowrap overflow-hidden
             inline-block max-w-full'
            >
              {a.name}
            </h2>
            {a.distance > 0 && (
              <div className='flex justify-between'>
                <span>{a.paceString}</span>
                <span>{(a.distance / 1000).toFixed(2)} km</span>
              </div>
            )}
            <span>{a.calories} kcal</span>
          </div>
        </div>
      </Wrapper>
    );
  }

  if (d.activities.length > 1) {
    return (
      <Wrapper>
        <div className='card-body'>
          <p>{dayjs(d.date).format("DD.MM")}</p>
          {d.activities.map((a) => (
            <div key={a.id}>
              <div className='flex justify-between'>
                <SportIcon
                  activityId={a.id}
                  calories={a.calories}
                  sportType={a.sportType}
                  onClick={() => {
                    hydrateActivity.mutate({
                      id: a.id,
                      sourceId: a.sourceId,
                    });
                  }}
                />
                <span>{a.calories} kcal</span>
              </div>
              {a.distance > 0 && (
                <div className='flex justify-between'>
                  <span>{(a.distance / 1000).toFixed(2)} km</span>
                  <span>{a.paceString}</span>
                </div>
              )}
            </div>
          ))}
        </div>
      </Wrapper>
    );
  }
};

const Wrapper = ({ children }: { children: React.ReactNode }) => (
  <div className='card bg-base-100 shadow-xl border border-gray-700  min-h-10'>
    {children}
  </div>
);

const SportIcon = ({ calories, onClick, sportType }: IconProps) => {
  let icon;

  switch (sportType) {
    case SportType.STRun:
      icon = <i className='fa-solid fa-person-running'></i>;
      break;
    case SportType.STRide:
      icon = <i className='fa-solid fa-bicycle'></i>;
      break;
    case SportType.STXCSki:
      icon = <i className='fa-solid fa-person-skiing-nordic'></i>;
      break;
    default:
      icon = <i className='fa-solid fa-dog'></i>;
  }

  return (
    <div
      className={classNames("badge badge-outline", {
        "hover:-translate-y-0.5 badge-secondary hover:scale-125 transform cursor-pointer":
          calories === 0,
        "badge-primary": calories > 0,
      })}
      onClick={onClick}
    >
      {icon}
    </div>
  );
};

type IconProps = {
  activityId: number;
  calories: number;
  sportType: SportType;
  onClick: () => void;
};

export default ActivityDayCard;

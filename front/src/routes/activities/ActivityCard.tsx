import dayjs from "dayjs";
import { ActivityResponse, SportType } from "@/types/activity";

type Props = {
  a: ActivityResponse | null;
};

const ActivityCard = ({ a }: Props) => {
  return (
    <div className='card bg-base-100 shadow-xl border border-gray-700 cursor-pointer min-h-10'>
      {a && (
        <div className='card-body'>
          <div className='flex'>
            <p>{dayjs(a.date).format("DD.MM")}</p>
            <div className='badge badge-outline badge-primary'>
              {sportIcon(a.sportType)}
            </div>
          </div>
          <h2 className='card-title'>{a.name}</h2>
          {a.distance > 0 && (
            <div className='flex justify-between'>
              <>
                <span>{a.paceString}</span>
                <span>{(a.distance / 1000).toFixed(2)} km</span>
              </>
            </div>
          )}
          <span>{a.calories} kcal</span>
        </div>
      )}
    </div>
  );
};

const sportIcon = (sportType: SportType) => {
  switch (sportType) {
    case SportType.STRun:
      return <i className='fa-solid fa-person-running'></i>;
    case SportType.STRide:
      return <i className='fa-solid fa-bicycle'></i>;
    case SportType.STXCSki:
      return <i className='fa-solid fa-person-skiing-nordic'></i>;
    default:
      return <i className='fa-solid fa-dog'></i>;
  }
};

export default ActivityCard;

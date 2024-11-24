import dayjs from "dayjs";

export function getDay(date: dayjs.Dayjs): number {
  const day = date.day();
  if (day === 0) {
    return 7;
  }
  return day;
}

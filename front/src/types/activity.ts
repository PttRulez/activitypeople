export type ActivityResponse = {
  calories: number;
  description: string;
  distance: number;
  date: string;
  elevate: number;
  heartrate: number;
  id: number;
  name: string;
  pace: number;
  paceString: string;
  source: ActivitySource;
  sourceId: number;
  sportType: SportType;
  totalTime: number;
};

export enum ActivitySource {
  Strava = "strava",
}

export enum SportType {
  STOther = "Other",
  STRide = "Ride",
  STRollerSkis = "RollerSkis",
  STRun = "Run",
  STStrength = "Strength",
  STXCSki = "XCSki",
}

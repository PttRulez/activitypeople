import axios from "axios";

const BASE_URL = process.env.REACT_APP_API_BASE_URL;

export default axios.create({
  baseURL: BASE_URL,
});

export const axiosPrivate = axios.create({
  baseURL: BASE_URL,
  headers: { "Content-Type": "application/json" },
  withCredentials: true,
});

// class ActivityPeopleAPi {
//   private static instance: ActivityPeopleAPi;
//   private readonly axios: AxiosInstance;

//   constructor() {
//     const activityAxiosInstance: AxiosInstance = axios.create({
//       baseURL: process.env.REACT_APP_API_BASE_URL,
//       withCredentials: true,
//     });

//     activityAxiosInstance.interceptors.request.use((req) => {
//       return req;
//     });

//     this.axios = activityAxiosInstance;
//   }

//   public static get(): ActivityPeopleAPi {
//     if (!ActivityPeopleAPi.instance) {
//       ActivityPeopleAPi.instance = new ActivityPeopleAPi();
//     }
//     return ActivityPeopleAPi.instance;
//   }

//   public login(
//     data: LoginData
//   ): Promise<AxiosResponse<{ accessToken: string }, any>> {
//     return this.axios.post("/login", data);
//   }

//   public setToken(token: string): void {
//     this.axios.interceptors.request.use((config) => {
//       config.headers["Authorization"] = `Bearer ${token}`;
//       return config;
//     });
//   }
// }

// const activityApi = ActivityPeopleAPi.get();

// export default activityApi;

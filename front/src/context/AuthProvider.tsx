import { createContext, useEffect, useState } from "react";
import { userKey } from "..";
import { Role } from "src/types/enums";

export type UserInfo = {
  accessToken?: string;
  user: {
    bmr: number;
    caloriesPer100Steps: number;
    email: string;
    name: string;
    role: Role;
    stravaLinked: boolean;
  };
};

export type AuthInfo = {
  auth: UserInfo;
  setAuth: React.Dispatch<React.SetStateAction<UserInfo>>;
};
const AuthContext = createContext<AuthInfo>({} as AuthInfo);

export const emptyUser: UserInfo = {
  accessToken: undefined,
  user: {
    bmr: 0,
    caloriesPer100Steps: 0,
    stravaLinked: false,
  },
} as UserInfo;

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const localItem = localStorage.getItem(userKey);
  let user = emptyUser;
  if (localItem) {
    user = JSON.parse(localItem);
  }

  const [auth, setAuth] = useState<UserInfo>(user);

  useEffect(() => {
    if (auth) {
      localStorage.setItem(userKey, JSON.stringify(auth));
    } else {
      localStorage.setItem(userKey, auth);
    }
  }, [auth]);

  return (
    <AuthContext.Provider value={{ auth, setAuth }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;

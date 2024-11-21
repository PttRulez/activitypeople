import { createContext, useState } from "react";
import { userKey } from "..";

export type UserInfo = {
  accessToken?: string;
};

export type AuthInfo = {
  auth: UserInfo;
  setAuth: React.Dispatch<React.SetStateAction<UserInfo>>;
};
const AuthContext = createContext<AuthInfo>({} as AuthInfo);

export const AuthProvider = ({ children }: { children: JSX.Element }) => {
  let user: UserInfo = {} as UserInfo;
  const localItem = localStorage.getItem(userKey);
  if (localItem) {
    user = JSON.parse(localItem);
  }

  const [auth, setAuth] = useState<UserInfo>(user);

  return (
    <AuthContext.Provider value={{ auth, setAuth }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthContext;

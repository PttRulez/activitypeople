import AuthContext, { AuthInfo } from "@/context/AuthProvider";
import { useContext } from "react";

const useAuth = (): AuthInfo => {
  return useContext(AuthContext);
};

export default useAuth;

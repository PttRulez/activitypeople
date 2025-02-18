import useAuth from "src/hooks/useAuth";
import { Navigate, Outlet, useLocation } from "react-router-dom";

const AuthRequired = () => {
  const { auth } = useAuth();
  const location = useLocation();

  return auth?.accessToken ? (
    <Outlet />
  ) : (
    <Navigate to='/login' state={{ from: location }} />
  );
};

export default AuthRequired;

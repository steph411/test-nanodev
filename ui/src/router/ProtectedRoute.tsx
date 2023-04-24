import { Navigate } from "react-router-dom";
import { useMe } from "../hooks/useMe";
import AppLoader from "../components/AppLoader";
import { useStore } from "zustand";
import { UserStore } from "../store";

interface ProtectedRouteProps {
  redirectTo?: string;
  roles?: string[];
  children: any;
}

export const ProtectedRoute = ({
  redirectTo = "/login",
  roles: allowedRoles,
  children,
}: ProtectedRouteProps) => {
  const { me, isLoading, error } = useMe();
  const { isLoggedIn } = useStore(UserStore);

  if (!isLoggedIn) {
    return <Navigate to={redirectTo} />;
  }

  if (!me && !isLoading) {
    console.log({ aboutToRedirect: { me, isLoading, error } });
    return <Navigate to={redirectTo} />;
  }
  if (isLoading) {
    return <AppLoader />;
  }

  if (allowedRoles && allowedRoles?.length > 0) {
    // if (intersection(allowedRoles, me?.roles)?.length === 0) {
    //   return <Navigate to={redirectTo} />;
    // }
  }

  return children;
};

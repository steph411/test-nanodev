import { createBrowserRouter } from "react-router-dom";
import { Suspense } from "react";
import {
  LoginPage,
  ConfirmEmailPage,
  RequestsPage,
  RegisterPage,
} from "./routes";
import AppLoader from "../components/AppLoader";
import { ProtectedRoute } from "./ProtectedRoute";

export const userRoles = {
  ADMIN: "ADMIN",
  CITIZEN: "CITIZEN",
};

export const router = createBrowserRouter([
  {
    path: "/login",
    element: (
      <Suspense fallback={<AppLoader />}>
        <LoginPage />
      </Suspense>
    ),
  },
  {
    path: "/signup",
    element: (
      <Suspense fallback={<AppLoader />}>
        <RegisterPage />
      </Suspense>
    ),
  },
  {
    path: "/register",
    element: (
      <Suspense fallback={<AppLoader />}>
        <RegisterPage />
      </Suspense>
    ),
  },
  {
    path: "/confirm-email",
    element: (
      <Suspense fallback={<AppLoader />}>
        <ConfirmEmailPage />
      </Suspense>
    ),
  },
  {
    path: "/requests",
    element: (
      <Suspense fallback={<AppLoader />}>
        <ProtectedRoute roles={[userRoles.ADMIN, userRoles.CITIZEN]}>
          <RequestsPage />
        </ProtectedRoute>
      </Suspense>
    ),
  },
  {
    path: "/",
    element: (
      <Suspense fallback={<AppLoader />}>
        <ProtectedRoute roles={[userRoles.ADMIN, userRoles.CITIZEN]}>
          <RequestsPage />
        </ProtectedRoute>
      </Suspense>
    ),
  },
]);

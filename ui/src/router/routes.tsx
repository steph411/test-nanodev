import React from "react";

type ImportFn = () => Promise<{ default: React.FC<{}> }>;

export const ReactLazyPreload = (importStatement: ImportFn) => {
  const Component = React.lazy(importStatement);
  // @ts-ignore
  Component.preload = importStatement;
  return Component;
};

export const RequestsPage = ReactLazyPreload(() => import("../pages/Requests"));
export const LoginPage = ReactLazyPreload(() => import("../pages/Login"));
export const RegisterPage = ReactLazyPreload(() => import("../pages/Register"));
export const ConfirmEmailPage = ReactLazyPreload(
  () => import("../pages/ConfirmEmail")
);

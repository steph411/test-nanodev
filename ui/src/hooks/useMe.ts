import { useQuery } from "react-query";
import { getMe } from "../libs/api";

export const useMe = (options?: any) => {
  const { data, status, ...rest } = useQuery<
    { name: string; email: string },
    "me"
  >("me", getMe, {
    retry: false,
    ...options,
  });
  const loading = status === "loading";
  return { me: data, loading, status, ...rest };
};

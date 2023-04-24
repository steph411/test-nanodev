import { useQuery } from "react-query";
import { getRequests } from "../libs/api";

export const useRequests = (userId?: string | undefined) => {
  const { data, status, ...rest } = useQuery<any[], "requests">(
    "requests",
    () => getRequests(userId),
    {
      retry: false,
    }
  );
  const loading = status === "loading";
  return { data, loading, status, ...rest };
};

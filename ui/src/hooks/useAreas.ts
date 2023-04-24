import { useQuery } from "react-query";
import { getExpertiseAreas, getMe } from "../libs/api";

export const useAreas = (options?: any) => {
  const { data, status, ...rest } = useQuery("areas", getExpertiseAreas, {
    retry: false,
    ...options,
  });
  const loading = status === "loading";
  return { data, loading, status, ...rest };
};

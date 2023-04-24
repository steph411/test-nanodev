import { useMutation, useQueryClient } from "react-query";
import { createExpertiseArea } from "../libs/api";

interface MutationOptions {
  onError?: (arg: Error) => void;
  onSuccess?: (data: any) => void;
}

export const useCreateExpertiseAreas = ({
  onError,
  onSuccess,
}: MutationOptions) => {
  const mutation = useMutation(createExpertiseArea, {
    onSuccess: (data, variables) => {
      const client = useQueryClient();
      client.refetchQueries(["areas"]);
      onSuccess && onSuccess(data);
    },
    onError: (error) => {
      onError && onError(error as Error);
    },
  });

  return mutation;
};

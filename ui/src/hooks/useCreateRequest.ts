import { useMutation, useQueryClient } from "react-query";
import { createRequest } from "../libs/api";

interface MutationOptions {
  onError?: (arg: Error) => void;
  onSuccess?: (data: any) => void;
}

export const useCreateRequest = ({ onError, onSuccess }: MutationOptions) => {
  const mutation = useMutation(createRequest, {
    onSuccess: (data, variables) => {
      const client = useQueryClient();
      client.refetchQueries(["requests"]);
      onSuccess && onSuccess(data);
    },
    onError: (error) => {
      onError && onError(error as Error);
    },
  });

  return mutation;
};

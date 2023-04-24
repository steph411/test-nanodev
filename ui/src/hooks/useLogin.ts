import { useMutation } from "react-query";
import { login } from "../libs/api";

interface MutationOptions {
  onError?: (arg: Error) => void;
  onSuccess?: (data: any) => void;
}

export const useLogin = ({ onError, onSuccess }: MutationOptions) => {
  const mutation = useMutation(login, {
    onSuccess: (data, variables) => {
      onSuccess && onSuccess(data);
    },
    onError: (error) => {
      onError && onError(error as Error);
    },
  });

  return mutation;
};

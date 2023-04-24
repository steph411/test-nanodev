import { useMutation } from "react-query";
import { register } from "../libs/api";

interface MutationOptions {
  onError?: (arg: Error) => void;
  onSuccess?: (data: any) => void;
}

export const useRegister = ({ onError, onSuccess }: MutationOptions) => {
  const mutation = useMutation(register, {
    onSuccess: (data, variables) => {
      onSuccess && onSuccess(data);
    },
    onError: (error) => {
      onError && onError(error as Error);
    },
  });

  return mutation;
};

import { useMutation } from "react-query";
import { logout } from "../libs/api";

interface MutationOptions {
  onError?: (arg: Error) => void;
  onSuccess?: (data: any) => void;
}

export const useLougout = ({ onError, onSuccess }: MutationOptions) => {
  const mutation = useMutation(logout, {
    onSuccess: (data, variables) => {
      onSuccess && onSuccess(data);
    },
    onError: (error) => {
      onError && onError(error as Error);
    },
  });

  return mutation;
};

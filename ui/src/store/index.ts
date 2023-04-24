import { getMe } from "../libs/api";
import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

type User = any;

type UserStore = {
  isLoggedIn: boolean;
  token: string;
  user?: User;
};

export const UserStore = create(
  persist<UserStore>(
    (set, get) => ({
      isLoggedIn: false,
      token: "",
      setAuthData: (data: { token: string; user: any }) => {
        set({ user: data.user, token: data.token });
      },
      getUser: async () => {
        const userData = await getMe();
        set({ user: userData });
      },
    }),
    {
      name: "auth",
    }
  )
);

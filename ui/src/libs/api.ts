import { UserStore as userStore } from "../store";
import ApiCall from "./request";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:1337";
const ApiService = new ApiCall();

export const login = async (params: {
  identifier: string;
  password: string;
}) => {
  const url = `${API_URL}/auth/local`;
  const data = await ApiService.post(url, params);
  console.log({ logindata: data });
  userStore.setState({ user: data.user, token: data.jwt });
  return data;
};

export const logout = async () => {
  userStore.setState({ user: undefined, token: "" });
  return;
};

export const register = async (params: {
  username: string;
  email: string;
  password: string;
}) => {
  const url = `${API_URL}/auth/local/register`;
  const data = await ApiService.post(url, params);
  console.log({ logindata: data });
  userStore.setState({ user: data.user, token: data.jwt });
  return data;
};

export const confirmEmail = async () => {};

export const getMe = async (): Promise<{ name: string; email: string }> => {
  const url = `${API_URL}/users/me`;
  const data = await ApiService.get(url);

  return { name: "", email: "", ...data };
};

export const getRequests = async (userId: string | undefined = undefined) => {
  let url = `${API_URL}/api/requests`;
  if (userId) url += `?userId=${userId}`;
  const data = await ApiService.get(url);
  return data;
};

export const createRequest = async (params: {
  name: string;
  content: string;
  areaId: string;
}) => {
  let url = `${API_URL}/api/requests`;
  const data = await ApiService.post(url, params);
  return data;
};

export type Status = "IN_PROGRESS" | "REJECTED" | "PROCESSED";
export const updateRequest = async (params: { status: Status }) => {
  let url = `${API_URL}/api/requests`;
  const data = await ApiService.put(url, params);
  return data;
};

export const getExpertiseAreas = async () => {
  let url = `${API_URL}/api/expertise-areas`;
  const data = await ApiService.get(url);
  return data;
};

export const createExpertiseArea = async (params: { name: string }) => {
  let url = `${API_URL}/api/expertise-area`;
  const data = await ApiService.post(url, params);
  return data;
};

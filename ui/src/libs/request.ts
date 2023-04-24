import { UserStore } from "../store";

export default class ApiCall {
  private token: string;
  private headers: Headers;

  constructor() {
    // this.url = url;
    this.token = UserStore.getState().token;
    this.headers = new Headers();
    this.headers.set("Authorization", `Bearer ${this.token}`);
  }

  async get(url: string) {
    const response = await fetch(url, {
      headers: this.headers,
    });
    const data = await response.json();
    return data;
  }

  async post(url: string, body: any) {
    const response = await fetch(url, {
      headers: this.headers,
      method: "POST",
      body: JSON.stringify(body),
    });
    const data = await response.json();
    return data;
  }

  async put(url: string, body: any) {
    const response = await fetch(url, {
      headers: this.headers,
      method: "PUT",
      body: JSON.stringify(body),
    });
    const data = await response.json();
    return data;
  }

  async delete(url: string, body: any) {
    const response = await fetch(url, {
      headers: this.headers,
      method: "DELETE",
      body: JSON.stringify(body),
    });
    const data = await response.json();
    return data;
  }
}

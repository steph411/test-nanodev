import * as axios from "axios";
export const AUTHORIZATION = "Authorization";
export interface WsConfig {
  url: string;
  prefix?: boolean;
}

export class WS {
  private instance: axios.AxiosInstance;
  private prefix?: string;
  public config: WsConfig;

  constructor(config: WsConfig) {
    this.config = config;
    this.prefix = config.prefix == undefined ? "" : `${config.prefix} `;
    this.instance = axios.default.create({ baseURL: config.url });
  }

  public request(config: axios.AxiosRequestConfig): axios.AxiosPromise {
    if (config.headers == undefined) config.headers = {};

    return this.instance
      .request(config)
      .then((resp: axios.AxiosResponse) => {
        return resp;
      })
      .catch((err: any) => {
        throw err;
      });
  }

  public get(
    url: string,
    config: axios.AxiosRequestConfig = {}
  ): axios.AxiosPromise {
    config.url = url;
    config.method = "GET";

    return this.request(config);
  }

  public post(
    url: string,
    config: axios.AxiosRequestConfig
  ): axios.AxiosPromise {
    config.url = url;
    config.method = "POST";

    return this.request(config);
  }

  public put(
    url: string,
    config: axios.AxiosRequestConfig
  ): axios.AxiosPromise {
    config.url = url;
    config.method = "PUT";

    return this.request(config);
  }

  public patch(
    url: string,
    config: axios.AxiosRequestConfig
  ): axios.AxiosPromise {
    config.url = url;
    config.method = "PATCH";

    return this.request(config);
  }

  public delete(
    url: string,
    config: axios.AxiosRequestConfig
  ): axios.AxiosPromise {
    config.url = url;
    config.method = "DELETE";

    return this.request(config);
  }
}

import axios from "axios";
import storage from "../storage";

const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_ENDPOINT,
  headers: {
    "Content-Type": "application/json"
  },
  responseType: "json"
});

let refreshTokenPromise;

const getRefreshToken = () =>
  apiClient
    .post(
      "/private/refresh",
      {},
      {
        headers: { Authorization: "Bearer " + storage.getRefreshToken() }
      }
    )
    .then(response => {
      storage.setAccessToken(response.data.access_token);
      storage.setRefreshToken(response.data.refresh_token);
      console.log(response);
    });

apiClient.interceptors.response.use(
  r => r,
  error => {
    if (
      error.config &&
      error.config.url !== "/signin" &&
      error.response &&
      error.response.status === 401 &&
      !error.config._retry
    ) {
      if (!refreshTokenPromise) {
        error.config._retry = true;
        refreshTokenPromise = getRefreshToken().then(response => {
          refreshTokenPromise = null;
          return response;
        });
      }

      return refreshTokenPromise.then(response => {
        error.config.headers["Authorization"] =
          "Bearer " + response.data.access_token;
        return apiClient.request(error.config);
      });
    }
    return Promise.reject(error);
  }
);

export default apiClient;

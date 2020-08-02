import axios from "axios";
import Axios from "axios";
import router from "../router";

const apiClient = axios.create({
  baseURL: process.env.VUE_APP_API_ENDPOINT,
  headers: {
    "Content-Type": "application/json",
  },
  responseType: "json"
});

// TODO: 動作確認ちゃんとしていないので後でする
apiClient.interceptors.response.use(
  res => {
    // if the api request is successful, return response as it is.
    return res;
  },
  error => {
    // if the api request failes, refresh the token.
    if (error.config && error.response && error.response.status === 401) {
      apiClient
        .post(
          "/refresh",
          {},
          {
            headers: {
              'Authorization': "Bearer " + localStorage.getItem("refresh_token"),
            }
          }
        )
        .then(res => {
          if (res.status === 200) {
            const config = error.config;
            localStorage.setItem("access_token", res.data.access_token);
            localStorage.setItem("refresh_token", res.data.refresh_token);
            config.headers["Authorization"] = res.data.refresh_token;

            // Retry the api request.
            return Axios.request(error.config);
          } else {
            router.push({
              path: "/signin"
            });
          }
        })
        .catch(error => {
          console.log(error);
          router.push({
            path: "/signin"
          });
        });
    }
    return Promise.reject(error);
  }
);
export default apiClient;

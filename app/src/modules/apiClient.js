import axios from "axios";

const apiClient = axios.create({
    baseURL: process.env.VUE_APP_API_ENDPOINT,
    headers: {
        "Content-Type": "application/json"
    },
    responseType: "json"
});

export default apiClient;
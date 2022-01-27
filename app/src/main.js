import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "sea.css/dist/sea.min.css";

const app = createApp(App);

app.use(router);
app.mount("#app");

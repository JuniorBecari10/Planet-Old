import { createPinia } from "pinia";
import { createApp } from "vue";
import "./styles.css";
import App from "./App.vue";

createApp(App).use(createPinia()).mount("#app");

// Custom Code

var anchors = document.querySelectorAll('a');

for (let i = 0; i < anchors.length; i++) {
  anchors[i].addEventListener('click', e => e.preventDefault(), false);
}

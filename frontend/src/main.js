import { createApp } from 'vue'
import App from './App.vue'

export const global = {
  data() {
    return {
      baseFolder: "",
    }
  }
};

createApp(App).mixin(global).mount('#app');

// custom code

var anchors = document.querySelectorAll('a');

for (let i = 0; i < anchors.length; i++) {
  anchors[i].addEventListener('click', e => e.preventDefault(), false);
}
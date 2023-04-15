import {createApp} from 'vue'
import App from './App.vue'

createApp(App).mount('#app')

// custom code

var anchors = document.querySelectorAll('a');

for (let i = 0; i < anchors.length; i++) {
  anchors[i].addEventListener('click', e => e.preventDefault(), false);
}
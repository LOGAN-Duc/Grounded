import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // 👈 cần import router

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';

const app = createApp(App);
app.use(router); // 👈 gắn router vào app
app.mount('#app');

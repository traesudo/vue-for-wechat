import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import Antd from 'ant-design-vue';
const app = createApp(App);
app.use(Antd).mount('#app');

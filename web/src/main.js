import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import en from 'element-plus/es/locale/lang/en'
import { createPinia } from 'pinia'
const app = createApp(App).use(createPinia())
app.use(store).use(router).use(ElementPlus, { locale: en }).mount('#app')

export default app
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import idsrvAuth from '@/lib/auth/auth';

idsrvAuth.startup().then((ok) => {
  if (ok) {
    const app = createApp(App);

    app.config.globalProperties.$oidc = idsrvAuth;

    app.use(store);
    app.use(router);

    app.mount('#app');
    console.log('Startup was ok');
  } else {
    console.log('Startup was not ok');
  }
});

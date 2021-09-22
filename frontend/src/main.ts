import { createApp,h } from 'vue';
import {
  // create naive ui
  create,
  // component
  NButton,
  NMenu,
} from 'naive-ui';
import 'vfonts/Lato.css';
import router from './router';
import store from './store';
import idsrvAuth from '@/lib/auth/auth';
import App from './App.vue';

idsrvAuth.startup().then((ok) => {
  if (ok) {
    const naive = create({
      components: [
        NButton,
        NMenu,
      ],
    });
    const app = createApp({
      render: ()=>h(App)
    });

    app.config.globalProperties.$oidc = idsrvAuth;

    app.use(store);
    app.use(router);
    app.use(naive);

    app.mount('#app');
  } else {
    console.log('Startup was not ok');
  }
});

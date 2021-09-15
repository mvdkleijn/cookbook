import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import idsrvAuth from '@/auth/auth';

import Home from '@/views/Home.vue';
import authedView from '@/views/authedView.vue';
import LogoutView from '@/views/LogoutView.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/authed',
    name: 'authedView',
    component: authedView,
    meta: {
      authName: 'main',
    },
  },
  {
    path: '/logout',
    name: 'LogoutView',
    component: LogoutView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

idsrvAuth.useRouter(router);

export default router;

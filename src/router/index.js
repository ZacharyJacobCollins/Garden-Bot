import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import HomeView from '../views/HomeView.vue'

export default new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    { path: '/home', component: HomeView },
    { path: '/', redirect: '/home' }
  ]
})

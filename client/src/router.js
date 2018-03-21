import Vue from 'vue'
import Router from 'vue-router'
import Lists from './views/Lists'
import Settings from './views/Settings'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'lists',
      component: Lists
    },
    {
      path: '/settings',
      name: 'settings',
      component: Settings
    }
  ]
})

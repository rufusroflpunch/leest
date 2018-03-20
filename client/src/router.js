import Vue from 'vue'
import Router from 'vue-router'
import Lists from './views/Lists.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'lists',
      component: Lists
    },
    {
      path: '/lists/:id',
      name: 'list',
      component: Lists
    }
  ]
})

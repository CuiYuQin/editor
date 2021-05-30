import Vue from 'vue'
import VueRouter from 'vue-router'
// import Home from '../views/Home.vue'

import Layout from "@/views/layout/layout"
import Login from "@/views/login/login"

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'login',
    component: Login
  },
  {
    path: '/editor',
    name: 'editor',
    component: Layout
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import('../views/layout/layout.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router

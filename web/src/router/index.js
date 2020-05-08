import Vue from 'vue'
import VueRouter from 'vue-router'
import Add from "../components/Add";
import List from "../components/List";

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'Add',
    component: Add
  },
  {
    path: '/list',
    name: 'List',
    component: List
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

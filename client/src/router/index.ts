import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import InsertUserView from '../views/InsertUserView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   component: () => import('../views/AboutView.vue')
    // },
    // {
    //   path: '/users',
    //   name: 'users',
    //   component: () => import('../views/UsersView.vue')
    // },
    // {
    //   path: '/users/:id',
    //   name: 'user',
    //   component: () => import('../views/UserView.vue')
    // },
    // {
    //   path: '/:pathMatch(.*)*',
    //   name: 'not-found',
    //   component: () => import('../views/NotFoundView.vue')
    // },
    {
      path: '/edit',
      name: 'edit',
      component: InsertUserView,
      props: route => ({ query: route.query.q })
    }
  ]
})

export default router

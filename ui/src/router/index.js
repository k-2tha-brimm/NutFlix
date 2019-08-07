import Vue from 'vue'
import Router from 'vue-router'
import Splash from '@/components/Splash'
import Movie from '@/components/Movie'
import LogIn from '@/components/Login';
import Browse from '@/components/Browse';

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'Splash',
      component: Splash
    },
    {
      path: `/movie/:id`,
      name: 'Movie',
      component: Movie
    },
    {
      path: '/login',
      name: 'LogIn',
      component: LogIn
    },
    {
      path: '/browse',
      name: 'Browse',
      component: Browse,
      meta: {
        requiresAuth: true
      }
    }
  ]
})

// router.beforeEach((to, from, next) => {
//   if(to.matched.some(record => record.meta.requiresAuth)) {
//     if(local)
//   }
// })

export default router;
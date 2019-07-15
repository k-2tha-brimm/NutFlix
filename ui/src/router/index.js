import Vue from 'vue'
import Router from 'vue-router'
import Splash from '@/components/Splash'
import Movie from '@/components/Movie'

Vue.use(Router)

export default new Router({
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
    }
  ]
})
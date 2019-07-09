import Vue from 'vue'
import Router from 'vue-router'
import Splash from '@/components/HelloWorld'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'splash',
      component: Splash
    }
  ]
})

import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import InputPenyakitView from '../views/InputPenyakitView.vue'
import TesDNAView from '../views/TesDNAView.vue'
import HasilPrediksiView from '../views/HasilPrediksiView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/inputpenyakit',
    name: 'inputpenyakit',
    component: InputPenyakitView
  },
  {
    path: '/tesdna',
    name: 'tesdna',
    component: TesDNAView
  },
  {
    path: '/hasilprediksi',
    name: 'hasilprediksi',
    component: HasilPrediksiView
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

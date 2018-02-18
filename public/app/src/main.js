// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Buefy from 'buefy'
import VueCookie from 'vue-cookie'
import VueResource from 'vue-resource'
import 'buefy/lib/buefy.css'

Vue.config.productionTip = false

Vue.use(Buefy)
Vue.use(VueCookie)
Vue.use(VueResource)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router: router,
  template: '<App/>',
  components: { App }
})

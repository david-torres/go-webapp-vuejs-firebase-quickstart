import Vue from 'vue'
import Firebase from 'firebase'
import VueFire from 'vuefire'
import { config } from '../config/firebase'

Vue.use(VueFire)

export default Firebase.initializeApp(config)

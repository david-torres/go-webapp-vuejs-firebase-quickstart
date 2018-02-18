<template>
  <section class="section">
    <div class="container">
      <h1 v-if="authenticated" class="title" v-html="message"></h1>
      <h1 v-else class="title">Unauthorized</h1>
      <button v-if="authenticated" class="button is-primary" @click="logout">Logout</button>
      <router-link v-else class="button is-primary" v-bind:to="'/auth'">Register/Login</router-link>
    </div>
  </section>
</template>

<script>
import firebase from '../helpers/firebase'

export default {
  name: 'Home',
  created () {
    let app = this
    let csrf = app.$cookie.get('_csrf')

    // if we authed, fetch the user
    firebase.auth().onAuthStateChanged((user) => {
      if (!user) {
        return
      }

      user.getIdToken().then((token) => {
        app.authenticated = true
        app.$http.get('/hello/world', { headers: { 'X-CSRF-Token': csrf, 'Authorization': 'Bearer ' + token } }).then(response => {
          response.json().then((body) => {
            console.log('got response from /hello/world')
            console.log(body)
            app.message = body
          })
        }, response => {
          // error callback
          console.log(response)
        })
      }).catch((error) => {
        // Handle error
        console.log(error)
      })
    })
  },
  methods: {
    logout () {
      let app = this
      firebase.auth().signOut().then(() => {
        app.authenticated = false
        app.message = ''
        console.log('logged out')
      }, (error) => {
        console.log(error)
      })
    }
  },
  data () {
    return {
      authenticated: false,
      message: ''
    }
  }
}
</script>

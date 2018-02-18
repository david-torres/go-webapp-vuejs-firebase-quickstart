<template>
  <section class="section">
    <div id="app" class="container">
      <router-view/>
    </div>
  </section>
</template>

<script>
import firebase from './helpers/firebase'

export default {
  name: 'App',
  created () {
    let app = this

    // if we authed, fetch the user
    firebase.auth().onAuthStateChanged((user) => {
      if (!user) {
        return
      }

      user.getIdToken().then((token) => {
        // listen for token events
        firebase.auth().onIdTokenChanged((user) => {
          if (user) {
            console.log('got user')
            console.log(user)
            user.getIdToken().then((token) => {
              console.log('got token')
              console.log(token)
            }).catch((error) => {
              // Handle error
              console.log(error)
            })
          }
        })

        // go home if we've authed but we're still on the auth page
        if (location.hash === '#/auth') {
          app.$router.push('/')
        }
      }).catch((error) => {
        // Handle error
        console.log(error)
      })
    })
  }
}
</script>

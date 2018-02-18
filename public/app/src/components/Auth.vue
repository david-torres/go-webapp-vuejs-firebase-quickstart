<template>
  <section class="section">
    <div class="container">
      <div id="firebaseui-auth-container"></div>
    </div>
  </section>
</template>

<script>
import Firebase from 'firebase'
import firebaseui from 'firebaseui'
import 'firebaseui/dist/firebaseui.css'

export default {
  name: 'Auth',
  mounted () {
    // start firebase auth ui
    let ui = firebaseui.auth.AuthUI.getInstance()
    if (!ui) {
      ui = new firebaseui.auth.AuthUI(Firebase.auth())
    }

    let uiConfig = {
      signInSuccessUrl: '/#/',
      signInOptions: [
        Firebase.auth.EmailAuthProvider.PROVIDER_ID
      ],
      credentialHelper: firebaseui.auth.CredentialHelper.NONE
    }

    ui.start('#firebaseui-auth-container', uiConfig)
  }
}
</script>

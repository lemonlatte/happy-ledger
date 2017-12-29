<style>
  .login {
    text-align: center;
  }

  .fb-login {
    display: inline-block;
    background-color: #31589c;
    color: white;
    text-align: center;
    padding: 8px;
    font-size: 16px;
    width: 200px;
    height: 40px;
    cursor: pointer;
  }
</style>

<template>
  <div class="login">
    <p>Sign in using:</p>
    <p>
      Email:
      <input type="text" v-model="email"> Password:
      <input type="password" v-model="password">
      <button @click="login">Login</button>
    </p>
    <hr>
    <div class="fb-login" @click="fbLogin">
      <i class="fa fa-facebook-official"></i> Facebook</div>
  </div>
</template>

<script>
  const firebase = require("firebase/app");
  require("firebase/auth");
  const provider = new firebase.auth.FacebookAuthProvider();
  export default {
    methods: {
      async fbLogin() {
        try {
          let u = await firebase.auth().signInWithRedirect(provider);
          //   let u = await firebase.auth().signInWithPopup(provider);
          let result = await firebase.auth().getRedirectResult()
          if (result.credential) {
            // This gives you a Facebook Access Token. You can use it to access the Facebook API.
            var token = result.credential.accessToken;
          }
          // The signed-in user info.
          var user = result.user;
          console.log(user)
          let redirect = this.$route.query.redirect || "/"
          this.$router.push({
            path: redirect
          })
        } catch (err) {
          console.log(err)
          // if (err.code === "auth/account-exists-with-different-credential") {
          //   await firebase.auth().signInWithCredential(err.credential)
          // }
        }
      },

      async login() {
        try {
          let u = await firebase.auth().signInWithEmailAndPassword(this.email, this.password);
          let redirect = this.$route.query.redirect || "/"
          this.$router.push({
            path: redirect
          })
        } catch (err) {
          if (err.code === 'auth/user-not-found') {
            try {
              await firebase.auth().createUserWithEmailAndPassword(this.email, this.password);
              this.login()
            } catch (err) {
              console.log(err)
            }
          } else {
            console.log(err)
          }
        }
      }
    },

    data() {
      return {
        email: "",
        password: "",
      }
    }
  }
</script>

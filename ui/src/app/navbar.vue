<template>
  <div>
    <router-link to="/history" tag="a">History</router-link>
    <router-link to="/summary" tag="a">Summary</router-link>
    <a href="#" @click="logout">Logout</a>
  </div>
</template>

<script>
  import firebase from 'firebase'
  export default {
    methods: {
      async logout() {
        await firebase.auth().signOut()
      }
    },
    data() {
      return {}
    }
  }
</script>

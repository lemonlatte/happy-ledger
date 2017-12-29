<template>
  <div>
    <router-link to="/history" tag="a">History</router-link>
    <router-link to="/summary" tag="a">Summary</router-link>
    <template v-if="!!user">
      <div style="float: right;">{{ user.email }} <a href="#" v-if="!!user" @click="logout">Logout</a> </div>
    </template>
  </div>
</template>

<script>
  import firebase from 'firebase'
  export default {
    props: {
      uid: String,
      user: Object
    },

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

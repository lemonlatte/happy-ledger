<template>
  <div>
    <div style="float: left;">Happy Ledger</div>
    <template v-if="!!user">
      <div style="float: right;">{{ user.email }}
        <a href="#" v-if="!!user" @click="logout">Logout</a>
      </div>
    </template>
    <div style="clear: both"></div>
    <hr>
  </div>
</template>

<script>
  import firebase from 'firebase'
  export default {
    props: {
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

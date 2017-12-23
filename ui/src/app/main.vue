<style lang="scss">

</style>

<template>
  <div id="main">
    <NavBar ></NavBar>
    <router-view v-on:error="this.handleError" :uid="uid" :basePrices="basePrices"></router-view>
  </div>
</template>

<script>
  import firebase from "firebase"
  import axios from "axios"

  const NavBar = require('./navbar.vue')

  export default {
    components: {
      NavBar: NavBar,
    },

    created() {
      this.getBaseCoinPrice()
      this.priceTask = setInterval(this.getBaseCoinPrice, 5000)
      firebase.auth().onAuthStateChanged((user) => {
        if (user) {
          this.uid = user.uid
        } else {
          this.uid = null
        }
      });
    },

    methods: {
      getBaseCoinPrice() {
        axios.get("https://happy-ledger.lemonlatte.tw/v1/price")
          .then((response) => {
            this.basePrices = response.data
            this.basePrices.usd = 1
          })
          .catch((error) => {
            console.log(error)
          })
      },

      handleError(message) {
        this.errorMsg = message
      },
    },

    data() {
      return {
        uid: null,
        basePrices: {},
        nodeInfo: {}
      }
    }
  }
</script>

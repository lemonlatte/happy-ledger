<style lang="scss">

</style>

<template>
  <div id="main">
    <NavBar ></NavBar>
    <router-view v-on:error="this.handleError" :uid="uid" :basePrices="basePrices"></router-view>
  </div>
</template>

<script>
  import axios from "axios"

  const NavBar = require('./navbar.vue')

  export default {
    components: {
      NavBar: NavBar,
    },

    created() {
      this.getBaseCoinPrice()
      this.priceTask = setInterval(this.getBaseCoinPrice, 5000)
    },

    methods: {
      getBaseCoinPrice() {
        axios.get("https://happy-ledger.lemonlatte.tw/v1/price")
          .then((response) => {
            this.basePrices = response.data
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
        basePrices: {},
        nodeInfo: {}
      }
    }
  }
</script>

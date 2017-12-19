<style lang="scss">
</style>

<template>
  <div>
    <router-view v-on:error="this.handleError"></router-view>
  </div>
</template>

<script>
  import axios from "axios"

  export default {
    methods: {
      handleError(message) {
        this.errorMsg = message
      },
    },

    data() {
      return {
        nodeInfo: {}
      }
    }
  }
</script>

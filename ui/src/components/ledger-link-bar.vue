<style lang="scss" scoped>
  .active {
    border: solid 1px red;
  }
</style>

<template>
  <div>
    <div>
      <router-link to="/" tag="a">Home</router-link> > Ledger > {{ ledgerName }}
    </div>
    <router-link active-class="active" :to="{name: 'history', params: { ledgerId: ledgerId }}" tag="a">History</router-link>
    <router-link active-class="active" :to="{name: 'summary', params: { ledgerId: ledgerId }}" tag="a">Summary</router-link>
  </div>
</template>

<script>
  export default {
    props:{
      ledgerId: String,
      ledgerName: String
    }
  }
</script>

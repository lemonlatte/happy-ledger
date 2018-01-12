<style lang="scss" scoped>
.active {
  border: solid 1px red;
}
</style>

<template>
  <div>
    <router-link active-class="active" :to="{name: 'history', params: { ledgerName: ledgerName }}" tag="a">History</router-link>
    <router-link active-class="active" :to="{name: 'summary', params: { ledgerName: ledgerName }}" tag="a">Summary</router-link>
  </div>
</template>

<script>
export default {
  props: {
    ledgerName: String
  }
}
</script>


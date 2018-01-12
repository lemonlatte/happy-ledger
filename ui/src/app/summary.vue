<style lang="scss">

</style>

<template>
  <div>
    <LedgerLinkBar :ledgerName="ledgerName"></LedgerLinkBar>
    <table border="1" width="100%" v-if="basePrices.btc">
      <tr>
        <td>幣別</td>
        <td>數量</td>
        <td>持有成本</td>
        <td>現今估值</td>
        <td></td>
      </tr>
      <tr v-for="(value, key) in balance" :key="key">
        <td>{{ key }}</td>
        <td>{{ value.amount }}</td>
        <td @click="() => {toggleDetailPrice(key)}">
          <p>{{ value.usdCost }}USD</p>
          <p v-if="collapseFlag[key]">{{ value.btcCost }}BTC</p>
          <p v-if="collapseFlag[key]">{{ value.ethCost }}ETH</p>
        </td>
        <td>
          <p>{{ value.usdCost }}USD</p>
          <p v-if="collapseFlag[key]">{{ value.btcCost * basePrices.btc }}BTC</p>
          <p v-if="collapseFlag[key]">{{ value.ethCost * basePrices.eth }}ETH</p>
        </td>
        <td>
          <button>Open Chart</button>
        </td>
      </tr>
    </table>
  </div>
</template>

<script>
  import Vue from 'vue'
  import firebase from 'firebase'
  require("firebase/firestore")
  import axios from 'axios'

  const db = firebase.firestore();

  import LedgerLinkBar from '../components/ledger-link-bar.vue'

  export default {

    components: {
      LedgerLinkBar: LedgerLinkBar
    },

    props: {
      ledgerName: String,
      uid: String,
      basePrices: Object
    },

    computed: {
      balance() {
        let balance = {};
        this.records.forEach((v, i) => {
          let coinBalance = balance[v.toCoin] || {};
          let _amount = coinBalance.amount || 0;
          let _usdCost = coinBalance.usdCost || 0;
          let _btcCost = coinBalance.btcCost || 0;
          let _ethCost = coinBalance.ethCost || 0;
          coinBalance.amount = _amount + v.buySell * v.toAmount;
          coinBalance.usdCost = (_usdCost * _amount + v.buySell * v.baseCoinCost.usd) / coinBalance.amount
          coinBalance.btcCost = (_btcCost * _amount + v.buySell * v.baseCoinCost.btc) / coinBalance.amount
          coinBalance.ethCost = (_ethCost * _amount + v.buySell * v.baseCoinCost.eth) / coinBalance.amount

          if (v.baseCoin !== 'usd') {
            let baseCoinBalance = balance[v.baseCoin] || {};
            let _baseCoinAmount = baseCoinBalance.amount || 0;
            let _baseCoinUsdCost = baseCoinBalance.usdCost || 0;
            let _baseCoinBtcCost = baseCoinBalance.btcCost || 0;
            let _baseCoinEthCost = baseCoinBalance.ethCost || 0;
            baseCoinBalance.amount = _baseCoinAmount - v.buySell * v.baseCoinAmount
            baseCoinBalance.usdCost = (_baseCoinUsdCost * _baseCoinAmount - v.buySell * v.baseCoinCost.usd) /
              baseCoinBalance.amount
            baseCoinBalance.btcCost = (_baseCoinBtcCost * _baseCoinAmount - v.buySell * v.baseCoinCost.btc) /
              baseCoinBalance.amount
            baseCoinBalance.ethCost = (_baseCoinEthCost * _baseCoinAmount - v.buySell * v.baseCoinCost.eth) /
              baseCoinBalance.amount
            balance[v.baseCoin] = baseCoinBalance;
          }

          balance[v.toCoin] = coinBalance;
        })
        return balance
      }
    },

    methods: {
      toggleDetailPrice(key) {
        Vue.set(this.collapseFlag, key, !this.collapseFlag[key])
      },
      currentBasePrice() {
        return this.basePrices || {}
      },
      loadData(uid) {
        db.collection("users").doc(uid).collection('ledgers').doc(this.ledgerName)
          .collection("history").orderBy("timestamp", "desc")
          .onSnapshot(snapshot => {
            let records = snapshot.docs.map(doc => {
              let d = doc.data()
              d.id = doc.id
              return d
            })
            this.records = records
          });
      }
    },
    watch: {
      uid(newVal, oldVal) {
        if (newVal) {
          this.loadData(newVal)
        }
      }

    },
    mounted() {
      if (this.uid) {
        this.loadData(this.uid)
      }
    },
    data() {
      return {
        lastBalance: {},
        records: [],
        collapseFlag: {}
      }
    }
  }
</script>

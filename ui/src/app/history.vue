<style lang="scss">

</style>

<template>
  <div>
    <table border="1">
      <tr>
        <td>Time</td>
        <td>Trade Pair</td>
        <td>Price</td>
        <td>Amount</td>
        <td>Trade Price(Base Coin)</td>
        <td>Trade Value(USD)</td>
        <td></td>
      </tr>
      <tr v-for="r in records" v-bind:key="r.id">
        <td>{{ r.timestamp | moment }}</td>
        <td>{{ r.toCoin }}/{{ r.baseCoin }}</td>
        <td>{{ r.baseCoinAmount / r.toAmount}} {{ r.baseCoin }}</td>
        <td>{{ r.toAmount }}</td>
        <td>{{ r.baseCoinAmount }} {{ r.baseCoin }}</td>
        <td>{{ r.baseCoinCost.usd }}</td>
        <td>
          <button @click="() => {remove(r.id)}">Edit</button>
          <button @click="() => {remove(r.id)}">Delete</button>
        </td>
      </tr>
    </table>

    <form>
      <p>Base Coin:
        <select v-model="baseCoin">
          <option value="usd">USD</option>
          <option value="btc">BTC</option>
          <option value="eth">ETH</option>
        </select>
        <input type="number" v-model.number="baseCoinAmount">
      </p>
      <p v-if="baseCoin!='usd'">Current Price:
        <input type="number" :placeholder="currentBaseCoinPrice" v-model.number="baseCoinPrice">
      </p>
      <p>To:
        <select v-model="toCoin">
          <option value="">-</option>
          <option value="btc" v-if="baseCoin!=='btc'">BTC</option>
          <option value="eth" v-if="baseCoin!=='eth'">ETH</option>
          <option value="usd" v-if="baseCoin!=='usd'">USD</option>
          <option value="bcc">BCC</option>
          <option value="ltc">LTC</option>
          <option value="neo">NEO</option>
          <option value="omg">OMG</option>
          <option value="iota">IOTA</option>
          <option value="xrp">XRP</option>
        </select>
        <input type="number" v-model.number="toAmount">
      </p>
    </form>
    <button @click="buy">買</button>
    <button @click="sell">賣</button>
  </div>
</template>

<script>
  import firebase from 'firebase'
  require("firebase/firestore")
  import moment from 'moment'
  import axios from 'axios'

  const db = firebase.firestore();

  export default {

    props: {
      uid: String,
      basePrices: Object
    },

    destroyed() {
      clearInterval(this.priceTask)
    },

    computed: {
      currentBaseCoinPrice() {
        if (this.baseCoin === 'usd') {
          return 1
        }
        return this.baseCoinPrice || this.basePrices[this.baseCoin]
      }
    },

    methods: {
      buy() {
        this.newRecord(1)
      },

      sell() {
        this.newRecord(-1)
      },

      async newRecord(buySell) {
        if (!this.baseCoinAmount || !this.toAmount) {
          alert("invalid amount")
          return
        }

        // calculate the base coin costs
        let baseCoinCost = {}
        Array("usd", "eth", "btc").forEach((coin) => {
          let cost;
          if (this.baseCoin === coin) {
            cost = this.baseCoinAmount
          } else {
            if (coin === 'usd') {
              cost = this.baseCoinAmount * this.basePrices[this.baseCoin]
            } else {
              cost = this.baseCoinAmount * this.basePrices[this.baseCoin] / this.basePrices[coin]
            }
          }
          baseCoinCost[coin] = cost
        })

        let r = {
          buySell: buySell,
          timestamp: Date.now(),
          baseCoin: this.baseCoin,
          baseCoinAmount: this.baseCoinAmount,
          baseCoinCost: baseCoinCost,
          baseCoinPrice: this.currentBaseCoinPrice,
          unitPrice: this.baseCoinAmount / this.toAmount,

          toCoin: this.toCoin,
          toAmount: this.toAmount,
        }

        try {
          let ref = await db.collection("ledger").doc(this.uid).collection("history").add(r)
        } catch (error) {
          console.error("Error adding document: ", error);
        }

        // clean manual base coin price
        this.baseCoinPrice = null;
      },

      async remove(id) {
        try {
          await db.collection("ledger").doc(this.uid).collection("history").doc(id).delete()
        } catch (error) {
          console.log(error)
        }
      },

      loadData(uid){
        db.collection("ledger").doc(uid).collection("history").orderBy("timestamp", "desc")
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

    filters: {
      moment: function (timestamp) {
        return moment(timestamp).format('MMMM Do YYYY, hh:mm:ss');
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
        priceTask: null,
        baseCoin: "usd",
        toCoin: "",
        baseCoinAmount: 0,
        baseCoinPrice: null,
        toAmount: 0,

        records: []
      }
    }
  }
</script>

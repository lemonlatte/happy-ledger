<style>

</style>

<template>
  <div>

    <table border="1">
      <tr>
        <td>交易時間</td>
        <td colspan="2">購買數量</td>
        <td colspan="3">交易貨幣</td>
        <td colspan="2">交易當下現值(USD)</td>
        <td rowspan="2">刪除</td>
      </tr>
      <tr>
        <td></td>
        <td>幣別</td>
        <td>數量</td>
        <td>幣別</td>
        <td>單價</td>
        <td>數量</td>
        <td>單價</td>
        <td>總值</td>
      </tr>
      <!-- <template ></template> -->
      <tr v-for="(r, index) in records" v-bind:key="index">
        <td>{{ r.timestamp }}</td>
        <td>{{ r.toCoin }}</td>
        <td>{{ r.toAmount }}</td>
        <td>{{ r.baseCoin }}</td>
        <td>{{ r.fromAmount }}</td>
        <td>{{ r.baseCoinAmount }}</td>
        <td>{{ r.baseCoinPrice * r.fromAmount }}</td>
        <td>{{ r.baseCoinPrice * r.baseCoinAmount }}</td>
        <td @click="() => {remove(index)}">X</td>
      </tr>
    </table>

    <form>
      <p>Base Coin:
        <select v-model="baseCoin">
          <option value="usd">USD</option>
          <option value="btc">BTC</option>
          <option value="eth">ETH</option>
        </select>
        <input type="text" v-model="baseCoinAmount">
      </p>
      <p v-if="baseCoin!='usd'">Current Price:
        <input type="number" :placeholder="currentBaseCoinPrice" v-model="baseCoinPrice">
      </p>
      <p>To:
        <select v-model="toCoin">
          <option value="btc">BTC</option>
          <option value="eth">ETH</option>
          <option value="usd">USD</option>
          <option value="iota">IOTA</option>
          <option value="xrp">XRP</option>
        </select>
        <input type="text" v-model="toAmount">
      </p>
    </form>
    <button @click="newRecord">新增</button>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {

    mounted() {
      this.getBaseCoinPrice()
      this.priceTask = setInterval(this.getBaseCoinPrice, 5000)
    },

    destroyed() {
      clearInterval(this.priceTask)
    },

    computed: {
      currentBaseCoinPrice(){
        return this.baseCoinPrice || this.basePrices[this.baseCoin]
      }
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

      newRecord() {
        let r = {
          id: 0,
          timestamp: new Date(),
          baseCoin: this.baseCoin,
          baseCoinAmount: this.baseCoinAmount,
          baseCoinPrice: this.currentBaseCoinPrice,
          fromAmount: this.baseCoinAmount / this.toAmount,
          toCoin: this.toCoin,
          toAmount: this.toAmount,
        }
        this.baseCoinPrice = null;
        console.log(r)
        this.records.push(r)
        console.log("買")
      },

      remove(index) {
        console.log(index)
        this.records.splice(index, 1)
      }
    },

    data() {
      return {
        basePrices: {
          "btc": 0,
          "eth": 0
        },

        priceTask: null,
        baseCoin: "usd",
        toCoin: "btc",
        baseCoinAmount: 0,
        baseCoinPrice: null,
        toAmount: 0,

        records: [{
          id: 0,
          timestamp: new Date(),
          baseCoin: "usd",
          baseCoinAmount: 200,
          baseCoinPrice: 1,
          fromAmount: 2000,
          toCoin: "btc",
          toAmount: 0.1,
        }, {
          id: 0,
          timestamp: new Date(),
          baseCoin: "usd",
          baseCoinAmount: 700,
          baseCoinPrice: 1,
          fromAmount: 700,
          toCoin: "eth",
          toAmount: 1,
        }, {
          id: 0,
          timestamp: new Date(),
          baseCoin: "eth",
          baseCoinAmount: 0.14,
          baseCoinPrice: 700,
          fromAmount: 0.007,
          toCoin: "iota",
          toAmount: 20,
        }, {
          id: 0,
          timestamp: new Date(),
          baseCoin: "eth",
          baseCoinAmount: 0.08,
          baseCoinPrice: 700,
          fromAmount: 0.008,
          toCoin: "iota",
          toAmount: -10,
        }, {
          id: 0,
          timestamp: new Date(),
          baseCoin: "usd",
          baseCoinAmount: 56,
          baseCoinPrice: 1,
          fromAmount: 5.6,
          toCoin: "iota",
          toAmount: 10,
        }, {
          id: 0,
          timestamp: new Date(),
          baseCoin: "eth",
          baseCoinAmount: 0.05,
          baseCoinPrice: 700,
          fromAmount: 0.0005,
          toCoin: "xrp",
          toAmount: 100,
        }]
      }
    }
  }
</script>

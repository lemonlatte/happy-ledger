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
        <td>{{ r.toValue }}</td>
        <td>{{ r.fromCoin }}</td>
        <td>{{ r.fromCost}}</td>
        <td>{{ r.fromValue }}</td>
        <td>{{ usdRate(r.fromCoin) * r.fromCost }}</td>
        <td>{{ usdRate(r.fromCoin) * r.fromValue }}</td>
        <td @click="() => {remove(index)}">X</td>
      </tr>
    </table>

    <form>
      <p>From:
        <select v-model="fromCoin">
          <option value="usd">USD</option>
          <option value="btc">BTC</option>
          <option value="eth">ETH</option>
          <option value="iota">IOTA</option>
          <option value="xrp">XRP</option>
        </select>
        <input type="text" v-model="fromValue">
      </p>
      <p>To:
        <select v-model="toCoin">
          <option value="btc">BTC</option>
          <option value="eth">ETH</option>
          <option value="usd">USD</option>
          <option value="iota">IOTA</option>
          <option value="xrp">XRP</option>
        </select>
        <input type="text" v-model="toValue">
      </p>
    </form>
    <button @click="newRecord">新增</button>
  </div>
</template>

<script>
  export default {

    methods: {
      usdRate(coinType) {
        switch (coinType) {
          case "btc":
            return 20000
          case "eth":
            return 700
          case "usd":
            return 1
          default:
            throw new Error("unsupported coin")
        }
      },

      newRecord() {
        this.records.push({
          id: 0,
          timestamp: new Date(),
          fromCoin: this.fromCoin,
          fromValue: this.fromValue,
          fromCost: this.fromValue / this.toValue,
          toCoin: this.toCoin,
          toValue: this.toValue,
        })
        console.log("買")
      },

      remove(index) {
        console.log(index)
        this.records.splice(index, 1)
      }
    },

    data() {
      return {
        fromCoin: "usd",
        toCoin: "btc",
        fromValue: 0,
        toValue: 0,

        records: [{
          id: 0,
          timestamp: new Date(),
          fromCoin: "usd",
          fromValue: 200,
          fromCost: 2000,
          toCoin: "btc",
          toValue: 0.1,
        }, {
          id: 0,
          timestamp: new Date(),
          fromCoin: "usd",
          fromValue: 700,
          fromCost: 700,
          toCoin: "eth",
          toValue: 1,
        }, {
          id: 0,
          timestamp: new Date(),
          fromCoin: "eth",
          fromValue: 0.14,
          fromCost: 0.007,
          toCoin: "iota",
          toValue: 20,
        }, {
          id: 0,
          timestamp: new Date(),
          fromCoin: "eth",
          fromValue: 0.08,
          fromCost: 0.008,
          toCoin: "iota",
          toValue: -10,
        }, {
          id: 0,
          timestamp: new Date(),
          fromCoin: "usd",
          fromValue: 56,
          fromCost: 5.6,
          toCoin: "iota",
          toValue: 10,
        }, {
          id: 0,
          timestamp: new Date(),
          fromCoin: "eth",
          fromValue: 0.05,
          fromCost: 0.0005,
          toCoin: "xrp",
          toValue: 100,
        }]
      }
    }
  }
</script>

<style lang="scss" scoped>
  .mask {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: rgba(15, 15, 15, 0.7);
    text-align: center;
  }

  .dialog {
    background-color: white;
    border: solid 1px black;
    display: inline-block;
    width: 600px;
    height: 300px;

    .close {
      float: right;
    }
  }
</style>

<template>
  <div v-if="action != ''">
    <div class="mask">
      <div class="dialog">
        <button class="close" @click="close">Close</button>

        <button @click="buy">Buy</button>
        <button @click="sell">Sell</button>
        <form>
          <p>
            Date
            <input type="date" v-model="date">
          </p>
          <p>
            Time
            <input type="time" v-model="time">
          </p>
          <p>Base Coin:
            <select v-model="baseCoin">
              <option value="usd">USD</option>
              <option value="usdt">USDT</option>
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
            <p>
              Trading Value
              <input type="number" v-model.number="baseCoinCost['eth']">
            </p>
            <p>
              Trading Price
              <input type="number" v-model.number="baseCoinCost['usd']">
            </p>
          </p>
        </form>
        <button @click="ok">確定</button>
        <button @click="close">取消</button>
      </div>
    </div>
  </div>
</template>
<script>
  import moment from 'moment'
  import firebase from 'firebase'
  require("firebase/firestore")
  const db = firebase.firestore();

  export default {

    props: {
      action: "",
      inputData: {},
      basePrices: {}
    },

    computed: {
      currentBaseCoinPrice() {
        if (this.baseCoin === 'usd') {
          return 1
        }

        // currentBaseCoinPrice will search the following in order and return the first non-empty one:
        // 1. baseCoinPrice
        // 2. historyPrices
        // 3. basePrices
        return this.baseCoinPrice || this.historyPrices[this.baseCoin] || this.basePrices[this.baseCoin]
      },

      baseCoinCost() {
        let baseCoinCost = {
          usd: 0,
          usdt: 0,
          eth: 0,
          btc: 0,
        }

        if (!this.baseCoinAmount || !this.toAmount) {
          // alert("invalid amount")
          return baseCoinCost
        }

        let basePrices = this.basePrices
        if (this.historyPrices.eth) {
          basePrices = this.historyPrices
        }

        // calculate the base coin costs
        Array("usd", "eth", "btc", "usdt").forEach((coin) => {
          let cost;
          if (this.baseCoin === coin) {
            cost = this.baseCoinAmount
          } else {
            if (coin === 'usd') {
              cost = this.baseCoinAmount * basePrices[this.baseCoin]
            } else {
              cost = this.baseCoinAmount * basePrices[this.baseCoin] / basePrices[coin]
            }
          }
          baseCoinCost[coin] = cost
        })
        return baseCoinCost
      }
    },

    methods: {
      async getHistoryPrice(d, t) {
        let datetime = moment(d + " " + t)
        // if datetime is invalid or within 5 mins, return {}
        if (!datetime.isValid() || moment().diff(datetime) < 300000) {
          return {}
        }

        let lastHistoryDate = moment(Math.floor(+datetime.toDate() / 300000) * 300000).utc()

        try {
          let ref = await db.collection("crypto-price")
            .doc(lastHistoryDate.format("YYYYMMDD"))
            .collection(lastHistoryDate.format("HH"))
            .doc(lastHistoryDate.format("mm"))
            .get()
          let data = ref.data()
          return {
            btc: data.BTC.USD,
            eth: data.ETH.USD,
            usdt: data.USDT.USD,
            usd: 1
          }
        } catch (error) {
          console.log(error)
          console.log("can not get price for time:",
            lastHistoryDate.format("YYYYMMDD"), lastHistoryDate.format("HH"), lastHistoryDate.format("mm"))
          return {}
        }
      },

      ok() {
        let r = {
          buySell: this.buySell,
          timestamp: Date.now(),
          baseCoin: this.baseCoin,
          baseCoinAmount: this.baseCoinAmount,
          baseCoinCost: this.baseCoinCost,
          baseCoinPrice: this.currentBaseCoinPrice,
          unitPrice: this.baseCoinAmount / this.toAmount,

          toCoin: this.toCoin,
          toAmount: this.toAmount,
        }
        this.$emit(this.action, this.inputData.id, r)
        this.baseCoinPrice = null;
      },
      close() {
        this.$emit('close')
      },

      buy() {
        this.buySell = 1
      },

      sell() {
        this.buySell = -1
      }
    },

    watch: {
      inputData(newValue, oldVaule) {
        let now = moment()
        this.buySell = this.inputData.buySell || 1
        this.baseCoin = this.inputData.baseCoin || "usd"
        this.baseCoinAmount = this.inputData.baseCoinAmount || 0
        this.baseCoinPrice = this.inputData.baseCoinPrice || null
        this.toCoin = this.inputData.toCoin || ""
        this.toAmount = this.inputData.toAmount || 0
        this.date = now.format("YYYY-MM-DD")
        this.time = now.format("HH:mm")
      },

      async date(newValue, oldVaule) {
        this.historyPrices = await this.getHistoryPrice(newValue, this.time)
      },

      async time(newValue, oldVaule) {
        this.historyPrices = await this.getHistoryPrice(this.date, newValue)
      }
    },

    data() {
      let now = moment()
      return {
        timerTask: null,
        now: now,
        date: now.format("YYYY-MM-DD"),
        time: now.format("HH:mm"),

        buySell: 1,
        baseCoin: "usd",
        baseCoinAmount: 0,
        baseCoinPrice: null,
        historyPrices: {},
        toCoin: "",
        toAmount: 0
      }
    }
  }
</script>

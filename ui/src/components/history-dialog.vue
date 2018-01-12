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
            <input type="text" :placeholder="currentDate">
          </p>
          <p>
            Time
            <input type="text" :placeholder="currentTime">

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

  export default {

    props: {
      action: "",
      basePrices: {}
    },

    computed: {
      currentDate() {
        return moment(this.now).format('YYYY-MM-DD')
      },
      currentTime() {
        return moment(this.now).format('hh:mm:ss')
      },
      currentBaseCoinPrice() {
        if (this.baseCoin === 'usd') {
          return 1
        }
        return this.baseCoinPrice || this.basePrices[this.baseCoin]
      },

      baseCoinCost() {
        let baseCoinCost = {
          usd: 0,
          eth: 0,
          btc: 0,
        }

        if (!this.baseCoinAmount || !this.toAmount) {
          // alert("invalid amount")
          return baseCoinCost
        }

        // calculate the base coin costs
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
        return baseCoinCost
      }
    },

    methods: {
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
        this.$emit(this.action, r)
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

    mounted() {
      this.timerTask = setInterval(() => {
        this.now = Date.now()
      }, 2000)
    },

    destroyed() {
      clearInterval(this.timerTask)
    },

    data() {
      return {
        timerTask: null,
        now: Date.now(),

        buySell: 1,
        baseCoin: "usd",
        baseCoinAmount: 0,
        baseCoinPrice: null,
        toCoin: "",
        toAmount: 0
      }
    }
  }
</script>

<style lang="scss">

</style>

<template>
  <div>
    <LedgerLinkBar :ledgerId="ledgerId" :ledgerName="ledgerName"></LedgerLinkBar>
    <button @click="() => {this.action = 'new', this.inputData = {}}">新增</button>
    <table border="1" width="100%">
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
        <td>{{ r.buySell
          < 0 ? "-" : " " }}{{ r.toAmount }}</td>
            <td>{{ r.baseCoinAmount }} {{ r.baseCoin }}</td>
            <td>{{ r.baseCoinCost.usd }}</td>
            <td>
              <button @click="() => {editRecord(r)}">Edit</button>
              <button @click="() => {removeRecord(r.id)}">Delete</button>
            </td>
      </tr>
    </table>

    <HistoryDialog :action="action" :inputData="inputData" :basePrices="basePrices" v-on:close="closeDialog" v-on:new="onNewRecord"
      v-on:edit="onEditRecord"></HistoryDialog>
  </div>
</template>

<script>
  import firebase from 'firebase'
  require("firebase/firestore")
  import moment from 'moment'
  import axios from 'axios'

  import HistoryDialog from '../components/history-dialog.vue'
  import LedgerLinkBar from '../components/ledger-link-bar.vue'

  const db = firebase.firestore();

  export default {

    props: {
      ledgerId: String,
      uid: String,
      basePrices: Object
    },

    components: {
      HistoryDialog: HistoryDialog,
      LedgerLinkBar: LedgerLinkBar
    },

    destroyed() {},

    methods: {

      closeDialog() {
        this.action = ""
      },

      async onNewRecord(r) {
        console.log(r)
        try {
          let ref = await db.collection("users").doc(this.uid).collection('ledgers').doc(this.ledgerId)
            .collection("history").add(r)
        } catch (error) {
          console.error("Error adding document: ", error);
        }
        this.closeDialog()
      },

      async editRecord(r) {
        this.action = "edit"
        this.inputData = r
      },

      async onEditRecord(id, r) {
        console.log("edit record", id, r)
        try {
          let ref = await db.collection("users").doc(this.uid).collection('ledgers').doc(this.ledgerId)
            .collection("history").doc(id).update(r)
        } catch (error) {
          console.error("Error adding document: ", error);
        }
        this.closeDialog()
      },

      async removeRecord(id) {
        let yes = confirm("really")
        if (!yes) {
          return
        }
        try {
          await db.collection("users").doc(this.uid).collection('ledgers').doc(this.ledgerId)
            .collection("history").doc(id).delete()
        } catch (error) {
          console.log(error)
        }
      },

      loadData(uid) {
        this.unsubscribe1 = db.collection("users").doc(uid).collection('ledgers').doc(this.ledgerId)
          .onSnapshot(snapshot => {
            this.ledgerName = snapshot.data().name
          })

        this.unsubscribe2 = db.collection("users").doc(uid).collection('ledgers').doc(this.ledgerId)
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

    destroyed() {
      this.unsubscribe1();
      this.unsubscribe2();
    },

    data() {
      return {
        action: "",
        inputData: {},
        records: []
      }
    }
  }
</script>

<style lang="scss" scoped>
  .ledger-box {
    border: 1px solid black;
    width: 100px;
    height: 100px;
    display: flex;
    align-items: center;
    text-align: center;
    .inner {
      text-align: center;
      width: 100%;
    }
  }
</style>

<template>
  <div>
    <div>Ledgers</div>
    <div class="ledger-box" v-for="l in ledgers" v-bind:key="l.id">
      <div class="inner">
        <router-link :to="'/ledger/' + l.id" tag="a">{{ l.id }}</router-link>
        <hr>
        <button type="button" @click="editLedger">Edit</button>
      </div>
    </div>
    <hr>
    <div>
      Ledger Name
      <input type="text" v-model="newLedgerName">
      <button type="button" @click="addLedger">Add</button>
    </div>
  </div>
</template>


<script>
  import firebase from 'firebase';
  require("firebase/firestore")


  const db = firebase.firestore();

  export default {

    props: {
      uid: String
    },

    methods: {
      addLedger() {
        if (!this.newLedgerName) {
          return;
        }
        db.collection("users").doc(this.uid).collection('ledgers').doc(this.newLedgerName).set({
          createdAt: new Date()
        })
      },

      editLedger() {
        console.log("edit")
      },

      getLedgers() {
        if (!this.uid) {
          return
        }
        db.collection("users").doc(this.uid).collection('ledgers').orderBy("createdAt", "desc")
          .onSnapshot(snapshot => {
            let ledgers = snapshot.docs.map(doc => {
              let d = doc.data()
              d.id = doc.id
              return d
            })
            this.ledgers = ledgers
          });
      }
    },

    watch: {
      uid(newVal, oldVal) {
        this.getLedgers()
      }
    },

    mounted() {
      if (this.uid) {
        this.getLedgers()
      }
    },

    data() {
      return {
        ledgers: null,
        newLedgerName: ""
      }
    }
  }
</script>

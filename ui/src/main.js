var Vue = require('vue')
var VueRouter = require('vue-router')
Vue.use(VueRouter)

var firebase = require("firebase/app");
require("firebase/auth");

var config = {
  apiKey: "AIzaSyDKs8ugaxIPshjQNKGG3CvxwDuSQQRqMSk",
  authDomain: "happy-ledger.firebaseapp.com",
  databaseURL: "https://happy-ledger.firebaseio.com",
  projectId: "happy-ledger",
  storageBucket: "happy-ledger.appspot.com",
  messagingSenderId: "644049233724"
};
firebase.initializeApp(config);

var Main = require('./app/main.vue')
var History = require('./app/history.vue')
var Login = require('./app/login.vue')

var routes = [
  {path: '/login', component: Login},
  {path: '/', component: History},
  {path: '*', component: History}
]

var router = new VueRouter({routes})

firebase.auth().onAuthStateChanged(user => {
  if (user) {
    console.log('logged in as user:', user.uid)
    localStorage.setItem("uid", user.uid)
    if (app.$route.path === "/login") {
      app.$router.push({path: '/'})
    }
  } else {
    console.log('not logged in')
    localStorage.removeItem("uid")
    app.$router.push({path: '/login', redirect: app.$route.path})
  }
})

var app = new Vue({
  router,
  el: '#main',
  render: h => h(Main)
})

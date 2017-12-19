var Vue = require('vue')
var VueRouter = require('vue-router')
Vue.use(VueRouter)

var Main = require('./app/main.vue')
var History = require('./app/history.vue')

var routes = [
  {path: '/', component: History},
  {path: '*', component: History}
]

var router = new VueRouter({routes})

var app = new Vue({
  router,
  el: '#main',
  render: h => h(Main)
})

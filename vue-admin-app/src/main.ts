import Vue from 'vue';

import 'normalize.css';
import ElementUI from 'element-ui'
import App from './App.vue';
import router from './router';
import store from './vuex/store';

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');

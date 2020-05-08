import Vue from "vue";

import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-default.css';

import App from "./App";
import router from "./router/index";

import VueAxios from "vue-axios";
import axios from "axios";
Vue.use(VueAxios, axios);

import Davinci from "./plugins/davinciEnviroment";
import "vue-notifyjs/themes/default.css";

Vue.use(Davinci);

var pagination = require("vuejs-uib-pagination");
Vue.use(pagination);

window.$ = require("jquery-ui");
window.JQuery = require("jquery-ui");
import "jquery-ui";
import "jquery-ui/ui/widgets/datepicker";
import "jquery-ui/themes/base/core.css";
import "jquery-ui/themes/base/theme.css";
import "jquery-ui/themes/base/datepicker.css";

//Metodo comunes en las distintas Vistas
Vue.mixin({
  methods: {

    alertSuccess: function (message) {
      Vue.use(VueToast);
      Vue.$toast.open({
        message: message,
        type: 'success',
      });
    },

    alertError: function (message) {
      Vue.use(VueToast);
      Vue.$toast.open({
        message: message,
        type: 'error',
      });
    },

    alertInfo: function (message) {
      Vue.use(VueToast);
      Vue.$toast.open({
        message: message,
        type: 'info',
      });

    },

    alertCustom: function (custom) {
      Vue.use(VueToast);
      Vue.$toast.open(
        custom
      );
    },
    

    screen: function(theSwitch) {
      var all = document.getElementsByClassName("fixedmodal");
      for (var i = 0; i < all.length; i++) {
        console.log("OCULTA!!");
        all[i].style.display = theSwitch == "on" ? "block" : "none";
      }
    },
    formatDays: (date) => {
      //2019-09-23 12:02:17
      var createDate = new Date(date);

      if (Object.prototype.toString.call(createDate) === "[object Date]") {
        var today = new Date();
        var times = today.getTime() - createDate.getTime();
        var days = Math.round(times / (1000 * 3600 * 24));

        //If Hoy
        if (days == 0) {
          return "Hoy";
        }
        //if Ayer
        if (days == 1) {
          return "Ayer";
        }
        //if Hace mas de 1 dia
        if (days > 1 && days <= 6) {
          return `hace ${days} dias`;
        }
        //if Hace mas una semana
        if (days == 7) {
          return `hace una semana`;
        }
        //if Hace mas de una Semana
        if (days > 7 && days < 30) {
          return `hace ${Math.round(days / 7)} semanas`;
        }
        //if Hace un Mes
        if (days >= 30 && days < 60) {
          return `hace un mes`;
        }
        //if Hace mas de un Mes
        if (days >= 60 && days < 365) {
          return `hace ${Math.round(days / 30)} meses`;
        }
        //if Hace un año
        if (days >= 365 && days < 600) {
          return `hace un año`;
        }
        //if Hace N años
      }
      return "";
    },
    isUnique(value, key, list) {
      for (var i in list) {
        var item = list[i];
        if (item[key] == value) {
          return true;
        }
      }
      return false;
    },
    makeId: (length) => {
      var result = "";
      var characters =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
      var charactersLength = characters.length;
      for (var i = 0; i < length; i++) {
        result += characters.charAt(
          Math.floor(Math.random() * charactersLength)
        );
      }
      return result;
    },
    isEmptyOrSpaces: (str) => {
      if (typeof str === "undefined") {
        return true;
      }
      str += "";
      return str === null || str.match(/^ *$/) !== null;
    },
    cloneObject: (any) => {
      return JSON.parse(JSON.stringify(any));
    },
  },
});

/* eslint-disable no-new */
new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");

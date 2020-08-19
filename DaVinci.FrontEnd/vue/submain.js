import Vue from "vue";
import SubApp from "./SubApp";
import router from "./router/subRoutes";

import VueToast from 'vue-toast-notification';
import 'vue-toast-notification/dist/theme-default.css';

import VueAxios from "vue-axios";
import axios from "axios";
Vue.use(VueAxios, axios);

import Davinci from "./plugins/davinciEnviroment";

Vue.use(Davinci);

import VueMenu from "@hscmap/vue-menu";
Vue.use(VueMenu);

var pagination = require("vuejs-uib-pagination");
Vue.use(pagination);

window.$ = require("jquery-ui");
window.JQuery = require("jquery-ui");
import "jquery-ui";
import "jquery-ui/ui/widgets/datepicker";
import "jquery-ui/themes/base/core.css";
import "jquery-ui/themes/base/datepicker.css";
import "jquery-ui/themes/base/theme.css";

//Metodo comunes en las distintas Vistas
Vue.mixin({
  methods: {

    formatDays: date => {
      //2019-09-23 12:02:17
      var createDate = new Date(date);


      if (Object.prototype.toString.call(createDate) === "[object Date]") {
        var today = new Date();
        var times = today.getTime() - createDate.getTime();
        var days = Math.round(times / (1000 * 3600 * 24));

        //If Hoy
        if (days == 0) {
          return "Hoy a las " + date.substring(11, 19);
        }
        //if Ayer
        if (days == 1) {
          return "Ayer a las " + date.substring(11, 19);
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
    async getInventionDefByCode(inventionCode) {
      var projectCode = await this.getCodeProject();
      var url = `/api/project/${projectCode}/${inventionCode}/get`;

      var inventionByCodeCode = `${projectCode}#${inventionCode}.precondition`;

      var inventionDef = JSON.parse(
        sessionStorage.getItem(inventionByCodeCode)
      );

      if (inventionDef == null) {
        var inventionDB = {};
        await this.axios.get("/api/invention/all/").then(rs => {
          rs.data.forEach(invention => {
            if (invention.code == inventionCode) {
              inventionDB = invention;
              return inventionDB;
            }
          });
        });
        var list = [];
        var wareHouse = {};
        await this.axios.get(url).then(rs => {
          for (var i in rs.data) {
            var item = rs.data[i];
            var option = {};
            option[inventionDB.keyLabel] = item[inventionDB.keyLabel][0];
            option[inventionDB.keyValue] = item[inventionDB.keyValue][0];
            list.push(option);
            wareHouse[item[inventionDB.keyValue][0]] = item;
          }
        });

        inventionDef = {
          list: list,
          WareHouse: wareHouse,
          artifacts: inventionDB.artifacts,
          keyLabel: inventionDB.keyLabel,
          keyValue: inventionDB.keyValue
        };

        sessionStorage.setItem(
          inventionByCodeCode,
          JSON.stringify(inventionDef)
        );
      }
      return inventionDef;
    },

    getProjectDomain() {
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      return projectName;
    },

    async getCodeProject() {
      var code = "";
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];

      var varCodeProject = projectName + "_CODE";

      var code = JSON.parse(sessionStorage.getItem(varCodeProject));
      console.log(code);
      if (code == null) {
        await this.axios.get(`/davinci/${projectName}/code`).then(rs => {
          code = rs.data;
        });
        sessionStorage.setItem(varCodeProject, JSON.stringify(code));
      }
      return code;
    },
    async getTypesOfEpics() {
      //TODO : Guardar como sesion stored
      var epicTypes = [];
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      await this.axios.get(`/davinci/${projectName}/epicTypes`).then(rs => {
        epicTypes = rs.data == null ? [] : rs.data;
      });
      return epicTypes;
    },

    getUserOnline() {
      return "admin";
    },

    async getTypesOfUserStories() {
      //TODO : Guardar como sesion stored
      var userStoriesTypes = [];
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      await this.axios.get(`/davinci/${projectName}/userStories`).then(rs => {
        userStoriesTypes = rs.data == null ? [] : rs.data;
      });
      return userStoriesTypes;
    },

    getContextCSS() {
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      ////console.log("LOAD CONTEXT >" + projectName);
      var context = JSON.parse(sessionStorage.getItem(projectName));
      var style = document.createElement("style");
      style.type = "text/css";
      style.innerHTML = context;
      document.getElementsByTagName("head")[0].appendChild(style);

      return context;
    },

    setContextCSS: (projectName, context) => {
      ////console.log(`Contexto de proyect ${projectName} Seteado.`);
      sessionStorage.setItem(projectName, JSON.stringify(context));
    },
    async getUserStoryByREF(code, indexVersion) {
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      var us = {};
      var usRef = {
        code: code,
        indexVersion: indexVersion
      };
      await this.axios
        .post(`/davinci/${projectName}/userStories/ref/`, usRef)
        .then(rs => {
          us = rs.data;
        });
      return us;
    },
    async getAllDataTypes() {
      var dataTypes = [];
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      await this.axios.get(`/davinci/${projectName}/datatype/getAll/`).then(rs => {
        dataTypes = rs.data;
      });
      return dataTypes;
    },
    //trae los inv en contexto del proeycto
    async getAllInventions() {
      var invents = [];
      console.log("Go to back")
      var projectCode = await this.getCodeProject();
      console.log(`/davinci/invention/${projectCode}/all/`)
      await this.axios.get(`/davinci/invention/${projectCode}/all/`).then(rs => {
        invents = rs.data;
        console.log("getAll submain",invents)
        console.log("rs",rs)
      });
      return invents;
    },
    async saveEpic(epic) {
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      var projectEpics = pathname.split("/")[2] + "_EPICS";
      var epics = JSON.parse(sessionStorage.getItem(projectEpics));
      if (epics == null || Object.keys(epics).length == 0) {
        epics = {};
        epics[projectEpics] = {};
        epics[projectEpics][epic.code] = epic;
      } else {
        epics[projectEpics][epic.code] = epic;
      }
      var epicReturn = {};
      await this.axios.post(`/davinci/${projectName}/epic`, epic).then(rs => {
        //TODO : Formatear la respuesta de guardado.
        epicReturn = rs.data;
      });

      sessionStorage.setItem(projectEpics, JSON.stringify(epics));
      return epicReturn;
    },

    async getEpics() {
      var epics = [];
      /*  
      http://localhost:1123/davinci/TestingBECH/epics
      */
      var pathname = window.location.pathname;
      var projectName = pathname.split("/")[2];
      await this.axios.get(`/davinci/${projectName}/epics`).then(rs => {
        epics = rs.data;
        console.log("epics",epics)
      });
      return epics;
    },

    makeId: length => {
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
    isEmptyOrSpaces: str => {
      if (typeof str === "undefined") {
        return true;
      }
      str += "";
      return str === null || str.match(/^ *$/) !== null;
    },
    cloneObject: any => {
      return JSON.parse(JSON.stringify(any));
    }
  }
});

/* eslint-disable no-new */
new Vue({
  router,
  render: h => h(SubApp)
}).$mount("#app");

import Notify from "vue-notifyjs";
import SideBar from "../components/UserInterface/SidebarPlugin";
import GlobalComponents from "./globalComponents";
import GlobalDirectives from "./globalDirectives";

import "es6-promise/auto";

//css assets
import "bootstrap/dist/css/bootstrap.css";
import "@/assets/sass/davinci-dashboard.scss";
import "@/assets/css/themify-icons.css";
import "@/assets/css/font-awesome.css";
import 'vue-form-wizard/dist/vue-form-wizard.min.css'

export default {
  install(Vue) {
    Vue.use(GlobalComponents);
    Vue.use(GlobalDirectives);
    Vue.use(SideBar);
    Vue.use(Notify);
  }
};

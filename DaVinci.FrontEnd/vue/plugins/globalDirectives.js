import { directive as vClickOutside } from "vue-clickaway";

/**
 * Directivas Globales
 */

const GlobalDirectives = {
  install(Vue) {
    Vue.directive("click-outside", vClickOutside);
  }
};

export default GlobalDirectives;

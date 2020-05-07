<template>
  <div :class="{ 'nav-open': $sidebar.showSidebar }">    
    <img src="@/assets/img/davinci-logo.png" id="favicon" style="display:none" />
    <session-watch></session-watch>
    <notifications></notifications>
    <router-view></router-view>
  </div>
</template>
<script>
  export default {
    name: 'app',
    mounted() {
      var link = document.querySelector("link[rel*='icon']") || document.createElement('link');
      link.type = 'image/x-icon';
      link.rel = 'shortcut icon';
      link.href = document.getElementById("favicon").src;
      document.getElementsByTagName('head')[0].appendChild(link);
      document.title = "DaVinci : Orquestador de Diseños.";
    }
  }
</script>
<style lang="scss">
  .vue-notifyjs.notifications {
    .alert {
      z-index: 10000;
    }

    .list-move {
      transition: transform 0.3s, opacity 0.4s;
    }

    .list-item {
      display: inline-block;
      margin-right: 10px;
    }

    .list-enter-active {
      transition: transform 0.2s ease-in, opacity 0.4s ease-in;
    }

    .list-leave-active {
      transition: transform 1s ease-out, opacity 0.4s ease-out;
    }

    .list-enter {
      opacity: 0;
      transform: scale(1.1);
    }

    .list-leave-to {
      opacity: 0;
      transform: scale(1.2, 0.7);
    }
  }
</style>
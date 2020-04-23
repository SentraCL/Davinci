<template>
  <nav class="navbar navbar-expand-lg navbar-light navback" :style="{'background-image': 'url(' + require('@/assets/img/toolbar.png') + ') !important', 'background-size': '100% 100%','background-repeat': 'no-repeat'}" >
    <div class="container-fluid">
      <h1 class="route-name">{{ routeName }} </h1>
      <button class="navbar-toggler navbar-burger" type="button" @click="toggleSidebar" :aria-expanded="$sidebar.showSidebar" aria-label="Toggle navigation">
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
        <span class="navbar-toggler-bar"></span>
      </button>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto">
          <!-- 
          <drop-down class="nav-item" :title="notifications.length + ' Notificaciones'" title-classes="nav-link" icon="ti-bell">
            <a class="dropdown-item" :title="notification.subject" :href="notification.url" v-for="notification,index in notifications" :key="index">{{notification.subject | just25Char}}</a>
          </drop-down>
          -->

          <li class="nav-item">
            <a href="#" @click="logout()" class="nav-link">
              <i class="ti-user"></i>
              <p>
                Sesión
              </p>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script>
  export default {
    props: {
      notifications: Array
    },
    computed: {
      routeName() {
        const {
          name
        } = this.$route;
        //return this.capitalizeFirstLetter(name);
        return name;
      }
    },
    filters: {
      just25Char: function (value) {
        var content = value;
        if (typeof value === "string" || value instanceof String) {
          if (value != null) {
            if (value.length <= 25) {
              content = value;
            } else {
              content = value.substring(0, 25) + "...";
            }
          }
        }
        return content;
      }
    },
    data() {
      return {
        activeNotifications: false,

      }
    },

    methods: {

      logout() {
        this.axios
          .post("/api/logout/").then(rs => {
            sessionStorage.setItem("loginHash", null);
            window.location.href = "/";
          });

      },
      /*
      capitalizeFirstLetter(string) {
        try {
          return string.charAt(0).toUpperCase() + string.slice(1);
        } catch {
          return ""
        }

      },
      */
      toggleNotificationDropDown() {
        this.activeNotifications = !this.activeNotifications;
      },
      closeDropDown() {
        this.activeNotifications = false;
      },
      toggleSidebar() {
        this.$sidebar.displaySidebar(!this.$sidebar.showSidebar);
      },
      hideSidebar() {
        this.$sidebar.displaySidebar(false);
      }
    }

  };
</script>
<style scoped>
  .navback {
  
  }
</style>
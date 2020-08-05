<template>
  <div class="wrapper">
    <span>
      <side-bar>
        <!-- Seteo de arreglo para SideBar -->
        <template slot="links">
          <sidebar-link :to="link.to" :name="link.name" :icon="link.icon" v-for="(link, index) in pages"  v-if="role=='admin' && link.name=='Empresas'" :key="index" />
          <sidebar-link :to="link.to" :name="link.name" :icon="link.icon" v-for="(link, index) in pages" v-if="link.name!='Empresas'"   :key="index" />
        </template>
        <mobile-menu>
          <!-- MENU TOP PARA MOBIL -->
          <li class="nav-item">


            <drop-down class="nav-item" :title="notifications.length + ' Notificaciones'" title-classes="nav-link" icon="ti-bell">
              <a class="dropdown-item" :title="notification.subject" :href="notification.url" v-for="notification,index in notifications" :key="index">{{notification.subject | just25Char}}</a>
            </drop-down>


          <li class="divider"></li>
        </mobile-menu>
      </side-bar>
      <div class="main-panel">
        <top-navbar :notifications="notifications"></top-navbar>

        <dashboard-content @click.native="toggleSidebar"> </dashboard-content>

        <content-footer>
          <template slot="links">
            <sidebar-link :to="link.to" :name="link.name" :icon="link.icon" v-for="(link, index) in pages" v-if="role=='admin' && link.name=='Empresas'" :key="index" />
            <sidebar-link :to="link.to" :name="link.name" :icon="link.icon" v-for="(link, index) in pages" v-if="link.name!='Empresas'"  :key="index" />
          </template>
        </content-footer>
      </div>
    </span>
  </div>
</template>
<style lang="scss"></style>
<script>
  import TopNavbar from "./TopNavbar.vue";
  import ContentFooter from "./ContentFooter.vue";
  import DashboardContent from "./Content.vue";
  import MobileMenu from "./MobileMenu";
  export default {
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
    components: {
      TopNavbar,
      ContentFooter,
      DashboardContent,
      MobileMenu
    },

    data() {
      let rol=this.getRole();
      return {
        role:rol,
        notifications: [{
          url: '#',
          subject: 'Ayuda con un nuevo Tipo de Plan de Prueba.',
          author: 'analista@cliente.cl',
          message: 'Necesito crear un nuevo tipo de plan de prueba.',
          fecha: '2019/05/29 19:49:11'
        }, {
          url: '#',
          subject: 'Ayuda con un nuevo Tipo de Plan de Prueba.',
          author: 'analista@cliente.cl',
          message: 'Necesito crear un nuevo tipo de plan de prueba.',
          fecha: '2019/05/29 19:49:11'
        }, {
          url: '#',
          subject: 'Ayuda con un nuevo Tipo de Plan de Prueba.',
          author: 'analista@cliente.cl',
          message: 'Necesito crear un nuevo tipo de plan de prueba.',
          fecha: '2019/05/29 19:49:11'
        },

        ],
        pages: [{
          to: '/enterprise',
          name: 'Empresas',
          icon: 'ti-home',
          rol:'admin'
        },{
          to: '/users',
          name: 'Usuarios',
          icon: 'ti-user',
          rol:'admin'
        },{
          to: '/dashboard',
          name: 'Proyectos',
          icon: 'ti-blackboard',
          rol:'any'
        }, {
          to: '/inventions',
          name: 'Inventos',
          icon: 'ti-ruler-pencil',
          rol:'any'
        }, {
          to: '/portfolio',
          name: 'Portafolio',
          icon: 'fa fa-suitcase',
          rol:'any'
        }, {
          to: '/design',
          name: 'Diseño',
          icon: 'ti-write',
          rol:'any'
        }]
      }
    },
    methods: {
      async getRole(){
        return await this.getUserRoleByHash(sessionStorage.getItem("loginHash"));
      },
      async getUserRoleByHash(hash) {
        var url = `/api/user/${hash}/get`;
        await this.axios.get(url,{"loginHash": hash}).then(rs => {
            this.role=rs.data;
            return rs.data;
        });
        return "any";
      },

      toggleSidebar() {
        if (this.$sidebar.showSidebar) {
          this.$sidebar.displaySidebar(false);
        }
      }
    }
  };
</script>
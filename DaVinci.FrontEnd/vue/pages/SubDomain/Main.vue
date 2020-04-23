<template>
  <div id="WORKSPACE" class="wrapper ">
    <hsc-menu-style-white class="wallpaper st-container">
      <hsc-menu-bar id="SUBDOMAIN-MENU" class="st-menu main-menu always-on-top">
        <img src="@/assets/img/davinci-logo.png" width="24px" />
        <hsc-menu-bar-item label="Epicos">
          <hsc-menu-item :label="createTitleEpic">
            <hsc-menu-item :label="typeEpic.name" v-for="typeEpic,key in typeOfEpics" @click="createEpic(typeEpic)" :key="key" />
          </hsc-menu-item>
          <hsc-menu-item :label="epicRepositoryTitle" @click="doFindEpic()" />
          <!--
          <hsc-menu-separator />
          
          <hsc-menu-item label="Guardar" @click="showLog('Save')" :disabled="true" />
         
          <hsc-menu-item label="Agregar Historia de Usuario" @click="showLog('Save')" :disabled="true" />
          <hsc-menu-item label="Exportar" :disabled="true">
            <hsc-menu-item label="XML (TestLink)" />
            <hsc-menu-item label="CSV" />
          </hsc-menu-item>
          -->
        </hsc-menu-bar-item>

        <hsc-menu-bar-item label="Historias de Usuarios">
          <hsc-menu-item :label="createTitleUH">
            <hsc-menu-item :label="typeUserStory.title" v-for="typeUserStory,key in typeOfUserStoriesList" @click="createUserStories(typeUserStory)" :key="key" />
          </hsc-menu-item>
          <hsc-menu-item :label="findTitle" @click="doFindUserStory()" />
          <hsc-menu-separator v-if="existUserTask" />
          <hsc-menu-item :title="usWindows.userStory.title + ' : ' + usWindows.userStory.fields.description" :label="usWindows.userStory.title" v-if="usWindows.isUserStory && usWindows.exist" v-for="usWindows in tasks" @click="goto(usWindows.userStory.title)" :key="usWindows.userStory.title" />
        </hsc-menu-bar-item>

        <hsc-menu-bar-item label="Datos">
          <hsc-menu-item label="Cargar" @click="showLog('Cargar Datos')" />
          <hsc-menu-item label="Crear Tipos de Datos" @click="showLog('Crear Datos')" />
          <hsc-menu-item label="Buscar" @click="showLog('Buscar Datos')" />
        </hsc-menu-bar-item>


        <hsc-menu-bar-item label="Sesión">
          <hsc-menu-item label="Estadisticas" @click="showLog('Estadisticas')" />
          <hsc-menu-item label="Cambiar Contraseña" @click="showPassword()" />
          <hsc-menu-item label="Cerrar Sesion" @click="logout()" />
        </hsc-menu-bar-item>
      </hsc-menu-bar>
      <br />
      <br />
      <div class="row">
        <div class="col-md-1"></div>
        <div class="col-md-10">
          <workspace ref="workspace" :tasks.sync="tasks">
          </workspace>

        </div>
        <div class="col-md-1"></div>
      </div>
    </hsc-menu-style-white>


    <side-panel :show.sync="sidePanel.show">
      <div slot="content">
        <span v-if="sidePanel.isUS">
          <find-user-story :idUS.sync="idUS" v-on:toClose="closeFindUser()" v-on:found="getUserStory()"></find-user-story>
        </span>
        <span v-if="sidePanel.isEpic">
          <repository-epic :idEpic.sync="idEpic" v-on:found="getEpic()" v-on:toClose="closeRepositoryEpic()"></repository-epic>
        </span>
      </div>
    </side-panel>



    <m-dialog id="showPasswordDlg" :title="showPasswordDlg.title" :show.sync="showPasswordDlg.show" :isClose.sync="showPasswordDlg.close">
      <div slot="dialog">
        <password></password>
      </div>
    </m-dialog>

  </div>
</template>

<script>
  import Workspace from "./Workspace.vue";
  import Epic from "./Epic/Epic.vue";
  import UserStory from "./UserStories/UserStory.vue";
  import RepositoryEpic from "./Epic/Repository.vue";
  import Password from "./Profile/Password.vue";
  import FindUserStory from "./UserStories/FindUserStory.vue";

  export default {
    name: "subbdomain",
    computed: {},
    components: {
      Workspace,
      Epic,
      RepositoryEpic,
      UserStory,
      FindUserStory,
      Password
    },
    props: {},
    watch: {
      tasks(newValue, oldValue) {
        this.existUserTask = false;
        this.existEpicTask = false;
        newValue.forEach(task => {
          if (task.exist) {
            if (task.isUserStory) {
              this.existUserTask = true;
            }
            if (task.isEpic) {
              this.existEpicTask = true;
            }
          }
        })
      }
    },

    async mounted() {
      await this.loadEpics();
      await this.getContextCSS();
      this.typeOfEpics = await this.getTypesOfEpics();
      this.typeOfUserStoriesList = await this.getTypesOfUserStories();
    },

    data() {
      return {
        idEpic: "",
        idUS: "",
        existEpicTask: false,
        existUserTask: false,
        sidePanel: {
          isUS: false,
          isEpic: false,
          show: false
        },


        projectName: "",
        epics: [],
        tasks: [],
        epicRepositoryTitle: `<i class="ti-agenda"></i> Repositorio de Epicos`,
        createTitleEpic: `<i class="ti-bookmark-alt"></i> Crear`,
        createTitleUH: `<i class="ti-notepad"></i> Crear`,
        findTitle: `<i class="ti-direction-alt"></i> Buscar`,

        findEpicsDlg: {
          title: `<span class="ti-server"> </span> Repositorio de Epicos.`,
          show: false,
          close: true,
          artifact: {}
        },
        showPasswordDlg: {
          title: `<span class="ti-lock"> </span> Cambiar Contraseña`,
          show: false,
          close: true,
          artifact: {}
        },

        //Data que vendria de Davinci:
        typeOfEpics: [],
        typeOfUserStoriesList: []
      };
    },

    methods: {

      async getEpic() {
        await this.$refs.workspace.takeEpic(this.idEpic);
      },

      async getUserStory() {
        await this.$refs.workspace.takeUserStory(this.idUS);
      },

      goto(nameLink) {
        window.location.href = '#' + nameLink;
      },

      doFindUserStory() {
        this.sidePanel.isEpic = false;
        this.sidePanel.isUS = true;
        this.sidePanel.show = true;
      },


      closeRepositoryEpic() {
        //console.log("cierro desde el parent");
        this.sidePanel.show = false;
      },
      closeFindUser() {
        //console.log("cierro desde el parent");
        this.sidePanel.show = false;
      },

      async loadEpics() {
        this.epics = (await this.getEpics()) == null ? [] : this.getEpics();
      },

      createEpic(itemType) {
        this.$refs.workspace.addNewEpic(itemType);
      },

      createUserStories(itemType) {
        this.$refs.workspace.addNewUserStyory(itemType);
      },

      doFindEpic() {
        var epics = this.getEpics();
        //console.log(JSON.stringify(epics));
        //this.findEpicsDlg.show = true;
        this.sidePanel.isEpic = true;
        this.sidePanel.isUS = false;
        this.sidePanel.show = !this.sidePanel.show;
      },
      showPassword() {
        this.showPasswordDlg.show = true;
        this.showForm = true;
      },


      showLog(msg) {
        console.log(msg);
      },
      logout() {
        var pathname = window.location.pathname;
        this.projectName = pathname.split("/")[2];
        this.axios.get(`/davinci/${this.projectName}/logOut`).then(rs => {
          this.$router.push("login");
        });
      }
    }
  };
</script>
<style>
  .main-menu {
    border-radius: 0 0 4pt 0;
    position: absolute;
    z-index: 9999;
    cursor: pointer !important;
  }

  .wallpaper {
    overflow-y: scroll;
    overflow-x: hidden;
    position: fixed;

    background: linear-gradient(217deg,
        var(--sub-card),
        var(--sub-wallpaper) 70.71%),
      linear-gradient(127deg, var(--sub-background), var(--sub-toolbar) 70.71%),
      linear-gradient(336deg, var(--sub-text), var(--sub-button) 70.71%);
    height: 100%;
    width: 100%;
  }

  div.ftl-vertical-tab-content {
    float: right;
    background-color: #fff !important;

    padding-left: 10px;
    padding-right: 10px;
    padding-top: 10px;
    min-width: calc(100% - 180px);
    max-width: 600px;
  }

  div.ftl-vertical-tab-menu div.list-group>span.active:after {
    border-left: 10px solid var(--sub-dark) !important;
  }

  .always-on-top {
    -webkit-touch-callout: none;
    /* iOS Safari */
    -webkit-user-select: none;
    /* Safari */
    -khtml-user-select: none;
    /* Konqueror HTML */
    -moz-user-select: none;
    /* Firefox */
    -ms-user-select: none;
    /* Internet Explorer/Edge */
    user-select: none;
    position: fixed;
    top: 0px;
  }
</style>
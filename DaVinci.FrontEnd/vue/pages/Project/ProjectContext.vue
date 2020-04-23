<template>
  <div class="card">
    <div class="col-md-12">
      <div class="row">
        <div class="card-body">
          <h4 class="card-title"><i class="fa fa-users"></i> Administración de Usuario</h4>
          <!-- {{project}} -->
          <v-tabs :tabs="userTabs">
            <div :slot="tab.slots" v-for="(tab, index) of userTabs" :key="index" class="user-panel">
              <span v-if="tab.slots=='NewUser'">
                <!--Nuevo usuario-->

                <input-text type="text" label="Nombre" v-model="tab.usuarioVO.name"></input-text>
                <input-text type="password" label="Password" v-model="tab.usuarioVO.pass"></input-text>
                <hr />

                <d-button type="info" round @click.native.prevent="addNewUserTab(tab.usuarioVO)">
                  <span style="color:darkgreen" class="ti-back-left"></span> Nuevo Usuario
                </d-button>
              </span>
              <span v-if="tab.slots==tab.usuarioVO.name">
                <input-text type="text" label="Nombre" :isInactive="true" v-model="tab.usuarioVO.name"></input-text>
                <input-text type="password" label="Password" v-model="tab.usuarioVO.pass"></input-text>
                <hr />
                <div class="btn-toolbar">
                  <d-button type="info" round @click.native.prevent="saveTab(tab.usuarioVO)">
                    <span style="color:darkgreen" class="ti-back-left"></span> Guardar Usuario
                  </d-button>
                  <d-button type="danger" round @click.native.prevent="deleteTab(tab.usuarioVO)">
                    <span style="color:darkgreen" class="ti-back-left"></span> Eliminar Usuario
                  </d-button>
                </div>
              </span>
            </div>
            <br />
            <div class="clearfix"></div>
          </v-tabs>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  export default {
    name: "project-context",
    computed: {},
    props: {
      project: {}
    },

    created() {
      this.userTabs = this.getTemplateNewUserTab();
    },

    data() {
      return {
        userTabs: []
      };
    },

    methods: {
      deleteTab(userVO) {
        var status = false;
        this.axios
          .delete("/api/project/" + this.project.code + "/user/" + userVO.code)
          .then(rs => {
            status = rs.data;
          });
        this.$emit("delete");
      },
      saveTab(userVO) {
        var userUpdate = {};
        userUpdate.name = userVO.name;
        userUpdate.pass = userVO.pass;
        userUpdate.code = userVO.code;

        this.axios
          .put("/api/project/" + this.project.code + "/user/update/", userUpdate)
          .then(rs => {
            status = rs.data;
          });
        this.$emit("update");
      },
      addNewUserTab(_newUser) {
        var cloneUser = this.cloneObject(_newUser);
        cloneUser.slots = cloneUser.name;
        _newUser.name = "";
        _newUser.pass = "";

        this.userTabs.push({
          id: -1,
          slots: cloneUser.slots,
          title: cloneUser.name,
          subtitle: "Llena el formulario",
          usuarioVO: cloneUser
        });

        var status = false;
        this.axios
          .post("/api/project/" + this.project.code + "/user/save/", cloneUser)
          .then(rs => {
            status = rs.data;
          });
        this.$emit("save");
      },
      getTemplateNewUserTab() {
        var templateTab = [];
        templateTab.push({
          id: -1,
          slots: "NewUser",
          title: "Nuevo Usuario",
          subtitle: "Llena el formulario",
          usuarioVO: {
            name: "",
            pass: ""
          }
        });
        if (this.project && this.project.Users) {
          this.project.Users.forEach(element => {
            templateTab.push({
              id: -1,
              slots: element.UserName,
              title: element.UserName,
              usuarioVO: {
                code: element.Code,
                name: element.UserName,
                pass: element.Password
              }
            });
          });
        }
        return templateTab;
      }
    }
  };
</script>
<style>
  .user-panel {
    min-height: 200px;
  }
</style>
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
                <input-text alphabetic removeSpace  type="text" id="inputNombre" label="Nombre" v-model="tab.usuarioVO.name"></input-text>
                <input-text removeSpace  type="password" id="inputPass" label="Password" v-model="tab.usuarioVO.pass"></input-text>
                <check-box :disableonoff="true" :value.sync="tab.usuarioVO.isDesign" id="inputDesign" label="Diseñador"></check-box>
                <hr />

                <d-button type="info" round @click.native.prevent="addNewUserTab(tab.usuarioVO)">
                  <span style="color:darkgreen" class="ti-back-left"></span> Nuevo Usuario
                </d-button>
              </span>
              <span v-if="tab.slots==tab.usuarioVO.name">
                <input-text  type="text" label="Nombre" :isInactive="true" v-model="tab.usuarioVO.name"></input-text>
                <input-text type="password" label="Password" v-model="tab.usuarioVO.pass"></input-text>
                <check-box :disableonoff="true" v-model="tab.usuarioVO.isDesign" label="Diseñador"></check-box>
                <hr />
                <div class="btn-toolbar">
                  <d-button type="info" round @click.native.prevent="saveTab(tab.usuarioVO)">
                    <span style="color:darkgreen" class="ti-back-left"></span> Guardar Usuario
                  </d-button>
                  <d-button type="danger" round @click.native.prevent="deleteUser()">
                    <span style="color:darkgreen" class="ti-back-left"></span> Eliminar Usuario
                  </d-button>
                  <m-dialog id="deleteDialog"  :show.sync="delAlert.show" :isClose.sync="delAlert.close">
            <span slot="dialog">
                <div>
                   Estas seguro que quieres eliminar usuario <strong>{{tab.usuarioVO.name}}</strong> ?<br /> 
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeTab()">
                        No , Cancelar
                    </d-button>
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="deleteTab(tab.usuarioVO)">
                        Si, Eliminar
                    </d-button>
                </span>
            </span>
        </m-dialog>
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
        delAlert: {
          mensaje:"",
          tiempo: 10,
          show: false
        },
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
        this.userTabs.forEach((value,index)=>{
          if(value.usuarioVO.code==userVO.code){
            this.userTabs.splice(index,1);
          }
        });
        this.$emit("delete");
        this.delAlert.show = false;
      },
      saveTab(userVO) {
        var userUpdate = {};
        userUpdate.name = userVO.name;
        userUpdate.pass = userVO.pass;
        userUpdate.code = userVO.code;
        userUpdate.isDesign = userVO.isDesign;
        this.axios
          .put("/api/project/" + this.project.code + "/user/update/", userUpdate)
          .then(rs => {
            status = rs.data;
          });
        this.$emit("update");
      },
      async addNewUserTab(_newUser) {
        if(document.getElementById("inputNombre").value == '' || document.getElementById("inputPass").value ==''){
          this.alertError("No puedes dejar campos sin valor.");
          return false;
        }
        if(document.getElementById("inputNombre").value.length< 5){
          this.alertError("El nombre de usuario debe de ser mayor a 4 caracteres.");
          return false;
        }
         await this.axios
         .get("/api/project/"+ this.project.code + "/users")
         .then(rs => {
            var cloneUser = this.cloneObject(_newUser);
            var userExist = false;
            status = rs.data;
            var users = rs.data;
            for(var i = 0; i < users.length; i++){
             console.log(users[i].UserName);
             
              if(users[i].UserName.localeCompare(cloneUser.name) == 0 ){
               userExist = true;
             }
            }
            
            if(!userExist){
              if(cloneUser.isDesign){
                cloneUser.isDesign="1";
              }else{
                cloneUser.isDesign="0";
              }
               cloneUser.slots = "";
              _newUser.name = "";
              _newUser.pass = "";
              _newUser.isDesign = false;

                      var status = false;
                      this.axios
                      .post("/api/project/" + this.project.code + "/user/save/", cloneUser)
                      .then(rs => {
                      status = rs.data;
                      if(status){
                        this.userTabs.push( {
                        id: -1,
                        slots: cloneUser.slots,
                        title: cloneUser.name,
                        subtitle: "Llena el formulario",
                        usuarioVO: cloneUser
                        });
                        this.$emit("save");
                        this.$emit("update");
                      }else{
                        this.alertError("Usuario "+cloneUser.name+" ya existe como usuario principal!")
                      }
                      });
                      userExist = false;
            }else{
              this.alertError("Usuario "+cloneUser.name+" ya existe como usuario de subdominio!")
              userExist = true;
            }
          });
        
      }, 
      deleteUser() {
       this.delAlert.show = true;
       this.delAlert.close = true;
       this.$emit("update");
      },
      closeTab(){
        this.delAlert.show = false;
        this.delAlert.close = false;
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
            pass: "",
            isDesign:false
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
                pass: element.Password,
                isDesign: element.IsDesign
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
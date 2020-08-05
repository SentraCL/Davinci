<template>
    <span>
    <h6 >
      <input-text v-if="!isEdit" v-on:input="filterUser()" label="<span class='ti-home'></span> Crear Usuario" v-model="userName" :value.sync="userName" autocomplete="off"></input-text>
    </h6>
    <div class="card-group">
      <div class="col-xl-12 col-lg-12 col-md-12" v-if="isWorkingInAnUser">
        <div class="card card">
          <div class="card-header">
            <h2  v-if="!isNew" class="card-title"><span class="ti-panel"></span> Modificar Usuario</h2>
            <h2  v-if="isNew" class="card-title"><span class="ti-panel"></span> Crear Usuario</h2>
          </div>
          <div class="card-body custom-card">
            <form v-if="!reloadForm" @submit.prevent>
                <div class="row">
                    <div class="col-md-12">
                      
                      <div class="card supercard">
                        <div class="row">
                            <div class="col-md-6">
                                <input-text type="text" v-on:input="reload" :isInactive="isEdit" label="Usuario" title="Nombre del usuario" v-model="selectedUser.UserName">
                                </input-text>
                            </div>
                            <div class="col-md-6">
                              <combo-simple label="Rol" :list="roles" keyValue="name" keyLabel="name" :value.sync="selectedUser.Role"></combo-simple>
                            </div>
                            <div v-if="isEdit" class="col-md-12 positionpadding">
                              <check-box :value.sync="cambioPass"  label="¿Desea modificar la contraseña?" ></check-box>
                            </div>
                            <div class="col-md-6 positionpadding">
                                <input-text :isInactive="!cambioPass" type="password" label="Password" title="Clave del usuario" v-model="selectedUser.Password">
                                </input-text>
                            </div>
                            <div class="col-md-6">
                                <input-text :isInactive="!cambioPass" type="password" label="Confirmar Password" title="Confirmar clave" v-model="password">
                                </input-text>
                            </div>
                        </div>
                    </div>
                        <div class="card supercard">

                          <div class="row">
                            <div class="col-md-12 positionpadding">
                              <h3><span class="ti-home"></span>Empresas</h3>
                            </div>
                            <div class="col-md-6 positionpadding">
                              <combo-simple label="Empresa" :list="enterprises" keyValue="EnterpriseId" keyLabel="Name" :value.sync="selectedEnterprise"></combo-simple>
                              <div class="btn btn-success" @click="addEnterprise()"><span class="ti-plus"></span></div>
                            </div>
                              <div class="col-md-6">
                              <ul class="list-group">
                                <li class="list-group-item" v-for="(enterprise,index) in selectedUser.showEnterprises" :key="index" >{{enterprise}} <span  @click="deleteEnterprise(enterprise)" class="ti-trash moveright"></span></li>
                              </ul>
                              </div>
                          </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-12 left">
                    <d-button type="info" :disabled="edit?true:false" round @click.native.prevent="back">
                        <span style="color:darkgreen" class="ti-back-left"></span> Volver
                    </d-button>
                    <d-button v-if="!isNew" type="success" round @click.native.prevent="updateUser">
                        <span style="color:white" class="ti-save"></span> Editar
                    </d-button>
                    <d-button v-if="isNew" type="success" round @click.native.prevent="saveUser">
                        <span style="color:white" class="ti-save"></span> Guardar
                    </d-button>
                    <d-button v-if="!isNew" type="danger" round @click.native.prevent="deleteUser">
                        <span style="color:yellow" class="ti-eraser"></span> Eliminar
                    </d-button>
                </div>

            </form>
          </div>
        </div>
      </div>
      <div class="col-lg-12" v-if="!isWorkingInAnUser&&usersList" >
        <table class="table">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Usuario</th>
              <th scope="col">Rol</th>
              <th scope="col">Ultima conexion</th>
              <th scope="col">Empresas</th>
            </tr>
          </thead>
          <tbody>
            <tr class="seleccion"  v-for="(user,index) in usersList" :key="index" @click="showUserForm(user)">
              <th>{{index+1+(pages*limit)}}</th>
              <td>{{user.UserName}}</td>
              <td>{{user.Role!=""?user.Role:"diseñador"}}</td>
              <td>{{user.LastTime}}</td>
              <td>{{user.showEnterprises}}</td>
            </tr>
          </tbody>
        </table>
        <div class="row">
            <div class="btn btn-success col-md-1" :class="actual>0?'':'divoff'" @click="changeTable('atras')"><span class="ti-angle-double-left"></span>Anterior</div><div class="col-md-1"><span class="pagina">Pagina {{pages+1}}</span></div><div class="btn btn-success col-md-1" :class="ultima<actual||numRegistros<limit?'divoff':''" @click="changeTable('siguiente')">Siguiente<span class="ti-angle-double-right"></span></div>
        </div>
      </div>
    </div>
  </span>
</template>
<script>

import UserShow from "./UserShow.vue";
import UserForm from "./UserForm.vue";

export default {
    name:"users",
    props:{
    },
    data(){
        this.getAllEnterprise()
        this.getAllUser()
        return{
            selectedUser:{},
            users:[],
            userName:"",
            isNew:false,  
            isWorkingInAnUser:false,
            usersList:[],
            enterprises:[],
            actual:0,
            limit:20,
            ultima:0,
            pages:0,
            numRegistros:0,
            selectedEnterprise:{},
            edit:false,
            roles:[{name:"admin"},{name:"diseñador"}],
            reloadForm:false,
            password:"",
            isEdit:false,
            cambioPass:false
        }
    },
    methods:{
        filterUser(){
          if(this.userName!=""){
              let tempList=[]
              this.users.forEach(user=>{
                if(this.filterName(user.UserName)){
                    tempList.push(user)
                }
              });
            if(tempList.length==1){
              this.selectedUser=Object.assign({}, tempList[0])
              this.isWorkingInAnUser=true;
              this.isNew=false
            }else if(tempList.length==0){
              this.reloadForm=true;
              this.isEdit=false;
              this.isNew=true;
              this.isWorkingInAnUser=true;
              let user={
                UserName:this.userName,
                Password:"",
                Role:"",
                LastTime:"",
                Enterprises:[]
              }
              this.selectedUser=Object.assign({}, user);
              this.reloadForm=false;
            }else{
              this.usersList=tempList.slice(0,this.limit);
              this.numRegistros=this.usersList.length;
            }
          }else{
            this.usersList=this.users.slice(0,this.limit);
              this.numRegistros=this.usersList.length;
          }

        },
        filterName(user) {
          return (user.toUpperCase().indexOf(this.userName.toUpperCase()) > -1)
        },
        async getAllUser(){
              await this.axios
         .get("/api/users/getAllDomainUsers" )
         .then(rs => {
              let userList=[]
              let userTemp={}
              rs.data.forEach(user=>{
                userTemp=user
                userTemp.showEnterprises=[]
                if(user.Enterprises.length>0){
                  user.Enterprises.forEach(enterpriseId=>{
                    this.enterprises.forEach(enterprise=>{
                      if(enterpriseId==enterprise.EnterpriseId){
                        userTemp.showEnterprises.push(enterprise.Name)
                      }
                    })
                  });
                }
                userList.push(userTemp)
              })
              this.ultima=userList.length/this.limit;
              this.users=userList;
              this.usersList=userList.slice(0,this.limit);
              this.numRegistros=this.usersList.length;
              })
        },
        async getAllEnterprise(){
              await this.axios
         .get("/api/enterprise/getAll" )
         .then(rs => {
              this.enterprises=rs.data;
              })
            },
        reload(){
          if(!this.isEdit){
            let encontrado=false;
            this.users.forEach(user=>{
              if(user.UserName==this.selectedUser.UserName){
                encontrado=true;
                this.selectedUser=Object.assign({}, user);
                this.isNew=false;
              }
            })
            if(!encontrado){
              this.isNew=true;
              let user={
                UserName:this.selectedUser.UserName,
                Password:"",
                Role:"",
                LastTime:"",
                Enterprises:[]
              }
              this.selectedUser=Object.assign({}, user);
            }
          }
        },
        changeTable(tipo){
          if(tipo=="atras"){
            this.actual-=this.limit;
            this.pages--;
          }else if(tipo=="siguiente"){
            this.actual+=this.limit;
            this.pages++;
          }else{
            this.actual=0;
          }
          if(this.actual>0){
            this.usersList=this.users.slice(this.actual,(this.limit+1)*this.pages);
              this.numRegistros=this.usersList.length;
          }else{
            this.usersList=this.users.slice(0,this.limit);
              this.numRegistros=this.usersList.length;
          }
        },
          changeLimit(limit){
            this.limit=limit;
            changeTable("");
        },
        showUserForm(user){
          this.selectedUser=Object.assign({}, user);
          this.isWorkingInAnUser=true;
          this.isNew=false;
          this.isEdit=true;
        },
        back(){
          this.userName="";
          this.isWorkingInAnUser=false;
          //this.filterUser();
          this.isEdit=false;
          this.getAllUser();
          this.cambioPass=false;
        },
        formatEnterprise(){
          let enterprises=""
          this.selectedUser.showEnterprises.forEach((enterprise,index)=>{
            this.enterprises.forEach(ent=>{
              if(ent.Name==enterprise){
                enterprises+=ent.EnterpriseId
                if(index<this.selectedUser.showEnterprises.length-1){
                  enterprises+=";"
                }
              }
            })
          });
          this.selectedUser.Enterprises=enterprises;
        },
        async updateUser(){
            if(this.cambioPass&&this.selectedUser.Password.trim()==""){
              return this.alertInfo("La contraseña no puede estar vacia")
            }
            if(this.cambioPass&&this.selectedUser.Password.length<6){
              return this.alertInfo("La contraseña debe de tener un largo minimo de 6 caracteres")
            }
            if(this.cambioPass&&this.password!=this.selectedUser.Password){
              return this.alertInfo("La contraseña y su confimacion deben de ser iguales");
            }
            if(this.selectedUser.UserName.trim()==""){
              return this.alertInfo("El nombre no puede estar vacio");
            }
            if(this.selectedUser.Role==""){
              return this.alertInfo("El usuario debe tener un rol asignado");
            }
            if(!this.cambioPass){
              this.selectedUser.Password="";
            }
            this.formatEnterprise();
            await this.axios.post("/api/user/edit", this.selectedUser).then(rs => {
              status = rs.data;
              this.edit=false;
              if(status){
                this.back();
              }              
          });
        },
        async saveUser(){
          if(this.selectedUser.Password.trim()==""){
            return this.alertInfo("La contraseña no puede estar vacia");
          }
          if(this.selectedUser.UserName.trim()==""){
            return this.alertInfo("El nombre no puede estar vacio");
          }
          if(this.selectedUser.Password.length<6){
            return this.alertInfo("La contraseña debe de tener un largo minimo de 6 caracteres")
          }
          if(this.selectedUser.Role==""){
            return this.alertInfo("El usuario debe tener un rol asignado");
          }
          if(this.password!=this.selectedUser.Password){
            return this.alertInfo("La contraseña y su confimacion deben de ser iguales");
          }
          
          this.formatEnterprise();
          await this.axios.post("/api/user/saveForm", this.selectedUser).then(rs => {
            status = rs.data;
                this.edit=false;
              if(status){
                this.back();
              }         
          })
        },
        async deleteUser(){
            await this.axios.post("/api/user/delete", this.selectedUser).then(rs => {
              status = rs.data;
                this.edit=false;
              if(status){
                this.back();
              }         
          })
        },
        deleteEnterprise(enterprise){
          let temp=[];
          this.selectedUser.showEnterprises.forEach(name=>{
            if(name==enterprise){
                this.edit=true;
            }else{
              temp.push(name);
            }
          })
          this.selectedUser.showEnterprises=temp;
        },
        addEnterprise(){
          if(!this.selectedUser.showEnterprises){
            this.selectedUser.showEnterprises=[]
          }
          this.enterprises.forEach(enterprise=>{
            if(enterprise.EnterpriseId==this.selectedEnterprise){    
              if(!this.selectedUser.showEnterprises.includes(enterprise.Name)){
                this.selectedUser.showEnterprises.push(enterprise.Name)
                this.edit=true;
              }else{
                return this.alertInfo("El usuario ya se encuentra en dicha empresa");
              }
            }
          })
        }
    }   
}
</script>
<style scoped>
.seleccion:hover{
  background-color:#9eb6ce3d ;
}
.moveright{
  position: absolute;
  right: 0px;
}

.supercard{
    box-shadow: 0 -4px 22px rgba(204, 197, 185, 0.5);
}

.divoff
{
  pointer-events: none;
  opacity: 0.7;
}

.pagina{
  position: absolute;
    margin-top: 14px;
    margin-left: 14px;
}

.positionpadding{
  padding-left: 15px!important;
}

</style>
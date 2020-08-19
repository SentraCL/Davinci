<template>
<span>
    <m-dialog id="deleteDialog" :title="titleRecovery()" :show.sync="getAlert.show" :isClose.sync="getAlert.close">
            <span slot="dialog">
                <div>
                    Mientras la empresa <strong>{{selectedEnterprise.Name}}</strong> se encuentre inhabilitada, no podra ver todos los proyectos enlazados a esta.<br /> ¿Desea recuperar la empresa?
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDialog">
                        No , Cancelar
                    </d-button>
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="recoveryEnterprise">
                        Si, Recuperar
                    </d-button>
                </span>
            </span>
        </m-dialog>
    <h6 >
      <input-text v-on:input="filterEnterprise()" label="<span class='ti-home'></span> Crear Empresa" v-model="enterpriseName" :value.sync="enterpriseName" autocomplete="off"></input-text>
    </h6>
    <div class="card-group">
      <div class="col-lg-3" v-if="isWorkingInAEnterprise">
        <div class="card card-user">
          <div class="image">
            <img src="/img/toolbar.ad2e8fae.jpg" alt="...">
          </div><div class="content">
          <div class="author">
            <a href="#">
              <img v-if="selectedEnterprise.Avatar" :src="selectedEnterprise.Avatar" alt="..." class="avatar border-white">
              <img v-if="!selectedEnterprise.Avatar" src="@/assets/img/davinci-logo.png" alt="..." class="avatar border-white">
            </a>
            <h4 class="title">{{selectedEnterprise.Name}}<br>
              <span style="color: greenyellow;">
                <small>{{selectedEnterprise.Rut}}</small>
              </span>
            </h4>
          </div>
          <div class="col-lg-12">
            <span>
              <p :title="selectedEnterprise.Description" class="prettySummary">{{selectedEnterprise.Description}}</p>
              
            </span>
            <p></p>
            </div>
          </div>
          <hr>
          <div class="text-center fixpadding">
          <div class="col-lg-12">
            <div class="row">
            </div>
            <div class="row">
              <div class="col-md-6">
                
                </div>
                <div class="col-md-6">
                  
                </div>
              </div>
            </div>
          <br>
        </div>
        </div>
      </div>
      <div class="col-xl-8 col-lg-7 col-md-6" v-if="isWorkingInAEnterprise">
        <div class="card card">
          <div class="card-header">
            <h2  v-if="!isNew" class="card-title"><span class="ti-panel"></span> Modificar Empresa</h2>
            <h2  v-if="isNew" class="card-title"><span class="ti-panel"></span> Crear Empresa</h2>
          </div>
          <div class="card-body custom-card">
            <enterprise-form :enterprise="selectedEnterprise" :isNew="isNew" v-on:getAll="getAllEnterprise()" v-on:back="reload()"></enterprise-form>
          </div>
        </div>
        <div class="col-md-11">

        </div>
      </div>
      <div class="col-lg-12" v-if="!isWorkingInAEnterprise" >
        <enterprise-item :enterprises="enterprisesList" @selectedEnterprise="onChildClick" v-on:back="selectedEnterprise"></enterprise-item>
      </div>
    </div>
  </span>
</template>
<script>
    
import EnterpriseItem from "./EnterpriseItem.vue";
import EnterpriseForm from "./EnterpriseForm.vue";
export default {
    name: "Enterprise",
    components: {
        EnterpriseItem,
        EnterpriseForm
    },
    computed: { 
    },
    props: {
    },
    data() {
        this.getAllEnterprise()
        return{
            getAlert: {
                    mensaje: "",
                    tiempo: 10,
                    show: false
                },
            isWorkingInAEnterprise:false,
            enterprises: {},
            enterprisesList: {},
            selectedEnterprise: {},
            enterpriseName:"",
            isNew: false
        }
    },
    watch: {

    enterpriseName() {
      if(this.enterpriseName==""){
        this.reload();
      }
      }
    },
    methods:{
        titleRecovery(){
          let salida="";
          if (this.selectedEnterprise){
            salida=`<span class="fa fa-question-circle-o"></span> Confirmar restauración de la Empresa <strong>${this.selectedEnterprise.Name}</strong>`
          }else{
            salida="Confirmar la restauración de la Empresa"
          }
          return salida;
        },
        async getAllEnterprise(){
              await this.axios
         .get("/api/enterprise/getAll" )
         .then(rs => {
              this.enterprises=rs.data;
              this.enterprisesList=rs.data;
              })
            },
      equalsName(enterprise) {
        return (enterprise.name == this.enterpriseName)
      },
      filterName(enterprise) {
        return (enterprise.Name.toUpperCase().indexOf(this.enterpriseName.toUpperCase()) > -1)
      },
      onChildClick(value){
        this.selectedEnterprise=value;
        this.enterpriseName=value.Name
        if(this.selectedEnterprise.Status==1){
          this.isWorkingInAEnterprise=true;
          this.isNew=false;
          this.getAlert.show = false;
        }else{
          this.isWorkingInAEnterprise=false;
          this.isNew=false;
          this.getAlert.show = true;
          this.getAlert.close = true;
        }
      },
      reload(){
        this.isWorkingInAEnterprise=false
        this.enterpriseName=""
        this.enterprisesList=this.enterprises
        this.getAlert.show = false;
      },
      
      closeDialog() {
          this.getAlert.show = false;
      },
      async recoveryEnterprise() {
          var status = false;
          await this.axios.post("/api/enterprise/recovery/", this.selectedEnterprise).then(rs => {
              status = rs.data;
              this.getAllEnterprise();
              this.reload();
          });
      },

      filterEnterprise(){
        if(this.enterpriseName!=""){
          let tempEnterprises=[]
            this.enterprises.forEach(value=>{
              if(this.filterName(value)){
                tempEnterprises.push(value)
              }
          });
          this.enterprisesList=tempEnterprises
          if(this.enterprisesList.length==1){
            this.selectedEnterprise=this.enterprisesList[0]
            this.isWorkingInAEnterprise=true;
            this.isNew=false
          }else if(this.enterprisesList.length==0){
            this.isNew=true;
            this.isWorkingInAEnterprise=true;
            let enterprise={
              Name:this.enterpriseName,
              Rut:"",
              Direction:"",
              Description:"",
              Avatar:"",
              Status:1
            }
            this.selectedEnterprise=enterprise;
          }
        }else{
          this.reload()
        }
      }
    }
}
</script>
<style scoped>
  .fixpadding{
    padding-bottom: 25px;
  }
</style>
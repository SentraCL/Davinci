<template>
  <div class="container">
    <form-wizard color="#A47B31" title="Asistente de Importación de Proyecto"
      subtitle="Modulo de importacion de proyecto, donde se evalua y si hiciste bien tu pega.">
      <tab-content title="Validar Proyecto" :before-change="ready">
        <div class="large-12 medium-12 small-12 cell">
          <span v-if="nameExist">Importando proyecto para: <strong>{{nameProject}}</strong></span>
          <input v-if="step==0" accept=".dvc" type="file" id="DavinciFile" name="DavinciFile" ref="fileProject"
            v-on:input="handleFileUpload()" hidden />


          <div class="row">
            <div class="col-md-10" v-if="error">
              <h5><i class="ti-comment"></i> Tenemos que renombrar algunos inventos antes de continuar..</h5>
              <div class="error row" readonly v-if="error">
                <div class="col-md-4" v-for="(obj, key) of invToRename" :key="key">
                  <input-text removeSpace :label="key" v-model="invToRename[key]"></input-text>
                </div>
              </div>
            </div>

            <div v-if="projectOK && !error">
              <img class="avatar border-white" :src="importProject.project.Avatar64">
              {{importProject.project.Name}} listo!
            </div>

            <div class="col-md-2">
              <img src="@/assets/img/leo.png" v-if="error" alt=" Trata de hacerlo bien" width="128px" />
            </div>
            <div class="col-md-12">
              <h5 class="well" v-if="error">o bien subir otro archivo..</h5>
              <button class="btn btn-success" @click="openFile">
                <span style="color:blue" class="ti-upload"></span> {{labelUpload}}
              </button>


            </div>
          </div>


        </div>
      </tab-content>
      <tab-content title="Subir Proyecto" style="white-space: nowrap;">
        <span v-if="step==1" class="row">

          <div class="col-md-10" style="color: right;">
            <div class="row">
              <div class="col-4">
                <h4 style="overflow: hidden;text-overflow: ellipsis;"><i
                    class="ti-blackboard"></i>&nbsp;<strong>{{importProject.project.Name}}</strong></h4>
                <img class="avatar border-white" :src="importProject.project.Avatar64"><br />

              </div>
              <div class="col-sm-4">
                <i class="ti-user"></i>&nbsp;<strong>{{importProject.project.Author}}</strong><br />
                <i class="ti-email"></i>&nbsp;<strong>{{importProject.project.Administrator.Email}}</strong><br />

                <!--<input-text :label="nameProject" v-model="importProject.project.Name"></input-text><br />-->

                <!-- <i class="ti-briefcase"></i><strong>{{importProject.project.Enterprise}}</strong>-->


                <i class="ti-wand"></i>&nbsp; Inventos Cargados:
                {{Object.keys(importProject.project.Repository).length}}<br />
                <!--<p v-for="repo in importProject.project.Repository">
                  ({{Object.keys(repo.WareHouse).length}})
                </p>-->
                <i class="ti-user"></i>&nbsp; Usuarios Cargados:
                {{Object.keys(importProject.project.Users).length}}<br />

                <!--  <i class="ti-book"></i>&nbsp; Epicos Cargados:
                {{Object.keys(importProject.project.Epics.Types).length}}<br />
                <i class="ti-clipboard"></i>&nbsp; Historias de usuario Cargados:
                {{Object.keys(importProject.project.UserStories.Types).length}}<br />
              -->
              </div>
              <div class="col-md-12">





              </div>
              <div class="col-md-6">

              </div>
            </div>



            <!--<div class="col-md-12">
              <div class="form-group">
                <span v-html="importProject.project.Resume"></span>
              </div>
            </div>-->
          </div>
        </span>
        <!--<button class="btn btn-success" v-if="step==1" @click="finishStep">Cargar Archivo</button>-->
      </tab-content>
      <tab-content title="Finalisszar" style="white-space:nowrap;">
        <span v-if="step==1" class="row">

          <div class="col-4">
            <h4 style="overflow: hidden;text-overflow: ellipsis;"><i
                class="ti-blackboard"></i>&nbsp;<strong>{{importProject.project.Name}}</strong></h4>
            <img class="avatar border-white" :src="importProject.project.Avatar64"><br />

          </div>

          <div class="col-md-6"
            style="white-space:pre-line; padding:10px;border:4px double #A47B31; border-radius:16px 40px 0px 32px;">
            <strong>
              <p>{{importProject.project.Resume}}</p>
            </strong>
          </div>

          <div style="margin-top: 20px;" class="col-sm-4">
            <h5><i class="ti-comment"></i>&nbsp;Este contenido sera insertado en el sistema</h5>

            <i class="ti-user"></i>&nbsp; Total de registros de Inventos:
            {{inventsLength}}
            <!--<strong  v-for="repo in importProject.project.Repository">({{Object.keys(repo.WareHouse).lenght}})</strong>--><br />

            <i class="ti-book"></i>&nbsp; Epicos Cargados:
            {{epicLength}}<br />
            <i class="ti-clipboard"></i>&nbsp; Historias de usuario Cargados:
            {{userStorieLength}}<br />
            <!-- {{Object.keys(importProject.project.Epics.Types)}}-->
            <!--<p v-for="repo in importProject.project.Repository">
                  ({{Object.keys(repo.WareHouse).length}})
                </p>-->

          </div>





          <!--<div class="col-md-12">
              <div class="form-group">
                <span v-html="importProject.project.Resume"></span>
              </div>
            </div>-->

        </span>

      </tab-content>
      <button class="btn btn-info" type="primary" slot="finish" @click="finishStep()">Importar</button>
      <button class="btn btn-info" type="primary" slot="next" @click="nextStep()">Siguiente</button>
      <button class="btn btn-info" type="primary" slot="prev" @click="backStep">Volver</button>
    </form-wizard>
    <!--

    -->
  </div>
</template>

<script>
  export default {
    props: {
      nameProject: String,
      isNew: Boolean
    },
    data() {
      return {
        projectOK: false,
        file: "",
        doContinue: false,
        step: 0,
        labelUpload: "Subir Proyecto",
        error: false,
        errorMessage: "",
        importProject: {},
        invToRename: {},
        nameExist: !this.isEmptyOrSpaces(this.nameProject),
        invents: {},
        epicLength: 0,
        userStorieLength: 0,
        inventsLength: 0

      };
    },

    methods: {
      ready() {
        return this.doContinue;
      },
      close() {
      },
      openFile() {
        this.invToRename = {}
        this.error = false;
        if (this.step == 0) {
          this.$refs.fileProject.value = null;
          document.getElementById("DavinciFile").click();
        }

      },

      async backStep() {
        if (this.step == 1) {
          this.error = false;
          this.projectOK = false;
        }
      },

      isDiffInv(invA, invB) {
        console.log("isDiffInv");
        return !((invA.code == invB.code) &&
          (invA.name == invB.name) &&
          (JSON.stringify(invA.artifacts) == JSON.stringify(invB.artifacts)) &&
          (invA.keyLabel == invB.keyLabel) &&
          (invA.keyValue == invB.keyValue));
      },

      async nextStep() {
        if (this.step == 0) {
          // this.doContinue=false;          
          //Begin : Validaciones
          if (!this.$refs.fileProject.files[0]) {
            this.alertError(
              "No ha seleccionado Archivo."
            );
            this.error = true;
          } else {
            if (Object.keys(this.invToRename).length > 0) {
              this.error = false;
              Object.keys(this.invToRename).forEach(key => {
                if (Object.keys(this.invToRename).indexOf(this.invToRename[key]) > -1) {
                  if (this.isDiffInvent(this.invToRename[key])) {
                    this.alertError(
                      `No renombraste ${this.invToRename[key]}`
                    );
                    this.error = true;
                  }
                }
              });
              if (!this.error) {
                for (var i = 0; i < this.importProject.inventions.length; i++) {
                  if (Object.keys(this.invToRename).indexOf(this.importProject.inventions[i].name) > -1) {
                    this.importProject.inventions[i].name = this.invToRename[this.importProject.inventions[i].name];
                    if (this.isEmptyOrSpaces(this.importProject.inventions[i].code)) {
                      this.importProject.inventions[i].code = await this.EncriptCode(this.importProject.inventions[i].name);
                    }
                  }
                }

                for (var r = 0; r < this.importProject.project.Repository.length; r++) {
                  if (Object.keys(this.invToRename).indexOf(this.importProject.project.Repository[r].Invention) > -1) {
                    this.importProject.project.Repository[r].Invention = this.invToRename[this.importProject.project.Repository[r].Invention];
                    if (this.isEmptyOrSpaces(this.importProject.project.Repository[r].InventionCode)) {
                      this.importProject.project.Repository[r].InventionCode = await this.EncriptCode(this.importProject.project.Repository[r].Invention);
                    }
                  }
                }
              }
            } else {
              this.error = false;
            }
          }

          this.doContinue = !this.error;

          //End : Validaciones
          if (this.doContinue) {
            this.step = 1;
            if (this.nameExist) {
              this.importProject.project.Name = this.nameProject;
            }
          }
        }
        if (this.step == 1) {


          this.epicLength = Object.keys(this.importProject.project.Epics.Types).length;
          this.userStorieLength = Object.keys(this.importProject.project.UserStories.Types).length;
          var count = 0;

          this.importProject.project.Repository.forEach(repo => {
            count = count + Object.keys(repo.WareHouse).length;
          })

          /*
          for (var i=0;i < (this.importProject.project.Repository).length; i++){
               count = count + Object.keys(this.importProject.project.Repository[i].WareHouse).length;
          }
          */
          this.inventsLength = count;
        }
      },
      async finishStep() {
        var me = this;
        this.importProject.project.Alias = this.importProject.project.Name.replace(/ /g, "");
        var dcode = await this.EncriptCode(this.importProject.project.Alias );;        
        this.importProject.project.Code = dcode;                
        await this.axios
          .post("/api/project/import/", this.importProject).then(function (rs) {
            me.alertInfo(
              "Proyecto " + rs.data + " Importado."
            );
            me.$emit("importSuccess");
          })
          .catch(function () {
            me.alertError(
              "Ocurrio un error Terrible llame urgentemente a Nika"
            );
          });

      },

      handleFileUpload(event) {
        this.file = this.$refs.fileProject.files[0];
        this.labelUpload = "Desea subir otro proyecto.";
        if (this.file) {
          var readFile = new FileReader();
          this.error = false;
          var me = this;
          readFile.onload = async function (event) {
            var content = event.target.result;
            try {
              var jsonImport = JSON.parse(content);
            } catch{
              me.alertError(`El Formato no es el correcto.`);
              return false;
            }
            //TODO: Ser mas eficiente al consultar si el JSON es o no Un proyecto
            me.projectOK = (Object.keys(jsonImport).indexOf("project") > -1 && Object.keys(jsonImport).indexOf("inventions") > -1)
            if (!me.projectOK) {
              me.alertError(`El Formato no es el correcto.`);
              return false;
            }

            me.alertInfo(`Evaluando proyecto ${jsonImport.project.Name}`);
            me.importProject = jsonImport;
            me.invToRename = {}
            //Itera json.inventions , y valida si existen en nuestro ambiente
            //Compara ese listado con el GET que te retorne esta url /api/invention/all/
            //rs.data[] == json.inventions
            await me.axios.get("/api/invention/all/").then(rs => {
              var jsonInvents = rs.data;

              for (var i = 0; i < jsonInvents.length; i++) {
                for (var y = 0; y < jsonImport.inventions.length; y++) {
                  if (jsonInvents[i].code == jsonImport.inventions[y].code && me.isDiffInv(jsonInvents[i], jsonImport.inventions[y])) {
                    console.log(jsonInvents[i].name);
                    me.invToRename[jsonInvents[i].name] = jsonInvents[i].name;
                    me.error = true;
                    /*
                    me.errorMessage += `Invento ${jsonInvents[i].name} ya existe debe ser renombrado<br/>`// "Invento " + jsonInvents[i].name + " ya existe..\n"
                    
                    me.alertInfo(
                      "Invento " + jsonInvents[i].name + " ya existe.."
                    );
                    */
                  }
                }
              }
            });
          };
          readFile.readAsText(this.file);
        }
        this.step = this.error ? 1 : 0;
      }
    }
  };
</script>
<style scoped>
  .error {
    width: 100%;
    border: solid thin #fff;
    background-color: #fff;
    overflow-y: scroll;
    overflow-x: hidden;
    max-height: 450px;
  }
</style>
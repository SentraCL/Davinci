<template>
  <div class="container">
    <form-wizard title="Asistente de Importación de Proyecto" subtitle="Modulo de importacion de proyecto, donde se evalua y si hiciste bien tu pega.">
      <tab-content title="Validar Proyecto" :before-change="nextStep">
        <div class="large-12 medium-12 small-12 cell">

          <input v-if="step==0" accept=".dvc" type="file" id="DavinciFile" name="DavinciFile" ref="fileProject" v-on:input="handleFileUpload()" hidden />


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
      <tab-content title="Subir Proyecto">
        <span v-if="step==1" class="row">

          <div class="col-md-10">
            <div class="row">
              <div class="col-md-2">
                <img class="avatar border-white" :src="importProject.project.Avatar64">
              </div>
              <div class="col-md-10">
                <h4>{{importProject.project.Name}}</h4>
                <strong>{{importProject.project.Author}}</strong><br />
                {{importProject.project.Administrator.FullName}}<br />
                <i class="ti-email"></i> {{importProject.project.Administrator.Email}}<br />
                {{importProject.project.Company}}
              </div>
              <div class="col-md-6">
                <!--http://jsonviewer.stack.hu/ -->
                <h5>Inventos Cargados..</h5>
                <span v-for="repo in importProject.project.Repository">
                  ({{Object.keys(repo.WareHouse).length}}) <div>{{repo.Invention}}</div>
                </span>
                <h5>Usuarios Cargados..</h5>
              </div>
              <div class="col-md-6">

              </div>
            </div>



            <div class="col-md-12">
              <div class="form-group">
                <span v-html="importProject.project.Resume"></span>
              </div>
            </div>
          </div>
        </span>
        <!--<button class="btn btn-success" v-if="step==1" @click="submitFile">Cargar Archivo</button>-->
      </tab-content>
      <tab-content title="Finalizar"></tab-content>
      <button class="btn btn-info" type="primary" slot="next" @click="nextStep">Siguiente</button>
      <button class="btn btn-info" type="primary" slot="prev" @click="backStep">Volver</button>
    </form-wizard>
    <!--

    -->
  </div>
</template>

<script>
  export default {
    data() {
      return {
        projectOK: false,
        file: "",
        step: 0,
        labelUpload: "Subir Proyecto",
        error: false,
        errorMessage: "",
        importProject: {},
        invToRename: {}

      };
    },

    methods: {
      close() {
      },
      openFile() {
        this.invToRename = {}
        this.error = false;
        this.$refs.fileProject.value = null;
        document.getElementById("DavinciFile").click();
      },
      backStep() {
        this.step = this.step - 1;
        if (this.step == 0) {
          this.error = false;
          this.projectOK = false;
        }
      },
      async nextStep() {
        if (this.step == 0) {
          this.error = false;
          if (!this.$refs.fileProject.files[0]) {
            this.alertError(
              "No ha seleccionado Archivo."
            );
            return false;
          }
          //Si hay elementos que renombre entre al IF
          if (Object.keys(this.invToRename).length > 0) { //+            
            Object.keys(this.invToRename).forEach(key => {
              if (Object.keys(this.invToRename).indexOf(this.invToRename[key]) > -1) {
                this.alertError(
                  `No renombraste ${this.invToRename[key]}`
                );
                this.error = true;
              }
            })
          }
          if (!this.error) {
            //ESta instancia donde todos los nombres repetidos fueron corregidos            

            for (var i = 0; i < this.importProject.inventions.length; i++) {
              if (Object.keys(this.invToRename).indexOf(this.importProject.inventions[i].name) > -1) {
                this.importProject.inventions[i].name = this.invToRename[this.importProject.inventions[i].name];
                this.importProject.inventions[i].code = await this.EncriptCode(this.importProject.inventions[i].name);
              }
            }

            for (var r = 0; r < this.importProject.project.Repository.length; r++) {
              if (Object.keys(this.invToRename).indexOf(this.importProject.project.Repository[r].Invention) > -1) {
                this.importProject.project.Repository[r].Invention = this.invToRename[this.importProject.project.Repository[r].Invention];
                this.importProject.project.Repository[r].InventionCode = await this.EncriptCode(this.importProject.project.Repository[r].Invention);
              }
            }

            console.log(JSON.stringify(this.importProject));
            this.step = 1;
          }
          return !this.error;
        }
        return true;
      },
      submitFile() {
        let formData = new FormData();
        if (!this.error) {
          formData.append("DavinciFile", this.file);
          this.axios
            .post("/api/project/import/", formData, {
              headers: {
                "Content-Type": "multipart/form-data"
              }
            })
            .then(function (rs) {
              console.log("SUCCESS!!" + rs.data);
            })
            .catch(function () {
              console.log("FAILURE!!");
            });
        } else {
          this.alertError(
            "Hemos encontrado algunos erroes"
          );
        }
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
            var jsonImport = JSON.parse(content);
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
                  if (jsonInvents[i].code == jsonImport.inventions[y].code) {
                    console.log(jsonInvents[i].name);
                    me.invToRename[jsonInvents[i].name] = jsonInvents[i].name;
                    me.error = true;
                    me.errorMessage += `Invento ${jsonInvents[i].name} ya existe debe ser renombrado<br/>`// "Invento " + jsonInvents[i].name + " ya existe..\n"
                    /*
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
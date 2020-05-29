<template>
  <div class="container">
    <a class="btn btn-xs e-btn" title="Cerrar" @click="close"><i class="ti-close"></i></a>
    <form-wizard v-if="showWizard" title="Asistente de Importación de Proyecto" subtitle="Modulo de importacion de proyecto, donde se evalua y si hiciste bien tu pega.">
      <tab-content title="Validar Proyecto" :before-change="checkFile">
        <div class="large-12 medium-12 small-12 cell">

          <input v-if="step==0" accept=".dvc" type="file" id="DavinciFile" name="DavinciFile" ref="fileProject" v-on:input="handleFileUpload()" hidden />


          <div class="row">
            <div class="col-md-10">
              <h5 v-if="error">Vaia vaia parece que se equivoco!, para continuar
                usted deberia renombrar los siguiente
                inventos.</h5>

              <hr />
              <div class="error" readonly v-if="error">
                <span v-for="(obj, key) of invToRename" :key="key" class="col-md-2">
                  <input-text removeSpace :label="key" v-model="invToRename[key]"></input-text><br />
                </span>
              </div>
            </div>
            <div class="col-md-2">
              <img src="@/assets/img/leo.png" v-if="error" alt=" Trata de hacerlo bien" width="128px" />
            </div>
            <div class="col-md-12">
              <h5 class="well" v-if="error">o bien subir otro archivo..</h5>
              <button class="btn btn-success" v-if="step==0" @click="openFile">
                <span style="color:blue" class="ti-upload"></span> {{labelUpload}}
              </button>
              <button class="btn btn-success" v-if="step==1" @click="submitFile">Cargar Archivo</button>

            </div>
          </div>


        </div>
      </tab-content>
      <tab-content title="Subir Proyecto" :before-change="reset"></tab-content>
      <tab-content title="Finalizar"></tab-content>
      <button class="btn btn-info" type="primary" slot="next" @click="checkFile">Siguiente</button>
      <button class="btn btn-info" type="primary" slot="prev" @click="reset">Volver</button>
    </form-wizard>
    <!--

    -->
  </div>
</template>

<script>
  export default {
    /*
          Defines the data used by the component
        */
    data() {
      return {
        file: "",
        step: 0,
        xfile: {},
        showWizard: true,
        labelUpload: "Subir Proyecto",
        error: false,
        errorMessage: "",

        invToRename: {}

      };
    },

    methods: {
      close() {
        this.showWizard = false;
      },
      openFile() {

        this.invToRename = {}
        this.error = false;
        this.$refs.fileProject.value = null;
        document.getElementById("DavinciFile").click();
      },
      reset() {
        this.$refs.fileProject.value = null;
        this.alertInfo("Me ejecute!");
        console.log("ye")
        return true;
      },
      checkFile() {

        if (!this.$refs.fileProject.files[0]) {
          this.alertError(
            "No ha seleccionado Archivo."
          );
          return false;
        }
        //Si hay elementos que renombre entre al IF
        if (Object.keys(this.invToRename).length > 0) { //+
          this.error = false;
          Object.keys(this.invToRename).forEach(key => {
            if (Object.keys(this.invToRename).indexOf(this.invToRename[key]) > -1) {
              this.alertError(
                `No renombraste ${this.invToRename[key]}`
              );
              this.error = true;
            }
          })

        }

        return !this.error;
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
        this.alertInfo("Intento");
        this.file = this.$refs.fileProject.files[0];
        if (this.file) {
          var readFile = new FileReader();
          this.error = false;
          var me = this;
          readFile.onload = async function (event) {
            var content = event.target.result;
            var jsonImport = JSON.parse(content);
            me.alertInfo(`Evaluando proyecto ${jsonImport.project.Name}`);

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
                    me.labelUpload = "Desea subir otro proyecto.";
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
    height: 300px;
    width: 100%;
    border: solid thin #fff;
    background-color: #a1a1a1;
    overflow-y: scroll;
    overflow-x: hidden;
    max-height: 450px;
  }
</style>
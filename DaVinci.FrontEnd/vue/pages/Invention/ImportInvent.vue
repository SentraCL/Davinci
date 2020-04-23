<template>
  <div class="row">
    <div class="col-md-12" v-if="csvtoJson.length==0 && stepIndex==0">
      <upload-text @upload="contentCsv = $event" label="IMPORTAR CSV" accept=".csv"></upload-text>
      <upload-text @upload="contentJson = $event" label="IMPORTAR JSON" accept=".json"></upload-text>
      <hr />
      <small><span class="ti-alert"></span> Recuerde que si carga un {{name}} existente, este actualizara.</small>
    </div>
    <div class="col-md-12" v-if="csvtoJson.length>0 && stepIndex==0">
      <button class="btn btn-success" @click="toImport"><span style="color:blue" class="ti-upload"></span> Cargar {{name}} </button>
      <button class="btn btn-info" @click="cancel"> Cancelar </button>

    </div>
    <div class="col-md-12">
      </br>
    </div>

    <div class="col-md-12" v-if="upOK">
      <davinci-table :pagination="page" v-if="csvtoJson.length>0" :data.sync="csvtoJson"></davinci-table>
      <progress v-if="stepIndex==1" :value="progressImport" max="100" class="progress-bar-success"></progress>
      <small v-if="obs && stepIndex==0"><span class="ti-eye" style="color:red"></span> {{obs}}</small>
    </div>

    <!-- Alertas de error -->
    <m-dialog id="alertDialog" :title="alertDialog.title" :show.sync="alertDialog.show" :isClose.sync="alertDialog.close">
      <span slot="dialog">
        <div class="row">
          <span class="col-md-6">
            <span v-html="alertDialog.html"></span>
          </span>
          <span class="col-md-6">
            <img src="@/assets/img/leo.png">
          </span>
        </div>
      </span>
      <span slot="actions">
        <span class="btn-group">
          <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDialog">
            Cerrar
          </d-button>
        </span>
      </span>
    </m-dialog>

  </div>


</template>

<script>
  import UploadText from "./UploadTextFile";
  import io from 'socket.io-client'

  export default {
    name: "import-invent",
    data() {
      return {
        socket: io(`${window.location.origin}`, {
          forceNew: true
        }),
        alertDialog: {
          show: false,
          close: false,
          title: "",
          html: ""
        },
        contentCsv: "",
        contentJson: "",
        csvtoJson: [],
        page: {
          itemsPerPage: 6
        },
        obs: "",
        columns: [],
        columnsUp: [],
        upOK: false,
        stepIndex: 0,
        progressImport: 0,
        progressMax: 0,
        upCount: 0,
      }
    },
    props: {
      project: {},
      format: {},
      code: String,
      name: String
    },
    mounted() {
      for (var a in this.format.artifacts) {
        var c = this.format.artifacts[a];
        this.columns.push(c.name);
        this.columnsUp.push(c.name.toUpperCase());
      }

    },
    watch: {
      stepIndex(stepN, stepO) {
        if (stepN == 2) {
          this.$emit("finish");
        }
      },
      contentCsv(newValue, oldValue) {
        this.csvtoJson = [];
        var lines = newValue.split('\n')
        var index_columns = Object.keys(lines[0].split(';'));
        var index = 0;
        for (var l in lines) {
          var line = lines[l];
          var prettyArtifact = {}
          for (var k in line.split(';')) {
            var data = line.split(';')[k];
            //SOlo Mostrar columnas que si existen como referencias al invento            
            if (this.columnsUp.indexOf(lines[0].split(';')[k].toUpperCase()) > -1) {
              prettyArtifact[lines[0].split(';')[k]] = data
            } else {
              if (lines[0].split(';')[k] != "") {
                this.obs = "Atención, el archivo cargado tenias mas columnas de las necesarias, considere si efectivamente es el documento que desea cargar.."
              }
            }
          }
          //prettyArtifact[`<span class="ti-time"></span>`] = "Pendiente.."
          if (index > 0) {
            if (prettyArtifact[this.format.keyValue] != "") {
              this.csvtoJson.push(prettyArtifact)
            }

          }
          index++;
        }
        this.checkImport();
      },
      contentJson(newValue, oldValue) {
        this.csvtoJson = [];
        try {
          var artifacts = JSON.parse(newValue);
          for (var a in artifacts) {
            var prettyArtifact = {}
            var artifact = artifacts[a];
            for (var k in Object.keys(artifact)) {
              var key = Object.keys(artifact)[k]
              if (key.endsWith("_date")) {
                key = "Fecha";
              }
              if (key != "") {
                if (this.columnsUp.indexOf(key.toUpperCase()) > -1) {
                  var value = (artifact[key] == "1" || artifact[key] == "0") ? artifact[key] == "1" ? "Si" : "No" : artifact[key];
                  prettyArtifact[key] = value
                } else {
                  this.obs = "el Archivo cargado tenias mas columnas de las necesarias, considere si efectivamente es el documento que desea cargar.."
                }
              }
            }
            //prettyArtifact[`<span class="ti-time"></span>`] = "Pendiente.."
            if (prettyArtifact[this.format.keyValue] != "") {
              this.csvtoJson.push(prettyArtifact)
            }
          }
          this.checkImport();
        }
        catch (error) {
          //console.error(error);
          // expected output: ReferenceError: nonExistentFunction is not defined
          // Note - error messages will vary depending on browser
        }

      }
    },

    components: {
      UploadText,

    },
    methods: {
      closeDialog() {
        this.cancel();
        this.alertDialog.show = false;
      },
      checkImport() {
        this.upOK = false;
        if (this.csvtoJson.length > 0) {
          //Tiene las columnas del invento.
          var count = 0;
          var upColumns = [];
          for (var c in this.csvtoJson[0]) {
            upColumns.push(c);
            this.columns.forEach(column => {
              if (column.toUpperCase() == c.toUpperCase()) {
                count++
              }
            });
          }
          if (count != this.columns.length) {
            this.alertDialog.show = true;
            this.alertDialog.close = false;
            this.alertDialog.title = `<strong><span class="ti-face-sad"></span> No es el Formato indicado.</strong>`
            this.alertDialog.html = `El archivo subido no posee las columnas necesarias para la carga masiva:<span style="color:green"> <br/> ${this.columns} <hr/></span>Usted Intento subir lo siguiente:<span style="color:red"> <br/> ${upColumns} <hr/></span>`
            return false;
          }
          //Existen codigo repetidos en la carga
          var keyValue = this.format.keyValue;
          var keyCollector = []
          var keyDuplicate = {};
          for (var i in this.csvtoJson) {
            var invention = this.csvtoJson[i];
            var key = invention[keyValue] === undefined ? invention[keyValue.toUpperCase()] : invention[keyValue];
            if (keyCollector.indexOf(key) == -1) {
              keyCollector.push(key);
            } else {
              keyDuplicate[key] = invention;
            }
          }
          if (Object.keys(keyDuplicate).length > 0) {
            this.alertDialog.show = true;
            this.alertDialog.close = false;
            this.alertDialog.title = `<strong><span class="ti-face-sad"></span> Existen incongruencia de datos.</strong>`
            this.alertDialog.html = `El archivo subido posee campos claves repetidos <br/><span style="color:red">${Object.keys(keyDuplicate)}</span></span>`
            return false;
          }
          this.upOK = true;
        }
        return true;

      },
      cancel() {
        this.csvtoJson = [];
        this.contentCsv = "";
        this.contentJson = "";
      },
      toInventionRequestDTO(artifacts) {
        var requestDTO = {};
        for (var a in artifacts) {
          var artifact = artifacts[a]
          var key = "";
          //Da el formato correcto de case sensity..
          this.columns.forEach(column => {
            if (column.toUpperCase() == a.toUpperCase()) {
              key = column;
            }
          });
          if (key != "") {
            //Convierte un JSON comun, al formato que recepcion el back-end.
            if (artifact.constructor === Array) {
              requestDTO[key] = artifact;
            } else {
              requestDTO[key] = [];
              var value = artifact.constructor === String ? artifact : JSON.stringify(artifact)
              requestDTO[key].push(value)
            }
          }
        }
        return requestDTO;
      },
      toImport() {
        this.stepIndex = 1;
        this.progressMax = this.csvtoJson.length;
        this.upCount = 0;
        for (var i in this.csvtoJson) {
          var preDto = this.toInventionRequestDTO(this.csvtoJson[i]);
          if (preDto[this.format.keyValue] != "") {
            var dto = JSON.stringify(preDto);
            this.socket.emit("import", dto, this.project.code, this.format.code);
          }
        }
        this.csvtoJson = [];
        var me = this;
        this.socket.on("importResult", function (response) {
          me.upCount++;
          me.progressImport = Math.trunc((me.upCount * 100) / me.progressMax);
          //console.log(me.progressImport);
          var jsonResponse = JSON.parse(response)
          var content = JSON.parse(jsonResponse.Container[0].ContentJSON);
          var jsonArtifact = {}
          for (var k in content) {
            if (content[k].constructor === Array && content[k].length == 1) {
              jsonArtifact[k] = content[k][0];
            } else {
              jsonArtifact[k] = content[k];
            }
            jsonArtifact["Resultado"] = `<span style="color:green" class="ti-thumb-up"></span>`;
          }
          me.csvtoJson.unshift(jsonArtifact);
          me.stepIndex = me.progressImport == 100 ? 2 : 1;
        })

      }
    },
  };
</script>
<style scoped>
  progress[value] {
    width: 100%;
    height: 20px;
  }
</style>
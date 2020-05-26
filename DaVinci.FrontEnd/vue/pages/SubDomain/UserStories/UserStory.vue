<template>
  <card v-if="userStoryForm" style="min-height: auto;background: var(--transparent);z-index:0 !important;" :id="idHTML">
    <template slot="header">
      <div class="topnav">
        <div class="user-stories-title" :name="userStoryCode">
          <span class="us-title">
            <span  v-if="userStoryCode==''" v-html="userStoryForm.title"></span>
            {{userStoryCode}}
          </span>
          <small id="vers" class="us-ver"> 
            <combo-simple  v-if="isUserStoryExist && versions.length>0" v-on:change="changeVersion()" :showKey="true"
              label class="input-item" :list="versions" keyValue="version" keyLabel="ref" :value.sync="currentVer">
            </combo-simple>
          </small>
          <a class="btn btn-xs us-btn" title="Cerrar" @click="remove">
            <i class="ti-close"></i>
          </a>
          <a class="btn btn-xs us-btn" v-if="currentVer!=lastVersion" title="No es posible Editar">
            <i class="ti-lock"></i>
          </a>
          <a class="btn btn-xs us-btn" v-if="isUserStoryExist && currentVer==lastVersion && versions.length>0"
            title="Guardar" @click="save">
            <i class="ti-save"></i>
          </a>
          <a class="btn btn-xs us-btn" v-if="currentVer==lastVersion" title="Crear Nueva Version" @click="create">
            <i class="ti-wand"></i>
          </a>
        </div>
      </div>
    </template>
    <div style="padding-left:16px">
      <h-tabs :tabs="UsTabs" viewScoped v-if="showForm" v-on:loadTab="similarUserStories()">
        <span slot="dat">
          <h5>Descripcion del {{userStory.usType.title}}</h5>
          <textarea :disabled="currentVer!=lastVersion" class="ta8" placeholder="Descripcion" id="txtDescripcion"
            v-model="userStoryForm.fields.description"></textarea>
          <br />
          <br />
          <h5>Resultado Esperado del {{userStory.usType.title}}</h5>
          <textarea :disabled="currentVer!=lastVersion" class="ta8" placeholder="Resultado Espeado del Caso"
            id="txtResultadoHU" v-model="userStoryForm.fields.expectedResult"></textarea>
          <small v-if="isUserStoryExist" style="position: relative;float:right">
            <i>Ultima Modificación {{dateFormat}}</i>
          </small>
        </span>
        <span slot="pre">
          <div style="min-height: 300px;">
            <pre-condition v-on:update="similarUserStories()" :inactive="currentVer!=lastVersion"
              :codeFormat.sync="userStory.usType.codeFormat" :preconditions.sync="userStoryForm.preConditions"
              :values.sync="values"></pre-condition>
            <hr />
            <div v-if="loadSimil & userStoriesSimil.length>0">
              <h5>Historias de Usuario Similares...</h5>
              <user-stories :data="userStoriesSimil" v-on:select="getUserStory()" :idUS.sync="idUS"
                v-on:found="sendToWorkpace()"></user-stories>
            </div>
            <div v-if="!loadSimil && userStoriesSimil.length==0">
              <i class="ti-thumb-up"></i> No se registran registros similares..
            </div>
          </div>
        </span>
        <span slot="step">
          <div style="min-height: 300px;">
            <us-step :inactive="currentVer!=lastVersion" :steps.sync="steps"></us-step>
          </div>
        </span>
        <span slot="epic">
          Pendiente..
          <br />
          <!--
          <vue-json-pretty :path="'res'" :data="userStoryForm.us">
          </vue-json-pretty>
          -->
        </span>
      </h-tabs>
      <span class="btn btn-xs btn-success" @click="exportXML">Exportar XML</span>
    </div>
  </card>
</template>
<script>
  import UserStories from "./UserStories.vue";
  import PreCondition from "./Preconditions.vue";
  import UsStep from "./UserStoriesSteps.vue";
  import VueJsonPretty from "vue-json-pretty";

  export default {
    name: "user-story",
    components: {
      PreCondition,
      UserStories,
      UsStep,
      VueJsonPretty
    },
    computed: {},
    props: {
      userStory: {} // Posee el contexto de la historia de usuario.
    },

    watch: {
      values(newValue, oldValue) {
        this.similarUserStories("");
      }
    },

    data() {
      var _userStoryCode = "";
      var _userStoryForm = {};
      var _codeOK = false;
      var _steps = [];
      _userStoryForm = this.cloneObject(this.userStory);
      if (Object.keys(this.userStory).indexOf("us") > -1) {
        _userStoryCode = this.userStory.us.id;
        _steps = this.cloneObject(this.userStory.steps);
        _codeOK = true;
      }
      return {
        CREATE: 0,
        SAVE: 1,

        currentVer: "",
        lastVersion: "",
        showForm: true,

        codeOK: _codeOK,

        versions: [],
        keys: [],

        idHTML: "",
        isUserStoryExist: _codeOK,
        steps: _steps,
        idUS: "",
        dateFormat: "",
        userStoryForm: _userStoryForm,
        loadSimil: false,
        codeSimil: "",
        userStoryCode: _userStoryCode,
        userStoriesSimil: [],
        values: {},
        UsTabs: [
          {
            title: `<i class="ti-clipboard"></i> Atributos`,
            slot: "dat",
            id: 0
          },
          {
            title: `<i class="ti-check-box"></i> Pre-condicion`,
            slot: "pre",
            id: 1
          },
          {
            title: `<i class="ti-menu-alt"></i> Paso a paso`,
            slot: "step",
            id: 2
          }
        ]
      };
    },

    mounted() {
      this.userStoryForm = this.cloneObject(this.userStory);
      //console.log(JSON.stringify(this.userStoryForm.usType.codeFormat));
      this.keys = this.userStoryForm.usType.codeFormat;

      this.userStoryForm.preConditions.forEach(pre => {
        this.values[pre.id] = pre.value;
      });

      this.idHTML = "US." + this.makeId(10);

      if (this.isUserStoryExist) {
        var lastIndex = this.userStoryForm.us.versions.length - 1;
        this.dateFormat = this.formatDays(
          this.userStoryForm.us.versions[lastIndex].date
        );
        this.idHTML = this.userStoryForm.us.id;
        for (var v in this.userStoryForm.us.versions) {
          var ver = this.userStoryForm.us.versions[v];
          var reference = `Creada por ${ver.author}, ${this.formatDays(
            ver.date
          )}`;
          this.versions.push({
            version: ver.version,
            ref: reference
          });
          this.currentVer = ver.version;
        }
        this.lastVersion = this.currentVer;

        this.UsTabs.push({
          title: `<i class="ti-folder"></i> Epicos`,
          slot: "epic",
          id: 3
        });
      }
    },

    methods: {
      exportXML() {      
        var  preConditions = ``;
        var lastVersionIndex = this.userStory.us.lastVersionIndex;
        var num = this.currentVer.match(/\d+/g);
        var currentVer = num[0] - 1;
        var userStoryExport = this.userStory.us.versions[currentVer];
        userStoryExport.preConditions.forEach(pc=>{
          preConditions += `${pc.name} : ${pc.value} \n`
        });
        
        var xml_steps = "";

        userStoryExport.steps.forEach(step=>{
          var execution_type = "1"      ;
          var xml_step = `<step>
            <step_number><![CDATA[${step.id}]]></step_number>
            <actions>${step.step}
              
            </actions>
            <expectedresults><![CDATA[${step.result}]]></expectedresults>
            <execution_type><![CDATA[${execution_type}]]></execution_type>
          </step>`;
          xml_steps += xml_step;
        });
        var status = 1;
        var xmltext = `<?xml version="1.0" encoding="UTF-8"?>
        <testcases>
          <testcase internalid="${this.userStory.title}" name="${this.userStory.title}">
            <node_order><![CDATA[1]]></node_order>
            <externalid><![CDATA[1776]]></externalid>
            <version><![CDATA[${userStoryExport.version}]]></version>
            <summary><![CDATA[${userStoryExport.fields.description}]]></summary>
            <preconditions><![CDATA[${preConditions}]]></preconditions>
            <execution_type><![CDATA[2]]></execution_type>
            <importance><![CDATA[3]]></importance>
            <estimated_exec_duration></estimated_exec_duration>
            <status>${status}</status>
            <is_open>1</is_open>
            <active>1</active>
          <steps>${xml_steps}
        </steps>
        <custom_fields>
          <custom_field>
            <name><![CDATA[@key]]></name>
            <value><![CDATA[@value]]></value>
          </custom_field>	
        </custom_fields></testcase>
        </testcases>`;

        var pom = document.createElement("a");

        var filename = `${this.userStory.title}.${this.currentVer}.xml`;
        var pom = document.createElement("a");
        var bb = new Blob([xmltext], { type: "text/plain" });
        pom.setAttribute("href", window.URL.createObjectURL(bb));
        pom.setAttribute("download", filename);
        pom.dataset.downloadurl = ["text/plain", pom.download, pom.href].join(
          ":"
        );
        pom.draggable = true;
        pom.classList.add("dragout");
        pom.click();
        this.alertInfo(
          `Se ha descargado ${filename}.`
        );        
      },

      code() {
        if (!this.codeOK) {
          var code = "";
          var emptyCount = this.userStoryForm.usType.codeFormat.length;
          this.userStoryForm.usType.codeFormat.forEach(id => {
            var value = "NNNNN.";
            if (Object.keys(this.values).indexOf(id) > -1) {
              if (
                !(
                  typeof this.values[id] === "undefined" ||
                  this.values[id] === null ||
                  this.values[id] === ""
                )
              ) {
                value = `${this.values[id]}_`;
                emptyCount--;
              }
            }
            code += value;
          });
          if (emptyCount == this.userStoryForm.usType.codeFormat.length) {
            this.userStoryCode = "";
          } else {
            this.codeOK = emptyCount == 0;

            this.userStoryCode = code + "????";
          }
        }
      },
      changeVersion() {
        //this.showForm = false;

        var index = 0;
        this.versions.forEach(ver => {
          if (ver.version == this.currentVer) {
            //console.log(`Cambio a version : ${this.currentVer} - INDEX = ${index}`);
            var userStoryForm = this.cloneObject(this.userStory);
            console.log(">>>>>>>>>>>> " + JSON.stringify(userStoryForm));
            var usVer = userStoryForm.us.versions[index];
            var userStoryForm = {};

            userStoryForm.preConditions = usVer.preConditions;
            userStoryForm.fields = usVer.fields;

            //console.log(`Inputs : ${JSON.stringify(userStoryForm)}`);
            /*
              this.$emit("update:steps", usVer.steps)
              this.$emit("update:userStoryForm", userStoryForm)
              */

            //this.$emit("update:userStoryForm", userStoryForm)
            this.steps = usVer.steps;
            this.userStoryForm = userStoryForm;

            usVer.preConditions.forEach(pre => {
              this.values[pre.id] = pre.value;
            });

            /*
              setTimeout(() => {
              this.showForm = true;
              }, 60);
              */
            //console.log(`Inputs : ${JSON.stringify(usVer.fields)}`);
            //this.showForm = true;
            return index;
          }
          index++;
        });
      },

      sendToWorkpace() {
        //console.log(`UserStory > ${this.idUS} ejecutando sendToWorkSpace`);
        this.$emit("update:idUS", this.idUS);
        this.$emit("sendToWorkSpace");
      },
      async similarUserStories(codeSimil) {
        this.code();
        this.loadSimil = false;
        var pathname = window.location.pathname;
        var projectName = pathname.split("/")[2];
        var url = `/davinci/${projectName}/userStories/similar`;
        this.userStoriesSimil = [];
        var preconditionQuery = [];
        var full = 0;
        for (var i in this.values) {
          if (!this.isEmptyOrSpaces(this.values[i])) {
            preconditionQuery.push({
              id: i,
              value: this.values[i]
            });
            full++;
          }
        }
        if (full > 0) {
          await this.axios.post(url, preconditionQuery).then(rs => {
            this.codeSimil = codeSimil;
            for (var i in rs.data) {
              var us = rs.data[i];
              //consolelog("US =>", us.id);
              if (us.id != this.userStoryCode) {
                this.userStoriesSimil.push(us);
              }
            }
          });
        }

        this.loadSimil = true;
      },

      remove() {
        this.$emit("remove");
      },

      async create() {
        await this.sendToDB(this.CREATE);
      },
      async save() {
        await this.sendToDB(this.SAVE);
      },
      async sendToDB(action) {
        var newPreConditions = [];
        var inventionSequence = "";

        if (!this.codeOK) {
          this.alertInfo(
            "No ha determinado todos los valores que forman el codigo identificador de esta historia de usuario."
          );
          return false;
        }

        if (this.isEmptyOrSpaces(this.userStoryForm.fields.description)) {
          this.alertInfo(
            "Las Historias de Usuarios, debe poseer si o si una descripcion."
          );
          return false;
        }

        this.userStory.preConditions.forEach(pre => {
          var preUS = {
            id: pre.id,
            code: pre.code,
            name: pre.name,
            label: pre.label,
            parent: pre.parent,
            value: this.values[pre.id]
          };

          if (this.userStory.usType.codeFormat.indexOf(pre.id) > -1) {
            inventionSequence += pre.code;
          }

          newPreConditions.push(preUS);
        });

        var codeEmpty = this.userStoryCode.replace("????", "");

        var versions = [];
        var version = -1;

        if (action == this.SAVE) {
          version = this.userStory.us.versions.length - 1;
          //Importante para las versiones!
          versions = this.userStory.us.versions;
          versions[version] = {
            version: "V" + (version + 1),
            fields: this.userStoryForm.fields,
            preConditions: newPreConditions,
            steps: this.steps,
            author: this.getUserOnline()
          };
        }

        if (action == this.CREATE) {
          if (this.isUserStoryExist) {
            version = this.userStory.us.versions.length;
            versions = this.userStory.us.versions;
            versions.push({
              version: "V" + (version + 1),
              fields: this.userStoryForm.fields,
              preConditions: newPreConditions,
              steps: this.steps,
              author: this.getUserOnline()
            });
          } else {
            versions = [
              {
                version: "V1",
                fields: this.userStoryForm.fields,
                preConditions: newPreConditions,
                steps: this.steps,
                author: this.getUserOnline()
              }
            ];
            version = 0;
          }
        }

        var userStoryTO = {
          id: version == 0 ? "" : this.userStoryCode,
          type: this.userStory.usType.code,
          LastVersionIndex: version,
          code: {
            value: version == 0 ? codeEmpty : this.userStoryCode,
            format: this.userStory.usType.codeFormat,
            inventionSequence: inventionSequence
          },
          versions: versions
        };
        var pathname = window.location.pathname;
        var projectName = pathname.split("/")[2];

        var url = `/davinci/${projectName}/userStories`;

        //console.log(JSON.stringify(userStoryForm));
        await this.axios.post(url, userStoryTO).then(rs => {
          //Guardo temporalmente la definicion del tipo de historia de usuario, para conservarla despues de guardar
          var usType = this.userStory.usType;
          //consolelog(JSON.stringify(rs.data.id));
          this.userStoryCode = rs.data.id;
          this.dateFormat = this.formatDays(rs.data.date);
          this.versions = [];
          for (var v in rs.data.versions) {
            var ver = rs.data.versions[v];
            var reference = `Creada por ${ver.author}, ${this.formatDays(
              ver.date
            )}`;
            this.versions.push({
              version: ver.version,
              ref: reference
            });
            this.lastVersion = ver.version;
            this.currentVer = ver.version;
          }
          this.isUserStoryExist = true;
          this.userStoryForm.us = rs.data;
          this.userStoryForm.usType = usType;
          this.$emit("update:userStory", this.userStoryForm);
        });
        var code = "";
        this.userStoryForm.usType.codeFormat.forEach(id => {
          var value = "";
          if (Object.keys(this.values).indexOf(id) > -1) {
            value = `${this.values[id]}_`;
          }
          code += value;
        });
        this.codeSimil = "";
        this.alertInfo("Historia de usuario guardada");
      }
    }
  };
</script>
<style scoped>
  .topnav {
    overflow: hidden;
    color: #000;
  }

  .user-stories-title {
    position: relative;
    top: 10px;
    left: 10px;
    font-size: medium;
  }

  .topnav a {
    float: right;
    color: var(--sub-text);
    text-align: center;
    padding: 14px 16px;
    text-decoration: none;
    font-size: 17px;
  }

  .topnav a:hover {
    background-color: var(--sub-panel);
    color: black;
  }

  .topnav a.active {
    background-color: #fff;
    color: 000;
  }

  .us-title {
    font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
      "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
    color: #fff;
    border-block-color: #000;
    padding: 5px;
    background-color: #000;
    text-shadow: 1px 1px var(--sub-panel);
  }

  .us-ver {
    top: -12px;
    position: absolute;
  }

  .us-btn {
    background-color: #fff;
    color: #000;
  }

  .us-btn:hover {
    top: -5px;
    background-color: var(--sub-button) !important;
  }
</style>
<template>
  <form @submit.prevent v-if="showForm" :id="epic.id">
    <card style="min-height: auto;background: var(--transparent);">
      <template slot="header">
        <div class="topnav">
          <div class="epic-title">
            <span class="epic-reference" v-html="reference"></span>
            <!--<a title="Cerrar" @click="close"><i class="ti-close"></i></a>-->
            <a class="btn btn-xs e-btn" title="Cerrar" @click="remove"><i class="ti-close"></i></a>
            <a class="btn btn-xs e-btn" title="Deshacer" @click="undo"><i class="ti-reload"></i></a class="btn btn-xs e-btn">
            <a class="btn btn-xs e-btn" title="Crear" @click="save"><i class="ti-save"></i></a class="btn btn-xs e-btn">
          </div>
        </div>
      </template>
      <div style="padding-left:16px">
        <h-tabs :tabs="epicTab" viewScoped>
          <span slot="epic" style="min-height: 400px;">
            <invention-form v-if="showEpicAtrs" :projectCode="projectCode" :invention.sync="epicForm.attributes" :values.sync="attributes"></invention-form>

            <!--<small>{{epic.definition}}</small>-->
            <h6>Exportar</h6>
          </span>
          <span slot="act" style="min-height: 400px;">
            <card>
              <invention-form :projectCode="projectCode" :invention.sync="epicForm.inventions" :values.sync="inventions"></invention-form>
              <div style="float: right;">
                <h4 class="btn btn-success" @click="createActivity" href="#">
                  <i style="color:yellow" class="ti-bolt"></i> CREAR ACTIVIDAD
                </h4>
              </div>
            </card>


            <accordion :title="activity.title" animation='bottomToTop' v-for="activity,index in epicForm.activityList" :key="index">
              <activity-form :epic="epicForm" v-on:saveUserStories="setUserStories(index)" :activity.sync="epicForm.activityList[index]" :idUS.sync="idUS" v-on:found="sendToWorkpace()"></activity-form>
            </accordion>
          </span>
        </h-tabs>
        <small class="epic-reference" style="float:right" v-if="epic.formatDate">Ultima Modificacion {{epic.formatDate}}</small>
      </div>
    </card>
  </form>
</template>
<script>
  import InventionForm from "../../Invention/InventionForm.vue";
  import ActivityForm from "./ActivityForm.vue";
  export default {
    name: "epic",
    components: {
      InventionForm,
      ActivityForm
    },
    computed: {
      reference() {
        if (this.isEmptyOrSpaces(this.attributes[this.keyValue])) {
          return `${this.epic.name} sin referencia <small>(Ingrese el ${this.keyValue})</small>`;
        } else {
          return `<strong style="color:yellow">${this.epic.name}</strong> : ${this.attributes[this.keyValue]}`;
        }
      }
    },
    props: {
      epic: {},
      form: {}
    },

    data() {
      return {
        idUS: "",
        showEpicAtrs: true,
        keyValue: "",
        showActivity: false,
        attributes: [],
        inventions: [],
        epicTab: [
          {
            id: 0,
            title: `<span class="ti-package"></span> ${this.epic.name}`,
            slot: "epic"
          },
          {
            id: 1,
            title: `<span class="ti-dropbox"></span> Actividades`,
            slot: "act"
          }

        ],

        epicForm: {},
        values: {},
        projectCode: "",
        showForm: false
      }
    },

    watch: {
      epicForm(newValue, oldValue) {
        this.$emit("update:form", newValue)
      }
    },

    async created() {
      //1.485.000‬
      /*
      if (this.epic['values'] != null) {
        this.epicForm = this.cloneObject(this.epic.values);
      }
      */
      await this.createForm();
    },

    methods: {
      sendToWorkpace() {
        //console.log(`UserStory > ${this.idUS} ejecutando sendToWorkSpace`);
        this.$emit("update:idUS", this.idUS)
        this.$emit("sendToWorkSpace")
      },
      getDetailArtifactByName(name) {
        for (var a in this.epicForm.inventions.artifacts) {
          if (this.epicForm.inventions.artifacts[a].name == name) {
            return this.cloneObject(this.epicForm.inventions.artifacts[a]);
          }
        }
      },
      createActivity() {

        //console.log(JSON.stringify(this.inventions));
        for (var key in this.inventions) {
          if (this.inventions[key] == "") {
            alert("No ha Ingresado el valor para " + key);
            return false;
          }
        }

        if (this.isEmptyOrSpaces(this.attributes[this.keyValue])) {
          alert("No ingreso el Codigo de la Iniciativa");
          return false;
        }

        this.addActivity("", this.inventions, [])

      },

      addActivity(activity, inventions, userStoriesRef) {
        var inventionDetail = {}
        var epic = this.attributes[this.keyValue];
        if (activity == "") {
          activity = epic;
        }
        var continua = false;
        for (var k in inventions) {
          if (activity == epic || continua) {
            continua = true;
            activity += " " + inventions[k];
          }
          inventionDetail[k] = this.getDetailArtifactByName(k);
          inventionDetail[k].value = [inventions[k]]
        }

        var unique = true;

        this.epicForm.activityList.forEach(act => {
          if (act.activityForm.activity == activity) {
            alert("Esta actividad ya fue ingresada.");
            unique = false;
            return false;
          }
        });

        if (unique) {
          var activityTab = {
            id: this.makeId(10),
            title: ` <h5 class="activity-title"><strong>${activity}</strong></h5>`,
            slots: activity,
            activityForm: {
              epic: epic,
              activity: activity,
              inventions: this.cloneObject(inventionDetail),
              userStories: userStoriesRef
            }
          }

          this.epicForm.activityList.push(activityTab);
          this.showActivity = true;
        }

      },
      async createForm() {
        var artifacts = []



        if (this.form == null) {

          var allInventions = []
          await this.axios.get("/api/invention/all/").then(rs => {
            rs.data.forEach(inventRS => {
              allInventions.push(inventRS);
            })
          })


          for (var ref in this.epic.epicForm.attributes) {
            var artifactType = this.epic.epicForm.attributes[ref].artifactType;
            var name = this.epic.epicForm.attributes[ref].name
            this.keyValue = this.epic.epicForm.reference;
            this.projectCode = await this.getCodeProject();
            var artifactVO = {}
            var idAtrib = this.epic.epicForm.attributes[ref].artifactType;

            await this.axios.get(`/davinci/typeref/${artifactType}/`).then(rs => {
              var typeRef = rs.data;

              artifactVO.code = typeRef.Code;
              artifactVO.name = name;
              artifactVO.typeName = typeRef.Name;
              artifactVO.isEssential = true
              /*
              artifactVO.value = [
                this.epic.epicForm.attributes[ref].value
              ];
              */
            });

            if (this.isEmptyOrSpaces(artifactVO.code)) {
              allInventions.forEach(invention => {
                if (invention.code == artifactType) {
                  artifactVO = {
                    "name": name,
                    "code": invention.code,
                    "typeName": "json",
                    "nickName": invention.name,
                    "typeRef": invention.code,
                    "isEssential": false,
                    "isList": false,
                    "isJson": true,
                    "icon": invention.icon,
                    "keyValue": invention.keyValue,
                    "keyLabel": invention.keyLabel,
                    "value": [
                      this.epic.epicForm.attributes[ref].value
                    ],
                    "label": name
                  }
                }
              })
            }
            this.attributes[name] = this.cloneObject(this.epic.epicForm.attributes[ref].value);
            artifacts.push(artifactVO);
          }

          var inventions = []
          for (var ref in this.epic.epicForm.inventions) {
            var invention = this.epic.epicForm.inventions[ref]
            var artifact = {}
            var code = await this.getCodeProject();
            allInventions.forEach(inventRS => {
              if (inventRS.code == invention.code) {
                artifact = {
                  "name": invention.name,
                  "code": invention.code,
                  "typeName": "json",
                  "nickName": invention.name,
                  "typeRef": invention.code,
                  "isEssential": false,
                  "isList": false,
                  "isJson": true,
                  "icon": inventRS.icon,
                  "keyValue": inventRS.keyValue,
                  "keyLabel": inventRS.keyLabel,
                  "value": [
                    invention.value
                  ],
                  "label": invention.name
                }
                return true;
              }
            })

            if (artifact != {}) {
              await this.axios.post(`/api/project/${code}/${artifact.code}/get`).then(rs => {
                var comboInv = []
                for (var key in rs.data) {
                  var data = rs.data[key];
                  var select = {}
                  select[artifact.keyValue] = data[artifact.keyValue][0]
                  select[artifact.keyLabel] = data[artifact.keyLabel][0]
                  comboInv.push(select)
                }
                sessionStorage.setItem(artifact.code + code, JSON.stringify(comboInv));
                return true;
              });
            }


            inventions.push(artifact);
          }

          //console.log(JSON.stringify(this.epic.epicForm.inventions));

          //TODO: ACA VER COMO LEER DESDE BASE DE DATOS TODO
          this.epicForm['keyValue'] = this.keyValue;
          // {"Canal":"ANF","Tipo Ejecución":"CONT","Ambiente":"QA"}

          this.epicForm =
          {
            isNew: true,
            keyValue: this.keyValue,
            activityList: [],
            attributes: {
              //Setea como KEY que se Asigno desde diseño al EPICO.
              keyValue: this.keyValue
            },
            inventions: {},
          }





          this.epicForm.attributes['artifacts'] = artifacts;
          this.epicForm.inventions['artifacts'] = inventions;

          if (this.epic.epicForm.activities != null) {

            this.epic.epicForm.activities.forEach(act => {
              var inventionsForm = {}
              act.inventions.forEach(inv => {
                inventionsForm[inv.name] = inv.value[0]
              })
              this.addActivity(act.code, inventionsForm, act.userStoriesRef);
            })
          }
          //console.log(JSON.stringify(this.epicForm));
          this.$emit("update:form", this.epicForm)
        }
        this.showForm = true
      },

      remove() {
        this.$emit("remove");
        //TODO: Agregar remover a base de datos
      },

      close() {
        this.$emit("remove");
      },

      undo() {

      },


      makeAttributesDB() {
        //<invention-form :projectCode="projectCode" :invention.sync="epicForm.attributes" :values.sync="attributes"></invention-form>
        this.attributes
        var attributesDB = [];
        this.epicForm.attributes.artifacts.forEach(atr => {
          //console.log("Value  => " + JSON.stringify(this.attributes[atr.name]));
          atr.inactive = true;
          var attributeDB = {
            name: atr.name,
            code: atr.code,
            value: this.attributes[atr.name]
          }
          attributesDB.push(attributeDB);
        });
        return attributesDB;
      },

      setUserStories(index) {
        var userStories = this.epicForm.activityList[index].userStories;
        var userStoriesDB = this.formatUSDB(userStories);
      },

      formatUSDB(userStories) {
        var usDBs = [];
        var order = 1;
        userStories.forEach(us => {
          var usDB = {
            code: us.id,
            version: us.version,
            indexVersion: us.indexVersion,
            executeOrder: order
            /*
            create: us.date,
            steps: us.steps,
            fields: us.fields,
            preConditions: us.preConditions,
            version: us.version,
            */
          }
          order++;
          usDBs.push(usDB);
        })
        return usDBs;
      },

      async save() {

        //this.epicForm.update = new Date().format("%d-%m-%Y")
        var activities = []


        if (this.attributes[this.keyValue] == "") {
          alert("No ha ingresado el Codigo del Epico");
          return false;
        }

        this.epicForm.activityList.forEach(act => {
          var form = act.activityForm;
          var inventions = []
          Object.keys(form.inventions).forEach(key => {
            var invention = form.inventions[key]
            var inv = {
              name: invention.name,
              code: invention.code,
              value: invention.value
            }
            inventions.push(inv);
          })
          var activity = {
            code: form.activity,
            inventions: inventions,
            userStoriesRef: this.formatUSDB(act.userStories),
          }
          activities.push(activity);
        });

        var attributesDB = this.makeAttributesDB();
        var epicId = this.attributes[this.keyValue];

        var epicDB = {
          id: this.epic.id,
          code: epicId,
          attributes: attributesDB,
          type: this.epic.type,
          author: this.getUserOnline(),
          activities: activities
        }
        this.showEpicAtrs = false;
        var epicResult = await this.saveEpic(epicDB);
        var updatedEpic = this.cloneObject(this.epic);
        updatedEpic.id = epicResult.id;
        updatedEpic.date = epicResult.date;
        updatedEpic['formatDate'] = this.formatDays(epicResult.date);
        this.$emit("update:epic", updatedEpic)
        this.showEpicAtrs = true;
        alert('Epico Guardado.');
      }

    }
  };
</script>
<style>
  .topnav {
    overflow: hidden;
    color: #000;
  }

  .epic-title {
    position: relative;
    top: 10px;
    left: 10px;
    font-size: large;
  }

  .topnav a {
    float: right;
    right: 10px;
    color: var(--sub-text);
    text-align: center;
    padding: 14px 8px;
    text-decoration: none;
    font-size: 17px;
  }

  .topnav a:hover {
    color: var(--wallpaper);
    cursor: pointer;
  }

  .topnav a.active {
    background-color: var(--transparent);
    color: 000;
  }

  .epic-reference {
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
    color: #fff;
    border-block-color: #000;
    padding: 5px;
    background-color: #000;
    text-shadow: 1px 1px var(--sub-panel);
  }

  .e-btn {
    background-color: #FFF !important;
    color: #000
  }

  .e-btn:hover {
    top: -5px;
    background-color: var(--sub-button) !important;
  }
</style>
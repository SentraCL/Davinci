<template>
  <div class="row">
    <div class="col-1 scrolling">
      <draggable class="dragArea  " :list="itemsInvention" :group="{ name: 'form-design', pull: 'clone', put: false }" @change="log">
        <center class="list-group-item invention" v-for="item in itemsInvention" :key="item.name">
          <img :src="item.icon" class="invention-item" :title="item.name" width="48px" /><br />
          <small style="z-index:10000">{{item.name}}</small>
        </center>
      </draggable>
    </div>

    <br /><br />

    <div class="col-11" id="design-panel">
      <card>
        <drop-menu :title="userTitle" style="float:right;position:relative;" v-on:change="userStoriesOptions(user.option)" :options="user.options" :option.sync="user.option"></drop-menu>
        <h3><span class="ti-notepad"></span> {{userStoriesForm.title}} </h3>
        <card class="card-body" style="background-color: #fafafa;">
          <span>
            <span v-if="showFormat">
              <h5><i class="fa fa-barcode"></i> Codigo : <span v-if="Object.keys(testCaseCode).length>0" v-for="code,index in testCaseCode"><strong>{{values[code.id]==null?'????':values[code.id]}}</strong>_</span>0001
              </h5>
              <small><strong>Formato :</strong> <span v-if="Object.keys(testCaseCode).length>0" v-for="code,index in testCaseCode"><i>{{code.name}}</i> + </span> [Correlativo] </small>
            </span>
          </span>
          <hr />
          <draggable style="min-height: 400px;" class="dragArea list-group" :list="inputsUserStories" group="form-design" @change="log">
            <span class="row col-mb-12" v-for="input,index in inputsUserStories">
              <!--<img :src="input.icon" height="32px" width="32px">-->
              <combo-simple style="position:relative;top:5px;" :label="input.label" class="input-item" :list="input.list" :keyValue="input.keyValue" v-on:update="loadChildren(input, index)" :keyLabel="input.keyLabel" :value.sync="values[input.id]"></combo-simple>
              <drop-menu :title="taskTitle" v-if="refreshMenu" v-on:change="editInput(input,input.option, index)" :options.sync="input.options" :option.sync="input.option"></drop-menu>
            </span>
          </draggable>
        </card>
      </card>
      <small><i class="ti-announcement"></i>Los datos son solo referenciales para tener un diseño mas realista. <br />
        <span style="color:red" v-if="Object.keys(testCaseCode).length==0">No asigno un formato para el codigo que identificara este esta historia de usuario. </span>
        <!--
        <span style="color:red" v-if="Object.keys(testCaseCode).length==0">No asigno un formato para el codigo que identificara este esta historia de usuario. </span>
        -->
      </small>
    </div>



    <span class="fondo">
      <!-- 
    <small><span class="ti-comment"></span> Considera que solo veras los inventos en los cuales ya ingresaste su contenido.</small>
    <br/>
    -->
      <!--<img width="256px" src="@/assets/img/disign.png" />-->
    </span>
    <!-- DIALOG's -->
    <m-dialog id="editDialog" :title="userStoriesDialog.title" :show.sync="userStoriesDialog.show" :isClose.sync="userStoriesDialog.close">
      <div slot="dialog">

        <input-text label="Titulo" name="titulo" v-model="userStoriesForm.title" autocomplete="off"> </input-text>
      </div>
      <div slot="actions">
        <div class="row">
          <div class="col-md-2">
            <button class="btn btn-round btn-success" @click="changeTitle()">
              Renombrar
            </button>
          </div>
          <div class="col-md-2">
            <button class="btn btn-round btn-warning" @click="cancelTitle()">
              Cancelar
            </button>
          </div>
        </div>
      </div>
    </m-dialog>

    <m-dialog id="editDialog" :title="editDialog.title" :show.sync="editDialog.show" :isClose.sync="editDialog.close">
      <div slot="dialog">
        <div style="width: 600px;min-width: 400px !important;">
          <img :src="editDialog.input.icon" width="64">
          <input-text label="Nombre" name="name" v-model="editDialog.input.label" autocomplete="off" autofocus> </input-text>
        </div>
      </div>
      <div slot="actions">
        <div class="row">
          <div class="col-md-12" style="float:right;position: relative;">
            <button class="btn btn-round btn-success" @click="saveEdit()">
              Renombrar
            </button>
            <button class="btn btn-round btn-warning" @click="cancelEdit()">
              Cancelar
            </button>
          </div>
        </div>
      </div>
    </m-dialog>
  </div>
</template>

<script>
  import draggable from "vuedraggable";
  export default {
    name: "user-stories-item",
    display: "user-stories-item",
    order: 2,
    props: {
      project: {},
      inventions: {},
      userStories: {}
    },

    data() {
      return {
        newname: "",
        showFormat: true,
        refreshMenu: true,

        RENAME: 1,
        DELETE: 2,
        CODECODE: 3,
        CLEAN: 4,
        SAVE: 5,
        BACK: 6,

        testCaseCode: {},

        itemsInvention: [],
        inputsUserStories: [],
        taskTitle: `<small> Editar</small>`,
        userTitle: `<i class="ti-settings"></i> <span> Ajustar</span>`,
        values: {},

        user: {
          option: -1,
          options: [
            { html: `<span><i style="color:green" class="ti-pencil-alt"></i> Cambiar Titulo</span>`, option: 1 },
            { html: `<span><i style="color:blue" class="ti-save"></i> Guardar</span>`, option: 5 },
            { html: `<span><i style="color:black" class="ti-eraser"></i> Borrar Todo</span>`, option: 4 },
            { html: `<span><i style="color:brown" class="ti-control-backward"></i> Volver</span>`, option: 6 },
          ]
        },

        userStoriesForm: {
          title: `Historia de Usuario #1`,
          preConditions: [],
        },

        userStoriesDialog: {
          title: `<span class="ti-pencil-alt"> </span> Historias de Usuario <small>(Casos de uso)</small>.`,
          original: {},
          show: false,
          close: false,
        },

        editDialog: {
          title: `<span class="ti-more-alt"> </span> Editar Atributos.`,
          show: false,
          close: false,
          index: -1,
          input: {}
        },
      };
    },
    components: {
      draggable
    },

    watch: {
      inputsUserStories(newValue, oldValue) {
        this.setIdInventions();
      }
    },



    async mounted() {

      for (var k in this.inventions) {
        var invention = this.inventions[k];
        invention.id = "";//invention.code;
        invention.code = invention.code;
        invention.originalKey = k;
        invention.label = invention.name;
        invention.keyValue = invention.keyValue;
        invention.keyLabel = invention.keyLabel;
        invention.option = 0;
        invention.options = [
          { html: `<small><i style="color:green" class="ti-pencil-alt2"></i> Renombrar</small>`, option: this.RENAME },
          { html: `<small><i style="color:red" class="ti-eraser"></i> Borrar</small>`, option: this.DELETE },
          { html: `<small><i style="color:blue" class="fa fa-barcode"></i> Ser Parte del Codigo de Caso</small>`, option: this.CODECODE },
        ]

        invention.artifacts.forEach(artifact => {
          if (!artifact.isEssential) {
            var subArtefact = {
              html: `<small><i style="color:black" class="ti-plus"></i> Agregar <strong>${artifact.name}</strong> filtrado por <strong>${invention.name}</strong></small>`,
              option: artifact.typeRef
            }
            invention.options.push(subArtefact);
          }
        });

        var defaultValue = "";
        var list = [];

        invention.WareHouse.forEach(data => {
          var stringJSON = data.ContentJSON;
          var jsonContent = JSON.parse(stringJSON)
          var item = {}
          item[invention.keyValue] = jsonContent[invention.keyValue][0];
          item[invention.keyLabel] = (invention.keyLabel == invention.keyValue) ? jsonContent[invention.keyValue][0] : jsonContent[invention.keyValue][0] + " : " + jsonContent[invention.keyLabel][0];
          list.push(item);
          defaultValue = item[invention.keyValue]
        });

        invention.list = list;
        this.itemsInvention.push(invention);



      }

      if (this.userStories != null) {
        this.user.options.shift();
        this.userStoriesForm.title = this.userStories.title;
        this.userStories.preConditions.forEach(usDB => {
          var inputUS = this.getItemInventionByCode(usDB.code);
          inputUS.id = usDB.id;

          if (usDB.parent.id != "") {
            inputUS.parent = {
              input: this.getParentById(usDB.parent.id),
              artifactName: usDB.parent.artifactName,
              id: usDB.parent.id
            };
            inputUS.allData = this.cloneObject(inputUS.list);;
            inputUS.list = [];
          }
          inputUS.label = usDB.label;
          this.inputsUserStories.push(inputUS);
          if (this.userStories.codeFormat.indexOf(inputUS.id) > -1) {
            this.testCaseCode[inputUS.id] = inputUS;
          }
        })

      }

    },

    methods: {
      getParentById(id) {
        var theRealParent = {}
        this.inputsUserStories.forEach(parent => {
          if (parent.id == id) {
            theRealParent = parent;//this.getItemInventionByCode(parent.code);
          }
        });
        return theRealParent;
      },

      getItemInventionByCode(code) {
        var object = {}
        this.itemsInvention.forEach(item => {
          if (code == item.code) {
            object = item;
          }
        })
        return object;
      },

      setIdInventions() {
        if (this.inputsUserStories.length > 0) {

          var lastInv = this.cloneObject(this.inputsUserStories[this.inputsUserStories.length - 1])
          if (lastInv.id == "" || lastInv.id == null) {
            var d = new Date();
            lastInv.id = `${this.project.code}.${d.getTime()}`;
            this.inputsUserStories[this.inputsUserStories.length - 1] = lastInv;
          }
        }
      },

      saveEdit() {
        this.inputsUserStories[this.editDialog.index] = this.cloneObject(this.editDialog.input);
        this.editDialog.show = false;
      },

      cancelEdit() {
        this.editDialog.show = false;
      },

      log: function (evt) {
        /*
        var d = new Date();
        var invention = this.cloneObject(evt.added.element);
        invention.label = d.getTime();          
        invention.id = d.getTime();          

        evt.added.element = invention;
        //console.log(JSON.stringify(evt.added.element.id));
        */
      },

      async userStoriesOptions(option) {
        //console.log("OPTION >> " + option);
        if (option == this.RENAME) {
          this.userStoriesDialog.original.title = this.cloneObject(this.userStoriesForm.title)
          this.userStoriesDialog.show = true;
        }
        if (option == this.BACK) {
          this.$emit("back");
        }


        if (option == this.SAVE) {

          if (Object.keys(this.testCaseCode).length == 0) {
            alert("Aun no asigna la formula del codigo que hara unica esta historia de usuario.");
            return false;
          }

          //console.log(JSON.stringify(this.userStoriesForm,null,4));
          //console.log(JSON.stringify(this.inputsUserStories,null,4));
          //console.log(JSON.stringify(this.inputsUserStories));
          this.inputsUserStories.forEach(inputUH => {
            var input = {
              id: inputUH.id,
              code: inputUH.code,
              artifacts: inputUH.artifacts,
              name: inputUH.name,
              label: inputUH.label,
              parent: Object.keys(inputUH).indexOf('parent') == -1 ? {} : {
                id: inputUH.parent.id, artifactName: inputUH.parent.artifactName
              }
            }
            this.userStoriesForm.preConditions.push(input);
          });
          this.userStoriesForm.codeFormat = Object.keys(this.testCaseCode);
          //console.log(JSON.stringify(this.userStoriesForm));
          //console.log(JSON.stringify(this.userStoriesForm.preConditions, null, 4));
          await this.axios.post(`/api/project/${this.project.code}/design/userstories`, this.userStoriesForm).then(rs => {
            alert("Historia de Usuario Guardada..");
            this.$emit("back");
          });
        }


        if (option == this.CLEAN) {
          this.inputsUserStories = [];
        }

      },

      changeTitle() {
        this.userStoriesDialog.show = false;
      },

      cancelTitle() {
        this.userStoriesForm.title = this.userStoriesDialog.original.title;
        this.userStoriesDialog.show = false;
      },

      loadChildren(parent, index) {
        var value = this.values[parent.id];
        this.inputsUserStories.forEach(input => {
          if (Object.keys(input).indexOf('parent') > -1 && (input.parent.input.code == parent.code && input.parent.id == parent.id)) {
            input.list = [];
            //console.log(input.name);
            //Considerar que toda la data sin filtro esta en input.allData
            //los Filtros estan en warehouse
            for (var r in parent.WareHouse) {
              var repository = JSON.parse(parent.WareHouse[r].ContentJSON);
              var key = input.parent.artifactName;
              //console.log(value);
              //Se genera el Filtro.
              if (value == parent.WareHouse[r].KeyValue) {
                //Se identidica el artefacto para ser filtrado
                if (Object.keys(repository).indexOf(key) > -1) {
                  //console.log(JSON.stringify(repository[key]));
                  for (var a in input.allData) {
                    var item = input.allData[a]
                    //var unique = (input.keyValue != input.keyLabel) ? item[input.keyValue] : item[input.keyValue].split(':')[0];
                    //unique = unique.trim();
                    var unique = item[input.keyValue];
                    if (repository[key].indexOf(unique) > -1) {
                      input.list.push(item);
                    }
                  }
                  break;
                }
              }
            }
          }
        })
      },

      async editInput(input, option, index) {
        //console.log(option);
        var id = input.originalKey;

        var itemIndex = -1;
        for (var i in this.inputsUserStories) {
          if (this.inputsUserStories[i].label == input.label) {
            itemIndex = i;
            break;
          }
        }
        this.editDialog.index = itemIndex;

        if (option == this.DELETE) {
          if (itemIndex > -1) {
            var labelDel = this.inputsUserStories[itemIndex].label;
            //console.log(JSON.stringify(this.userStoriesForm.preConditions));
            delete this.userStoriesForm.preConditions.labelDel
            this.inputsUserStories.splice(itemIndex, 1)
            return true;
          }
        }

        if (option == this.RENAME) {
          this.editDialog.show = true
          this.editDialog.input = this.cloneObject(input);
          return true;
        }

        //Switch , agrear o quitar de Codigo de Caso.
        if (option == this.CODECODE || option == this.CODECODE * -1) {
          this.refreshMenu = false;
          input.options[2] = (option == this.CODECODE) ? {
            html: `<small><i style="color:red" class="fa fa-barcode"></i>Quitar del Codigo</small>`,
            option: this.CODECODE * -1
          } : {
              html: `<small><i style="color:blue" class="fa fa-barcode"></i> Ser Parte del Codigo</small>`,
              option: this.CODECODE
            };
          this.showFormat = false;
          setTimeout(() => {
            if (option == this.CODECODE) {
              this.testCaseCode[input.id] = input;
            } else {
              delete this.testCaseCode[input.id];
            }
            this.refreshMenu = true;
            this.showFormat = true;
          }, 60);
          return true;
        }




        this.inventions.forEach(original_inv => {
          var invention = this.cloneObject(original_inv);
          //console.log(JSON.stringify(invention));
          if (invention.code == option) {
            //Obtener Filtro desde warehouse de input.
            var filter = []
            var artifact = {};
            for (var a in input.artifacts) {
              artifact = input.artifacts[a];
              if (artifact.typeRef == option) {
                break;
              }
            }

            //TODO : Setear Parent , tambien ver sub inventos.
            invention.parent = {
              input: this.cloneObject(input),
              artifactName: artifact.name,
              id: input.id
            };

            invention.allData = this.cloneObject(invention.list);
            invention.list = [];
            this.inputsUserStories.push(invention);
          }

        });

      }
    }
  };
</script>
<style scoped>
  #design-panel {
    position: absolute;
    top: 120;
    left: 30;
    min-width: 60%;
    min-height: 100px;
    margin: 8em;
    width: 80%;
  }

  .fondo {
    -webkit-touch-callout: none;
    /* iOS Safari */
    -webkit-user-select: none;
    /* Safari */
    -khtml-user-select: none;
    /* Konqueror HTML */
    -moz-user-select: none;
    /* Firefox */
    -ms-user-select: none;
    /* Internet Explorer/Edge */
    user-select: none;
    top: 560px;
    position: fixed;
    float: right;
    z-index: 1;
    right: 40px;
  }

  .invention {
    min-width: 100px;
    margin-top: 3px;
    transition: .3s;
    font-size: 12px;
  }

  .invention:hover {
    cursor: pointer;
    font-size: 1.2em;
    margin-top: -5px;
  }


  .scrolling {
    overflow-x: hidden;
    overflow-y: scroll;
    white-space: nowrap;
    max-height: 600px;
    min-width: 120px !important;
  }
</style>
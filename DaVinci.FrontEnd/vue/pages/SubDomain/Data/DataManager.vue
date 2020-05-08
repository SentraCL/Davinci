<template>
  <div class="row" width="%80">
    <div class="col-md-12">
      <button style="position:absolute;float:right;right: 10px;" class="btn btn-success btn-xs" @click="toClose"><i class="ti-close"></i> </button>
      <h3> <i class="ti-harddrive"></i> Administración de Datos de Prueba.</h3>

      <div class="card">
        <!-- WIZARD Asistente de Datos -->
        <form-wizard title="Asistente de Creación de Dato" subtitle="Un tipo de Dato, categoriza la data de prueba que puede ser utilizada para una historia de usuario.">
          <tab-content title="Comenzar" :before-change="setForm">
            <card>
              <label><input type="checkbox" id="cbox1" v-model="newType"> Crear un nuevo tipo de dato</label><br>
              <div v-if="!newType">
                Seleccione el tipo de dato que desea modificar.
                <hr />
                <div class="data-types-content row">
                  <div v-for="dt in dataTypes" class="col-md-3">
                    <!-- dt.name==dataType.name -->

                    <span class="iconOK" @click="select(dt)" v-if="dt.id==dataTypeID">
                      <img v-if="dt.icon64" width="68px" :src="dt.icon64" />
                      <img v-if="!dt.icon64" width="68px" :src="defaultIcon" />
                      <br />
                      <strong :title="dt.description" class="data-type-label">{{dt.name}}</strong>
                    </span>

                    <span class="icon" @click="select(dt)" v-if="dt.id!=dataTypeID">
                      <img v-if="dt.icon64" width="64px" :src="dt.icon64" />
                      <img v-if="!dt.icon64" width="64px" :src="defaultIcon" />
                      <br />
                      <span :title="dt.description" class="data-type-label">{{dt.name}}</span>
                    </span>
                  </div>
                </div>
              </div>
            </card>
          </tab-content>
          <tab-content title="Configurar el Dato" :before-change="validateSave">
            <card>
              <div class="row">
                <div class="col-md-10">
                  <input-text :label="nameLabel" v-model="dataType.name"></input-text>
                  <input-text :label="principalKeyLabel" v-model="dataType.principalKey" :title="principalKeyTips"></input-text>
                </div>
                <div class="col-md-2">
                  <img v-if="dataType.icon64" width="96px" :src="dataType.icon64">
                  <img v-if="!dataType.icon64" width="96px" :src="defaultIcon" />
                  <icon ref="icon" :prefill="urlIcon" @change="changeIcon" :removable="false" :autoToggleAspectRatio="true" margin="2" radius="50" size="16" accept="image/jpeg,image/png,image/gif" buttonClass="btn btn-xs" :customStrings="settingIcon">
                  </icon>

                </div>
                <div class="col-md-12">
                  <div class="form-group">
                    <label>Descripción</label>
                    <textarea rows="5" class="form-control border-input" v-model="dataType.description">
                  </textarea>
                  </div>
                </div>
              </div>
            </card>
          </tab-content>
          <tab-content title="Crear Dato">
            <card>
              ¿Desea Confirmar lo ingresado?
              <hr />
              <div class="row">
                <div class="col-md-4"><strong>Tipo de Dato</strong></div>
                <div class="col-md-8"><i class="ti-pin"></i> {{dataType.name}}</div>
                <div class="col-md-4"><strong>Atributo Unico de Referencia</strong></div>
                <div class="col-md-8"><i class="ti-key"></i> {{dataType.principalKey}}</div>
                <div class="col-md-8">
                  {{dataType.description}}
                </div>
                <div class="col-md-4">
                  <img v-if="dataType.icon64" width="64px" :src="dataType.icon64">
                </div>
              </div>
            </card>
          </tab-content>

          <button class="btn btn-info" type="primary" slot="next">Siguiente</button>
          <button class="btn btn-info" type="primary" slot="prev">Volver</button>
          <button class="btn btn-success" type="primary" slot="finish" @click="save">{{newType?'Crear nuevo Tipo de Dato':'Guardar Cambios'}}</button>
        </form-wizard>
      </div>
      <!-- Listado de Datos  -->
    </div>
  </div>

</template>
<script>
  import icon from "vue-picture-input";
  import { FormWizard, TabContent } from 'vue-form-wizard'
  import 'vue-form-wizard/dist/vue-form-wizard.min.css'
  export default {

    name: "data-manager",
    components: {
      FormWizard,
      icon,
      TabContent
    },
    props: {
      artifactForm: {}
    },
    computed: {
      urlIcon() {
        var projectName = this.getProjectDomain();
        var urlIcon = this.dataType.id != "" ? `/davinci/${this.projectName}/datatype/${this.dataType.id}/icon.png` : this.defaultIcon;
        var avatarChecker = new XMLHttpRequest(),
          CheckThisUrl = urlIcon;
        avatarChecker.open('get', CheckThisUrl, true);
        if (avatarChecker.readyState === 1) {
          if (avatarChecker.status == 0) {
            return urlIcon
          }
          else {
            return this.defaultIcon
          }
        }

        return urlIcon
      }
    },
    data() {
      return {

        dataTypeID: "",
        projectName: "",
        defaultIcon: "",

        settingIcon: {
          upload: "<h1>Fallo!</h1>",
          selected: "<p>Icono Seleccionado</p>", // HTML allowed
          fileSize: "El tamaño excede los limites", // Text only
          fileType: "Este archivo de imagen no es soportado.",
          drag: "<h5>Arrastre o seleccione una imagen</h5>", // HTML allowed
          tap: "Seleccione una imagen desde su galeria", // HTML allowed
          change: "Elegir icono", // Text only
          remove: "Eliminar icono", // Text only
          select: "Seleccionar foto", // Text only
          selected: "<p>Su avatar de presentacion a sido seleccionado</p>", // HTML allowed
          fileSize: "el archivo pesa demasiado", // Text only
          fileType: "Este tipo de archivo no es soportado." // Text only
        },

        dataTypes: [],

        newType: true,
        nameLabel: `<i class="ti-pin"></i> Nombre`,
        principalKeyLabel: `<i class="ti-key"></i> Atributo Unico del Dato`,
        principalKeyTips: "como por ejemplo el RUT",
        dataType: {
          id: "",
          name: "",
          principalKey: "",
          description: "",
          icon64: []
        }
      }
    },

    async mounted() {
      this.projectName = this.getProjectDomain();
      this.dataTypes = await this.loadDataTypes();
      this.defaultIcon = `/api/project/avatar/${this.projectName}.png`
    },

    methods: {

      async setForm() {
        this.dataTypes = await this.loadDataTypes();
        if (this.newType) {
          this.dataType = {
            id: "",
            name: "",
            principalKey: "",
            description: "",
            icon64: this.defaultIcon
          }
          return true;
        }
        return true;
      },

      select(dt) {
        this.dataType = dt;
        this.dataTypeID = dt.id;
      },

      validateSave() {
        if (this.isEmptyOrSpaces(this.dataType.name) || this.isEmptyOrSpaces(this.dataType.principalKey)) {
          this.alertInfo('Debe Completar todos los datos indispensables!.');
          return false;
        }
        return true;
      },
      toClose() {
        this.$emit("toClose");
      },

      changeIcon() {
        if (this.$refs.icon.image) {
          this.dataType.icon64 = [];
          this.dataType.icon64 = this.$refs.icon.image;
        }
      },

      async loadDataTypes() {

        var dataTypes = []
        await this.axios.get(`/davinci/${this.projectName}/datatype/getAll/`).then(rs => {
          dataTypes = rs.data;
        });
        return dataTypes;
      },


      async save() {

        await this.axios.post(`/davinci/${this.projectName}/datatype/save/`, this.dataType).then(rs => {
          this.dataType.id = rs.data.id;
          this.dataTypeID = rs.data.id
          this.newType = false;
          this.$emit("saved");
        });
        this.dataTypes = await this.loadDataTypes();
      }
    }
  };
</script>
<style>
  .data-type-content {
    width: 200px;
    height: 100px;
    overflow-y: scroll;
  }

  .data-type-label {
    font-size: 12px;
    font-family: Arial, Helvetica, sans-serif;
  }

  .iconOK {
    border: 1px;
    cursor: pointer;
    margin: 0;
    opacity: 1;
    text-align: center;
    display: block;
    margin-left: auto;
    margin-right: auto;
  }

  .icon {
    cursor: pointer;
    margin: 0;
    opacity: 0.5;
    text-align: center;
    display: block;
    margin-left: auto;
    margin-right: auto;
  }

  .icon:hover {
    opacity: 1;
  }
</style>
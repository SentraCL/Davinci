<template>
  <div class="row" width="%80">
    <div class="col-md-12">
      <button style="position:absolute;float:right;right: 10px;" class="btn btn-success btn-xs" @click="toClose"><i class="ti-close"></i> </button>
      <h3> <i class="ti-harddrive"></i> Administración de Datos de Prueba.</h3>

      <div class="card">
        <!-- WIZARD Asistente de Datos -->
        <form-wizard  title="Asistente de Creación de Dato" subtitle="Un tipo de Dato, categoriza la data de prueba que puede ser utilizada para una historia de usuario.">
          <tab-content title="Comenzar" :before-change="setForm">
            <card>
              <label><input type="checkbox" id="cbox1" v-model="newType"> Crear un nuevo tipo de dato</label><br>
              <div v-if="!newType">
                Seleccione el tipo de dato que desea modificar.
              </div>
            </card>
          </tab-content>
          <tab-content title="Configurar el Dato" :before-change="validateSave">
            <card>
              <div class="row">
                <div class="col-md-12">
                  <input-text :label="nameLabel" v-model="dataType.name"></input-text>
                  <input-text :label="principalKeyLabel" v-model="dataType.principalKey" :title="principalKeyTips"></input-text>
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
                <div class="col-md-12">
                  <hr />{{dataType.description}}</div>
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
  import { FormWizard, TabContent } from 'vue-form-wizard'
  import 'vue-form-wizard/dist/vue-form-wizard.min.css'
  export default {

    name: "data-manager",
    components: {
      FormWizard,
      TabContent
    },
    props: {
      artifactForm: {}
    },

    data() {
      return {
        newType: true,
        nameLabel: `<i class="ti-pin"></i> Nombre`,
        principalKeyLabel: `<i class="ti-key"></i> Atributo Unico del Dato`,
        principalKeyTips: "como por ejemplo el RUT",
        dataType: {
          id: "",
          name: "",
          principalKey: "",
          description: ""
        }
      }
    },

    methods: {

      setForm: function () {
        if (this.newType) {
          this.dataType = {
            id: "",
            name: "",
            principalKey: "",
            description: ""
          }
          return true;
        }else{
          this.dataType = {
            id: "",
            name: "XXXX",
            principalKey: "XXXX",
            description: "XXXX"
          }
          return true;
        }
        
      },

      validateSave() {
        if (this.isEmptyOrSpaces(this.dataType.name) || this.isEmptyOrSpaces(this.dataType.principalKey)) {
          alert('Debe Completar todos los datos indispensables!.');
          return false;
        }
        return true;
      },
      toClose() {
        this.$emit("toClose");
      },
      async save() {
        var pathname = window.location.pathname;
        var projectName = pathname.split("/")[2];
        await this.axios.post(`/davinci/${projectName}/datatype/save/`, this.dataType).then(rs => {
          this.$emit("saved");
        });
      }
    }
  };
</script>
<style></style>
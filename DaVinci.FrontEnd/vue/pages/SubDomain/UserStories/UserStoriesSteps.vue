<template>
  <div>
    <span v-if="!inactive">
      <div class="row">
        <div class="col-12">
          <textarea
            class="ta8"
            placeholder="Ingrese información del paso (Si presiona Shift + Enter se agregara automaticamente)"
            id="txtPaso"
            v-model="step.step"
            @keydown.enter.shift.exact.prevent="addStep()"
          ></textarea>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <br />
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <textarea
            class="ta8"
            placeholder="Resultado del Paso"
            id="txtResultado"
            v-model="step.result"
            @keydown.enter.shift.exact.prevent="addStep()"
          ></textarea>
        </div>
      </div>
      <div class="row">
        <div class="col-10"></div>
        <div class="col-2">
          <button type="button" @click="addStep()" style="width: 100%" class="btn sub-button">
            <i class="ti-notepad"></i> Agregar
          </button>
        </div>
      </div>
      <br />
      <br />
    </span>
    <div class="scrollingSteps" v-if="localSteps.length>0">
      <table class="table scroll" id="tblPasos">
        <tbody style="background-color: #fff;">
          <tr height="32px">
            <td width="5%">
              <strong>Orden</strong>
            </td>
            <!--<th>ID</th>-->
            <td width="40%">
              <strong>Descripcion del Paso</strong>
            </td>
            <td width="40%">
              <strong>Resultado Esperado</strong>
            </td>
            <td width="10%" class="lastItem" v-if="!inactive">
              <strong>Opción</strong>
            </td>
          </tr>

          <tr v-for="item,step in localSteps" height="32px">
            <td width="5%" class="stepCol">{{step+1}}</td>
            <!-- <th width="5%" class="stepCol">P-{{item.id}}</th> -->
            <td width="40%">
              <small :title="item.step">{{item.step | readMore(140, '...') }}</small>
            </td>
            <td width="45%">
              <small :title="item.result">{{item.result | readMore(140, '...')}}</small>
            </td>
            <td width="5%" class="lastItem" v-if="!inactive">
              <span class="up">
                <span class="ti-arrow-up" title="Subir Paso" v-on:click="up(item.id)"></span>
              </span>
              <span class="down">
                <span class="ti-arrow-down" title="Bajar Paso" v-on:click="down(item.id)"></span>
              </span>
              <span class="edit">
                <span class="ti-pencil-alt" title="Editar Paso" v-on:click="edit(item.id)"></span>
              </span>
              <span class="del">
                <span class="ti-trash" title="Eliminar Paso" v-on:click="deleteStep(item.id)"></span>
              </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <side-panel id="editDialog" :show.sync="editDialog.show">
      <div slot="content" style="min-width:800px;">
        <card>
          <h4 v-html="editDialog.title"></h4>
          <h6>Paso : P-{{editDialog.step.id}}</h6>
          <hr />
          <textarea class="ta8" v-model="editDialog.step.step"></textarea>
          <br />
          <br />
          <textarea class="ta8" v-model="editDialog.step.result"></textarea>
          <button type="button" @click="editFinish()" class="btn sub-button">Modificar</button>
          <button type="button" @click="cerrar()" class="btn sub-button">Cancelar</button>
          <hr />
        </card>
      </div>
    </side-panel>
  </div>
</template>
<script>
export default {
  name: "uh-steps",
  computed: {},
  props: {
    steps: {},
    inactive: {}
  },

  data() {
    var _steps = [];
    if (Array.isArray(this.steps)) {
      _steps = this.cloneObject(this.steps);
    }
    return {
      localSteps: _steps,
      editDialog: {
        title: `<span class="ti-pencil-alt"> </span> Editar Paso.`,
        show: false,
        close: true,
        step: {}
      },
      step: {
        step: "",
        result: "",
        id: 0
      }
    };
  },
  watch: {
    steps: function(n, o) {
      this.localSteps = n;
    },
    localSteps: function(n, o) {
      if (JSON.stringify(n) != JSON.stringify(this.steps)) {
        this.$emit("update:steps", n);
      }
    }
  },
  mounted() {},
  methods: {
    cambiar(pos1, pos2, vector) {
      var arreglo = new Array();
      for (let i = 0; i < vector.length; i++) {
        const element = vector[i];
        if (i == pos1) {
          arreglo.push(vector[i + 1]);
          i++;
        }
        arreglo.push(element);
      }
      return arreglo;
    },
    ordenar() {
      for (let i = 0; i < this.localSteps.length; i++) {
        this.localSteps[i].id = i + 1;
      }
    },
    addStep() {
      if (this.isEmptyOrSpaces(this.step.step)) {
        this.alertInfo("Como minimo debe describir el paso.");
        return false;
      }
      var step = this.cloneObject(this.step);
      step.id = this.localSteps.length + 1;
      this.localSteps.push(step);
      this.step.step = "";
      this.step.result = "";
    },
    deleteStep(id) {
      this.localSteps.splice(id - 1, 1);
      this.ordenar();
    },
    up(id) {
      if (id > 1) {
        this.localSteps = this.cambiar(id - 2, id - 1, this.localSteps);
        this.ordenar();
      }
    },
    down(id) {
      if (id < this.localSteps.length) {
        this.localSteps = this.cambiar(id - 1, id, this.localSteps);
        this.ordenar();
      }
    },
    edit(id) {
      this.editDialog.show = true;
      this.editDialog.step = this.cloneObject(this.localSteps[id - 1]);
    },
    cerrar() {
      this.editDialog.show = false;
    },
    editFinish() {
      this.editDialog.show = false;
      this.localSteps[this.editDialog.step.id - 1] = this.cloneObject(
        this.editDialog.step
      );
      this.$emit("update:steps", this.localSteps);
    }
  }
};
</script>
<style>
.ta8 {
  width: 100%;
  height: 80px;
  min-height: 80px;
  border: 0px;
  border-radius: 5px;
  padding-left: 10px;
  -webkit-box-shadow: 0px 2px 14px #000;
  box-shadow: 1px 1px 1px 1px #444;
  background-color: #fff;
  color: #222;
  font-weight: normal;
  font-size: 12px;
  background: linear-gradient(#fff, #fff 28px, #dddddd 28px);
  background-size: 30px 30px;
  line-height: 30px;
}

.stepCol {
  white-space: nowrap;
}

table {
  width: 100%;
}

table .lastItem {
  text-align: right;
  width: 10%;
  padding-right: 10px;
}

.del span:hover {
  color: red;
  font-size: 18px;
  cursor: pointer;
}

.edit span:hover {
  color: blue;
  font-size: 18px;
  cursor: pointer;
}

.up span:hover {
  color: green;
  font-size: 18px;
  cursor: pointer;
}

.down span:hover {
  color: green;
  font-size: 18px;
  cursor: pointer;
}

.scrollingSteps {
  overflow-x: scroll;
  overflow-y: hidden;
  white-space: nowrap;
  min-height: 90px;
}

table.scroll {
  /* width: 100%; */
  /* Optional */
  /* border-collapse: collapse; */
  border-spacing: 0;
  border: 2px solid black;
}

table.scroll tbody,
table.scroll thead {
  display: block;
}

thead tr th {
  max-height: 30px;
  line-height: 30px;
  /* text-align: left; */
}

tbody tr :hover {
  cursor: pointer;
  background-color: #000 !important;
  color: #fff;
}

table.scroll tbody {
  height: 370px;
  overflow-y: auto;
  overflow-x: hidden;
}

tbody {
  border-top: 2px solid black;
}

tbody td,
thead th {
  /* width: 20%; */
  /* Optional */
  border-right: 1px solid black;
  /* white-space: nowrap; */
}

tbody td:last-child,
thead th:last-child {
  border-right: none;
}
</style>
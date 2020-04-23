<template>

  <div>
    <div class="card" style="min-height: 150px !important;">
      <div class="card-header">
        <div class="typo-line">
          <p class="category">
            <img width="128px" :src="project.avatar" />
            <br />

          </p>
          <blockquote>
            <h4>
              <span class="ti-write"></span> {{project.name}}</h4>
            <p v-html="project.resume"></p>
          </blockquote>
        </div>
      </div>
      <div class="card-body">

      </div>
      <div class="card-footer">
        <span class="click" @click="createNew()"><span class="ti-new-window"></span> Crear nuevo tipo de <strong>Epico</strong></span>
      </div>
    </div>

    <epic-design v-if="task==NEW" :epicType="epicForm" v-on:back="listEpic();" :inventions="inventions" :project="project"></epic-design>

    <div class="card-group" v-if="task==NONE">
      <div class="col-md-3" v-for="(epicType, index) in epicTypes" :key="index">
        <card style="background-color:khaki;">
          <drop-menu :title="tasktitle" style="float:left;top:-10px;" v-on:change="menuEpicType(option,epicType,index)" :options="epicOptions" :option.sync="option"></drop-menu>
          <template slot="header">
            <span class="card-title">
              <h5><strong> {{epicType.name}}</strong></h5>
            </span>
            <span class="row">
              <span class="col-md-12">
                {{epicType.definition}}
              </span>
            </span>
          </template>

        </card>
      </div>

    </div>
    <epic-design v-if="task==EDIT" :epicType="epicForm" :inventions="inventions" v-on:back="listEpic();" :project="project"></epic-design>


  </div>
</template>
<script>

  import EpicDesign from "./EpicDesign.vue";

  export default {
    name: "user-stories",
    computed: {

    },

    components: {
      EpicDesign,
    },

    props: {
      project: {},
      inventions: {},
    },

    data() {
      return {

        EDIT: 1,
        NEW: 0,
        NONE: -1,

        task: -1,

        epicForm: {
          name: "",
          definition: "",
          attributes: [],
          inventions: [],
        },
        tasktitle: `<small style="color:green"><i class="ti-angle-double-down"></i> Menu</small>`,
        option: -1,
        epicOptions: [
          { html: `<div><i style="color:green" class="ti-panel"></i> Editar</div>`, option: 1 }
          // { html: `<div><i style="color:red" class="ti-thumb-down"></i> Descativar</d>`, option: 0 },
        ],

        selectedDate: "",
        name: "",
        inventionVOs: [],
        epicTypes: [],
        refresh: true
      }
    },
    async mounted() {
      await this.getInventions();
      await this.getEpicTypes();
    },
    methods: {
      async listEpic() {
        await this.getEpicTypes();
        this.task = this.NONE;
      },

      createNew() {
        this.epicForm = {
          name: "",
          definition: "",
          attributes: [],
          inventions: [],
        },
          this.task = this.NEW
      },

      menuEpicType(option, epicForm, index) {

        if (option == this.EDIT) {
          this.epicForm = epicForm
          this.task = this.EDIT;
          //console.log(JSON.stringify(epicForm));
        }
      },

      async  getEpicTypes() {
        await this.axios.get(`/api/project/${this.project.code}/design/epic`).then(rs => {
          this.epicTypes = rs.data == null ? [] : rs.data;
        });
      },
      async  getInventions() {
        await this.axios.get("/api/invention/all/").then(rs => {
          this.inventionVOs = rs.data;
        });
      },
    }
  };
</script>
<style></style>
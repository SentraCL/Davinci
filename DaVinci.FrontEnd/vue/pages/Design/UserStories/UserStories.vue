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
        <span class="click" @click="createNew()"><span class="ti-new-window"></span> Crear nuevo tipo de <strong>Historia de Usuario</strong></span>
      </div>
    </div>

    <user-stories-design v-if="task==NEW" :userStoriesType="userStoriesForm" v-on:back="listUserStories();" :inventions="inventions" :project="project"></user-stories-design>

    <div class="card-group" v-if="task==NONE">
      <div class="col-md-3" v-for="(userStoriesType, index) in userStories" :key="index">
        <card style="background-color:khaki;">
          <drop-menu :title="tasktitle" style="float:left;top:-10px;" v-on:change="menuEpicType(option,userStoriesType,index)" :options="epicOptions" :option.sync="option"></drop-menu>
          <template slot="header">
            <span class="card-title">
              <h5><strong> {{userStoriesType.title}}</strong></h5>
            </span>
            <span class="row">
              <span class="col-md-12">
                {{userStoriesType.definition}}
              </span>
            </span>
          </template>

        </card>
      </div>

    </div>
    <user-stories-design v-if="task==EDIT" :userStories.sync="userStoriesForm" :inventions="inventions" v-on:back="listUserStories();" :project="project"></user-stories-design>
  </div>
</template>
<script>

  import userStoriesDesign from "./userStoriesItem.vue";

  export default {
    name: "user-stories",
    computed: {

    },

    components: {
      userStoriesDesign,
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

        userStoriesForm: {
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
        userStories: [],
        refresh: true
      }
    },
    async mounted() {
      await this.getInventions();
      await this.getTypesOfUserStories();
    },
    methods: {
      async listUserStories() {
        await this.getTypesOfUserStories();
        this.task = this.NONE;
      },

      createNew() {
        this.userStoriesForm = {
          name: "",
          definition: "",
          attributes: [],
          inventions: [],
        },
          this.task = this.NEW
      },

      menuEpicType(option, userStoriesForm, index) {

        if (option == this.EDIT) {
          this.userStoriesForm = userStoriesForm
          this.task = this.EDIT;
          //console.log(JSON.stringify(userStoriesForm));
        }
      },

      async  getTypesOfUserStories() {
        await this.axios.get(`/api/project/${this.project.code}/design/userstories`).then(rs => {
          this.userStories = rs.data == null ? [] : rs.data;
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
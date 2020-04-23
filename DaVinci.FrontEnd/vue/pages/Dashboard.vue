<template>
  <div>
    <h6>
      <input-text :label="titles.projects" :value.sync="projectName" autocomplete="off" v-if="projects.length>0">
      </input-text>
    </h6>
    <div class="card-group" v-if="ready">
      <div class="col-lg-3" v-for="proItem in filterProject">
        <project-item :isNew="proItem.isNew" v-on:clickAvatar="selectProject(proItem)" :name="proItem.name" :epics="proItem.epics" :userStories="proItem.userStories" :data="proItem.data" :avatar="proItem.avatar">
          <span slot="description">
            <p v-html="proItem.resume"></p>

            <small><a v-if="!proItem.isNew" :href="'/davinci/' + proItem.alias + '/'" target="_blank"><i class="ti-share"></i> ir al sitio..</a></small>
            <center v-if="proItem.isNew && filterProject.length>1 ">
              <button class="btn btn-xs btn-round btn-info" @click="createProject(proItem)">Crear Proyecto.</button>
            </center>
            
            
            
          </span>
        </project-item>
      </div>
      <div class="col-xl-8 col-lg-7 col-md-6" v-if="filterProject.length<=1">
        <project-form v-on:save="notifySave()" v-on:back="reload()" :title="currentProject.title" :project="currentProject" :isMyFirtsProject="projects.length==0"> 

        </project-form>
      </div>

      <div class="col-md-11" v-if="filterProject.length==1">
        <project-context :project="currentProject"></project-context>
      </div>

      <div class="col-lg-3">
        <project-item :name="projectName" :resume="createText" v-if="filterProject.length==0">
          <p slot="description" class="description">
            <br />
            <center>
              <img src="@/assets/img/leo.png" alt=" te ayudo?" />
              <h4>Veo que estas por comenzar un nuevo Proyecto!</h4>
              <p slot="description">
                El Objetivo de un <strong>Projecto DaVinci</strong> es tener el poder completo del diseño, para la ejecucion de <i>historias de usuarios, epicos y temas.</i>
              </p>
            </center>
          </p>
        </project-item>
      </div>
    </div>
  </div>
</template>
<script>
  import ProjectForm from "./Project/ProjectForm.vue";
  import ProjectItem from "./Project/ProjectItem.vue";
  import ProjectContext from "./Project/ProjectContext.vue";
  export default {
    components: {
      ProjectForm,
      ProjectItem,
      ProjectContext
    },
    data() {
      return {
        titles: {
          inventions: `<span class="ti-ruler-pencil"></span> Inventos`,
          projects: `<span class="ti-blackboard"></span> Proyectos`,
        },
        projectName: "",
        ready: false,
        viewCreate: false,
        projects: [],

      }
    },
    watch: {

      projectName() {
        if (this.viewCreate) {
          this.viewCreate = false;
        }
      }

    },
    created() {
      this.loadProjects();
    },

    computed: {
      currentProject() {
        var project = {}
        var title = "";
        if (this.filterProject.length == 1) {
          project = this.cloneObject(this.filterProject[0]);
          if (!project.isNew) {
            project.title = `<span class="ti-panel"></span> Modificar Proyecto`;
          }
        } else {
          project = this.templateNewProject();
        }
        return project;
      },
      createText() {
        var html = `Crear Projecto <strong>${this.projectName}</strong><br> `;
        return html;
      },
      filterProject() {
        var fproject = this.projectName.length == 0 ? this.cloneObject(this.projects) : [];
        if (this.viewCreate) {
          return []
        }
        var namesProjects = []
        if (fproject.length == 0) {
          //Agregar proyecto nuevo
          if (this.isNameProjectUniq()){
            var newProject = this.templateNewProject()
            newProject.name=this.projectName;
            fproject.push(newProject);
          }
          for (var p in this.projects) {
            if (this.projects[p].name.toUpperCase().indexOf(this.projectName.toUpperCase()) > -1) {
              this.projects[p].isNew = false;
              fproject.push(this.projects[p]);
              namesProjects.push(this.projects[p].name.toUpperCase())
            }
          }
        }
        /*
        //Ver que paso aca...
        if (this.projectName.length > 2 && namesProjects.indexOf(this.projectName.toUpperCase()) == -1) {
          var newProject = this.templateNewProject();
          fproject.unshift(newProject);
        }
        */
        return fproject;
      }
    },
    methods: {
      isNameProjectUniq(){
        var unique=true;
        this.projects.forEach(project=>{
          if (project.name==this.projectName)
          {
            unique =  false;
          }
        })
        return unique;
      },
      templateNewProject() {
        var project = {}
        var title = "";
        var name = "";
        if (this.projects.length == 0) {
          title = `<span class="ti-package"></span> Crear mi Primer Proyecto!`
        } else {
          title = `<span class="ti-map"></span> Crear Proyecto`
        }
        project.isNew = true;
        project.title = title;
        project.name = this.projectName;
        project.company = "";
        project.admin = "";
        project.email = "";
        project.resume = "";
        project.alias = this.projectName.replace(/\s/g, '');
        project.avatar64 = "";
        return project
      },
      createProject(proItem) {
        this.viewCreate = true;
      },

      selectProject(proItem) {
        this.projectName = proItem.name
      },
      async notifySave() {
        await this.loadProjects();
        this.projectName = ""
      },

      async reload() {
        this.projectName = "";
        await this.loadProjects();
      },


      async loadProjects() {
        await this.axios.post("/api/project/getAll/").then(rs => {
          this.projects = rs.data
        });
        this.ready = true;

      }
    }
  };
</script>
<style></style>
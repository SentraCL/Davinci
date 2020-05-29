<template>
  <div>
    <h6>
      <!--
      <div class="pull-right">
        <button class="btn" @click="doImport()"><i class="fa fa-folder"></i></button>
      </div>
      -->
      <input-text :label="titles.projects" :value.sync="projectName" autocomplete="off" v-if="projects.length>0">
      </input-text>
    
    </h6>
    <br/>
    <div class="card-group" v-if="ready">
      <div class="col-lg-3" v-for="proItem in filterProject">
        <project-item :isNew="proItem.isNew" v-on:clickAvatar="selectProject(proItem)" :project="proItem" v-on:reload="reload()">
          <span slot="description">
            <!-- ToDo: hacer que se corte texto de resumen. -->
            
            <details>
              <summary>Ver Descripcion</summary>
            <p id="more" v-html="proItem.resume"></p>
            </details>
            <!--<button @click="showDescription()" id="readMore">Read more <span id="dots">...</span></button> -->
            <small><a v-if="!proItem.isNew" :href="'/davinci/' + proItem.alias + '/'" target="_blank"><i class="ti-share"></i> ir al sitio..</a></small>
            <center v-if="proItem.isNew">
              <button v-if="filterProject.length!=1" class="btn btn-xs btn-round btn-info" @click="createProject(proItem)">Crear Proyecto.</button>
            </center>



          </span>
        </project-item>
      </div>
      <div class="col-xl-8 col-lg-7 col-md-6" v-if="filterProject.length==1">
        <project-form ref="projectForm" v-on:save="notifySave()" v-on:back="reload()" :title="currentProject.title" :project="currentProject" :isMyFirtsProject="projects.length==0">

        </project-form>
      </div>

      <div class="col-xl-8 col-lg-7 col-md-6" v-if="filterProject.length==0">
        <project-form ref="projectNewForm" v-if="currentProject.isNew" v-on:save="notifySave()" v-on:back="reload()" :title="currentProject.title" :project="currentProject" :isMyFirtsProject="projects.length==0">

        </project-form>
      </div>

      <div class="col-md-11" v-if="filterProject.length==1">
        <project-context :project="currentProject"></project-context>
      </div>

      <div class="col-lg-3">
        <project-item :project="proItem" :resume="createText" v-if="filterProject.length==0">
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
          projects: `<span class="ti-blackboard"></span> Crear Proyecto`,
        },
        projectName: "",
        ready: false,
        viewCreate: false,
        projects: [],
        editProject: false
      }
    },
    watch: {

      projectName() {
        if (this.viewCreate) {
          this.viewCreate = false;
          this.editProject = false;
        }
      }

    },
    created() {
      this.loadProjects();
      this.proItem = this.templateNewProject();
    },
   
    computed: {
      

      currentProject() {
        var project = {}
        var title = "";
        
        if (this.filterProject.length == 1 ) {
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

        if (this.editProject) {
          return this.projects.filter(this.equalsName);
        }

        if (this.viewCreate) {
          return [this.templateNewProject()]
        }


        var fproject =this.cloneObject(this.projects.filter(this.filterName));

        //Agregar proyecto vacio y nuevo, al inicio          
        if (this.isNameProjectUniq() && this.projectName.length>0) {
          var newProject = this.templateNewProject()
          newProject.name = this.projectName;
          fproject.unshift(newProject);
        }



        return fproject;
      }
    },
    methods: {

       showDescription(){
        
        var dots = document.getElementById("dots");
        var moreText = document.getElementById("more");
        var btnText = document.getElementById("readMore");

        if (dots.style.display === "none") {
        dots.style.display = "inline";
        btnText.innerHTML = "Read more"; 
        moreText.style.display = "none";
        } else {
        dots.style.display = "none";
         btnText.innerHTML = "Read less"; 
         moreText.style.display = "inline";
    }
    },
      equalsName(project) {
        return (project.name == this.projectName)
      },
      filterName(project) {
        return (project.name.toUpperCase().indexOf(this.projectName.toUpperCase()) > -1)
      },
      isNameProjectUniq() {
        var unique = true;
        this.projects.forEach(project => {
          if (project.name == this.projectName) {
            unique = false;
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
        if (proItem.isNew){
          return false;
        }
        this.projectName = proItem.name
        this.editProject = true;
      },
      
      doImport(){
        this.ProjectItem.doImport();
      },
      async notifySave() {
        this.projectName = ""
        this.editProject = false;
        this.viewCreate = false;
        await this.loadProjects();

      },

      async reload() {
        this.projectName = "";
        this.editProject = false;
        await this.loadProjects();
      },


      async loadProjects() {
        await this.axios.post("/api/project/getAll/").then(rs => {
          this.projects = rs.data
        });
        this.ready = true;
      },
    }
  };
</script>
<style>
#more {display: inline;}
</style>
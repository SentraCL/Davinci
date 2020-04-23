<template>
  <span>
    <span v-if="filterInventions.length>0">
      <small><span class="ti-hand-open"> </span>Arrastre los inventos hacia los proyectos donde desea que persistan.</small>
      <h6>
        <input-text :label="titles.inventions" :value.sync="inventionName" autocomplete="off" v-if="projects.length>0">
        </input-text>
      </h6>
      <div class="row">
        <div class="col-md-12">
          <div class="scrolling">
            <span class="col-lg-6" v-on:mouseover="setInvention(invention)" :title="invention.resume" v-for="invention in filterInventions">
              <drag class="drag invention" :transfer-data="{ invention }">
                <img :src="invention.icon" width="64px" />
                <br />
                <span class="labelPrj">{{invention.name}}</span>
              </drag>
            </span>
          </div>
        </div>
      </div>
    </span>
    <div class="row" v-if="filterInventions.length==0">
      <div class="col-md-12">
        <!-- <h6><span class="ti-comment"></span> No existen inventos para agregar a los proyectos.</h6>-->
      </div>
    </div>
    <hr />
    <h6 v-if="!isWorkingInAProject">
      <input-text :label="titles.projects" :value.sync="projectName" autocomplete="off" v-if="projects.length>0">
      </input-text>
    </h6>
    <div class="card-group">
      <div class="col-xl-8 col-lg-7 col-md-6" v-if="isWorkingInAProject">
        <span class="card">
          <span class="card-body">
            <h2 v-html="titles.currentProject"></h2>
            <div class="col-md-12">
              <h3>{{projectContext.company}}</h3>
              <blockquote v-html="projectContext.resume">
              </blockquote>
            </div>
          </span>
          <span class="card-footer float-right">
            <strong>Gestor Responsable </strong> <span class="ti-user"></span> {{projectContext.admin}} - <span class="ti-email"></span> {{projectContext.email}}
          </span>
        </span>
        <span class="btn-group">
          <button class="btn btn-info" round @click="back">
            <i class="ti-blackboard"></i> Ver todos los Projectos
          </button>
        </span>
      </div>
      <div class="col-lg-4" v-for="proItem in filtersProjects">
        <!--
        <span class="col-md-12" v-if="isWorkingInAProject && loadInvention && projectContext.inventions.length==0"><span class="ti-settings"></span> Arrastre un Invento para ser poblado.</span>
        -->
        <drop class="drop" @drop="handleDrop(proItem)">
          <project-item v-on:clickAvatar="makeMyProject(proItem)" :name="proItem.name" :epics="proItem.epics" :userStories="proItem.userStories" :data="proItem.data" :avatar="proItem.avatar">
            <span slot="description" v-if="loadInvention">
              <span v-if="projectInventions[proItem.name]" v-for="inventionPrj in projectInventions[proItem.name]">
                <span class="miniPrj min-invention "><img :src="inventionPrj.icon" width="16px" /> {{inventionPrj.name}} <span class="fa fa-close close" @click="removeInvention(proItem.name,inventionPrj)"></span> </span>
              </span>
            </span>
          </project-item>
        </drop>
      </div>
    </div>
    <div class="row" v-if="isWorkingInAProject && loadInvention">
      <porfolio-manager v-if="projectContext.inventions.length>0" :project.sync="projectContext"></porfolio-manager>
    </div>
    <m-dialog id="errorDrop" :title="titles.delete" :show.sync="errDelAlert.show" :isClose.sync="errDelAlert.close">
      <span slot="dialog">
        <div class="row">
          <span class="col-md-12">
            No es posible eliminar el Invento <strong>{{errDelAlert.nameInvention}}</strong>, puesto este es necesario para los siguientes inventos:
          </span>
          <span class="col-md-12" v-for="artifacUse in errDelAlert.artifacts" v-html="artifacUse"></span>
        </div>
      </span>
      <span slot="actions">
        <span class="btn-group">
          <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDeleteDialog">
            Cerrar
          </d-button>
        </span>
      </span>
    </m-dialog>
  </span>
</template>
<script>
  import ProjectItem from "../Project/ProjectItem.vue";
  import PorfolioManager from "./PorfolioManager.vue";
  export default {
    name: "Porfolio",
    components: {
      ProjectItem,
      PorfolioManager
    },
    computed: {
      isWorkingInAProject() {
        if (this.filtersProjects.length == 1) {
          this.makeMyProject(this.filtersProjects[0]);
          return true;
        }
        if (Object.keys(this.projectContext).length > 0) {
          return this.projectName == this.projectContext.name;
        } else return false;
      },
      filtersProjects: function () {
        var filtersProjects = [];
        if (this.projectName == "") {
          return this.projects
        }
        for (var i in this.projects) {
          var pro = this.projects[i]
          if (pro.name.toUpperCase().indexOf(this.projectName.toUpperCase()) > -1) {
            filtersProjects.push(pro);
          }

        }
        return filtersProjects;
      },
      filterInventions: function () {
        var filterInventions = [];
        for (var i in this.inventionVOs) {
          var inv = this.inventionVOs[i]
          if (inv.name.toUpperCase().indexOf(this.inventionName.toUpperCase()) > -1 || this.inventionName == "") {
            var include = false;
            var invCount = 0;
            var artiNames = ""
            for (var a in inv.artifacts) {
              var artifact = inv.artifacts[a]
              if (!artifact.isEssential) {
                invCount++
                include = true;
                artiNames += `\n${artifact.name} : ${artifact.nickName}`
              }
            }
            if (include) {
              inv.resume = `${inv.name} : Depende de ${invCount} inventos para funcionar: ${artiNames}`
            }
            filterInventions.push(inv);
          }

        }
        return filterInventions;
      }
    },
    props: {
      artifactForm: {}
    },

    data() {
      return {
        errDelAlert: {
          show: false,
          close: false,
          nameInvention: "",
          artifacts: []
        },
        titles: {
          inventions: `<span class="ti-ruler-pencil"></span> Inventos.`,
          projects: `<span class="ti-blackboard"></span> Proyectos.`,
          delete: `<span class="ti-alert"></span> Acción Cancelada.`,
          currentProject: ""
        },
        projectName: "",
        projectContext: {},
        inventionName: "",
        loadInvention: false,
        inventionVOs: [],
        projects: [],
        currentInvention: {},
        projectInventions: {}
      }
    },
    async created() {
      await this.getInventions();
      await this.loadProjects();
    },
    methods: {


      isUsed(nameInvention, projectName) {
        var uses = []
        for (var i in this.projectInventions[projectName]) {
          var inv = this.projectInventions[projectName][i]
          for (var a in inv.artifacts) {
            var artifact = inv.artifacts[a]
            if (!artifact.isEssential && artifact.name == nameInvention) {

              uses.push(`<br/>El atributo <strong>${artifact.name}</strong> en <strong>${inv.name}</strong>, es de tipo <strong style="color:#4e1d00">${nameInvention}</strong> `);
            }
          }
        }
        return uses;
      },

      closeDeleteDialog() {
        this.errDelAlert.show = false;
      },
      removeInvention(projectName, invention) {
        this.loadInvention = false;
        var uses = this.isUsed(invention.name, projectName);
        if (uses.length > 0) {
          this.errDelAlert.show = true;
          this.errDelAlert.close = false;
          this.errDelAlert.artifacts = uses;
          this.errDelAlert.nameInvention = invention.name;
        } else {
          var removeIndex = 0;
          for (var i in this.projectInventions[projectName]) {
            if (this.projectInventions[projectName][i].name == invention.name) {
              removeIndex = i;
              break;
            }
          }
          this.projectInventions[projectName].splice(removeIndex, 1)

        }
        if (this.isWorkingInAProject) {
          this.makeMyProject(this.projectContext);
        }
        this.loadInvention = true;
      },
      setInvention(invention) {
        this.currentInvention = invention;
      },
      async loadProjects() {
        await this.axios.post("/api/project/getAll/").then(rs => {
          this.projects = rs.data
          for (var p in this.projects) {
            var project = this.projects[p]
            //console.log(project.code);
            if (project.repository != null) {

              for (var i in this.inventionVOs) {
                var invention = this.inventionVOs[i]

                for (var r in project.repository) {
                  var repo = project.repository[r]

                  if (invention.code == repo._id) {
                    //console.log(`Agregando al proyecto ${project.name} el invento ${invention.name} `);
                    this.addInventoToProject(project, invention)
                  }
                }
              }
              //this.projectInventions[project.name].push(inventionNew);
            }

          }

        });
      },
      async getInventions() {
        await this.axios.get("/api/invention/all/").then(rs => {
          this.inventionVOs = rs.data;
        });
      },
      back() {
        this.projectName = ""
      },

      getInventionByTypeRef(code) {
        for (var i in this.inventionVOs) {
          var inv = this.inventionVOs[i]
          if (inv.code == code) {
            return inv;
          }
        }
        return null;
      },
      addInventoToProject(project, inventionNew) {
        if (this.projectInventions[project.name] == null) {
          this.projectInventions[project.name] = [];
        }
        if (Object.keys(inventionNew).indexOf('artifacts') > -1)
          if (this.projectInventions[project.name].indexOf(inventionNew) == -1 && inventionNew.artifacts != null) {
            for (var a in inventionNew.artifacts) {
              var artifact = inventionNew.artifacts[a]
              if (!artifact.isEssential) {
                var artifactInv = this.getInventionByTypeRef(artifact.typeRef);
                if (artifactInv != null) {
                  this.addInventoToProject(project, artifactInv);
                }
              }
            }
            this.projectInventions[project.name].push(inventionNew);
          }
        this.loadInvention = true;
      },

      handleDrop(project, event) {
        this.loadInvention = false;
        this.addInventoToProject(project, this.currentInvention)
        if (this.isWorkingInAProject) {
          this.makeMyProject(this.projectContext);
        }
        this.loadInvention = true;
      },
      makeMyProject(project) {

        this.titles.currentProject = `<span class="ti-blackboard"></span> ${project.name}.`;
        this.projectName = project.name
        var upInventions = []
        upInventions = [];

        for (var i in this.projectInventions[this.projectContext.name]) {
          var invention = this.projectInventions[this.projectContext.name][i];
          upInventions.push(invention);
        }

        this.projectContext = {
          name: project.name,
          resume: project.resume,
          company: project.company,
          admin: project.admin,
          email: project.email,
          inventions: upInventions,
          code: project.code
        };

      }
    }
  }
</script>
<style scoped>
  /* width */
  ::-webkit-scrollbar {
    width: 20px;
  }

  /* Track */

  /* Handle */
  ::-webkit-scrollbar-thumb {
    background: #FFF;
    border-radius: 10px;
  }

  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background: #4e1d00;

    transform: translate3d(0, 0, 0);
    backface-visibility: hidden;
    perspective: 1000px;
  }

  .drag,
  .drop {

    font-family: sans-serif;
    display: inline-block;
    border-radius: 10px;
    background: transparent;
    color: #000;
    position: relative;
    padding: 2px;
    min-width: 80px;
    text-align: center;
    vertical-align: top;
  }

  .drag {
    cursor: grab !important;
    color: #000;
    font-size: 10px;
    text-transform: uppercase;

  }


  .drag:hover {
    animation: shake 0.82s cubic-bezier(.36, .07, .19, .97) both;
    transform: translate3d(0, 0, 0);
    backface-visibility: hidden;
    perspective: 1000px;
  }

  .drop {
    background: transparent;
    color: #000;
  }


  .scrolling {
    overflow-x: scroll;
    overflow-y: hidden;
    white-space: nowrap;
    min-height: 90px;
  }


  .close {
    cursor: pointer;
  }

  .labelPrj {
    font-size: 12px;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif;
    text-transform: uppercase;

  }

  .invention {
    margin-top: 3px;
    transition: .3s;
  }

  .invention:hover {
    font-size: 1.0em;
    margin-top: -5px;
  }

  .min-invention {
    cursor: default;
    margin-top: 3px;
    transition: .3s;
  }

  .min-invention:hover {
    font-size: 1.0em;
    margin-top: -5px;
  }


  .miniPrj {
    cursor: pointer;
    font-size: 12px;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif;
    text-transform: capitalize;
    color: #000;
    display: inline-block;
    border-radius: 5px;
    background: #E4D6C0;
    position: relative;
    padding: 8px;
    min-width: 20px;
    border-right: 1px solid #906538;
    border-bottom: 1px solid #4e1d00;
    text-align: center;
    vertical-align: top;
  }
</style>
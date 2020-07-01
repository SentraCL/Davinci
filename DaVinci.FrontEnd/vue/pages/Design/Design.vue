<template>
    <div>
        <h6>
        <div class="row">
            <div class="col-md-4">
                <input-text :label="titles.projects" :value.sync="projectName" autocomplete="off"></input-text>
            </div>
            
            <div class="col-md-8 text-right">
                <label-list @clicked="filterByEnterprise" label="Empresa" :list="enterprise" keyValue="EnterpriseId" keyLabel="Name"></label-list>
            </div>
        </div>
        </h6>


        <div v-if="ready">
            <div v-if="filterProject.length>=1 && option==-1">
              
                <card>
                    <h4>Objetivo</h4>
                    <p>Este modulo usted podra crear categorias de Epicos e Historias de Usuario, acorde al negocio de
                        su proyecto.</p>
                </card>
              
                <div class="card-group">
                    <div class="col-lg-4" v-for="proItem in filterProject">





                        <project-item :name="proItem.name" :project="proItem" hiddenTools>
                            <span slot="description">
                                <blockquote>
                                    <!--<p v-html="proItem.resume"></p>-->
                                    <h5>
                                        <i class="fa fa-magic"></i> Diseñar...
                                    </h5>
                                    <span class="row">
                                        <span class="icon-design col-md-6">
                                            <img width="56px" title="Bla bla bla EPICO"
                                                src="@/assets/img/SubDomain/d.epic.png" @click="CreateEpic(proItem)" />
                                            <br />Tipo de Epico
                                        </span>

                                        <span class="icon-design col-md-6">
                                            <img width="56px" title="Bla bla bla HISTO..."
                                                src="@/assets/img/SubDomain/d.us.png"
                                                @click="CreateUserStories(proItem)" />
                                            <br />Tipo de Historia de Usuario
                                        </span>
                                    </span>
                                </blockquote>
                            </span>
                        </project-item>
                    </div>
                </div>
            </div>
        </div>

        <div class="col-md-12" v-if="filterProject.length==1">
            <button class="btn btn-back btn-warning btn-xs rigth" @click="reload()" v-if="option>0">Ver todos los
                proyectos.</button>
            <div>
                <user-stories v-if="option==USER_HISTORY" :inventions="inventions[projectCode]" :project="project">
                </user-stories>
                <epic v-if="option==EPIC" :inventions="inventions[projectCode]" :project="project"></epic>
            </div>
        </div>
    </div>
</template>
<script>
    import ProjectForm from "../Project/ProjectForm.vue";
    import ProjectItem from "../Project/ProjectItem.vue";
    import Epic from "./Epic/Epic.vue";
    import UserStories from "./UserStories/UserStories.vue";
    export default {
        components: {
            ProjectForm,
            ProjectItem,
            Epic,
            UserStories
        },
        data() {
            return {
                option: -1,
                USER_HISTORY: 1,
                EPIC: 2,
                projectName: "",
                inventions: {},
                projectCode: String,
                ready: false,
                enterprise:[],
                filter:[],
                projects: [],
                project: {},
                titles: {
                    projects: `<span class="ti-write"></span> PROYECTO`
                }
            };
        },
        watch: {
            projectName() {
                if (this.filterProject.length == 1) {
                    this.project = this.filterProject[0];
                }
            }
        },
        async mounted() {
            await this.getAllEnterprise();
            await this.loadProjects();
        },

        computed: {
            currentProject() {
                var project = {};
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
                
                let projects=[]
                this.projects.forEach(project=>{
                if(!this.filter.includes(project.enterprise)){
                    projects.push(project);  
                }
                })
                var fproject =
                    this.projectName.length == 0 ? this.cloneObject(projects) : [];
                var namesProjects = [];
                if (fproject.length == 0) {
                    for (var p in projects) {
                        if (
                            this.projects[p].name
                                .toUpperCase()
                                .indexOf(this.projectName.toUpperCase()) > -1
                                && !this.filter.includes(projects[p].enterprise)
                        ) {
                            projects[p].isNew = false;
                            fproject.push(projects[p]);
                            namesProjects.push(projects[p].name.toUpperCase());
                        }
                    }
                }
                return fproject;
            }
        },
        methods: {
            async getAllEnterprise(){
              await this.axios
            .get("/api/enterprise/" )
            .then(rs => {
                this.enterprise=rs.data;
                })  
            },
            filterByEnterprise(list){
                this.filter=list;
            },
            async CreateUserStories(proItem) {
                this.projectCode = proItem.code;
                await this.loadInventionsByProjectCode(this.projectCode);
                this.projectName = proItem.name;
                this.option = this.USER_HISTORY;
            },
            async CreateEpic(proItem) {
                this.projectName = proItem.name;
                this.projectCode = proItem.code;
                await this.loadInventionsByProjectCode(this.projectCode);
                this.option = this.EPIC;
            },

            async loadInventionsByProjectCode(projectCode) {
                if (this.inventions[projectCode] == null) {
                    await this.axios
                        .post("/api/project/inventions/", {
                            projectCode: projectCode
                        })
                        .then(rs => {
                            this.inventions[projectCode] = rs.data;
                        });
                }
                return this.inventions[projectCode];
            },

            async reload() {
                this.option = -1;
                this.projectName = "";
                await this.loadProjects();
            },

            async loadProjects() {
                this.projects = [];
                await this.axios.post("/api/project/getAll/").then(rs => {
                    rs.data.forEach(project => {
                        if (project.repository && project.repository.length > 0) {
                            this.projects.push(project);
                        }
                    });
                });
                this.ready = true;
            }
        }
    };
</script>
<style>
    .btn-back {
        z-index: 99;
        right: 30px;
        position: absolute !important;
    }

    .icon-design {
        position: relative;
        min-height: 64px;
        cursor: pointer;
        font-size: 12px;
    }

    .icon-design:hover {
        top: -1;
        left: -1;
        font-size: 12px;
    }
</style>
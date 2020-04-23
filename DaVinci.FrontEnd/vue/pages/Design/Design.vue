<template>
    <div>
        <h6>
            <input-text :label="titles.projects" :value.sync="projectName" autocomplete="off">
            </input-text>
        </h6>
        <div class="card-group" v-if="ready">
            <div class="col-lg-3" v-for="proItem in filterProject" v-if="filterProject.length>=1 && option==-1">
                <project-item :name="proItem.name" :epics="proItem.epics" :userStories="proItem.userStories" :data="proItem.data" :avatar="proItem.avatar">
                    <span slot="description">
                        <blockquote>
                            <p v-html="proItem.resume"></p>
                            <button class="btn btn-xs btn-success col-md-12" @click="CreateEpic(proItem)">Crear Epico.</button>
                            <button class="btn btn-xs btn-success col-md-12" @click="CreateUserStories(proItem)">Historias de Usuario.</button>
                        </blockquote>
                    </span>
                </project-item>
            </div>
        </div>


        <div class="col-md-12" v-if="filterProject.length==1">
            <button class="btn btn-back btn-warning btn-xs rigth" @click="reload()"  v-if="option>0" >Ver todos los proyectos.</button>
            <div>
                <user-stories v-if="option==USER_HISTORY" :inventions="inventions[projectCode]" :project="project"></user-stories>
                <epic v-if="option==EPIC" :inventions="inventions[projectCode]"  :project="project"></epic>
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
                projects: [],
                project: {},
                titles: {
                    projects: `<span class="ti-write"></span> PROYECTO`,
                },
            }
        },
        watch: {
            projectName() {
                if (this.filterProject.length == 1) {
                    this.project = this.filterProject[0];
                }
            }

        },
        async mounted() {
            await this.loadProjects();
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
                var namesProjects = []
                if (fproject.length == 0) {
                    for (var p in this.projects) {
                        if (this.projects[p].name.toUpperCase().indexOf(this.projectName.toUpperCase()) > -1) {
                            this.projects[p].isNew = false;
                            fproject.push(this.projects[p]);
                            namesProjects.push(this.projects[p].name.toUpperCase())
                        }
                    }
                }
                return fproject;
            }
        },
        methods: {


            async CreateUserStories(proItem) {
                this.projectCode = proItem.code;
                await this.loadInventionsByProjectCode(this.projectCode)
                this.projectName = proItem.name;
                this.option = this.USER_HISTORY;
            },
            async CreateEpic(proItem) {
                this.projectName = proItem.name;
                this.projectCode = proItem.code;
                await this.loadInventionsByProjectCode(this.projectCode)
                this.option = this.EPIC;
            },

            async loadInventionsByProjectCode(projectCode) {
                if (this.inventions[projectCode] == null) {
                    await this.axios.post("/api/project/inventions/", {
                        projectCode: projectCode
                    }).then(rs => {
                        this.inventions[projectCode] = rs.data
                    });
                }
                return this.inventions[projectCode]
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
</style>
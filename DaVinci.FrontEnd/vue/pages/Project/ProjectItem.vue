<template>

    <div class="card card-user">
        <div class="image">
            <img src="@/assets/img/toolbar.jpg" alt="...">
        </div>
        <div class="content">
            <div class="author">
                <a href="#">
                    <img v-if="avatar" class="avatar border-white" @click="avatarClick()" :src="avatar" alt="...">
                    <img v-if="!avatar" class="avatar border-white" @click="avatarClick()" src="@/assets/img/davinci-logo.png" alt="...">
                </a>
                <h4 class="title" v-if="name">{{name}}<br>
                    <span style="color:greenyellow"><small>@{{alias}}</small></span>
                </h4>
            </div>
            
            <div class="col-lg-12">
                <slot name="description">

                </slot>
                </p>
            </div>
        </div>
        <hr>
        <div class="text-center">
            <div class="col-lg-12" v-if="isNew">
                <div class="row">
                    <div class="col-md-4">
                        <small title="Exportar proyecto para enviar a otro sistema Davinci"><i class="fa fa-briefcase"></i> Importar</small>
                    </div>
                    <div class="col-md-4">

                    </div>
                    <div class="col-md-4">

                    </div>
                </div>
            </div>
            <div class="col-lg-12" v-if="!isNew">
                <div class="row">
                    <div class="col-md-4" v-if="epics">
                        <h6>{{epics}}<br><small>Epicos</small></h6>
                    </div>
                    <div class="col-md-4" v-if="userStories">
                        <h6>{{userStories}}<br><small>Historias de Usuario</small></h6>
                    </div>
                    <div class="col-md-4" v-if="data">
                        <h6>{{data}}<br><small>Datos</small></h6>
                    </div>
                </div>

                <div class="row">
                    <div class="col-md-4">
                        <a href="#" style="color:black" @click="showExport" title="Exportar proyecto para enviar a otro sistema Davinci">
                            <small>
                                <i class="fa fa-briefcase"></i> Exportar
                            </small>
                        </a>
                    </div>

                    <div class="col-md-4">
                        <a href="#" style="color:black" title="Genera reporte de Epicos, Actividades e Historias de Usuario">
                            <small>
                                <i class="fa fa-file-excel-o"></i> Reporte
                            </small>
                        </a>
                    </div>
                    <div class="col-md-4">
                        <a href="#" style="color:black" @click="showCopy" title="Crea copia de Proyecto">
                            <small>
                                <i class="fa fa-copy"></i> Crear Copia
                            </small>
                        </a>
                    </div>
                </div>
            </div>
            <br />
        </div>

        <m-dialog :id="project.code" :title="copyDialog.title" :show.sync="copyDialog.show" :isClose.sync="copyDialog.close">
            <span slot="dialog">
                <div class="row">
                    <span class="col-md-8">
                        <div class="row">
                            <span class="col-md-2">
                                <img v-if="avatar" class="avatar border-white" @click="avatarClick()" :src="avatar" alt="...">
                                <img v-if="!avatar" class="avatar border-white" @click="avatarClick()" src="@/assets/img/davinci-logo.png" alt="...">
                            </span>
                            <span class="col-md-6">
                                <h3><i class="fa fa-copy"></i> Copiar Proyecto {{project.name}}</h3>

                            </span>
                        </div>


                        <div class="row">
                            <span class="col-md-12">
                                <input-text label="Nombre Copia" v-model="copyProject.name" autocomplete="off">
                                </input-text>
                                <hr />
                                <div>
                                    <check-box :value.sync="copyProject.users" label="Usuarios"></check-box>
                                    <check-box :value.sync="copyProject.epics" label="Tipos de Epicos"></check-box>
                                    <check-box :value.sync="copyProject.userStories" label="Tipos de Historias de Usuario"></check-box>
                                    <check-box :value.sync="copyProject.data" label="Datos"></check-box>
                                </div>
                            </span>
                        </div>
                    </span>
                    <span class="col-md-4">
                        <img src="@/assets/img/leo.png"><br />
                        Antes de crear una copia de <i>{{project.name}}</i>, Debes Asignarle un nombre Diferente, y seleccionar que atributos desea replicar.
                    </span>
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn" round @click.native.prevent="doCopy">
                        Crear Copia
                    </d-button>
                    <d-button type="warning" class="btn" round @click.native.prevent="closeDialog">
                        Cerrar
                    </d-button>
                </span>
            </span>
        </m-dialog>

         <m-dialog :id="project.code" :title="exportDialog.title" :show.sync="exportDialog.show" :isClose.sync="exportDialog.close">
            <span slot="dialog">
                <div class="row">
                    <span class="col-md-8">
                        <div class="row">
                            <span class="col-md-2">
                                <img v-if="avatar" class="avatar border-white" @click="avatarClick()" :src="avatar" alt="...">
                                <img v-if="!avatar" class="avatar border-white" @click="avatarClick()" src="@/assets/img/davinci-logo.png" alt="...">

                            </span>
                            <span class="col-md-6">
                                <h3><i class="fa fa-export"></i> Exportar Proyecto {{project.name}}</h3>
                            </span>
                        </div>
                        

                        <div class="row">
                            <span class="col-md-12">
                                <hr />
                                <div>
                                    <check-box :value.sync="exportProject.avatar" label="Logo"></check-box>
                                    <check-box :value.sync="exportProject.users" label="Usuarios"></check-box>
                                    <check-box :value.sync="exportProject.epics" label="Tipos de Epicos"></check-box>
                                    <check-box :value.sync="exportProject.userStories" label="Tipos de Historias de Usuario"></check-box>
                                    <check-box :value.sync="exportProject.data" label="Datos"></check-box>
                                </div>
                            </span>
                        </div>
                    </span>
                    <span class="col-md-4">
                        <img src="@/assets/img/leo.png"><br />
                        Antes de exportar el proyecto <i>{{project.name}}</i>, Debes Asignarle un nombre Diferente, y seleccionar que atributos desea replicar.
                    </span>
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn" round @click.native.prevent="doExport">
                        Exportar
                    </d-button>
                    <d-button type="warning" class="btn" round @click.native.prevent="closeDialog">
                        Cerrar
                    </d-button>
                </span>
            </span>
        </m-dialog>

    </div>


    

</template>
<script>

    export default {
        name: "project-item",
        computed: {
            alias() {
                return this.name.replace(/\s/g, '');
            }
        },
        props: {
            project: {},
            resume: String,
        },

        data() {
            return {
                name: this.project.name,
                epics: this.project.epics,
                userStories: this.project.userStories,
                data: this.project.data,
                avatar: this.project.avatar,
                isNew: this.project.isNew,
                copyDialog: {
                    show: false,
                    close: false,
                    title: "",
                    html: ""
                },
                 exportDialog: {
                    show: false,
                    close: false,
                    title: "",
                    html: ""
                },
                copyProject: {
                    code: this.project.code,
                    name: "Copia de " + this.project.name,
                    users: 0,
                    epics: 0,
                    userStories: 0,
                    data: 0
                },
                 exportProject: {
                    code: this.project.code,
                    name: this.project.name,
                    users: 0,
                    epics: 0,
                    userStories: 0,
                    data: 0,
                    avatar: 0
                }
            }
        },
        updated() {
            this.name = this.project.name;
            this.epics = this.project.epics;
            this.userStories = this.project.userStories;
            this.data = this.project.data;
            this.avatar = this.project.avatar;
            this.isNew = this.project.isNew;
            this.copyName = "Copia " + this.project.name;
        },
        methods: {
            closeDialog() {
                this.copyDialog.show = false;
                this.exportDialog.show = false;
            },
            showCopy() {
                this.copyDialog.show = true;
            },
            showExport() {
                this.exportDialog.show = true;
            },
            async doExport(){
                await this.axios.post("/api/project/export/", this.exportProject).then(rs => {
                        this.$emit("reload")
                        const data = JSON.stringify(rs.data);
                         const blob = new Blob([data], {type: 'text/plain'})
                         const e = document.createEvent('MouseEvents'),
                         a = document.createElement('a');
                         a.download = this.project.name+".json";
                         a.href = window.URL.createObjectURL(blob);
                         a.dataset.downloadurl = ['text/json', a.download, a.href].join(':');
                         e.initEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
                         a.dispatchEvent(e);
                         
                    });
            },
            async doCopy() {
                var projects = []
                await this.axios.post("/api/project/getAll/").then(rs => {
                    projects = rs.data
                });
                var copyName = this.copyProject.name;
                var filter = projects.filter(function (project) {
                    return project.name == copyName
                })

                if (filter.length == 0) {
                    await this.axios.post(`/api/project/copy/`, this.copyProject).then(rs => {
                        this.$emit("reload")
                        this.copyDialog.show = false;
                        this.alertInfo("Copia de Proyecto " + copyname + ", fue creado")
                    });
                }else{
                    this.alertInfo(`El Proyecto ${copyName} ya Existe.`);                    
                }
            },
            avatarClick() {
                this.$emit("clickAvatar");
            },
            editProject() {
                this.$router.push('adminProject');
            }
        }
    };
</script>
<style></style>
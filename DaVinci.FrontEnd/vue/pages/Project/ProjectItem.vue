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
                    <div class="col-md-4" >
                        <small title="Exportar proyecto para enviar a otro sistema Davinci"><i class="fa fa-briefcase"></i> Importar</small>
                    </div>
                    <div class="col-md-4" >
                        
                    </div>
                    <div class="col-md-4" >
                        
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
                    <div class="col-md-4" >
                        <a href="#" style="color:black" @click="exportPro" title="Exportar proyecto para enviar a otro sistema Davinci">
                            <small>
                                <i class="fa fa-briefcase"></i> Exportar
                            </small>
                        </a>
                    </div>
                    <div class="col-md-4" >
                        <a href="#" style="color:black"  title="Genera reporte de Epicos, Actividades e Historias de Usuario">
                            <small>
                                <i class="fa fa-file-excel-o"></i> Reporte
                            </small>
                        </a>
                    </div>
                    <div class="col-md-4" >
                        <a href="#" style="color:black"  @click="copyPro" title="Crea copia de Proyecto">
                            <small>
                                <i class="fa fa-copy"></i> Crear Copia
                            </small>
                        </a>
                    </div>
                </div>                            
            </div>
            <br/>    
        </div>
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
            project:{},
            resume: String,
        },

        data() {
            return {
                name:  this.project.name,
                epics:  this.project.epics,
                userStories:  this.project.userStories,
                data:  this.project.data,
                avatar:  this.project.avatar,
                isNew : this.project.isNew
            }
        },

        methods: {
            copyPro(){
                console.log("Copiar Proyecto");
            },
            exportPro(){
                console.log("Exportar Proyecto");
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
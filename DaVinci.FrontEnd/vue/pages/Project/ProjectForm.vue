<template>
    <card class="card" :title="title">

        <m-dialog id="deleteDialog" :title="titleDelete" :show.sync="delAlert.show" :isClose.sync="delAlert.close">
            <span slot="dialog">
                <div>
                    Al Eliminar el proyecto <strong>{{project.name}}</strong>, borrara todo el contenido ,tanto como Epicos e Historias de Usuario.<br /> ¿Desea continuar?
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDeleteDialog">
                        No , Cancelar
                    </d-button>
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="dropProject">
                        Si, Eliminar
                    </d-button>
                </span>
            </span>
        </m-dialog>
        <div>

            <form @submit.prevent>
                <div class="row">

                    <div class="col-md-10">
                        <div class="row">
                            <div class="col-md-6">
                                <input-text type="text" label="Proyecto" :disabled="!isMyFirtsProject" title="Nombre que identifica al Proyecto" v-model="project.name">
                                </input-text>
                            </div>
                            <div class="col-md-6">
                                <combo-simple label="Empresa" :list="enterprise" keyValue="EnterpriseId" keyLabel="Name" :value.sync="projectForm.enterprise"></combo-simple>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-md-6">
                                <input-text type="text" label="Gestor Responsable" title="Gestor Responsable" v-model="projectForm.admin">
                                </input-text>
                            </div>
                            <div class="col-md-4">
                                <input-text type="email" label="Email Responsable" title="Email" v-model="projectForm.email">
                                </input-text>
                            </div>
                        </div>

                        <div class="col-md-12">
                            <div class="form-group">
                                <label>Proposito del Proyecto <small>(o descripción del Cliente.)</small></label>

                                <textarea rows="5" class="form-control border-input" title="Acerca del Proyecto" v-model="projectForm.resume">

                            </textarea>
                            </div>
                        </div>
                    </div>
                    <div class="col-md-2">
                        <avatar ref="avatar" :prefill="urlAvatar" @change="changeAvatar" :removable="false" :autoToggleAspectRatio="true" margin="2" radius="50" size="16" accept="image/jpeg,image/png,image/gif" buttonClass="btn btn-xs" :customStrings="settingAvatar">
                        </avatar>
                        <br />
                        <div class="clearfix"></div>
                    </div>
                </div>
                <div class="col-md-12 left">
                    <d-button type="info" round @click.native.prevent="back">
                        <span style="color:darkgreen" class="ti-back-left"></span> Volver
                    </d-button>
                    <d-button type="success" round @click.native.prevent="updateProject">
                        <span style="color:white" class="ti-save"></span> Guardar
                    </d-button>
                    <d-button v-if="!projectForm.isNew" type="danger" round @click.native.prevent="deleteProject">
                        <span style="color:yellow" class="ti-eraser"></span> Eliminar
                    </d-button>
                </div>

            </form>
        </div>


    </card>
</template>
<script>

    import avatar from "vue-picture-input";
    export default {
        components: {
            avatar
        },
        props: {
            title: String,
            isMyFirtsProject: Boolean,
            project: {}
            
        },
        computed: {
            titleDelete() {
                return `<span class="fa fa-question-circle-o"></span> Confirmar eliminación proyecto <strong>${this.project.name}</strong>`;
            },
            urlAvatar() {
                var urlAvatar = this.project.isNew ? null : `/api/project/avatar/${this.projectForm.alias}.png`
                var avatarChecker = new XMLHttpRequest(),
                    CheckThisUrl = urlAvatar;
                avatarChecker.open('get', CheckThisUrl, true);
                if (avatarChecker.readyState === 1) {
                    if (avatarChecker.status == 0) {
                        return urlAvatar
                    }
                    else {
                        return null
                    }
                }

                return urlAvatar
            }
        },
        data() {
            return {
                delAlert: {
                    mensaje: "",
                    tiempo: 10,
                    show: false
                },
                projectForm: this.project,
                avatar64: [],
                enterprise:[],
                settingAvatar: {
                    upload: "<h1>Fallo!</h1>",
                    selected: "<p>Avatar Seleccionado</p>", // HTML allowed
                    fileSize: "El tamaño excede los limites", // Text only
                    fileType: "Este archivo de imagen no es soportado.",
                    drag: "<h5>Arrastre o seleccione una imagen</h5>", // HTML allowed
                    tap: "Seleccione una imagen desde su galeria", // HTML allowed
                    change: "Elegir logo", // Text only
                    remove: "Eliminar logo", // Text only
                    select: "Seleccionar foto", // Text only
                    selected: "<p>Su avatar de presentacion a sido seleccionado</p>", // HTML allowed
                    fileSize: "el archivo pesa demasiado", // Text only
                    fileType: "Este tipo de archivo no es soportado." // Text only
                }
            };
        },
        updated() {
            this.projectForm = this.project;
        },
        async mounted() {
            await this.getAllEnterprise();
        },
        methods: {
            async getAllEnterprise(){
              await this.axios
         .get("/api/enterprise/" )
         .then(rs => {
            this.enterprise=rs.data;
            })  
            },
            async updateProject() {
                var status = false;
                if(this.projectForm.admin==""){
                    this.alertError("El campo gestor responsable no debe estar vacio.");
                }
                if(this.projectForm.email==""){
                    this.alertError("El campo email no debe estar vacio y debe tener un formato valido.");
                }
                if(this.projectForm.enterprise!=""){    
                    await this.axios.post("/api/project/save/", this.projectForm).then(rs => {
                        status = rs.data;
                    });
                    this.$emit("save");
                    this.alertSuccess(`Proyecto ${this.projectForm.name} , Guardado.`);
                }else{
                    this.alertError(`Debe de elegir una empresa.`);
                }
                
            },

            back() {
                this.$emit("back");
            },
            deleteProject() {
                this.delAlert.show = true;
                this.delAlert.close = true;
            },
            closeDeleteDialog() {
                this.delAlert.show = false;
            },
            async dropProject() {
                var status = false;
                await this.axios.post("/api/project/drop/", this.projectForm).then(rs => {
                    status = rs.data;
                });
                this.$emit("back");
                
                this.alertInfo("Proyecto Eliminado")
            },
            changeAvatar() {
                if (this.$refs.avatar.image) {
                    this.projectForm.avatar = [];
                    this.projectForm.avatar = this.$refs.avatar.image;
                }
            }
        }
    };
</script>
<style scoped>
    .logo {
        border: 2px solid #4E1D00;
        border-radius: 100px;
        padding: 1px;
    }
</style>
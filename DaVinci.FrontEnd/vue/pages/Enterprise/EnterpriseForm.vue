<template>
    <div>
        <m-dialog id="deleteDialog" :title="titleDelete" :show.sync="delAlert.show" :isClose.sync="delAlert.close">
            <span slot="dialog">
                <div>
                    Al Eliminar la Empresa <strong>{{enterprise.Name}}</strong>, se inhabilitarán todos los proyectos enlazados a esta.<br /> ¿Desea continuar?
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDeleteDialog">
                        No , Cancelar
                    </d-button>
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="dropEnterprise">
                        Si, Eliminar
                    </d-button>
                </span>
            </span>
        </m-dialog>
        <form @submit.prevent>
                <div class="row">
                    <div class="col-md-10">
                        <div class="row">
                            <div class="col-md-6">
                                <input-text type="text" label="Empresa" :disabled="!isMyFirtsEnterprise" title="Nombre que identifica a la Empresa" v-model="enterprise.Name">
                                </input-text>
                            </div>
                            <div class="col-md-6">
                                <input-text type="text" label="Rut" title="Rol unico tributario" v-model="enterprise.Rut">
                                </input-text>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-md-6">
                                <input-text type="text" label="Direccion" title="Direccion comercial" v-model="enterprise.Direction">
                                </input-text>
                            </div>
                        </div>

                        <div class="col-md-12">
                            <div class="form-group">
                                <label>Descripcion de la empresa</label>

                                <textarea rows="5" class="form-control border-input" title="Acerca de la Empresa" v-model="enterpriseForm.Description">

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
                    <d-button v-if="enterprise.Status==1" type="success" round @click.native.prevent="updateEnterprise">
                        <span style="color:white" class="ti-save"></span> Guardar
                    </d-button>
                    
                    <d-button v-if="enterprise.Status==0" type="success" round @click.native.prevent="recoveryEnterprise">
                        <span style="color:white" class="ti-save"></span> Restaurar
                    </d-button>
                    <d-button v-if="!isNew&&enterprise.Status==1" type="danger" round @click.native.prevent="deleteEnterprise">
                        <span style="color:yellow" class="ti-eraser"></span> Eliminar
                    </d-button>
                </div>

            </form>
        </div>
</template>
<script>
import avatar from "vue-picture-input";
export default {
    components: {
        avatar
    },
    name:"enterprise-form",
    props:{
        enterprise:{},
        title: String,
        isMyFirtsEnterprise: Boolean,
        isNew:Boolean
    },
    computed: {
            titleDelete() {
                return `<span class="fa fa-question-circle-o"></span> Confirmar eliminación de la Empresa <strong>${this.enterprise.Name}</strong>`;
            },
            urlAvatar() {
                var urlAvatar = this.enterprise.Avatar ? null : `/api/project/avatar/${this.enterpriseForm.alias}.png`
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
                enterpriseForm: this.enterprise,
                avatar64: [],
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
            this.enterpriseForm = this.enterprise;
        },
        async mounted() {
        },
        methods: {
            async updateEnterprise() {
                var status = false;
                if(this.enterpriseForm.Name==""){
                    this.alertError("El nombre de la empresa no debe estar vacio.");
                }else{

                    await this.axios.post("/api/enterprise/save/", this.enterpriseForm).then(rs => {
                        status = rs.data;
                });
                this.$emit("getAll");
                this.alertSuccess(`Empresa ${this.enterpriseForm.Name} , Guardada.`);
                this.$emit("back");
                }
            },

            back() {
                this.$emit("back");
            },
            deleteEnterprise() {
                this.delAlert.show = true;
                this.delAlert.close = true;
            },
            closeDeleteDialog() {
                this.delAlert.show = false;
            },
            async dropEnterprise() {
                var status = false;
                await this.axios.post("/api/enterprise/drop/", this.enterpriseForm).then(rs => {
                    status = rs.data;
                });
                this.$emit("getAll");
                this.$emit("back");
                this.alertInfo("Empresa Eliminada")
            },
            async recoveryEnterprise() {
                var status = false;
                await this.axios.post("/api/enterprise/recovery/", this.enterpriseForm).then(rs => {
                    status = rs.data;
                });
                this.$emit("getAll");
                this.$emit("back");
                this.alertInfo("Empresa Recuperada")
            },
            changeAvatar() {
                if (this.$refs.avatar.image) {
                    this.enterpriseForm.avatar = [];
                    this.enterpriseForm.avatar = this.$refs.avatar.image;
                }
            }
        }
    };   
</script>
<style scoped>

</style>
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
export default {
    name:"user-form",
    props:{
        enterprise:{},
        title: String,
        isMyFirtsEnterprise: Boolean,
        isNew:Boolean
    },
    data(){
        return{
        }
    }
}
</script>

<style>

</style>
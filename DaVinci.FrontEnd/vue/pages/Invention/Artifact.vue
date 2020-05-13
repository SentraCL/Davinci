<template>

    <div class="inventionForm">
        <card :title="title">

            <div class="col-md-12 card-body">
                <div class="row">
                    <div class="col-md-10">
                        <span class="row">
                            <div class="col-md-2">
                                <img class="btn btn-success iconArtifactBig" :src="icon" title="Cambiar Icono" v-on:click="showIcon()" />
                            </div>
                            <div class="col-md-8">
                                <input-text label="Nombre" alphabetic removeSpace capitalize name="name" :value.sync="name" autocomplete="off"> </input-text>
                                <combo-simple label="Tipo Valor" v-on:change="getTypeAtr()" :list="defaultArtifact" keyValue="typeRef" keyLabel="nickName" :value.sync="artifactType"></combo-simple>
                            </div>
                        </span>

                        <div class="row">
                            <div class="col-md-6">
                                <button title="se Agrega como una referencia Unica en el invento" class="btn btn-xs btn-round btn-info addButton" v-on:click="addArtifact(false)">
                                    Agregar
                                </button>
                            </div>
                            <div class="col-md-6">
                                <button v-if="attribute" title="se Agrega como un Listado" class="btn btn-xs btn-round btn-info addButton" v-on:click="addArtifact(true)">
                                    Agregar Lista {{attribute}}
                                </button>
                            </div>

                        </div>
                        <span class="row">

                            <hr />
                        </span>
                    </div>
                    <div class="col-md-2">

                    </div>

                </div>
            </div>
        </card>
        <span class="row">
            <div v-if="artifactForm.artifacts.length>0" class="col-md-12">
                <h-tabs :tabs="artifactTab" :viewScoped="true">
                    <span slot="model">

                        <span class="row">

                            <div class="col-md-2">
                                <center>
                                    <img title="Cambiar Icono" v-on:click="showIcon()" :src="icon" width="32px" class="mona-lisa">
                                    <img src="@/assets/img/pintando.png">
                                    <hr />
                                </center>
                                <strong>Glosario</strong>
                                <br />
                                <ul class="list-unstyled small">
                                    <li>
                                        <i style="color:green" class="ti-eye"></i> : Referencia textual hacia <strong>{{artifactForm.name}}</strong>.
                                    </li>

                                    <li>
                                        <i style="color:brown" class="fa fa-key"></i> : Valor que identifica a <strong>{{artifactForm.name}}</strong>, como unico.
                                    </li>

                                    <li>
                                        <i style="color:red" class="fa fa-trash-o"></i> : Borra el Atributo.
                                    </li>

                                    <li>
                                        <i style="color:brown" class="fa fa-edit"></i> : Editar el atributo.
                                    </li>


                                </ul>
                            </div>
                            <div class="col-md-1"></div>

                            <ul class="list-unstyled team-members col-md-9">
                                <li>
                                    <div class="row" v-for="artifact in artifactForm.artifacts" :key="artifact.name">
                                        <div class="col-1"></div>
                                        <div class="col-4">
                                            <strong style="color : #4e1d00">{{ artifact.name }}</strong>
                                            <br />

                                            <span>
                                                <small v-if="!artifact.isEssential && artifact.isList">Lista de Artectos : {{ artifact.nickName }}</small>
                                                <small v-if="!artifact.isEssential && !artifact.isList">Artecto : {{ artifact.nickName }}</small>


                                                <small v-if="artifact.isEssential">{{ artifact.nickName }}</small>
                                            </span>
                                            <hr />
                                        </div>

                                        <div class="col-6">

                                            <d-radio v-if="artifact.isEssential" type="success" name="keyLabel" :value="artifact.name" :to.sync="keyLabel" v-on:click="setKeyLabel(artifact.name)" outline icon>
                                                <i class="ti-eye"></i>
                                            </d-radio>

                                            <d-radio v-if="artifact.isEssential" type="info" name="keyValue" :value="artifact.name" :to.sync="keyValue" v-on:click="setKeyValue(artifact.name)" outline icon>
                                                <i class="fa fa-key"></i>
                                            </d-radio>

                                            <d-button type="danger" @click.native.prevent="deleteMe(artifact)" outline icon>
                                                <i class="fa fa-trash-o"></i>
                                            </d-button>

                                            <d-button type="info" @click.native.prevent="edit(artifact)" outline icon>
                                                <i class="fa fa-edit"></i>
                                            </d-button>
                                        </div>

                                    </div>
                                </li>
                            </ul>

                        </span>


                    </span>
                    <span slot="view">
                        <span class="row">
                            <div class="col-md-2">
                                <center>
                                    <img src="@/assets/img/MonaLisa.png" width="120px">
                                    <hr />
                                    Esta es la vista que tendra al ser visualizada en la aplicacion final.
                                </center>
                            </div>


                            <div class="col-md-10">
                                <span class="row">
                                    <div class="col-md-2"></div>
                                    <div class="col-md-10">
                                        <artifact-form :title="titleForm" :inputs="artifactForm.artifacts"></artifact-form>
                                    </div>
                                </span>
                            </div>

                        </span>
                    </span>

                </h-tabs>


            </div>
        </span>
        <span class="row">
            <div class="col-md-4">
                <button class="btn btn-xs btn-round btn-danger" v-if="!isNew" round @click="deleteArtifact">
                    <i class="fa fa-eraser"></i> Borrar
                </button>
                <button v-if="artifactForm.artifacts.length>0" class="btn btn-xs btn-round btn-success" round @click="save">
                    <i class="fa fa-save"></i> Guardar
                </button>
                <button class="btn btn-xs btn-round btn-success" v-if="change" round @click="cancelChanges">
                    <i class="fa fa-undo"></i> Deshacer
                </button>
            </div>
            <div class="col-md-8"></div>
        </span>
        <hr />
        <!-- DIALOG's -->
        <m-dialog id="editDialog" :title="editDialog.title" :show.sync="editDialog.show" :isClose.sync="editDialog.close">
            <div slot="dialog">
                <div class="card" v-if="editDialog.show">
                    <input-text label="Nombre" alphabetic removeSpace capitalize name="name" :value.sync="editDialog.artifact.name"> </input-text>
                    <combo-simple label="Tipo Valor" :list="defaultArtifact" keyValue="typeRef" keyLabel="nickName" :value.sync="editDialog.artifact.code"></combo-simple>
                    <label v-if="defaultTypes.indexOf(editDialog.artifact.code) == -1" class="input"><input type="checkbox" v-model="editDialog.artifact.isList"> Crear como lista.</label><br>
                </div>
            </div>
            <div slot="actions">
                <div class="row">
                    <div class="col-md-2">
                        <button class="btn btn-round btn-info" @click="saveEdit()">
                            Cambiar
                        </button>
                    </div>
                    <div class="col-md-2">
                        <button class="btn btn-round btn-info" @click="cancelEdit()">
                            Cancelar
                        </button>
                    </div>
                </div>
            </div>

        </m-dialog>

        <m-dialog id="iconsGallery" :title="iconDialog.title" :show.sync="iconDialog.show" :isClose.sync="iconDialog.close">
            <div slot="dialog">
                <img class="col-xl-2 btn btn-success iconArtifact" :src="urlIcon" v-for="urlIcon,index in icons" @click="setIcon(index)" />
            </div>
        </m-dialog>

        <m-dialog id="errorSave" :title="titleSave" :show.sync="errorAlert.show" :isClose.sync="errorAlert.close">
            <span slot="dialog">
                <div v-if="errorAlert.isErrorDrop">
                    <span class="ti-alert"></span> No es posible Eliminar el Invento {{name}}, dado que este esta relacionado a los siguiente artefactos:
                    <hr />
                    <strong v-for="artifacUse in errorAlert.parents">{{artifacUse}}<br /></strong>
                </div>
                <div v-if="errorAlert.isErrorSave">
                    <span class="ti-alert"></span> Debes determinar los siguientes valores de referencia para tu invento.
                    <hr />
                    <div class="row">
                        <div class="col-md-6">
                            <ul class="list-unstyled ">
                                <li>
                                    <i style="color:green" class="ti-eye"></i> Referencia textual hacia <strong>{{artifactForm.name}}</strong>.
                                </li>

                                <li>
                                    <i style="color:brown" class="fa fa-key"></i>Valor que identifica a <strong>{{artifactForm.name}}</strong>, como unico.
                                </li>
                            </ul>
                        </div>
                        <div class="col-md-6">
                            <img src="@/assets/img/leo.png">
                        </div>
                    </div>
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

        <m-dialog id="deleteDialog" :title="titleDelete" :show.sync="delAlert.show" :isClose.sync="delAlert.close">
            <span slot="dialog">
                <div>
                    ¿Realmente Desea borrar este invento?
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDeleteDialog">
                        No , Cancelar
                    </d-button>
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="dropArtifact">
                        Si, Eliminar
                    </d-button>
                </span>
            </span>
        </m-dialog>


    </div>

</template>
<script>
    import PaperTable from "@/components/UserInterface/Tables/PaperTable.vue";
    import ArtifactForm from "../Invention/ArtifactForm.vue";
    export default {
        name: "artifact",

        props: {
            artifactForm: {}
        },
        components: {
            PaperTable,
            ArtifactForm
        },
        data() {
            return {
                inventionVOs: [],
                icons: [

                ],
                icon: {},
                editDialog: {
                    title: `<span class="ti-more-alt"> </span> Editar Atributos.`,
                    show: false,
                    close: false,
                    artifact: {}
                },
                iconDialog: {
                    title: `<span class="ti-layout-grid3-alt"> </span> Selecione un icono que pueda representar su invento.`,
                    show: false,
                    close: false,
                },
                attribute: "",
                change: false,
                isList: false,
                errorAlert: {
                    show: false,
                    close: false,
                    parents: []
                },
                isNew: false,
                delAlert: {
                    show: false,
                    close: false
                },
                overDafault: 0,
                defaultTypes: [],
                keyValue: "",
                keyLabel: "",
                titleSave: `<small><span style="color: #4e1d00" class="ti-delete"></span>Alto!!, no se puede continuar la acción con </small><strong> ${this.artifactForm.name}</strong>`,
                titleDelete: `<small><span style="color: #4e1d00" class="ti-delete"></span> Borrar </small><strong> ${this.artifactForm.name}</strong>`,
                title: `<strong> ${this.artifactForm.name}</strong><br/><small style="font-size:16px;"><span style="color : #4e1d00;" class="ti-settings"></span> Creador Atributos para </small>`,
                titleForm: `<span class="ti-layout-accordion-list"></span> ${this.artifactForm.name}`,
                artifactTab: [
                    {
                        id: 0,
                        title: `<span style="color:green" class="ti-palette"></span> Diseño`,
                        slot: "model"
                    },
                    {
                        id: 1,
                        title: `<span style="color:brown"  class="ti-blackboard"></span> Vista`,
                        slot: "view"
                    }

                ],
                artifactType: -1,
                name: "",
                defaultArtifact: []
            }
        },
        async mounted() {

            for (var i = 1; i <= 73; i++) {
                this.icons.push(require(`@/assets/img/artifacts/Artifact${i}.png`));
            }
            for (var a in this.artifactForm.artifacts) {
                var artifact = this.artifactForm.artifacts[a];
                artifact.edit = false;
            }
            this.icon = this.artifactForm.icon == "" ? this.icons[0] : this.artifactForm.icon;
            this.keyLabel = this.artifactForm.keyLabel;
            this.keyValue = this.artifactForm.keyValue;
            await this.getDefaultArtifact()
            this.inventionVOs = await this.getInventions()//JSON.parse(sessionStorage.getItem("inventos"));
            for (var i in this.inventionVOs) {
                var invention = this.inventionVOs[i]
                if (invention.code != this.artifactForm.typeRef) {
                    var artifact = {
                        name: invention.keyLabel,
                        typeName: "json",
                        nickName: invention.name,
                        typeRef: invention.code,
                        isEssential: false
                    }
                    var toAdd = true;

                    for (var a in invention.artifacts) {
                        var artifactX = invention.artifacts[a];
                        //Un artefacto no puede estar disponible si este posee una refencia al padre, asi se evita relacion Ciclica.
                        if (artifactX.nickName == this.artifactForm.name) {
                            toAdd = false;
                            break;
                        }
                    }
                    artifact.edit = false
                    if (toAdd) {
                        this.defaultArtifact.push(artifact);
                    }
                }
            }
        },

        methods: {

            saveEdit() {
                this.editDialog.artifact.typeName = this.getNameByKey(this.editDialog.artifact.code, "typeName")
                this.editDialog.artifact.nickName = this.getNameByKey(this.editDialog.artifact.code, "nickName")
                this.editDialog.artifact.isEssential = this.defaultTypes.indexOf(this.editDialog.artifact.code) > -1
                //console.log(JSON.stringify(this.editDialog.artifact));
                this.change = true;
                this.editDialog.show = false;
            },
            cancelEdit() {
                for (var i in this.artifactForm.artifacts) {
                    if (JSON.stringify(this.artifactForm.artifacts[i]) === JSON.stringify(this.editDialog.artifact)) {
                        this.artifactForm.artifacts[i] = this.cloneObject(this.editDialog.initialArtifact);
                        this.editDialog.artifact = this.cloneObject(this.artifactForm.artifacts[i]);
                        this.editDialog.initialArtifact = this.cloneObject(this.artifactForm.artifacts[i]);
                        this.editDialog.show = false;
                        return false;
                    }
                }
            },
            edit(artifact) {
                this.change = true;
                this.editDialog.artifact = artifact;
                this.editDialog.initialArtifact = this.cloneObject(artifact);
                this.editDialog.show = true;
                this.editDialog.close = false;
            },
            setIcon(indexIcon) {
                this.iconDialog.show = false;
                this.icon = this.icons[indexIcon];
            },
            showIcon() {
                this.iconDialog.show = true;
            },
            async getDefaultArtifact() {
                await this.axios.get("/api/invention/artifact/default/").then(rs => {

                    this.defaultArtifact = [];
                    //console.log("rs.data >>" + JSON.stringify(rs.data));
                    for (var a in rs.data) {
                        var defaultA = rs.data[a];
                        //console.log("defaultA >>" + JSON.stringify(defaultA));
                        var nickName = defaultA.typeName == "text" ? "Texto" :
                            defaultA.typeName == "numeric" ? "Numerico" :
                                defaultA.typeName == "date" ? "Fecha y Hora" :
                                    defaultA.typeName == "bool" ? "Si/No" :
                                        defaultA.typeName == "multitext" ? "Parrafo" : "No Determinado.";

                        this.defaultTypes.push(defaultA.code)

                        this.defaultArtifact.push(
                            { typeName: defaultA.typeName, nickName: nickName, isEssential: true, typeRef: defaultA.code }
                        )
                    }
                });

            },

            getTypeAtr() {
                var isCustomAtrib = this.defaultTypes.indexOf(this.artifactType) == -1
                if (isCustomAtrib) {
                    this.attribute = this.getNameByKey(this.artifactType, "nickName");
                } else {
                    this.attribute = null;
                }
            },
            async  getInventions() {
                return await this.axios.get("/api/invention/all/").then(rs => {
                    return rs.data;
                });
            },

            setKeyLabel(keyLabel) {
                this.change = true;
                this.keyLabel = keyLabel
            },
            setKeyValue(keyValue) {
                this.change = true;
                this.keyValue = keyValue
            },

            isValid(nameInvention) {
                for (var i in this.artifactForm.artifacts) {
                    if (this.artifactForm.artifacts[i].name == nameInvention) {
                        return true;
                    }
                }
                return false;
            },
            async save() {
                var status = false;
                //console.log("this.keyLabel : " + this.keyLabel);
                //console.log("this.keyValue : " + this.keyValue);
                if (!(this.isValid(this.keyLabel) && this.isValid(this.keyValue))) {
                    this.errorAlert.isErrorDrop = false
                    this.errorAlert.isErrorSave = true
                    this.errorAlert.show = true;
                    this.errorAlert.close = false;
                    return false;
                }
                this.artifactForm.keyLabel = this.keyLabel;
                this.artifactForm.keyValue = this.keyValue;
                this.artifactForm.icon = this.icon;
                this.artifactForm.author = this.getUserOnline()

                //console.log("Guardo : " + JSON.stringify(this.artifactForm));
                await this.axios.post("/api/invention/save/", this.artifactForm).then(rs => {
                    var response = rs.data;
                    this.change = false;
                    this.delAlert.show = false;
                    this.errorAlert.isErrorDrop = true
                    this.errorAlert.isErrorSave = false
                    if (response.status == "Error:Used") {
                        this.errorAlert.parents = response.datails.parents
                        this.errorAlert.show = true;
                        this.errorAlert.close = false;
                    }
                });
                this.$emit("save")
            },
            saveAll() {
                this.change = false;
                this.artifactForm.keyLabel = this.keyLabel;
                this.artifactForm.keyValue = this.keyValue;

                //console.log("SAVE >" + JSON.stringify(this.artifactForm))

                this.$emit("update:artifactForm", this.artifactForm);
                this.$emit("update")
                this.$emit("save")
            },
            getNameByKey(type, key) {
                //console.log("type :" + type)
                //console.log("key  :" + key)
                for (var t in this.defaultArtifact) {
                    if (this.defaultArtifact[t].typeRef == type) {
                        return this.defaultArtifact[t][key]
                    }
                }
                return ""
            },
            cancelChanges() {
                this.keyLabel = this.artifactForm.keyLabel;
                this.keyValue = this.artifactForm.keyValue;
                this.change = false;
                this.$emit("cancel");
            },
            deleteArtifact() {
                this.delAlert.show = true;
                this.delAlert.close = true;
            },
            closeDeleteDialog() {
                this.delAlert.show = false;
                this.errorAlert.show = false;
            },
            async dropArtifact() {
                var status = false;
                await this.axios.post("/api/invention/drop/", this.artifactForm).then(rs => {
                    var response = rs.data;
                    this.delAlert.show = false;
                    if (response.status == "Error:Used") {

                        this.errorAlert.isErrorDrop = true;
                        this.errorAlert.parents = response.datails.parents
                        this.errorAlert.show = true;
                        this.errorAlert.close = false;
                    }
                });
                this.$emit("drop")

            },

            deleteMe(artifact) {
                this.change = true;
                var newArtifacts = []
                var removeIndex = -1;
                for (var i in this.artifactForm.artifacts) {
                    if (this.artifactForm.artifacts[i].name == artifact.name) {
                        removeIndex = i;
                        break;
                    }
                }
                if (removeIndex > -1)
                    this.artifactForm.artifacts.splice(removeIndex, 1)

            },
            addArtifact(isList) {
                this.isList = isList;
                //Agregar ATRIBUTO:
                if (this.name != "" && this.artifactType != -1) {
                    this.change = true;
                    var typeName = this.getNameByKey(this.artifactType, "typeName")
                    var nickName = this.getNameByKey(this.artifactType, "nickName")

                    //console.log("this.artifactType : " + JSON.stringify(this.artifactType));
                    /* {"name":"Segmento","typeName":"json","nickName":"Segmento","keyValue":"Codigo","keyLabel":"Nombre","isEssential":false} */
                    //console.log("this.artifactType.isEssential : " + this.artifactType.isEssential)
                    //ACA CRITERIO DE LO ESENCIAL.
                    var isEssential = this.defaultTypes.indexOf(this.artifactType) > -1


                    var artifact = {
                        name: this.name,
                        code: this.artifactType,
                        typeRef: this.artifactType,
                        typeName: typeName,
                        nickName: nickName,
                        isEssential: isEssential,
                        isList: this.isList,
                        isJson: true
                    }


                    //console.log(JSON.stringify(artifact));

                    var index = -1
                    var newArtifacts = []
                    for (var i in this.artifactForm.artifacts) {
                        if (this.artifactForm.artifacts[i].name == this.name) {
                            this.isNew = false;
                        }
                        if (this.artifactForm.artifacts[i].name != this.name) {
                            newArtifacts.push(this.artifactForm.artifacts[i])
                        }
                    }
                    artifact.edit = false;
                    newArtifacts.push(artifact)
                    this.artifactForm.artifacts = newArtifacts;
                    this.name = ""
                    this.artifactType = -1
                    //console.log("Agregando a >>: " + JSON.stringify(artifact))
                }
                this.isList = false;
            }
        }
    }
</script>
<style scoped>
    .iconArtifact {
        padding: 2px;
        max-width: 48px;
        border-right: 2px solid #4e1d00;
        border-bottom: 2px solid #906538;
    }

    .iconArtifactBig {
        padding: 4px;
        max-width: 64px;
        border-right: 2px solid #4e1d00;
        border-bottom: 2px solid #906538;
    }

    .inventionForm {
        min-height: 600px;
    }

    .addButton {
        min-width: 200px;
    }

    .mona-lisa {
        cursor: pointer;
        opacity: 0.9;
        top: 48px;
        left: 146px;
        position: absolute;
        min-width: 32px
    }
</style>
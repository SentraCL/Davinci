<template>
    <div class="card ">
        <div class="col-md-12">
            <span class="row">
                <div class="col-md-12">
                    <h4><i style="color:black;background-color: yellow;" class="ti-layout-width-full"></i> EPICO : <span style="color:#555555">{{epicForm.name}}</span></h4>
                </div>
                <div class="col-md-12">
                    <input-text label="Titulo del Epico" alphabetic capitalize name="name" v-model="epicForm.name" autocomplete="off"> </input-text>
                    <drop-menu :title="epicTitle" class="epicMenu" v-on:change="epicManager(epicMenu.option)" :options="epicMenu.options" :option.sync="epicMenu.option"></drop-menu>
                </div>
                <div class="col-md-12" style="min-height: 300px;">
                    <h-tabs :tabs="epicInputsTab" viewScoped>
                        <span slot="des">
                            <span class="card-body">
                                <div class="row">
                                    <div class="col-md-12">
                                        <span><i class="ti-info"></i> Descripcion del proposito del EPICO, este texto se ocupara de referencia para conocer el propisito que busca este tipo de epico. </span>
                                    </div>
                                    <div class="col-md-12">
                                        <div class="col-md-12">
                                            <textarea rows="5" class="form-control border-input" title="Definicion del Epico" v-model="epicForm.definition">
                                            </textarea>
                                        </div>
                                    </div>
                                </div>
                            </span>
                        </span>
                        <span slot="atr">
                            <span class="card-body">
                                <div class="row">
                                    <div class="col-md-12">
                                        <span><i class="ti-comment-alt"></i> Los atributos de un <strong>Epico</strong>, establecen el contexto la tarea a evaluar. </span>
                                    </div>
                                    <div class="col-md-12">
                                        <input-text type="text" label="Atributo" v-model="input.name">
                                        </input-text>
                                        <combo-simple label="Tipo Valor" :list="epicAttributes" keyValue="typeRef" keyLabel="nickName" :value.sync="input.artifactType"></combo-simple>
                                        <d-button type="success" xs round @click.native.prevent="addAttribute()">
                                            Agregar
                                        </d-button>
                                    </div>
                                </div>
                                <hr />
                                <div class="row" v-for="attribute,index in epicForm.attributes" v-if="renderComponent">
                                    <div class="col-md-12">
                                        <input-text type="text" label="Atributo" v-model="attribute.name">
                                        </input-text>
                                        <combo-simple label="Tipo Valor" :list="epicAttributes" keyValue="typeRef" keyLabel="nickName" :value.sync="attribute.artifactType"></combo-simple>
                                        <drop-menu :title="optionTitle" v-on:change="menuAttribute(option,attribute,index)" :options="atrMenu" :option.sync="option"></drop-menu>
                                        <i v-if="epicForm.reference==attribute.name" style="color:brown;padding-left: 2px;top:4px;position:relative" class="ti-medall"></i>
                                    </div>
                                </div>
                            </span>
                        </span>
                        <span slot="act">
                            <span class="card-body">
                                <div class="row">
                                    <div class="col-md-12">
                                        <span><i class="ti-info"></i> Estos son los <strong>atributos comunes</strong> que tendran las <strong>historias de usuarios</strong> asociadas a este tipo de Epico. </span>
                                    </div>

                                    <div class="col-md-12">
                                        <input-text type="text" label="Nombre de Referencia" v-model="invention.name"></input-text>
                                        <combo-simple label="Invento" :list="inventions" keyValue="code" keyLabel="name" :value.sync="invention.code"></combo-simple>
                                        <combo-simple label="Valor por Defecto" :list="dataInvention[invention.code]" keyValue="val" keyLabel="key" :value.sync="invention.value"></combo-simple>
                                        <d-button type="success" xs round @click.native.prevent="addInvention()">
                                            Agregar
                                        </d-button>
                                    </div>
                                </div>
                                <hr />
                                <div class="row" v-for="epicInvention,index in epicForm.inventions">
                                    <div class="col-md-12">
                                        <input-text type="text" label="Atributo" v-model="epicInvention.name"></input-text>
                                        <combo-simple label="Invento" :list="inventions" keyValue="code" keyLabel="name" :value.sync="epicInvention.code"></combo-simple>
                                        <combo-simple label="Valor por Defecto" :list="dataInvention[epicInvention.code]" keyValue="val" keyLabel="key" :value.sync="epicInvention.value"></combo-simple>
                                        <drop-menu :title="epicInvTitle" v-on:change="menuInvention(option,epicInvention,index)" :options="invMenu" :option.sync="option"></drop-menu>
                                    </div>
                                </div>
                            </span>
                        </span>
                    </h-tabs>
                </div>
            </span>
        </div>
        <div class="card-footer">
            <span class="click" style="float:right;position:relative"> Epico <strong>{{epicForm.name}}</strong> , Referencia <strong><span style="color:brown">{{epicForm.reference}}</span></strong></span>
        </div>
    </div>

</template>
<script>

    export default {
        name: "epic-design",
        computed: {

        },
        props: {
            project: {},
            epicType: {},
            inventions: {}
        },

        data() {
            return {
                DROP: 0,
                EDIT: 1,
                UNDO: 2,
                NEW: 3,
                ADD: 4,
                BACK: 5,
                UNIQ: 6,

                renderComponent: true,

                dataInvention: {},

                optionTitle: `<small style="color:green" class="ti-arrow-circle-down"></small> Opciones`,
                option: -1,
                invMenu: [
                    { html: `<div><i style="color:red" class="ti-trash"></i> Eliminar</d>`, option: 0 },
                    { html: `<div><i style="color:blue" class="ti-save"></i> Guardar</div>`, option: 1 },
                ],

                atrMenu: [
                    { html: `<div><i style="color:red" class="ti-trash"></i> Eliminar</d>`, option: 0 },
                    { html: `<div><i style="color:blue" class="ti-save"></i> Guardar</div>`, option: 1 },
                    { html: `<div><i style="color:brown" title="Este atributo servira como referencia del EPICO" class="ti-medall"></i> Marcar Referencia</d>`, option: 6 },
                ],



                epicInvTitle: `<span><i style="color:green" class="ti-panel"></i> Editar</span>`,
                epicTitle: `<span><i style="color:blue" class="ti-panel"></i> Menu</span>`,
                epicMenu: {
                    option: -1,
                    options: [
                        { html: `<span><i style="color:green" class="ti-save"></i> Guardar Epico</span>`, option: 3 },
                        { html: `<span><i style="color:brown" class="ti-control-backward"></i> Volver</span>`, option: 5 },
                        //{ html: `<span><i style="color:blue" class="ti-save"></i> Mover</span>`, option: 2 },
                        //{ html: `<span><i style="color:red" class="ti-save"></i> Eliminar</span>`, option: 0 },
                    ]
                },

                epicForm: {
                    name: "",
                    definition: "",
                    attributes: [],
                    inventions: [],
                },

                epicInputsTab: [
                    {
                        id: 0,
                        title: `<span style="color:brown" class="ti-light-bulb"></span> Atributos`,
                        slot: "atr"
                    },
                    {
                        id: 1,
                        title: `<span style="color:blue"  class="ti-clip"></span> Actividades`,
                        slot: "act"
                    },
                    {
                        id: 2,
                        title: `<span style="color:green"  class="ti-clipboard"></span> Definición del Epico`,
                        slot: "des"
                    }

                ],

                defaultTypes: [],
                epicAttributes: [],
                title: "Nuevo",

                invention: {
                    code: "",
                    name: "",
                    //icon: "",
                    //type: "",
                    value: ""
                },

                input: {
                    name: "",
                    //code: "",
                    artifactType: ""
                }
            }
        },

        async mounted() {
            this.epicForm = this.epicType
            for (var index in this.inventions) {
                var inv = this.inventions[index];
                var values = []
                //Capturo lista de pk y descripcion
                inv.WareHouse.forEach(value => {
                    var key = value["KeyValue"];
                    var content = JSON.parse(value.ContentJSON)
                    var val = content[inv["keyLabel"]];
                    values.push({
                        key: key + ":" + val,
                        val: key
                    })
                });
                this.dataInvention[inv.code] = values;
            }
            await this.addDefaultArtifactToAttributes();
            await this.addInventionsToAttributes();
        },

        methods: {

            menuAttribute(option, attribute, index) {

                if (option == this.UNIQ) {
                    this.epicForm.reference = attribute.name;
                    this.renderComponent = false;

                    this.$nextTick(() => {
                        // Add the component back in
                        this.renderComponent = true;
                    });
                    return;
                }

                if (option == this.EDIT) {
                    this.epicForm.attributes[index] = attribute;
                    return;
                }
                if (option == this.UNDO) {
                    //TODO : Pendiente
                }
                if (option == this.DROP) {
                    this.epicForm.attributes.splice(index, 1);
                    return;
                }
            },

            menuInvention(option, detail, index) {
                // if(this.isJsonEmpty(detail)){
                //   return false;
                // }
                if (option == this.EDIT) {
                    this.epicForm.inventions[index] = detail;
                }
                if (option == this.UNDO) {
                    //TODO : Pendiente
                }
                if (option == this.DROP) {
                    this.epicForm.inventions.splice(index, 1);
                }
            },


            addAttribute() {
                if (!this.isEmptyOrSpaces(this.input.name) && !this.isEmptyOrSpaces(this.input.artifactType) && !this.isUnique(this.input.name, "name", this.epicForm.attributes)) {
                    var atribute = this.cloneObject(this.input);
                    this.epicForm.attributes.push(atribute);
                    this.input = {
                        code: "",
                        name: "",
                        artifactType: ""
                    };
                } else {
                    alert("NOOO!!!.");
                }
            },
            getNameInventionByCode(code) {
                for (var i in this.inventions) {
                    var inv = this.inventions[i]
                    if (inv.code == code) {
                        return inv.name;
                    }
                }
            },

            addInvention() {
                if (!this.isEmptyOrSpaces(this.invention.code) && !this.isUnique(this.invention.name, "name", this.epicForm.inventions)) {
                    //console.log(JSON.stringify(this.invention));
                    if (this.isEmptyOrSpaces(this.invention.name)) {
                        this.invention.name = this.getNameInventionByCode(this.invention.code);
                    }
                    var invention = this.cloneObject(this.invention);
                    this.epicForm.inventions.push(invention);
                    this.invention = {
                        code: "",
                        name: "",
                        //icon: "",
                        // type: "",
                        value: ""
                    };
                } else {
                    alert("NOOO!!!.");
                }
            },

            async epicManager(option) {
                if (option == this.BACK) {
                    this.$emit("back");
                }

                if (option == this.NEW) {
                    //console.log(JSON.stringify(this.epicForm, 4, null));
                    var epicType = {
                        name: this.epicForm.name,
                        projectCode: this.project.code,
                        reference: this.epicForm.reference,
                        definition: this.epicForm.definition,
                        attributes: this.epicForm.attributes,
                        inventions: this.epicForm.inventions
                    }
                    await this.axios.post(`/api/project/${this.project.code}/design/epic`, epicType).then(rs => {
                        alert("Epico Guardado");
                    });
                }
            },
            async addInventionsToAttributes() {
                this.inventions.forEach(inv => {
                    this.epicAttributes.push(
                        { typeName: inv.name, nickName: inv.title, isEssential: false, typeRef: inv.code }
                    )
                })
            },

            async addDefaultArtifactToAttributes() {
                await this.axios.get("/api/invention/artifact/default/").then(rs => {

                    this.epicAttributes = [];
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

                        this.epicAttributes.push(
                            { typeName: defaultA.typeName, nickName: nickName, isEssential: true, typeRef: defaultA.code }
                        )
                    }
                });

            },
        }
    };
</script>
<style scoped>
    .epicMenu {
        position: fixed;
        float: right;
        top: 0px;
        right: 80px;
    }
</style>
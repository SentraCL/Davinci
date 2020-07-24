<template>
    <form @submit.prevent>
        <h4 v-if="title" v-html="title"></h4>
        <div class="col-md-12">
            <span v-for="(input,index) in inputs" :key="input.name" v-if="input.isText || input.isNumeric || input.isDate || input.isBool">
                <input-text :removeSpace="keyValue==input.name" v-if="input.isText" :label="input.label" v-model="form[input.name]" :isInactive="isKeyValued && input.name==invention.keyValue || input.inactive"></input-text>
                <input-text :removeSpace="keyValue==input.name" v-if="input.isNumeric" type="number" :label="input.label" v-model="form[input.name]" :isInactive="isKeyValued && input.name==invention.keyValue  || input.inactive"></input-text>
                <input-text :removeSpace="keyValue==input.name" v-if="input.isDate" type="datetime-local" :label="input.label" v-model="form[input.name]" :isInactive="isKeyValued && input.name==invention.keyValue  || input.inactive"></input-text>
                <switch-box v-if="input.isBool" :label="input.name" :value.sync="form[input.name]"></switch-box>
            </span>
            <span v-for="(input,index) in inputs" :key="input.name" v-if="!input.isEssential && !input.isList" style="position: relative; top:-6px">
                <combo-simple :showKey="true" :label="input.label" class="input-item" :list="input.list" :keyValue="input.keyValue" :keyLabel="input.keyLabel" :value.sync="form[input.name]"></combo-simple>
            </span>
        </div>

        <div class="col-md-12" v-if="textAreaTabs.length>0">
            <h-tabs :tabs="textAreaTabs" viewScoped>
                <span :slot="input.name" v-for="(input,index) in inputs" v-if="input.isMultitext">
                    <textarea rows="5" class="form-control border-input" v-model="form[input.name]"></textarea>
                </span>
            </h-tabs>
        </div>


        <div class="col-md-12" v-if="selectTabs.length>0">
            <h4><i class="ti-control-shuffle"></i> Relacionar Inventos</h4>
            <h-tabs :tabs="selectTabs" viewScoped>
                <span :slot="input.name" v-for="(input,index) in inputs" v-if="input.isList ">
                    <div v-if="input.isList && input.list" class="input-item">
                        <select-table :data="input.list" :values.sync="form[input.name]" :varName="input.label"></select-table>
                    </div>
                    <div v-if="!input.list">
                        <span style="color:red" class="ti-light-bulb"></span> No existen items en <strong>{{input.label}}</strong> para asociar.
                    </div>
                </span>
            </h-tabs>
        </div>
    </form>
</template>
<script>
    //Ver toda clase de parametros para consumir micronaut
    export default {
        name: "invention-form",
        computed: {

        },
        props: {
            projectCode: String,
            invention: {},
            values: {},
            title: String
        },
        data() {
            var format = this.formatForm();
            var _form = format.form;
            var  _inputs= format.inputs;
            console.log("inputs",format.inputs)
            var _textTabs = format.textTabs;
            var _selectTabs = format.selectTabs;
            this.$emit("update:values", _form);
            return {
                form: _form,
                inputs: _inputs,
                isKeyValued: false,
                keyLabel: "",
                keyValue: "",
                textAreaTabs: _textTabs,
                selectTabs: _selectTabs
            }
        },
        watch: {
            invention() {
                var format = this.formatForm();
                this.form = format.form;
                this.inputs = format.inputs;
                console.log("inputs",format.inputs)
                this.textAreaTabs = format.textTabs;
                this.selectTabs = format.selectTabs;
            },
            async form(newForm, oldValue) {
                this.$emit("update:values", newForm);
                this.$emit("update");
                console.log(this.invention)
                console.log(this.values)
            }
        },
        created() {
            if (this.values[this.invention.keyValue]) {
                this.isKeyValued = true;
            }
        },
        methods: {
            formatForm() {
                var format = {
                    form: {},
                    inputs: [],
                    textTabs: [],
                    selectTabs: []
                }


                //console.log(JSON.stringify(this.invention.artifacts));

                var idTab = 0;

                for (var i in this.invention.artifacts) {
                    var input = this.invention.artifacts[i]
                    var conditionalKey = "is" + input.typeName.charAt(0).toUpperCase() + input.typeName.slice(1);
                    input[conditionalKey] = true
                    //DETECTA SI ES UN TIPO DE ARTEFACTO ESENCIAL O CUSTOM.
                    if (input.isJson) {
                        input.isEssential = false
                        var repository = JSON.parse(sessionStorage.getItem(input.typeRef + this.projectCode));
                        console.log("repository",repository)
                        if (repository == null) {
                            input['list'] = []
                        }
                        else {
                            input['list'] = []
                            if (repository.length > 0) {
                                //SETEA DATOS PARA SER LISTADOS.
                                input['list'] = repository
                                input.keyValue = Object.keys(repository[0])[0];
                                input.keyLabel = Object.keys(repository[0])[1];
                                input['value'] = []
                            }
                        }
                    } else {
                        input.isEssential = true
                    }

                    //ASOCIA EL VALOR AL ARTEFACTO.
                    if (this.values[input.name] == null || typeof this.values[input.name] === "undefined") {
                        //En caso que el artefacto tenga valores inexistentes
                        this.values[input.name] = input.isList ? [] : "";
                    }


                    //Segun sea el origen de la invocacion , los valores pueden venir vacios como string o un arreglo vacio.
                    if (input.isList) {
                        //en el caso de las listas, estas deben quedar como []
                        var values = this.values[input.name] == null || this.values[input.name].constructor === String ? [] : this.values[input.name];
                        format.form[input.name] = values;
                        format.selectTabs.push({
                            id: idTab,
                            title: `<span>${input.name}</span>`,
                            slot: input.name
                        })
                    } else {
                        //en el caso de las listas con un solo elemento, estas deben quedar como [][0]
                        var value = this.values[input.name].constructor === Array ? this.values[input.name][0] : this.values[input.name];
                        if (input.isBool) {
                            format.form[input.name] = value ? value : 0;
                        } else {
                            format.form[input.name] = value;
                            if (input.isMultitext) {
                                format.textTabs.push({
                                    id: idTab,
                                    title: `<span>${input.name}</span>`,
                                    slot: input.name
                                })
                                format.form[input.name] = value;
                                idTab++;
                            }
                        }
                    }
                    //ASOCIA EL VALOR AL ARTEFACTO.
                    var isImportant = false;
                    if ((input.name == this.invention.keyLabel) && (input.name == this.invention.keyValue)) {
                        input.label = `<span title="Referencia Textual de ${this.invention.name}" style="color:grey" class="ti-key"></span> <span style="color:green" class="ti-eye"></span> ${input.name}`;
                        format.inputs.splice(1, 0, input);
                        this.keyLabel = input.name;
                        isImportant = true;
                    } else if (input.name == this.invention.keyLabel) {
                        input.label = `<span title="Referencia Textual de ${this.invention.name}" style="color:green" class="ti-eye"></span> ${input.name}`;
                        format.inputs.splice(1, 0, input);
                        this.keyLabel = input.name;
                        isImportant = true;
                    } else if (input.name == this.invention.keyValue) {
                        input.label = `<span title="Referencia Unica de ${this.invention.name}" style="color:grey" class="ti-key"></span> ${input.name}`;
                        this.keyValue = input.name;
                        format.inputs.unshift(input)
                        isImportant = true;
                    }
                    if (!isImportant) {
                        input.label = input.name;
                        format.inputs.push(input)
                    }
                }
                return format;
            }
        }
    };
</script>
<style scoped>
    .grid-container {
        display: grid;
        grid-template-columns: auto auto auto;
        padding: 10px;
    }

    .input-item {
        padding: 5px;
    }
</style>
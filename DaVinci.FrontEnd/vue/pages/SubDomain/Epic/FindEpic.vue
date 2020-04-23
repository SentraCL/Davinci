<template>
    <form @submit.prevent class="find-epic">

        <div class="row">
            <div class=" col-md-12 ">
                <combo-simple label="Categoria" :list="typeOfEpics" keyValue="type" keyLabel="name" :value.sync="epicForm.type"></combo-simple>
                <input-text index="1" type="text" label="Codigo de Referencia" v-model="epicForm.code" required>
                </input-text>
                <input-text index="0" type="text" label="Titulo" title="Titulo que identifica al Epico" v-model="epicForm.title" required>
                </input-text>
            </div>
            <div class=" col-md-12 ">
                <button style="float:right" type="button" @click="find" class="btn sub-button">
                    Buscar
                </button>
            </div>
        </div>
        <br />
        <v-tabs :tabs="epicsTab" :activeIndex="0" v-if="epicsTab.length>0">
            <span :slot="epic.code" v-for="(epic, index) of epicsTab" :key="index" class="epic-content">
                <div style="padding-left:16px">

                    <div class="row" style="padding-left:0px">
                        <input-text index="1" type="text" label="Codigo de Referencia" v-model="epic.code" autofocus required>
                        </input-text>
                        <input-text index="0" type="text" label="Titulo" title="Titulo que identifica al Epico" v-model="epic.title" required>
                        </input-text>

                        <input-text index="0" type="email" label="Correo de Referencia" title="Correo referencia" v-model="epic.email">
                        </input-text>
                        <input-text index="1" type="date" label="Fecha de Termino" title="DeadLine" v-model="epic.deadline">
                        </input-text>
                        <input-text index="1" type="date" label="Ultima Modificacion" readonly title="DeadLine" v-model="epic.update">
                        </input-text>
                    </div>


                    <div style="left:-8px;position:relative;right:5px;" class="row form-group">
                        <div class=" col-md-12 ">
                            <label>Objetivo <small>(o descripción de {{epic.type.name}} .)</small></label>

                            <textarea rows="5" index="3" class="form-control border-input" v-model="epic.objective">

                                </textarea>
                        </div>
                    </div>

                </div>
            </span>
        </v-tabs>


    </form>

</template>
<script>
    export default {
        name: "find-epic",
        computed: {

        },

        data() {
            return {
                page: {
                    itemsPerPage: 16
                },
                typeOfEpics: [],
                epics: {},
                epicsTab: [],
                epicForm: {
                    title: "",
                    objective: "",
                    deadline: "",
                    email: "",
                    code: "",
                    type: ""
                }
            }
        },

        mounted() {
            this.typeOfEpics = this.getTypesOfEpics();
            this.epics = this.getEpics();
            //console.log(JSON.stringify(this.epics));
            if (this.epics != null) {
                for (var k in this.epics) {
                    var epic = this.epics[k];
                    epic['title'] = epic.code;
                    epic['id'] = epic.code;
                    epic['slots'] = epic.code;
                    this.epicsTab.push(epic);
                }
            }
        },

        methods: {
            remove() {
                this.$emit("remove");
            },

            undo() {
                if (this.epic['values'] != null) {
                    this.epicForm = this.cloneObject(this.epic.values);
                } else {
                    this.epicForm.title = "";
                    this.epicForm.objective = "";
                    this.epicForm.deadline = "";
                    this.epicForm.email = "";
                    this.epicForm.code = "";
                    this.epicForm.type = this.epic.type
                }
            },

            find() {

            },

            save() {

                this.saveEpic(this.epicForm);
            }

        }
    };
</script>
<style scoped>
    .find-epic {
        min-height: 650px;
        min-width: 800px !important;
    }
</style>
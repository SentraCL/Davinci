<template>
    <div class="panelfind">

        <h3><i class="ti-agenda"></i> Buscador de Historias de Usuario.</h3>
        <!--
        <combo-simple :showKey="false" label="Tipo de Historia de Usuario" class="input-item" :list="userStoryTypes" keyValue="code" keyLabel="title" :value.sync="usTypeCode"></combo-simple>
        -->
        <drop-menu :title="usMenu.title" class="usMenu" v-on:change="usSelect(usMenu.option)" :options="usMenu.options" :option.sync="usMenu.option"></drop-menu>
        <hr />
        <div v-if="showForm">
            <h4>Precondiciones</h4>
            <pre-condition :preconditions="preConditions" :values.sync="values" v-on:code="setCode()"></pre-condition>
        </div>
        <div class="row fa-pull-right">
            <button class="btn btn-success" v-if="showForm" @click="find"><i class="ti-search"></i> BUSCAR</button>
            <button class="btn btn-success" v-if="showForm" @click="clear"><i class="ti-eraser"></i> LIMPIAR</button>
            <button class="btn btn-success" @click="toClose"><i class="ti-close"></i> CERRAR</button>
        </div>
        <br />
        <br />
        <br />
        <!--
        <span v-for="result in userStories"><button class="btn btn-xs btn-outline-success">{{result.id}}</button></span>
        -->
        <card v-if="userStories.length>0" class="repository-us">
            <h4>{{userStories.length}} Resultados.</h4>
            <!--<user-stories :data="userStories" v-on:select="getUserStory()"></user-stories>-->
            <badger-accordion>
                <badger-accordion-item v-for="us in userStoriesRepository" :key="us.id" style="width: auto;">
                    <template slot="header"><strong>{{us.id}} </strong>, <small>Versión {{us.version}}, {{us.date}}</small></template>
                    <template slot="content">
                        <card style="background-color: transparent;" class="repository-us">
                            <div class="row">
                                <div class="col-md-8">
                                    <h4>{{us.id}}</h4>
                                    <br />
                                    {{us.date}}
                                </div>
                                <div class="col-md-4">
                                    <button class="btn btn-success" @click="take(us.id)"><i class="ti-share"></i> Abrir</button>
                                </div>

                            </div>
                            <hr />

                            <h-tabs :tabs="userStoryTab" viewScoped>
                                <span slot="precon">
                                    <span class="card-body">
                                        <h5>Definicion </h5>
                                        <div class="alert alert-info" style="margin-left:10px;margin-bottom: 5px;" :title="us.fields.description"><strong>Descripcion :</strong>{{us.fields.description | readMore(255, '...')}}</div>
                                        <div class="alert alert-info" style="margin-left:10px;margin-bottom: 5px;" :title="us.fields.expectedResult"><strong>Resultado Esperado :</strong>{{us.fields.expectedResult | readMore(255, '...')}}</div>
                                        <br />
                                        <h5>Pre-Condiciones</h5>
                                        <span v-for="prcon in us.preConditions">
                                            <br /><strong>{{prcon.label}} </strong> : {{prcon.value}}
                                        </span>
                                    </span>

                                </span>
                                <span slot="steps">
                                    <small class="card-body ">
                                        <div v-for="item in us.steps">
                                            <strong :title="item.step">Paso {{item.id}} </strong> : {{item.step | readMore(40, '...') }}
                                            <br />
                                            <span :title="item.result">{{item.result | readMore(40, '...') }}</span>
                                            <hr />
                                        </div>
                                    </small>

                                </span>
                            </h-tabs>
                        </card>
                    </template>
                </badger-accordion-item>
            </badger-accordion>
        </card>

    </div>
</template>
<script>
    import { BadgerAccordion, BadgerAccordionItem } from 'vue-badger-accordion'
    import PreCondition from "./Preconditions.vue";
    import UserStories from "./UserStories.vue";
    export default {
        name: "find-user-story",
        components: {
            PreCondition,
            UserStories,
            BadgerAccordion,
            BadgerAccordionItem,

        },
        computed: {

        },
        props: {
            idUS: {}
        },
        watch: {
            usTypeCode(newValue, oldValue) {
                this.userStoryTypes.forEach(usType => {
                    if (usType.code == newValue) {
                        this.preConditions = usType.preConditions;
                        this.showForm = true;
                    }
                });
            }
        },
        data() {
            return {
                userStoryTab: [
                    {
                        id: 0,
                        title: `<span style="color:brown" class="ti-check-box"></span> Atributos`,
                        slot: "precon"
                    },
                    {
                        id: 1,
                        title: `<span style="color:blue"  class="ti-menu-alt"></span> Paso a Paso`,
                        slot: "steps"
                    },


                ],
                userStoriesRepository: [],
                showForm: false,
                usTypeCode: "",
                userStoryTypes: [],
                userStories: [],
                preConditions: {},
                values: {},
                usMenu: {
                    option: -1,
                    title: `<span><i style="color:blue" class="fa fa-user"></i> Tipo de Historia de Usuario.</span>`,
                    options: []
                },
            }
        },
        async mounted() {
            this.userStoryTypes = await this.getTypesOfUserStories();
            this.userStoryTypes.forEach(option => {
                this.usMenu.options.push(
                    { html: `<span><i style="color:green" class="ti-angle-double-right"></i> ${option.title}</span>`, option: option.code, title: option.title }
                )
            });
        },
        methods: {

            usSelect(option) {
                this.usTypeCode = option;
                this.usMenu.options.forEach(us => {
                    if (us.option == option) {
                        this.usMenu.title = `<strong><i style="color:blue" class="fa fa-user"></i> ${us.title}.</strong>`;
                    }
                })
            },

            take(id) {
                this.$emit("update:idUS", id)
                this.$emit("found")
            },

            clear() {
                this.usTypeCode = ""
                this.showForm = false;
                this.values = {}
                this.userStories = []
            },
            toClose() {
                //this.findUserStories();
                this.$emit("toClose");
            },
            find() {
                this.findUserStories();
            },
            async findUserStories() {
                var pathname = window.location.pathname;
                var projectName = pathname.split("/")[2];
                var url = `/davinci/${projectName}/userStories/similar`
                this.userStories = [];
                var preconditionQuery = []
                for (var i in this.values) {
                    if (!this.isEmptyOrSpaces(this.values[i])) {
                        preconditionQuery.push({
                            id: i,
                            value: this.values[i]
                        })
                    }
                }
                await this.axios.post(url, preconditionQuery).then(rs => {
                    this.userStoriesRepository = [];
                    this.userStories = [];
                    for (var i in rs.data) {
                        var us = rs.data[i]
                        var index = us.lastVersionIndex;
                        var version = `V${index + 1}`
                        us.date = this.formatDays(us.versions[index].date);
                        us.version = version
                        var lastVersion = us;
                        for (var atri in us.versions[index]) {
                            lastVersion[atri] = us.versions[index][atri];
                        }
                        lastVersion.date = this.formatDays(us.versions[index].date);
                        this.userStoriesRepository.push(lastVersion)
                        this.userStories.push(us);
                    }
                })

            },
        }
    };
</script>
<style>
    .repository-us {
        left: -25;
        position: absolute;
        background-color: var(--sub-wallpaper);
        box-shadow: inset 0 0 5px grey;
        border-radius: 10px;
        margin-left: 30px;

        height: 500px;
        width: auto;
        overflow-y: scroll;
        overflow-x: none;
        margin-bottom: 25px;
    }

    .panelfind {
        max-width: 840px;
        width: 840px;
    }
</style>
<template>
    <div>
        <v-tabs :tabs="userStoriesTab" :activeIndex="0">
            <div :slot="us.slots" class="card-body" v-for="us,index of userStoriesTab" :key="index" style="min-height: 500px;">
                <h5><button class="btn btn-success" @click="take(us.slots)"><i class="ti-share"></i> {{us.slots}}</button></h5>
                <small><strong>{{us.date}}</strong> creado por <strong>{{us.userStory.author}}</strong></small>

                <hr />
                <!--<pre-condition :preconditions.sync="us.userStory.preConditions" :values.sync="us.values"></pre-condition>-->
                <span v-for="prcon in us.userStory.preConditions">
                    <strong>{{prcon.label}} </strong> : {{prcon.value}}<br />
                </span>
                <div v-if="us.userStory.steps.length>0">
                    <h3>Pasos de Ejecución</h3>
                    <hr />
                    <small class="scrolling">
                        <div v-for="item in us.userStory.steps">
                            <div class="row">
                                <div class="col-md-2" :title="item.step + '\n' + item.result">
                                    Paso {{item.id}}
                                </div>
                                <div class="col-md-8">
                                    <span :title="item.step">{{item.step | readMore(40, '...') }}</span>
                                    <br />
                                    <span :title="item.result">{{item.result | readMore(40, '...') }}</span>
                                </div>
                                <div class="col-md-2">

                                </div>
                            </div>
                        </div>
                    </small>
                </div>
            </div>
        </v-tabs>
        <uib-pagination class="pagination" v-if="data.length>itemsPerPage" @change="changePage()" firstText="«" lastText="»" previousText="‹" nextText="›" :total-items="data.length" v-model="pagination" :itemsPerPage="itemsPerPage" :max-size="10" :boundary-links="true">
        </uib-pagination>
    </div>
</template>
<script>
    import Vue from "vue";
    import PreCondition from "./Preconditions.vue";
    var pagination = require("vuejs-uib-pagination");
    Vue.use(pagination);


    export default {
        name: "user-stories",
        components: {
            PreCondition
        },

        props: {
            data: {},
            idUS: {}
        },

        data() {
            return {
                pagination: {
                    currentPage: 1,
                },
                indexs: [

                ],
                itemsPerPage: 10,
                userStoriesTab: []
            }
        },
        mounted() {
            this.changePage();
        },
        methods: {

            take(id) {
                //console.log(`UserStories.vue > ${id}, ejecutando FOUND`);
                this.$emit("update:idUS", id)
                this.$emit("found")
            },
            loadTabs() {
                var index = 0;
                this.userStoriesTab = [];

                for (var i in this.indexs) {
                    var index = this.indexs[i];
                    if (index < this.data.length) {
                        var us = this.data[index];
                        var lastIndex = us.versions.length - 1;
                        us.steps = us.versions[lastIndex].steps;
                        var tab = {
                            id: us.id,
                            title: `<small>${us.id}</small>`,
                            slots: us.id,
                            userStory: us,
                            values: {},
                            date: this.formatDays(us.date)
                        }
                        this.userStoriesTab.push(tab);
                    }
                }

            },
            changePage() {
                var from = (this.pagination.currentPage * this.itemsPerPage) - this.itemsPerPage;
                this.indexs = [];
                for (var i = 0; i < this.itemsPerPage; i++) {
                    this.indexs.push(from + i);
                }
                this.loadTabs();

            },
        }
    };
</script>
<style>
    .scrolling {
        background-color: aliceblue;
        overflow-x: hidden;
        overflow-y: scroll;
        white-space: nowrap;
        max-height: 250px;
    }
</style>
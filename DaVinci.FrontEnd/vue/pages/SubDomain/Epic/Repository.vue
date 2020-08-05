<template>
    <div style="width:auto;height:auto">
        <multipane class="horizontal-panes" layout="horizontal">
            <h4>Repositorio de Epicos.</h4>
            <button style="position:absolute;float:right;right: 10px;" class="btn btn-success btn-xs" @click="toClose"><i class="ti-close"></i> </button>
            <br />
            <multipane class="custom-resizer" layout="vertical">
                <div class="pane ">
                    <div>
                        <h6><i class="ti-package"></i> Epicos</h6>
                        <hr />
                        <span v-bind:key="epicType" v-for="(epicType,index) in types">
                            <a class="epicfilter" v-on:click="changeValue(epicType,index)" :value.sync="epicType"><strong>({{repository[epicType].length}}) {{epicType}}</strong></a>
                            <br />
                        </span>
                    </div>
                </div>
                <multipane-resizer></multipane-resizer>
                <div class="pane inbox">
                    <div class="category-epic">
                        <h5><i class="ti-dropbox"></i>Actividades </h5>
                        <card>
                            <epic-table  v-if="isActive" :reference="refColum" :value.sync="epicRef" :epics.sync="repository[epicIndex]" v-on:select="showActivities()"></epic-table>
                        </card>
                    </div>
                </div>
                <multipane-resizer></multipane-resizer>
            </multipane>
            <multipane-resizer></multipane-resizer>
            <br />
        </multipane>
    </div>
</template>

<script>
    import { Multipane, MultipaneResizer } from 'vue-multipane';
    import EpicTable from "./EpicTable.vue";
    export default {
        components: {
            Multipane,
            MultipaneResizer,
            EpicTable
        },
        props: {
            idEpic: String
        },
        data() {
            return {
                repository: {},
                types: [],
                epicIndex: "",
                epicRef: "",
                refColum: "",
                isActive:true
            }
        },
        async mounted() {
            var epicTypes = await this.getTypesOfEpics();
            var epics = await this.getEpics();

            epicTypes.forEach(typepic => {
                var type = typepic.Code;
                this.repository[typepic.name] = [];
                epics.forEach(epic => {
                    if (epic.type == type) {
                        this.refColum = typepic.reference;
                        this.repository[typepic.name].push(epic);
                    }
                });
            });
            this.types = Object.keys(this.repository);
            this.epicIndex = this.types[0]
            //console.log(JSON.stringify(this.repository));
        },
        methods: {
            async changeValue(epic,index){
                this.epicRef=this.types[index];
                this.refColum=this.types[index];
                this.isActive=false;
                this.epicIndex=epic;
                this.$emit("update:idEpic", this.epicRef);
                this.$emit("update");
                await this.reload();
            },
            async reload() {
                this.isActive=true;
            },
            toClose() {
                //this.findUserStories();
                this.$emit("toClose");
            },
            showActivities() {
                this.$emit("update:idEpic", this.epicRef);
                this.$emit("update");
                this.$emit("found");
            },
        }
    };
</script>

<style lang="scss">
    .custom-resizer {
        width: 100%;
        height: 300px;
    }

    .custom-resizer>.pane {
        text-align: left;
        padding: 15px;
        overflow: hidden;
    }
    
    .epicfilter{
        cursor: pointer;
    }
    .epicfilter:hover{
        color:#2634dbe6 !important;

    }

    .custom-resizer>.pane~.pane {}

    .custom-resizer>.multipane-resizer {
        margin: 0;
        left: 0;
        position: relative;

        &:before {
            display: block;
            content: "";
            width: 3px;
            height: 40px;
            position: absolute;
            top: 50%;
            left: 50%;
            margin-top: -20px;
            margin-left: -1.5px;

        }

        &:hover {
            &:before {
                border-color: #999;
            }
        }
    }

    .horizontal-panes {
        width: 100%;
        height: 90%;
    }

    .horizontal-panes>.pane {
        text-align: left;
        padding: 15px;
        overflow: hidden;
        background: #eee;
    }

    .inbox {
        overflow-y: scroll !important;
        width: 800px;
        background-color: #FAFAFA;
        height: auto !important;
        max-height: 98%
    }
</style>
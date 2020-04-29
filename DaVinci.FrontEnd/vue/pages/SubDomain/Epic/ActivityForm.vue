<template>
    <div>
        <h4><i style="color:red" class="ti-pin-alt"></i> Asociar Historias de Usuario.
            <span v-for="invention,label in activity.activityForm.inventions">
                <span class="alert alert-info" style="margin-left:10px">
                    <span style="color:#000"><b><img width="16px" :src="invention.icon"> {{label}}</b> :{{invention.value[0]}}</span>
                </span>
                <!--<combo-simple :inactive="true" :showKey="true" :label="label" class="input-item" :list="invention.list" :keyValue="invention.keyValue" :keyLabel="invention.keyLabel" :value.sync="invention.value[0]"></combo-simple>-->
            </span>
            <span @click="save()" style="float: right;position: relative;"> <i class="ti-save"></i> </span>
        </h4>
        <h-tabs :tabs="usActivityTab" viewScoped>
            <span slot="us">

                <small><i class="ti-exchange-vertical"></i> Orden de Ejecucion de las historias de usuario.(Usted puede modificar este orden arrastrando los items.)</small>
                <div class="activity-workspace">
                    <draggable ghost-class="ghost" v-model="userStories" @start="drag=true" @end="drag=false" draggable=".tusa">

                        <div v-for="us,index in userStories" :key="us.id" style="color:#000" class="tusa alert alert-success">
                            N°{{index+1}}) <strong :title="us.fields.description"> {{us.id}} </strong>
                            <i style="cursor: pointer;color:blue" @click="take(us.id)" title="Ver en detalle" class="ti-share"></i>
                            <br>
                            <small style="justify-content: flex-start;">{{us.fields.description | readMore(256, '...')}}</small>
                            <span class="check-us">
                                <button class="btn btn-danger" @click="removeUserStories(index)">Quitar</button>
                            </span>
                            <small>
                                <br>
                                Versión : <strong>{{us.version}}</strong>,<small>posee <strong>{{us.steps.length}}</strong> pasos de ejecución.</small>
                            </small>
                            <br>
                            <i>{{us.dateFormat}} ({{us.date}}) , creado por <strong>{{us.author}}.</strong>
                            </i>
                            </small>
                        </div>

                    </draggable>
                </div>
            </span>
            <span slot="find">
                <div class="activity-workspace">
                    <div class="row">
                        <div class="col-md-12">
                            <!--
                            <combo-simple :showKey="false" label="Tipo de Historia de Usuario" class="input-item" :list="userStoryTypes" keyValue="code" keyLabel="title" :value.sync="usTypeCode"></combo-simple>
                            -->
                            <drop-menu :title="usMenu.title" class="usMenu" v-on:change="usSelect(usMenu.option)" :options="usMenu.options" :option.sync="usMenu.option"></drop-menu>
                        </div>
                        <hr />
                        <div class="col-md-10" v-if="showFindBox">
                            <find-box :default="invention" :context.sync="context" :target.sync="target"> </find-box>
                        </div>
                        <div class="col-md-2" style="right:10" v-if="showFindBox">
                            <button @click="findUS" style="width:100%;top:0;position:relative" class="btn btn-success">
                                <i class="ti-search"></i> BUSCAR</button>
                        </div>
                    </div>
                    <small v-if="resultUserStories.length>1">{{resultUserStories.length}} Resultados</small>
                    <br />
                    <div v-for="us in resultUserStories" class="btn-group">
                        <div style=" padding: 5px;">
                            <div class="alert alert-info us-paper"><i style="cursor: pointer;color:blue" @click="take(us.id)" title="Ver en detalle" class="ti-share"> </i> <strong :title="us.fields.description"> {{us.id}} </strong>
                                <br><small style="justify-content: flex-start;">{{us.fields.description | readMore(256, '...')}}</small>
                                <span class="add-us">
                                    <button class="btn btn-success" v-if="!us.isAdded" @click="addUserStories(us)">Agregar</button>
                                    <button class="btn btn-danger" v-if="us.isAdded" @click="delUserStories(us)">Quitar</button>
                                </span>
                                <small>
                                    <br>
                                    Versión : <strong>{{us.version}}</strong><br>
                                    <i>{{us.dateFormat}} ({{us.date}}) , creado por <strong>{{us.author}}.</strong>
                                    </i>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
            </span>
            <span slot="dat">
                <div class="activity-workspace">
                    <attach-data :epic="epic" :activity="activity"></attach-data>
                </div>
            </span>
        </h-tabs>
    </div>
</template>
<script>
    import InventionForm from "../../Invention/InventionForm.vue";
    import Draggable from 'vuedraggable'
    import AttachData from './AttachData.vue'

    export default {
        name: "activity-form",
        components: {
            InventionForm,
            Draggable,
            AttachData
        },
        props: {
            epic: {},
            activity: {}
        },
        watch: {
            async usTypeCode(newValue, oldValue) {
                for (var x in this.userStoryTypes) {
                    var usType = this.userStoryTypes[x]
                    //this.userStoryTypes.forEach(usType => {
                    if (usType.code == newValue) {

                        //Desde la receta de como se genera un codigo de US, busco sus precondiciones para ver si algunos de sus atributos
                        //Coincide con los de la actividad
                        this.context = {}
                        this.showFindBox = false;
                        for (var cf in usType.codeFormat) {
                            var keyCF = usType.codeFormat[cf];
                            //console.log(`Key = ${keyCF}`);
                            for (var pr in usType.preConditions) {
                                var precondition = usType.preConditions[pr];
                                if (precondition.id == keyCF) {

                                    var def = await this.getInventionDefByCode(precondition.code);

                                    for (var ai in this.activity.activityForm.inventions) {
                                        var activityInvention = this.activity.activityForm.inventions[ai];
                                        var fixedValue = activityInvention.value;
                                        if (precondition.code == activityInvention.code) {
                                            //Guardar todos los inventos que dependan de este atributo.
                                            /*
                                            {   "ACRONIMOBECH":["DTF.BECH.ATMR"],
                                                "ACRONIMOCANAL":["ATMR"],
                                                "ATMR_date":["2020-01-03 11:57:37.188 -0300 -03"],
                                                "Descripcion":["ATM RedBanc"],
                                                "Funcionalidades":["DTF.BECH.ATMR_0001","DTF.BECH.ATMR_0002","DTF.BECH.ATMR_0003","DTF.BECH.ATMR_0005"],"Nombre":["ATM RedBanc"]}

                                                Buscar dentro de los key, si uno de estos es un invento o un JSOn, luego ver si estos tambien son parte del activity code
                                                si es asi setear el context
                                            */
                                            var typeRef = "";
                                            for (var nickName in def.WareHouse[fixedValue]) {
                                                //var nickName = def.WareHouse[fixedValue][l];
                                                for (var a in def.artifacts) {
                                                    var artifact = def.artifacts[a];
                                                    if (artifact.name == nickName && artifact.isList && artifact.isJson) {
                                                        typeRef = artifact.typeRef;
                                                        this.invention = artifact.nickName;
                                                        this.target[artifact.nickName] = "";

                                                        this.context[artifact.nickName] = await this.detailInvent(def.WareHouse[fixedValue][nickName], typeRef);
                                                        this.showFindBox = true;
                                                        break;
                                                    }
                                                }
                                                if (typeRef != "") {
                                                    break;
                                                }
                                            }
                                        } else {
                                            fixedValue = ""
                                            //Si el parent es un invento fijo, que el listado este filtrado por este.
                                            //Valdra la pena agregar atributos que no esten filtrados por atributos de la actividad? ahi la dejo...
                                            /*
                                            //console.log(`PreCondition = ${precondition.name} , Dinamico`);
                                            this.invention = precondition.name;
                                            this.target[precondition.name] = "";
                                            this.context[precondition.name] = ["DTF.BECH.ANF_0001", "DTF.BECH.ANF_0002", "DTF.BECH.ANF_0003"]
                                            */

                                        }
                                        if (Object.keys(this.typeRefs).indexOf(keyCF) == -1) {
                                            this.typeRefs[keyCF] = {
                                                name: precondition.name,
                                                code: precondition.code,
                                                //artifactName : "",
                                                id: precondition.id,
                                                value: fixedValue
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        /*
                        this.invention = "Funcionalidad"
                        this.target["Funcionalidad"] = "";
                        this.context = { "Funcionalidad": ["DTF.BECH.ANF_0001", "DTF.BEC....
                        */
                        this.preConditions = usType.preConditions;
                    }
                }
                //Agregar por EPICOS tambien    
            }
        },
        data() {
            return {
                userStories: [
                    //{ code: "1", description: "OK1" }
                ],
                usMenu: {
                    option: -1,
                    title: `<span><i style="color:blue" class="fa fa-user"></i> Tipo de Historia de Usuario.</span>`,
                    options: []
                },
                context: {},
                target: {},
                usTypeCode: "",
                invention: "",
                values: {},
                showFindBox: false,
                typeRefs: {},
                preConditions: {},
                resultUserStories: [],
                userStoryTypes: [],
                usActivityTab: [
                    {
                        id: 0,
                        title: `<span style="color:blue" class="ti-link"></span> Historias de Usuarios`,
                        slot: "us"
                    },
                    {
                        id: 1,
                        title: `<span style="color:brown"  class="ti-briefcase"></span> <strong>Asociar</strong> Historias de Usuario`,
                        slot: "find"
                    },
                    {
                        id: 2,
                        title: `<span style="color:green"  class="ti-clip"></span> Asociar <strong>Datos</strong>`,
                        slot: "dat"
                    }
                ],
            }
        },
        async mounted() {
            this.userStoryTypes = await this.getTypesOfUserStories();
            this.userStoryTypes.forEach(option => {
                this.usMenu.options.push(
                    { html: `<span><i style="color:green" class="ti-angle-double-right"></i> ${option.title}</span>`, option: option.code, title: option.title }
                )
            });

            if (this.activity.activityForm.userStories != null) {
                var me = this;
                this.activity.activityForm.userStories.forEach(async function (bo) {
                    var userStoryBO = await me.getUserStoryByREF(bo.code, bo.indexVersion);
                    var us = {}
                    us.id = userStoryBO.id;
                    us.code = userStoryBO.code;
                    console.log("Es Ultima Version? " + userStoryBO.lastVersionIndex == bo.indexVersion);
                    var myVersion = userStoryBO.versions[bo.indexVersion];
                    var version = `V${bo.indexVersion + 1}`
                    us["dateFormat"] = me.formatDays(myVersion.date);
                    us["date"] = myVersion.date.replace("T", " ").substring(0, 19);
                    us["author"] = myVersion.author;
                    us["version"] = version
                    us["isAdded"] = me.isAdded(us);
                    us["indexVersion"] = bo.indexVersion;
                    us["fields"] = myVersion.fields;
                    us["steps"] = myVersion.steps;
                    us["preConditions"] = myVersion.preConditions;
                    me.addUserStories(us);
                });
                var activity = this.cloneObject(this.activity);
                activity['userStories'] = this.userStories;
                this.$emit("update:activity", activity);
            }

        },
        methods: {
            async getInventionDef(code) {

                return await this.axios.get("/api/invention/all/").then(rs => {
                    for (var i in rs.data) {
                        var inv = rs.data[i]
                        if (inv.code == code) {
                            return inv;
                        }
                    }
                })
            },

            async getFullRef(code, list, def) {
                var projectCode = await this.getCodeProject();
                var url = `/api/project/${projectCode}/${code}/get`;
                var infoList = []
                await this.axios.get(url).then(rs => {
                    for (var code in rs.data) {
                        var inv = rs.data[code]
                        if (list.indexOf(code) > -1) {
                            var keylabel = inv[def.keyLabel];
                            var keyValue = inv[def.keyLabel];
                            //infoList.push(`${code} : ${keyValue}`);

                            infoList.push({
                                label: `${code} : ${keyValue}`,
                                value: code
                            });

                        }
                    }
                })

                return infoList;
            },

            async detailInvent(list, code) {
                //console.log(`${code} = ` + JSON.stringify(list));
                var def = await this.getInventionDef(code);
                console.log(`${code} = ` + JSON.stringify(def));
                var listDetail = await this.getFullRef(code, list, def);
                return listDetail;
            },

            save() {
                var activity = this.cloneObject(this.activity);
                activity['userStories'] = this.userStories;
                this.$emit("update:activity", activity);
                alert("Actividad Guardada.");
            },

            refreshTab() {
                var count = this.userStories.length;
                this.usActivityTab[0].title = `<span style="color:blue" class="ti-link"></span> Historias de Usuarios ${count == 0 ? '' : `<strong>(<span style="color:blue">${count}</span>)</strong>`}`;
            },

            addUserStories(us) {
                us.isAdded = true;
                if (this.userStories.indexOf(us) == -1) {
                    this.userStories.push(us);
                }
                this.refreshTab();
            },

            removeUserStories(index) {
                var us = this.userStories[index];
                this.resultUserStories.forEach(rus => {
                    if (rus.id == us.id) {
                        rus.isAdded = false;
                    }
                })
                this.userStories.splice(index, 1);
                this.refreshTab();
            },

            delUserStories(_us) {
                var index = 0;
                _us.isAdded = false;
                var found = false;
                this.userStories.forEach(us => {
                    if (us.id == _us.id) {
                        found = true;
                        return index;
                    }
                    if (!found) {
                        index++;
                    }
                });
                if (found) {
                    this.userStories.splice(index, 1);
                    this.refreshTab();
                }
            },

            usSelect(option) {
                this.usTypeCode = option;
                this.usMenu.options.forEach(us => {
                    if (us.option == option) {
                        this.usMenu.title = `<strong><i style="color:blue" class="fa fa-user"></i> ${us.title}.</strong>`;
                    }
                })
            },

            take(id) {
                //console.log(`UserStories.vue > ${id}, ejecutando FOUND`);
                this.$emit("update:idUS", id)
                this.$emit("found")
            },
            isAdded(_us) {
                var exist = false;
                this.userStories.forEach(us => {
                    if (us.id == _us.id) {
                        exist = true;
                    }
                })
                return exist
            },
            async  findUS() {
                //console.log(JSON.stringify(this.target));
                var id = this.typeRefs[Object.keys(this.target)[0]];
                var value = this.target[Object.keys(this.target)[0]];
                var preconditionQuery = []

                for (var t in this.typeRefs) {
                    var typeRef = this.typeRefs[t];
                    if (typeRef.value == "") {
                        if (this.target[typeRef.name]) {
                            preconditionQuery.push({
                                id: t,
                                value: this.target[typeRef.name]
                            })
                        }
                    } else {
                        preconditionQuery.push({
                            id: t,
                            value: typeRef.value[0]
                        })
                    }
                }

                var pathname = window.location.pathname;
                var projectName = pathname.split("/")[2];
                var url = `/davinci/${projectName}/userStories/similar`
                this.resultUserStories = [];

                await this.axios.post(url, preconditionQuery).then(rs => {
                    for (var i in rs.data) {
                        var us = {}
                        us.id = rs.data[i].id;
                        us.code = rs.data[i].code;
                        var index = rs.data[i].lastVersionIndex;
                        var lastVersion = rs.data[i].versions[index];
                        var version = `V${index + 1}`
                        us.dateFormat = this.formatDays(lastVersion.date);
                        us.date = lastVersion.date.replace("T", " ").substring(0, 19);
                        us.author = lastVersion.author;
                        us.version = version
                        us.isAdded = this.isAdded(us);
                        us.indexVersion = index
                        us.fields = lastVersion.fields;
                        us.steps = lastVersion.steps;
                        us.preConditions = lastVersion.preConditions;
                        this.resultUserStories.push(us);
                    }
                })
            }
        }
    };
</script>
<style scoped>
    .activity-title {
        text-shadow: 1px 1px var(--transparent);
    }

    .tusa:hover {
        background-color: cornflowerblue;
        cursor: grab;
    }

    .tusa:active:hover {
        cursor: grabbing;
    }

    .ghost:hover {
        cursor: grabbing;
    }

    .tusa {
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
    }

    .ghost {
        cursor: grabbing;
        background-color: cornflowerblue;
        border: 3px;
        border-color: #000;
        font-size: larger;
        opacity: 1;
        -webkit-box-shadow: 5px 5px 5px -1px rgba(0, 0, 0, 0.59);
        box-shadow: 5px 5px 5px -1px rgba(0, 0, 0, 0.59);
    }

    .tusa:active:focus {
        background-color: yellow;
    }

    .tusa:focus {
        background-color: aqua;
    }


    .usMenu {
        position: fixed;
        top: 0px;
    }

    .check-us {
        float: right;
        right: 20px;
        top: 5px;
    }

    .add-us {
        float: right;
        right: 5px;
        position: absolute;
        top: 190px;
    }

    .activity-workspace {
        /*
        overflow-y: auto;
        overflow-x: hidden;
        height: 450px;
        */
        border: 2px;
        width: 100%;

        min-height: 450px;
    }

    .us-paper {
        color: #000;
        max-width: 240px;
        max-height: 240px;
        height: 240px;
        min-height: 240px;
        padding: 5px;

    }

    .us-added {}
</style>
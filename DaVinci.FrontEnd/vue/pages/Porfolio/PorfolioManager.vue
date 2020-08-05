<template>

    <div class="col-md-12">

        <div class="content workspace " v-if="ready">

            <div class="col-md-12">

                <find-box :context.sync="find" :target.sync="inventionTarget" id="buscaInventos"> </find-box>
                <v-tabs :tabs.sync="inventionCRUDs" v-if="inventionCRUDs.length>0" :activeIndex="activeIndex">
                    <div :slot="NEW" class="card">
                        <!--GESTION DEL INVENTO-->
                        <div class="row" v-if="reloadForm">
                            <span style="position: absolute; right: 10;z-index:99999">
                                <drop-menu :title="taskTitle" v-on:change="manager()" :options="options" :option.sync="option"></drop-menu>
                            </span>
                        </div>

                        <div class="col-md-12 min-form">
                            <import-invent v-if="option==1" :project="project" :format="currentForm" :code="currentCode" :name="currentName" v-on:finish="loadCurrentInvention"></import-invent>
                            <export-invent v-if="option==2" :project="project" :inventions="inventionVOs" :code="currentCode" :name="currentName"></export-invent>
                            <invention-form v-if="option==0 && reloadForm " :title="createInv.subtitle" :projectCode="project.code" :invention.sync="createInv.form" :values.sync="createInv.artifact"></invention-form>
                        </div>
                        <span class="card-footer">
                            <p class="float-right" v-if="option==0">
                                <button class="btn btn-round btn-info" round @click="undo()">
                                    <i class="fa fa-undo"></i> Deshacer
                                </button>
                                <button class="btn btn-round btn-success" round @click="save(createInv)">
                                    <img :src="icon" width="14px" /> Crear {{currentName}}
                                </button>
                            </p>
                            <small>
                                <span v-html="options[option].html"></span>
                                <blockquote v-if="(Object.keys(repository[currentCode]).length - 1) > 0">Existen un total de <strong>{{Object.keys(repository[currentCode]).length - 1}}</strong> items para <strong>{{currentName}}</strong></blockquote>
                                <blockquote v-if="(Object.keys(repository[currentCode]).length - 1) == 0">Crear tu el Primer <strong>{{currentName}}</strong></blockquote>
                            </small>
                        </span>

                    </div>
                    <div :slot="invo.slots" v-if="invo.slots!=NEW" class="card" v-for="(invo, index) of inventionCRUDs" :key="index">
                        <!--LISTADO DE INVENTO Y GESTION PARTICULAR-->
                        <div class="col-md-12 min-form">
                            <invention-form v-if="reloadForm" :projectCode="project.code" :title="invo.subtitle" :invention.sync="invo.form" :values.sync="invo.artifact"></invention-form>
                        </div>
                        <span class="card-footer">
                            <p class="float-right">
                                <button class="btn btn-round btn-danger" @click="showDropDialog(invo)">
                                    <i class="fa fa-eraser"></i> Eliminar
                                </button>
                                <button class="btn btn-round btn-success" round @click="save(invo)">
                                    <img :src="icon" width="14px" /> Actualizar {{invo.slots}}
                                </button>
                            </p>
                            <small>
                                <blockquote> Ultima Modificación, {{ invo.update | aprox }} (<strong>{{invo.update}}</strong>)</blockquote>
                            </small>
                        </span>
                    </div>
                </v-tabs>
                <div style="width: auto" class="category">
                    <uib-pagination class="pagination" v-if="totalInv+1>itemsPerPage" @change="changePage()" firstText="«" lastText="»" previousText="‹" nextText="›" :total-items="totalInv" v-model="pagination" :itemsPerPage="itemsPerPage" :max-size="10" :boundary-links="true">
                    </uib-pagination>
                </div>
            </div>
        </div>
        <!-- Alertas de Grabado -->
        <m-dialog id="alertDialog" :title="alertDialog.title" :show.sync="alertDialog.show" :isClose.sync="alertDialog.close">
            <span slot="dialog">
                <div class="row">
                    <span class="col-md-6">
                        <span v-html="alertDialog.html"></span>
                    </span>
                    <span class="col-md-6">
                        <img src="@/assets/img/leo.png">
                    </span>
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDialog">
                        Cerrar
                    </d-button>
                </span>
            </span>
        </m-dialog>

        <!-- Alerta de Eliminacion -->
        <m-dialog id="dropDialog" :title="dropDialog.title" :show.sync="dropDialog.show" :isClose.sync="dropDialog.close">
            <span slot="dialog">
                <div class="row">
                    <span class="col-md-6">
                        Si elimina este invento, no podra recurar lo ingresado y esto podria provocar una paradoja en el futuro.
                    </span>
                    <span class="col-md-6">
                        <img src="@/assets/img/leo.png">
                    </span>
                </div>
            </span>
            <span slot="actions">
                <span class="btn-group">
                    <d-button type="danger" class="btn btn-xs" round @click.native.prevent="drop(dropDialog.inventionTO)">
                        Eliminar
                    </d-button>
                    <d-button type="success" class="btn btn-xs" round @click.native.prevent="closeDialog">
                        Cancelar
                    </d-button>
                </span>
            </span>
        </m-dialog>


    </div>
</template>
<script>
    import Vue from "vue";
    import InventionForm from "../Invention/InventionForm.vue";
    import ImportInvent from "../Invention/ImportInvent.vue";
    import ExportInvent from "../Invention/ExportInvent.vue";

    var pagination = require("vuejs-uib-pagination");
    Vue.use(pagination);

    const config = {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    }
    const qs = require('querystring')

    export default {
        filters: {
            aprox: function (date) {
                //2019-09-23 12:02:17
                var createDate = new Date(date);

                if (Object.prototype.toString.call(createDate) === '[object Date]') {
                    var today = new Date();
                    var times = today.getTime() - createDate.getTime();
                    var days = Math.round(times / (1000 * 3600 * 24));

                    //If Hoy    
                    if (days == 0) {
                        return "Hoy"
                    }
                    //if Ayer
                    if (days == 1) {
                        return "Ayer"
                    }
                    //if Hace mas de 1 dia
                    if (days > 1 && days <= 6) {
                        return `hace ${days} dias`
                    }
                    //if Hace mas una semana
                    if (days == 7) {
                        return `hace una semana`
                    }
                    //if Hace mas de una Semana
                    if (days > 7 && days < 30) {
                        return `hace ${Math.round(days / 7)} semanas`
                    }
                    //if Hace un Mes
                    if (days >= 30 && days < 60) {
                        return `hace un mes`
                    }
                    //if Hace mas de un Mes
                    if (days >= 60 && days < 365) {
                        return `hace ${Math.round(days / 30)} meses`
                    }
                    //if Hace un año
                    if (days >= 365 && days < 600) {
                        return `hace un año`
                    }
                    //if Hace N años                    
                }
                return ""
            }
        },
        props: {
            project: {}
        },
        components: {
            InventionForm,
            ImportInvent,
            ExportInvent
        },
        watch: {
            async project(newValue, oldValue) {
                if (newValue != oldValue) {
                    this.ready = false;
                    this.option = 0;
                    this.$emit("update:project", newValue);
                    this.activeIndex = 0;
                    await this.loadInventions()
                    this.ready = true;
                }
            }

            , async inventionTarget(newValue, oldValue) {
                for (var i in this.project.inventions) {
                    if (this.project.inventions[i].name == Object.keys(newValue)[0]) {
                        this.option = 0;
                        var inv = this.project.inventions[i];
                        //ACUTALIZA CONTEXTO DE INVENTOS
                        this.currentForm = inv;
                        this.keyValue = inv.keyValue;
                        this.keyLabel = inv.keyLabel;
                        this.currentName = inv.name ;
                        this.currentCode = inv.code;
                        this.icon = inv.icon;
                        this.inventionVOs = []

                        this.createInv = this.formatNewInvention();

                        //console.log(`this.currentName ${this.currentName} `);
                        //console.log(`this.currentCode ${this.currentCode} `);
                        //console.log(`this.keyValue ${this.keyValue} `);
                        //console.log(`this.currentCode ${this.currentCode} `);

                        var repository = this.repository[this.project.inventions[i].code]
                        for (var r in repository) {
                            var data = repository[r]
                            this.inventionVOs.push(data);
                        }
                        this.activeIndex = 0;
                        this.totalInv = Object.keys(this.repository[this.currentCode]).length;
                        break;
                    }
                }
                this.refreshTabs();

            }
        },
        computed: {
        },

        data: function () {
            return {
                pagination: {
                    currentPage: 1,
                },
                taskTitle: `<strong><i style="color:blue" class="ti-view-list-alt"></i> Acciones</strong>`,
                option: 0,
                options: [
                    { html: `<div><i style="color:yellow" class="ti-bolt"></i> Crear Nuevo</div>`, option: 0 },
                    { html: `<div><i style="color:green" class="ti-cloud-up"></i> Importar</div>`, option: 1 },
                    { html: `<div><i style="color:blue" class="ti-cloud-down"></i> Exportar</d>`, option: 2 }
                ],

                createInv: {},

                reloadForm: true,
                itemsPerPage: 10,
                activeIndex: 0,
                NEW: "NEW",
                inventionCRUDs: [],

                alertDialog: {
                    show: false,
                    close: false,
                    title: "",
                    html: ""
                },
                dropDialog: {
                    show: false,
                    close: false,
                    title: "",
                    html: ""
                },
                inventionTarget: {
                    invention: "",
                    value: ""
                },
                inventionVOs: {},
                ready: false,
                find: {},
                currentCode: "",
                currentName: "",

                icon: "",
                keyValue: "",
                keyLabel: "",
                artifact: {},
                repository: {}

            }
        }, async mounted() {
            await this.loadInventions();
        }, methods: {
            manager() {
                //console.log("Escogio opcion : " + this.option);
            },
            closeDialog() {
                this.alertDialog.show = false;
                this.dropDialog.show = false;
            },

            undo() {
                this.reloadForm = false;
                setTimeout(() => {
                    this.createInv = this.formatNewInvention();
                    this.reloadForm = true;
                }, 10);

            },
            formatNewInvention() {
                var inventionEmpty = {}
                inventionEmpty.slots = this.NEW;
                inventionEmpty.title = `<span style="min-width:120px;color:#88CC98">CREAR<br/><strong>${this.currentName}</strong> </span>`;
                inventionEmpty.subtitle = `<strong> Crear </strong> ${this.currentName} `;
                inventionEmpty.id = this.NEW;
                var cleanInputs = {}
                inventionEmpty.form = this.currentForm;
                for (var i in this.currentForm.artifacts) {
                    var input = this.currentForm.artifacts[i]
                    cleanInputs[input.name] = ""
                }
                inventionEmpty.artifact = cleanInputs;
                return inventionEmpty;
            },
            inventionWithData(data) {
                var inventionData = {}
                inventionData.artifact = data;
                inventionData.form = this.currentForm;
                return inventionData;
            },
            changePage() {
                this.reloadForm = false;
                setTimeout(() => {
                    this.refreshTabs();
                    this.activeIndex = 0;
                    this.reloadForm = true;
                }, 10);

            },
            refreshTabs() {
                this.ready = false;
                var findKey = this.inventionTarget[this.currentName];
                var inventionVOs = [];
                var find = this.inventionTarget.value;
                inventionVOs = this.inventionVOs;
                this.inventionCRUDs = []
                var inventions = this.repository[this.currentCode]
                if (findKey != "") {
                    this.ready = false;
                    inventions = {}
                    for (var k in this.repository[this.currentCode]) {
                        if (k.toUpperCase().indexOf(findKey.toUpperCase()) > -1) {
                            inventions[k] = this.repository[this.currentCode][k];
                        }
                    }
                    this.activeIndex = 0;
                }
                else {
                    inventions[this.NEW] = this.formatNewInvention()
                    this.inventionCRUDs.push(inventions[this.NEW]);
                }

                this.createInv = this.formatNewInvention();
                this.totalInv = Object.keys(inventions).length;

                var end = this.totalInv;
                var begin = (this.totalInv <= this.itemsPerPage) ? 0 : ((this.pagination.currentPage * this.itemsPerPage) - this.itemsPerPage);
                var index = 0;

                for (var key in inventions) {
                    var inventionCRUD = inventions[key];
                    if (index >= begin && key != this.NEW) {
                        this.inventionCRUDs.push(inventionCRUD);
                        if (this.inventionCRUDs.length == this.itemsPerPage + 1) {
                            break;
                        }
                    }
                    index++
                }

                this.ready = true;
            },
            async drop(inventionTO) {
                this.ready = false;
                await this.axios.post(`/api/project/${this.project.code}/${this.currentCode}/delete`, qs.stringify(inventionTO.artifact), config).then(rs => {
                    var keyValue = inventionTO.artifact[this.keyValue];
                    delete this.repository[this.currentCode][keyValue];
                    this.find[this.currentName] = Object.keys(this.repository[this.currentCode])
                    this.dropDialog.show = false;
                });
                /*
                await this.commitUI();
                */
                await this.loadCurrentInvention();
                this.activeIndex = 1;
                this.ready = true;
            },
            showDropDialog(inventionTO) {
                this.dropDialog.show = true;
                var keyValue = inventionTO.artifact[this.keyValue];

                this.dropDialog.close = false;
                this.dropDialog.title = `<span><span class="ti-trash"></span>¿ Realmente desea eliminar, ${this.currentName} <strong>${keyValue}</strong> ?</span>`
                this.dropDialog.inventionTO = inventionTO;

            },
            async save(inventionItem) {
                var inventionTO = {}
                if (inventionItem.slots == this.NEW) {
                    inventionTO = this.cloneObject(inventionItem)
                } else {
                    inventionTO = inventionItem
                }
                var keyValue = inventionTO.artifact[this.keyValue];
                var keyLabel = inventionTO.artifact[this.keyLabel];
                /*
                console.log("inventionTO  : " + JSON.stringify(inventionTO));
                console.log("keyValue     : " + JSON.stringify(keyValue));
                console.log("keyLabel     : " + JSON.stringify(keyLabel));
                */
                if (inventionTO.slots == this.NEW) {
                    if (this.isEmptyOrSpaces(keyLabel) || this.isEmptyOrSpaces(keyValue)) {
                        this.alertDialog.show = true;
                        this.alertDialog.close = false;
                        this.alertDialog.title = `<strong><span class="ti-face-sad"></span> Vaya parece que olvido algo importante!.</strong>`
                        this.alertDialog.html = `Los valores de referencia de los inventos son necesarios para ser identificados y ser distinguidos como unicos, es necesario que ingrese : <br/><hr/>`

                        if (this.isEmptyOrSpaces(keyLabel)) {
                            this.alertDialog.html += `<span style="color:grey" class="ti-key"></span>  El valor de Referencia Unico.<br/>`
                        }
                        if (this.isEmptyOrSpaces(keyValue)) {
                            this.alertDialog.html += `<span style="color:green" class="ti-eye"></span>  El valor de Referencia Textual.<br/>`
                        }
                        return false;
                    }

                    if (Object.keys(this.inventionVOs).indexOf(inventionTO.artifact[this.keyValue]) > -1) {
                        this.alertDialog.show = true;
                        this.alertDialog.close = false;
                        this.alertDialog.title = `<span><span class="ti-eraser"></span> El valor <strong>${keyValue}</strong> como <strong>${this.keyValue}</strong> de ${this.currentName}, ya Existe.</span>`
                        this.alertDialog.html = `Al parecer estamos poco creativos?, El valor ${keyValue} ya fue ocupado con anterioridad, un valor de Referencia identifica de manera única cada un invento. <br/><hr/>`

                        return false;
                    }
                }

                //var newInvention = this.cloneObject(inventionTO);
                var newInvention = inventionTO;
                var slotRef = inventionTO.artifact[this.keyValue];
                newInvention.slots = slotRef;
                newInvention.title = `<span style="min-width:120px" title=" ${keyLabel} ">${inventionTO.artifact[this.keyValue]} </span>`;
                newInvention.subtitle = `${this.currentName} : <strong>${inventionTO.artifact[this.keyValue]}</strong>`
                newInvention.id = slotRef;
                var today = new Date();
                var date = today.getFullYear() + '-' + ((today.getMonth() + 1) < 10 ? ("0" + (today.getMonth() + 1)) : (today.getMonth() + 1)) + '-' + (today.getDate() < 10 ? "0" + today.getDate() : today.getDate());
                var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
                var dateTime = date + ' ' + time;
                newInvention.update = dateTime
                this.inventionVOs[keyValue] = newInvention;
                var formPost = qs.stringify(newInvention.artifact);
                var postData = this.toInventionRequestDTO(newInvention.artifact);

                await this.axios.post(`/api/project/${this.project.code}/${this.currentCode}/save`, postData).then(rs => {

                    status = rs.data;
                    this.totalInv = Object.keys(this.repository[this.currentCode]).length;
                    this.repository[this.currentCode][keyValue] = newInvention;
                    var aprox = Object.keys(this.repository[this.currentCode]).length / this.itemsPerPage;

                    if (inventionTO.slots == this.NEW) {
                        this.find[this.currentName].push(keyValue);
                        this.activeIndex = 1;
                    }
                    this.undo();
                });
                await this.commitUI();
                this.activeIndex = 1;
                this.ready = true;
                //if (inventionTO.slots == this.NEW)
                //this.activeIndex = 1;
            },
            toInventionRequestDTO(artifacts) {
                var requestDTO = {};
                for (var a in artifacts) {
                    var artifact = artifacts[a]
                    //Convierte un JSON comun, al formato que recepcion el back-end.
                    if (artifact.constructor === Array) {
                        requestDTO[a] = artifact;
                    } else {
                        requestDTO[a] = [];
                        var value = artifact.constructor === String ? artifact : JSON.stringify(artifact)
                        requestDTO[a].push(value)
                    }
                }
                //console.log(JSON.stringify(requestDTO));
                return requestDTO;
            },
            async commitUI() {
                for (var i in this.project.inventions) {
                    if (this.currentCode == this.project.inventions[i].code) {
                        await this.setGlobalConext(this.project.inventions[i]);
                        this.refreshTabs();
                        break;
                    }
                }
            },

            async setGlobalConext(inv) {
                await this.axios.post(`/api/project/${this.project.code}/${inv.code}/get`).then(rs => {
                    var comboInv = []
                    for (var key in rs.data) {
                        var data = rs.data[key];
                        var select = {}
                        select[inv.keyValue] = data[inv.keyValue][0]
                        if (Object.keys(select).indexOf(inv.keyLabel) > -1) { select[inv.keyLabel] = data[inv.keyLabel][0] }
                        comboInv.push(select)
                        //{"010101_date":["2020-05-08 12:14:12.989 -0400 -04"],"Nombre":["Cepillo"],"Precio":["11000"],"SKU":["010101"]}

                        //console.log(JSON.stringify(data));

                        var labelInv= data[inv.keyLabel];

                        this.find[inv.name].push(key)
                        var invDataTAB = {}
                        invDataTAB.artifact = data;
                        invDataTAB.form = this.currentForm;
                        invDataTAB.slots = key;
                        invDataTAB.title = `<small><span style="min-width:120px" title="${key}:${labelInv}" <strong><p style="width:120px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis">${labelInv} </br> ${key} </strong></span><small>`;
                        //invDataTAB.title = `<small>${labelInv}</small>`
                        invDataTAB.subtitle = `${inv.name} : <strong> ${select[inv.keyValue]}</strong>`
                        invDataTAB.id = key;
                        invDataTAB.update = data[key + "_date"][0].substring(0, 19);
                        this.repository[inv.code][key] = invDataTAB;

                    }

                    sessionStorage.setItem(inv.code + this.project.code, JSON.stringify(comboInv));
                });
            },

            async loadCurrentInvention() {
                var upInventions = []

                //console.log("Actualizando inventos");

                var inv = {}
                this.project.inventions.forEach(invention => {
                    if (invention.code == this.currentCode) {
                        inv = invention;
                    }
                });
                this.ready = false;
                this.axios.post(`/api/project/${this.project.code}/${this.currentCode}/get`).then(rs => {
                    var comboInv = []
                    for (var key in rs.data) {
                        var data = rs.data[key];
                        var select = {}
                        select[inv.keyValue] = data[inv.keyValue][0]
                        if (Object.keys(select).indexOf(inv.keyLabel) > -1) { select[inv.keyLabel] = data[inv.keyLabel][0] }
                        comboInv.push(select)
                        this.find[inv.name].push(key)
                        var invData = {}
                        invData.artifact = data;
                        invData.form = this.currentForm;
                        invData.slots = key;
                        invData.title = `<span style="min-width:120px" >${key} </span>`;
                        invData.subtitle = `${inv.name} : <strong>${select[inv.keyValue]}</strong>`
                        invData.id = key;
                        invData.update = data[key + "_date"][0].substring(0, 19);
                        this.repository[this.currentCode][key] = invData;
                    }
                    this.currentName = inv.name;
                    this.inventionTarget = {}
                    //TODO : Pendiente que quede preseleccionado lo importado
                    this.inventionTarget[this.currentName] = "";

                });
                this.project.inventions[this.currentCode] = upInventions;
                this.ready = true;
            },

            async loadInventions() {
                this.ready = false;
                this.find = {}
                //llenar de inventos
                this.repository = {}
                for (var i in this.project.inventions) {
                    var inv = this.project.inventions[i]
                    this.find[inv.name] = []
                    this.inventionVOs = []
                    this.repository[inv.code] = []
                    this.currentForm = this.project.inventions[i];
                    await this.setGlobalConext(inv);
                    this.totalInv = Object.keys(this.repository[inv.code]).length;
                }
                this.ready = true;
            }
        }
    }

</script>
<style scoped>
    .content {
        z-index: 9999;
    }

    .min-form {
        min-height: 300px;
    }

    .workspace {
        min-height: 800px;
    }

    ul.pagination {
        z-index: 9000;
        display: inline-block;
        padding: 0;
        margin: 0;
    }

    ul.pagination li {
        display: inline;
    }

    ul.pagination li a {
        color: black;
        float: left;
        text-decoration: none;
    }


    ul.pagination {
        display: inline-block;
        padding: 0;
        margin: 0;
    }


    ul.pagination li {
        display: inline;
    }



    ul.pagination li a {
        color: black;
        float: left;
        text-decoration: none;
    }

    .invDataTAB.title{

    }
</style>
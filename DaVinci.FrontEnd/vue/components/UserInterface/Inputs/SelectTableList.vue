<template>
    <span class="card">

        <span class="card-header">            
            <div class="row">
                <find-box :context.sync="context" :target.sync="target" :id="varName + '_find_box'"> </find-box>
                <h5>
                    <i title="Seleccionar todos los visibles" @click="doSelectAll()" v-bind:class="{  'fa-check-square-o': selectAll, 'fa-square-o': !selectAll }" class="fa"></i>
                    <i :title="'Mostrar solo los Seleccionados ('+ totalSel+')'" @click="showSelects()" v-bind:class="{  'fa-eye-slash': onlySelects, 'fa-eye': !onlySelects }" class="fa"> </i>                    
                </h5>
            </div>
        </span>

        <table class="table scroll card-body" width="100%">
            <!--
            <thead>
                <slot name="columns">
                    <th v-for="column in columns" class="column" :key="column">
                        <h4>{{ column }}</h4>
                    </th>
                </slot>
            </thead>
            -->
            <tbody class="scroll">
                <tr v-for="(item, index) in filterdata" :key="index">
                    <slot :row="item">
                        <td class="column" v-for="(column, index) in columns" :key="index" v-if="index==0">
                            <label><input type="checkbox" :name="varName" :id="'chk_' + item[column]" :value="item[column]" @click="setValue(item)" :checked="Object.keys(selects).indexOf(item[column])>-1"> {{item[column]}}</label>
                        </td>
                        <td v-for="(column, index) in columns" class="description" :key="index" v-if="index==1">
                            {{item[column]}}
                        </td>
                    </slot>
                </tr>
            </tbody>
        </table>

        <span class="card-footer"><strong>{{totalSel}}</strong> {{varName}} seleccionados</span>
    </span>
</template>
<script>

    export default {
        name: "select-table",
        props: {
            data: {},
            values: Array,
            varName: String
        },
        computed: {
            totalSel() {
                return Object.keys(this.values).length;
            },

            id(){
                return this.makeId(10);
            },
            filterdata (){
                var filterdata = this.data;

                if (this.onlySelects){
                    filterdata = [];
                    for(var s in this.data){
                        var select = this.data[s];
                        var key = Object.keys(select)[0];
                        if( Object.keys(this.selects).indexOf(select[key])>-1){
                            filterdata.push(this.data[s]);
                        }
                    }
                }

                if (Object.keys(this.target).length>0){
                    var kf = Object.keys(this.target)[0];
                    var filter = this.target[kf];
                    if (filter!=""){
                        filterdata = [];
                        for(var d in this.data){
                            var stringData = JSON.stringify(this.data[d]);
                            if (stringData.toUpperCase().indexOf(filter.toUpperCase())>-1){
                                filterdata.push(this.data[d]);
                            }
                        }
                    }                    
                }        
                return filterdata;
            }
        },
        watch: {
            values(newValue, oldValue) {
                //console.log(JSON.stringify(newValue));
            }
        },
        data() {
            var _columns = Object.keys(this.data[0]);
            var _preContext = {}
            _preContext[this.varName] = []
            for (var i in this.data) {
                var key = Object.keys(this.data[i])[0];
                var name = this.data[i][key]
                _preContext[this.varName].push(name)
            }
            return {
                columns: [..._columns],
                context: _preContext,
                target: {},
                selectAll: true,
                selects: {},
                onlySelects:false
            };
        },
        mounted() {
            //PreSeleccion, Elimina tambien si esta seleccionado algun subitem que ya no este disponible o no exista.
            for (var v in this.values) {
                var keyValue = this.values[v];
                var item = this.getItemByKey(keyValue);
                if(item!=null){
                    this.selects[keyValue] = item;
                }
            }
            this.updateList();
        },
        methods: {

            showSelects(){
                this.onlySelects = !this.onlySelects;
            },

            getItemByKey(that) {
                for (var d in this.data) {
                    var data = this.data[d];
                    var find = data[Object.keys(data)[0]];
                    if (find == that) {
                        return data;
                    }
                }
                return null;
            },

            updateList(){
                var values = [];
                for (var k in this.selects) {
                    values.push(k);
                }
                this.$emit("update:values", values);
                this.$emit("update");
            },

            doSelectAll() {    
                var htmlOptions = document.getElementsByName(this.varName);
                for(var o in htmlOptions)
                if (htmlOptions[o].type == "checkbox") {
                    htmlOptions[o].checked = this.selectAll;
                    var item = this.getItemByKey(htmlOptions[o].value);
                    var key = item[Object.keys(item)[0]];
                    if (this.selectAll) {
                        var value = Object.keys(item)[1];
                        this.selects[key] = item[value]
                    } else {
                        delete this.selects[key];
                        //this.onlySelects = !this.onlySelects;
                    }
                }
                this.selectAll = !this.selectAll;
                this.updateList();              
            },
            setValue(item) {
                var isAdd = false;
                var key = item[Object.keys(item)[0]];
                var checkID = `chk_${key}`;
                var htmlOption = document.getElementById(checkID);
                if (htmlOption.type == "checkbox") {
                    isAdd = htmlOption.checked;
                }
                if (isAdd) {
                    var value = Object.keys(item)[1];
                    this.selects[key] = item[value]
                } else {
                    delete this.selects[key];
                    //this.onlySelects = !this.onlySelects;
                }
                this.updateList();
            }
        },
    };
</script>
<style scoped type="scss">
    .column  {
        min-width: 180px;
    }

    .description {
        min-width: 60%;
        white-space: nowrap;
    }

    table.scroll {
        font-size: 12px;
    }

    table.scroll tbody,
    table.scroll thead {
        display: block;
    }


    table.scroll tbody {
        max-height: 300px;
        min-height: 300px;
        overflow-y: auto;
        overflow-x: hidden;
    }

    tbody {
        border-top: 2px solid black;
    }

    tbody td,
    thead th {
        /* width: 20%; */
        /* Optional */
        border-right: 1px solid black;
        max-height: 80px !important;
        /* white-space: nowrap; */
    }

    tbody td:last-child,
    thead th:last-child {
        border-right: none;
    }
</style>
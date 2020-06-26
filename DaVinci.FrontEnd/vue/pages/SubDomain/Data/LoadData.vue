<template>

    <div class="load-windows">
        <button style="position:absolute;float:right;right: 10px;" class="btn btn-success btn-xs" @click="toClose"><i
                class="ti-close"></i> </button>
        <div class="row" width="%80">
            <div class="col-md-12">
                <h4>Cargar datos</h4>
            </div>
        </div>

        <span v-if="step==0">
            <div class="row" width="%80">
                <div class="col-md-12">
                    <h6>Seleccione el tipo de Archivo</h6>
                </div>
                <div class="col-md-2"> </div>
                <div class="col-md-2">
                    <button class="btn btn-success" @click="loadFile('csv')">
                        <img class="img" width="70px" src="@/assets/img/subDomain/DataExcel.png" alt="">
                        <center>CSV</center>
                    </button>
                </div>
                <div class="col-md-2">
                    <button class="btn btn-success" @click="loadFile('xlsx')">
                        <img class="img" width="70px" src="@/assets/img/subDomain/DataExcel.png" alt="">
                        <center>Excel</center>
                    </button>
                </div>
                <div class="col-md-2">
                    <button class="btn btn-success" @click="loadFile('txt')">
                        <img class="img" width="70px" src="@/assets/img/subDomain/DataExcel.png" alt="">
                        <center>Text</center>
                    </button>
                </div>
                <div class="col-md-2">
                    <button class="btn btn-success" @click="loadFile('xml')">
                        <img class="img" width="70px" src="@/assets/img/subDomain/DataExcel.png" alt="">
                        <center>XML</center>
                    </button>
                </div>
                <div class="col-md-2"> </div>
            </div>
        </span>

        <span v-if="step==1">

            <combo-simple label="Tipo de Datos" v-on:update="setDataType()" :value.sync="dataTypeId" class="input-item"
                :list="dataTypeList" keyValue="id" keyLabel="name"></combo-simple>
            Atributo unico es : {{dataType.principalKey}}
        </span>
    </div>

</template>
<script>
    export default {
        name: "load-data",
        computed: {

        },
        props: {

        },

        data() {
            return {
                step: 0,
                dataTypeList: {},
                dataType: {},
                dataTypeId: "",
                inventions: {}
            }
        },

        methods: {
            setDataType() {
              console.log("id" + this.dataTypeId);
                /*
                for dataTypeList
                encotrai el quetiene el id= dataTypeId, ese item
                lo igualai a dataType=item
                */
                this.dataTypeList.forEach(dList => {
                   if(this.dataTypeId === dList.id){
                       this.dataType = dList;
                   }
                })

                //console.log("typeList" + JSON.stringify(this.dataTypeList))
            },
            async loadFile(typeFile) {
                if (typeFile == "csv") {
                    this.step = 1;
                    if (Object.keys(this.dataTypeList).length == 0) {
                        this.dataTypeList = await this.getAllDataTypes();
                        this.inventions = await this.getAllInventions();
                        console.log(JSON.stringify(this.dataTypeList))
                        console.log(JSON.stringify(this.inventions))
                    }
                }
            },
            toClose() {
                this.$emit("toClose");
            },
        }
    };
</script>
<style scoped>
    .load-windows {
        width: 800px;
        min-height: 600px;
    }
    /*.img{
        width:  100px;
        height: 100px;
        object-fit:scale-down;
        margin-right: 10px;
    }*/
</style>
<template>
<div class="row" width="%80">
    <div class="col-md-12">
      <button style="position:absolute;float:right;right: 10px;" class="btn btn-success btn-xs" @click="toClose"><i class="ti-close"></i> </button>
      <h3> <i class="ti-harddrive"></i> Administración de carga de Datos.</h3>

      <div class="card">
        <!-- WIZARD Asistente de Datos -->
        <form-wizard title="Asistente de Carga de Datos" subtitle="">
          <tab-content title="Elegir formato">
            <card>
              <div class="row">
                    <div class="col-md-12">
                        <h6>Seleccione el tipo de Archivo</h6>
                    </div>
                    <div class="col-lg-3 col-sm-6">
                        <button class="btn btn-success" :class="accept.includes('csv')?'select':''" @click="loadFile('csv')">
                            <img class="img" width="70px" :class="accept.includes('csv')?'select':''" src="@/assets/img/SubDomain/DataExcel.png" alt="">
                            <center>CSV</center>
                        </button>
                    </div>
                    <div class="col-lg-3 col-sm-6 ">
                        <button class="btn btn-success" :class="accept.includes('xlsx')?'select':''" @click="loadFile('xlsx')">
                            <img class="img" width="70px" :class="accept.includes('xlsx')?'select':''" src="@/assets/img/SubDomain/DataExcel.png" alt="">
                            <center>EXCEL</center>
                        </button>
                    </div>
                    
                    <div class="col-lg-3  col-sm-6 ">
                        <button class="btn btn-success"  :class="accept.includes('txt')?'select':''" @click="loadFile('txt')">
                            <img class="img" width="70px" :class="accept.includes('txt')?'select':''" src="@/assets/img/SubDomain/DataExcel.png" alt="">
                            <center>TXT</center>
                        </button>
                    </div>
                    <div class="col-lg-3 col-sm-6 ">
                        <button class="btn btn-success" :class="accept.includes('xml')?'select':''" @click="loadFile('xml')">
                            <img class="img" width="70px" :class="accept.includes('xml')?'select':''" src="@/assets/img/SubDomain/DataExcel.png" alt="">
                            <center>XML</center>
                        </button>
                    </div>
                </div>
            </card>
          </tab-content>
          <tab-content title="Subir archivo">
            <card>
                
          <div class="row">
                <div class="col-md-12">
                    <h6>Seleccione el tipo de dato</h6>
                </div>
                <div class="col-md-6">
                <combo-simple label="Tipo de Datos" v-on:update="setDataType()" :value.sync="dataTypeId" class="input-item"
                :list="dataTypeList" keyValue="id" keyLabel="name"></combo-simple>
                </div>
                <div class="col-md-6">
                <h5>Atributo unico es : {{dataType.principalKey}}</h5>
            </div>
                <div class="col-md-12">
                    <h6>Seleccione el archivo a subir</h6>
                </div>
                <div class="col-md-12">
                <input class="form-control-file" :accept="accept"  type="file" id="DavinciFile" name="DavinciFile" ref="fileProject"
                    v-on:input="handleFileUpload()" />
                </div>
          </div>


            </card>
          </tab-content>
          <tab-content title="Validar datos">
            <card>
              <div class="col-md-12">
                  <div class="col-md-12"><h6>Tipo de dato:</h6></div>
                  <div class="col-md-6">{{this.dataType?this.dataType.name:""}}</div>
                  <div class="col-md-12"><h6>Cantidad de filas encontradas:</h6></div>
                  <div class="col-md-6">{{cantidad}}</div>
              </div>
            </card>
          </tab-content>

          <button class="btn btn-info" :disabled="nextBool" type="primary" @click="next" slot="next">Siguiente</button>
          <button class="btn btn-info" type="primary" @click="back" slot="prev">Volver</button>
          <button class="btn btn-success" type="primary" slot="finish" @click="save">Guardar Cambios</button>
        </form-wizard>
      </div>
      <!-- Listado de Datos  -->
    </div>
  </div>

</template>
<script>
    import xlsx from "xlsx";
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
                inventions: {},
                accept:"",
                cantidad:0,
                dataList:[],
                projectName:"",
                projectOK: false,
                file: "",
                doContinue: false,
                labelUpload: "Subir Data",
                error: false,
                errorMessage: "",
                importProject: {},
                invToRename: {},
                nameExist: !this.isEmptyOrSpaces(this.nameProject),
                invents: {},
                epicLength: 0,
                userStorieLength: 0,
                inventsLength: 0,
                nextBool:true
            }
        },
        async mounted() {
            this.projectName = this.getProjectDomain();
        },
        methods: {
            
            getHeader(sheet) {
            const XLSX = xlsx;
            const headers = [];
            const range = XLSX.utils.decode_range(sheet["!ref"]); // worksheet['!ref'] Is the valid range of the worksheet
            let C;
            /* Get cell value start in the first row */
            const R = range.s.r; //Line / / column C
            let i = 0;
            for (C = range.s.c; C <= range.e.c; ++C) {
                var cell =
                sheet[
                    XLSX.utils.encode_cell({ c: C, r: R })
                ]; /* Get the cell value based on the address  find the cell in the first row */
                var hdr = "UNKNOWN" + C; // replace with your desired default
                // XLSX.utils.format_cell Generate cell text value
                if (cell && cell.t) hdr = XLSX.utils.format_cell(cell);
                if(hdr.indexOf('UNKNOWN') > -1){
                if(!i) {
                    hdr = '__EMPTY';
                }else {
                    hdr = '__EMPTY_' + i;
                }
                i++;
                }
                headers.push(hdr);
            }
            return headers;
            },
            back(){
                this.nextBool=true;
                if(this.step>0){
                    this.step--;
                }
                if(this.step==0){
                    this.dataType={}
                    this.cantidad=0
                    this.dataList=[]
                    this.dataTypeList= {}
                    this.dataTypeId= ""
                    this.inventions= {}
                    this.accept=""
                    this.file=""
                    document.getElementById('DavinciFile').value="";
                }
                if(this.dataType!={}&&this.file!=""){
                    this.nextBool=false;
                }
            },
            next(){
                if(this.step==1&&(this.file==""||this.dataType=={})){
                    return alert("Debe de seleccionar un archivo a subir y un tipo de dato para continuar")
                }else
                if(this.step<2){
                    this.step++;
                    this.nextBool=true
                }
                if(this.step==2){
                   this.getData();
                }
            },
            dupCheck(o, val){
                var encontrado = false;
                for(var prop in o){
                    if(prop.toLowerCase() === val.toLowerCase()){
                    encontrado = true;
                    break;
                    }   
                }
                return encontrado;
            },
            getKeyCase(o,val){
                for(var prop in o){
                    if(prop.toLowerCase()===val.toLowerCase()){
                        return prop;
                    }
                }
            },
            getData(){
                if (this.typeFile == "xlsx"||this.typeFile=='xls') {
                    this.importExcel();
                }else if(this.typeFile=="csv"){
                    this.importCsv();
                }else if(this.typeFile=="txt"){
                    this.importTxt();
                }else if(this.typeFile=="xml"){
                    this.importXml();
                }
            },
            importXml(){
                var parseString = require('xml2js').parseString;
                const file=this.file;
                const dataKey=this.dataType.principalKey
                var reader = new FileReader();
                var self = this;
                let column=[]
                let position=0;
                let encontrado=false;
                this.dataList=[]
                this.cantidad=0
                let index=0

                reader.addEventListener('load', function(){
                    try{
                self.xml = reader.result;
                parseString(self.xml, function (err, result) {
                self.events = result.Workbook.Worksheet[0]["ss:Table"][0]["Row"]
                self.events.forEach(value=>{
                    if(encontrado && value.Cell[position]&& value.Cell[position].Data[0]["_"] && value.Cell[position].Data[0]["_"].trim()!=""){
                        self.dataList.push(value.Cell[position].Data[0]["_"]);
                        self.cantidad++;
                    }
                    if(value.Cell[position]&&index==0){
                        value.Cell.forEach(function (cell, cont){
                        if(cell.Data[0]["_"].toLowerCase==dataKey.toLowerCase){
                            position=cont;
                            encontrado=true;            
                            }
                        })
                    }
                    index++;
                })
                });
                } catch (e) {
                    return alert("Error de lectura");;
                    }
                });
                reader.readAsText(file);
            },
            importTxt(){
                const file=this.file;
                const fileReader = new FileReader();
                const dataKey=this.dataType.principalKey
                fileReader.onload= ev =>{
                    try{
                    const data= ev.target.result;
                    let row=data.split("\r\n")
                    let column=[]
                    let position=0;
                    let encontrado=false;
                    this.dataList=[]
                    this.cantidad=0
                    let index=0
                    row.forEach(value=>{
                        column=value.split("\t")
                        if(encontrado && column[position] && column[position].trim()!=""){
                            this.dataList.push(column[position]);
                            this.cantidad++;
                        }
                        if(index==0){
                            column.forEach(function(key,cont){
                                if(key.toLowerCase==dataKey.toLowerCase){
                                    position=cont;
                                    encontrado=true;
                                }
                            });
                        }
                        index++;
                    });
                    } catch (e) {
                    return alert("Error de lectura");;
                    }
                }
                fileReader.readAsBinaryString(file);
            },
            importCsv(){
                const file=this.file;
                const fileReader = new FileReader();
                const dataKey=this.dataType.principalKey
                fileReader.onload= ev =>{
                    try{
                    const data= ev.target.result;
                    let row=data.split("\r\n")
                    let column=[]
                    let position=0;
                    let encontrado=false;
                    this.dataList=[]
                    this.cantidad=0
                    let index=0
                    row.forEach(value=>{
                        column=value.split(",")
                        if(encontrado){
                            this.dataList.push(column[position]);
                            this.cantidad++;
                        }
                        if(index==0){
                            column.forEach(function(key,cont){
                                if(key.toLowerCase==dataKey.toLowerCase){
                                    position=cont;
                                    encontrado=true;
                                }
                            });
                        }
                        index++;
                    });
                    } catch (e) {
                    return alert("Error de lectura");;
                    }
                }
                fileReader.readAsBinaryString(file);
            },
            importExcel() {
                const files=this.file
                const data =[]
                if (!/\.(xls|xlsx)$/.test(files.name.toLowerCase())) {
                    return alert("El archivo no tiene formato xls o xlsx");
                }
                const fileReader = new FileReader();
                fileReader.onload = ev => {
                    try {
                    const data = ev.target.result;
                    const XLSX = xlsx;
                    const workbook = XLSX.read(data, {
                        type: "binary"
                    });
                    const wsname = workbook.SheetNames[0];
                    const ws = XLSX.utils.sheet_to_json(workbook.Sheets[wsname]); 
                    const excellist = []; 
                    for (var i = 0; i < ws.length; i++) {
                        excellist.push(ws[i]);
                    }
                    if(this.dupCheck(excellist[0],this.dataType.principalKey)){
                        let key=this.getKeyCase(excellist[0],this.dataType.principalKey);
                        this.dataList=[]
                        this.cantidad=0
                        excellist.forEach(value=>{
                            this.dataList.push(value[key]);
                            this.cantidad++;
                        });
                    }
                    
                    } catch (e) {
                    return alert("Error de lectura");;
                    }
                };
                fileReader.readAsBinaryString(files);
                
            },
            setDataType() {
                /*
                for dataTypeList
                encotrai el quetiene el id= dataTypeId, ese item
                lo igualai a dataType=item
                */
                this.dataTypeList.forEach(dList => {
                   if(this.dataTypeId === dList.id){
                       this.dataType = dList;
                   }
                   if(Object.keys(this.dataType).length!=0&&this.file!=""){
                        this.nextBool=false;
                    }
                })
                

                //console.log("typeList" + JSON.stringify(this.dataTypeList))
            },
            async loadFile(typeFile) {
                this.nextBool=false;
                this.typeFile=typeFile;
                if (typeFile == "csv") {
                    this.accept=".csv";
                }else if (typeFile == "txt") {
                    this.accept=".txt";
                }else if (typeFile == "xlsx" || typeFile == "xls") {
                    this.accept=".xlsx,.xls";
                }else if (typeFile == "xml") {
                    this.accept=".xml";
                }
                if (Object.keys(this.dataTypeList).length == 0) {
                        this.dataTypeList = await this.getAllDataTypes();
                        this.inventions = await this.getAllInventions();
                    }
            },
            async save(){
                var me = this;
                var json = {
                    atributes: this.dataList,
                    id:this.dataType.id,
                    dataType:this.dataType
                    }
                await me.axios.post(`/davinci/${this.projectName}/repository/save/`,json).then(rs => {
                    this.toClose()
                    return alert(rs.data)
                });
            },
            toClose() {
                this.$emit("toClose");
            },
            ready() {
        return this.doContinue;
      },
      openFile() {
        this.invToRename = {}
        this.error = false;
        if (this.step == 0) {
          this.$refs.fileProject.value = null;
          document.getElementById("DavinciFile").click();
        }

      },
            
      async handleFileUpload(event) {
        this.file = this.$refs.fileProject.files[0];
        if(Object.keys(this.dataType).length!=0&&this.file!=""){
            this.nextBool=false;
        }
        this.labelUpload = "Desea subir otro archivo.";
        if (this.file) {
            var readFile = new FileReader();
            let data={"data":this.file}
          };
        }

        }
    };
</script>
<style scoped>
    .load-windows {
        width: 800px;
        min-height: 600px;
    }
    .select{
        background-color:#218838 !important; 
    }
    .disable{
        pointer-events:none
    }
    /*.img{
        width:  100px;
        height: 100px;
        object-fit:scale-down;
        margin-right: 10px;
    }*/
</style>
<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label>
        <input type="file" id="DavinciFile" name="DavinciFile" ref="fileProject" v-on:change="handleFileUpload()"
          hidden />
      </label>

      <button class="btn btn-success" v-if="step==0" @click="openFile">
        <span style="color:blue" class="ti-upload"></span> Abrir archivo Davinci
      </button>
      <button class="btn btn-success" v-if="step==1" @click="submitFile">Cargar Archivo</button>
    </div>
  </div>
</template>

<script>
  export default {
    /*
        Defines the data used by the component
      */
    data() {
      return {
        file: "",
        step: 0,
        valid: {
          inventions: true,
          nameProject: true
        }
      };
    },

    methods: {
,
 
      openFile() {
        document.getElementById("DavinciFile").click();
      },
      /*
          Submits the file to the server
        */
      submitFile() {
        let formData = new FormData();
        if (this.valid.inventions) {
          formData.append("DavinciFile", this.file);
          this.axios
            .post("/api/project/import/", formData, {
              headers: {
                "Content-Type": "multipart/form-data"
              }
            })
            .then(function (rs) {
              console.log("SUCCESS!!" + rs.data);
            })
            .catch(function () {
              console.log("FAILURE!!");
            });
        } else {
          this.alertError(
            "Hemos encontrado algunos erroes eun dsofih dsajkhfkjsb"
          );
        }
      },

      handleFileUpload(event) {
        this.file = this.$refs.fileProject.files[0];
        if (this.file) {
          var readFile = new FileReader();
          var me = this;
          readFile.onload = async function (event) {
            var content = event.target.result;
            var jsonImport = JSON.parse(content);
            me.alertInfo(`Evaluando proyecto ${jsonImport.project.Name}`);
            me.valid.inventions = false;
            //Itera json.inventions , y valida si existen en nuestro ambiente
            //Compara ese listado con el GET que te retorne esta url /api/invention/all/
            //rs.data[] == json.inventions 
            await me.axios.get("/api/invention/all/").then(rs => {
              var jsonInvents = rs.data;
              for (var i = 0; i < jsonInvents.length; i++) {
                for (var y = 0; y < jsonImport.inventions.length; y++) {
                  if (jsonInvents[i].code == jsonImport.inventions[y].code) {
                    console.log(jsonInvents[i].name)
                    me.valid.invention = false;
                    me.alertInfo("Invento " + jsonInvents[i].name + " ya existe..");
                  }
                }
              }
            });
          };
          readFile.readAsText(this.file);
        }
        this.step = 1;
      }
    }
  };
</script>
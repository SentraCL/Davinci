<template>
  <div class="container">
    <div class="large-12 medium-12 small-12 cell">
      <label>
        <input type="file" id="DavinciFile" name="DavinciFile" ref="file" v-on:change="handleFileUpload()" hidden />
      </label>


      <button class="btn btn-success" v-if="step==0" @click="openFile"><span style="color:blue" class="ti-upload"></span> Abrir archivo Davinci </button>
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
        file: '',
        step: 0
      }
    },

    methods: {
      openFile() {
        console.log("XXXDDD");
        document.getElementById('DavinciFile').click();
      },
      /*
        Submits the file to the server
      */
      submitFile() {
        /*
                Initialize the form data
            */
        let formData = new FormData();

        /*
            Add the form data we need to submit
        */
        formData.append('DavinciFile', this.file);

        /*
          Make the request to the POST /single-file URL
        */
        this.axios.post('/api/project/import/',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
        ).then(function (rs) {

          console.log('SUCCESS!!' + rs.data);

        })
          .catch(function () {
            console.log('FAILURE!!');
          });
      },

      /*
        Handles a change on the file upload
      */
      handleFileUpload() {
        this.file = this.$refs.file.files[0];
        this.step = 1;
      }
    }
  }
</script>
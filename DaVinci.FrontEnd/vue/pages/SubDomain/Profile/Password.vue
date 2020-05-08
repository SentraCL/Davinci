<template>
  <form @submit.prevent class="password">
    <div class="row">
      <div class="col-md-1"></div>
      <div class="col-md-10">
        <input-text index="1" type="password" label="Nueva Contraseña" required v-model="datos.oldData"></input-text>
        <input-text index="2" type="password" label="Repita Contraseña" required v-model="datos.newData"></input-text>
        <br />
        <button style="float:right" type="button" class="btn" @click="validar()">Cambiar</button>
      </div>
    </div>
    <br />
  </form>
</template>
<script>

  export default {
    name: "password",
    computed: {},

    data() {
      return {
        datos: {
          oldData: "",
          newData: ""
        }
      };
    },

    mounted() { },


    methods: {
      validar() {
        if (this.datos.oldData != this.datos.newData) {
          this.alertInfo("Las claves no son iguales");
        } else {
          if (this.datos.newData.length < 5) {
            this.alertInfo("La clave debe poseer al menos 5 carácteres");
          } else {
            var uno = this.getCodeProject().then(rs => {
              //console.log("/davinci/" + rs + "/user/changePassword/")
              this.axios.put("/davinci/" + rs + "/user/changePassword/")
                .then(rs => {
                  //console.log("hola");
                });
            });

          }
        }
      }
    }
  }
</script>
<style scoped>
  .password {
    min-height: 200px;
    min-width: 800px !important;
  }
</style>
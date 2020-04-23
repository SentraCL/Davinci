<template>
  <div>
    <span v-if="inputs.length>0" v-for="input in inputs" :title="form[input.id]">
      <combo-simple v-if="input.visible" :value.sync="form[input.id]" :inactive="input.inactive" :showKey="true" v-on:update="change(input)" :label="input.label" class="input-item" :list="input.list" :keyValue="input.keyValue" :keyLabel="input.keyLabel"></combo-simple>
      <combo-simple v-if="!input.visible" :label="input.label" class="input-item"></combo-simple>
    </span>
  </div>

</template>
<script>
  export default {
    name: "pre-condition",
    computed: {

    },
    props: {
      preconditions: {},
      values: {},
      inactive: {},
      codeFormat: {}
    },

    watch: {
      async inactive(n, o) {
        //console.log("Cambio Inactive!!!");
        await this.load();
      },
      form(n, o) {
        this.$emit("update:values", n);
      }
      /*
      ,
      async values(n, o) {
        if (JSON.stringify(n) == JSON.stringify(o)) {
          await this.mounted();
        }
      }
      */
    },

    data() {
      return {
        codeValue: "",
        codeSelect: "",
        inputs: [],
        form: {},

      }
    },

    async mounted() {

      await this.load();
    },

    methods: {
      async load() {
        this.inputs = [];
        for (var p in this.preconditions) {
          var preCondition = this.preconditions[p];
          var def = await this.getInventionDefByCode(preCondition.code);
          var isInactive = false;
          if (this.inactive) {
            isInactive = true;
          } else {
            isInactive = this.codeFormat !== undefined;
            if (isInactive) {
              isInactive = this.codeFormat.indexOf(preCondition.id) > -1 && !this.isEmptyOrSpaces(this.values[preCondition.id]);
            }
          }

          var input = {
            id: preCondition.id,
            artifacs: def.artifacs,
            keyLabel: def.keyLabel,
            keyValue: def.keyValue,
            allData: def.list,
            WareHouse: def.WareHouse,

            visible: true,
            inactive: isInactive,

            label: preCondition.label,
            list: (preCondition.parent.id == "") ? def.list : [],
            parent: preCondition.parent,
            code: preCondition.code,
            idPreCondition: preCondition.id,
            value: this.values[preCondition.id]//preCondition.value
          }

          this.form[preCondition.id] = preCondition.value

          this.inputs.push(input);
        }
        this.inputs.forEach(input => {
          this.loadChildren(input);
        })

      },
      getUserStory() {
        //console.log("Escogio uno!");
      },

      change(input) {
        input.visible = false;
        setTimeout(() => {
          this.loadChildren(input);
        }, 60);
        this.$emit("update:values", this.form);
        this.$emit("update");
        //console.log(JSON.stringify(this.form));
        input.visible = true;
        input.value = this.form[input.id];
      },
      loadChildren(parent) {
        //this.ready = false;        
        var value = this.form[parent.id]//parent.value;
        this.inputs.forEach(input => {
          if (input.parent.id == parent.id) {
            try {
              var keyValue = input.keyValue;
              var data = parent.WareHouse[value];
              var artifactName = input.parent.artifactName;


              if (artifactName != null || data[artifactName] !== undefined) {
                var filter = data[artifactName];
                input.list = [];
                if (input.allData.length > 0 && (filter !== undefined || filter != null)) {
                  input.allData.forEach(item => {
                    if (filter.indexOf(item[keyValue]) > -1) {
                      input.list.push(item);
                    }
                  })
                }
                //this.form[preCondition.id] = "";

              }
            } catch (error) {
              //No se pudo leer la precondicion!
            }

          }
        })
        //console.log(JSON.stringify(this.form));
        //this.$emit("update:values", this.form);
        //this.form["_"] = this.makeId(100);



      },
    }
  };
</script>
<style></style>
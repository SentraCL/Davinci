<template>
  <div>

    <table class="table">
      <thead>

        <th v-for="column in columns" :key="column">{{ column }}</th>
        <th>Ultima Modificación</th>
        <th>Actividades</th>

      </thead>
      <tbody>

        <tr v-for="epic in epics" :key="epic.id">
          <!---->
          <td v-for="column,index in columns" :key="index">

            <small v-if="reference!=column" :title="epic.attributes[index].value">{{epic.attributes[index].value | readMore(140, '...')}}</small>
            <span style="display:flow-root" class="btn btn-success" @click="getEpic(epic.id)" v-if="reference==column"><i class="ti-share"></i> {{epic.attributes[index].value}}</span>
          </td>
          <td>
            <small>
              {{epic.author}}
              <br>
              {{formatDays(epic.date)}}
            </small>
          </td>
          <!--<td><small class="btn btn-xs btn-info" style="display:flex">Ver Actividades ({{epic.activities.length}})</small></td>-->
          <td>

          </td>
        </tr>

      </tbody>
    </table>
  </div>
</template>
<script>
  export default {
    name: "epic-table",
    props: {
      value: [String, Number],
      reference: String,
      epics: {}
    },
    data() {
      return {
        epicRef: "",
        columns: []

      }
    },
    watch: {
      epics() {
        if (this.epics) {
          if (this.epics.length > 0) {
            this.epics[0].attributes.forEach(atr => {
              this.columns.push(atr.name);
            })
          }
        }
      }
    },


    methods: {
      getEpic(id) {
        this.$emit("update:value", id);
        this.$emit("update");
        this.$emit("select");
      },
    }
  };
</script>
<style></style>
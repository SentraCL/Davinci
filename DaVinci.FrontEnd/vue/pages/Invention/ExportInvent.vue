<template>
  <div class="row">
    <div class="col-md-12">
      <d-button type="success" class="btn " round @click.native.prevent="exportCSV">
        EXPORTAR CSV
      </d-button>
      <d-button type="success" class="btn " round @click.native.prevent="exportJSON">
        EXPORTAR JSON
      </d-button>
    </div>
    <div class="col-md-12">
      </br>
      <a href="#" style="display:none" id="hiddenDownload"></a>
    </div>
    <div class="col-md-12">
      <davinci-table :pagination="page" v-if="artifacts.length>0" :data="artifacts"></davinci-table>
    </div>
  </div>
</template>
<script>
  export default {
    name: "import-invent",
    computed: {

    },
    props: {
      project: {},
      name: String,
      code: String,
      inventions: {}
    },

    data() {
      return {
        page: {
          itemsPerPage: 6
        }
      }
    },

    computed: {
      artifacts() {
        var artifacts = this.getArtifacts(this.inventions);
        var prettyArtifacts = []
        for (var a in artifacts) {
          var prettyArtifact = {}
          var artifact = artifacts[a]
          for (var k in Object.keys(artifact)) {
            var key = Object.keys(artifact)[k]
            if (key.endsWith("_date")) {
              key = "Fecha";
            }
            if (key != "") {
              var value = (artifact[key] == "1" || artifact[key] == "0") ? artifact[key] == "1" ? "Si" : "No" : artifact[key];
              prettyArtifact[key] = value
            }
          }

          prettyArtifacts.push(prettyArtifact)
        }
        return prettyArtifacts;
      }
    },

    methods: {
      exportCSV() {
        var artifacts = [];
        let csvContent = "data:text/csv;charset=utf-8,"
        var index = 0;
        for (var j in this.inventions) {
          var artifact = this.inventions[j].artifact;
          if (index == 0) {
            for (var col in Object.keys(artifact)) {
              var columnName = Object.keys(artifact)[col];
              if (columnName.endsWith("_date")) {
                columnName = "FECHA"
              }
              csvContent += `${columnName.toUpperCase()};`
            }
            csvContent += "\n";
          }
          for (var col in artifact) {
            csvContent += `${artifact[col]};`
          }
          csvContent += "\n";
          index++;

        }
        var dataStr = csvContent;
        var hiddenClick = document.getElementById('hiddenDownload');
        hiddenClick.setAttribute("href", dataStr);
        hiddenClick.setAttribute("download", `${this.project.name}_${this.name}.csv`);
        hiddenClick.click();
        this.$emit("success");;
      },
      exportJSON() {
        var artifacts = this.getArtifacts(this.inventions)

        var dataStr = "data:text/json;charset=utf-8," + encodeURIComponent(JSON.stringify(artifacts, null, 4));
        var hiddenClick = document.getElementById('hiddenDownload');
        hiddenClick.setAttribute("href", dataStr);
        hiddenClick.setAttribute("download", `${this.project.name}_${this.name}.json`);
        hiddenClick.click();
        this.$emit("success");;
      },
      getArtifacts() {
        var artifacts = [];
        for (var j in this.inventions) {
          var artifact = this.inventions[j].artifact;
          var jsonArtifact = {}
          //Convierte JSON Custom a JSON comun de descarga.
          for (var k in artifact) {
            if (artifact[k].constructor === Array && artifact[k].length == 1) {
              jsonArtifact[k] = artifact[k][0];
            } else {
              jsonArtifact[k] = artifact[k];
            }
          }
          artifacts.push(jsonArtifact);
        }
        return artifacts;
      },
      close() {
        this.$emit("success");;
      }
    }

  };
</script>
<style></style>
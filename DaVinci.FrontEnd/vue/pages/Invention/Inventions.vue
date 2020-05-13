<template>
  <div>
    <span class="row">
      <div class="col-md-2">
        <input-text label="Artefacto" alphabetic removeSpace capitalize :value.sync="name" autocomplete="off">
        </input-text>
      </div>
      <div class="col-md-2">
        <button type="button" @click="createInvention" class="btn btn-xs btn-round btn-info">
          <span style="color:black" class="fa fa-magic"></span> Inventar
        </button>
      </div>
    </span>

    <span class="row">
      <div class="col-md-12">
        <small>Los Inventos, son un conjuto de artefactos, los cuales dan contexto a las historias de Usuarios que seran usadas en cada Epico.</small>
      </div>
    </span>

    <div style="width: auto">
      <uib-pagination v-if="inventionVOs.length>itemsPerPage" firstText="«" lastText="»" previousText="‹" nextText="›" :total-items="inventionVOs.length" v-model="pagination" @change="pageChange()" :itemsPerPage="itemsPerPage" :max-size="10" :boundary-links="true">
      </uib-pagination>
    </div>
    <span v-if="inventionVOs.length >0">

      <v-tabs :tabs="inventionsPerPage" v-if="refresh">
        <span :slot="invo.slots" v-for="(invo, index) of inventionVOs" :key="index">
          <artifact :artifactForm.sync="invo" v-on:save="notifySave()" v-on:drop="getInventions()" v-on:cancel="getInventions()"></artifact>
        {{invo.origin}}
        </span>
      </v-tabs>

    </span>

  </div>
</template>
<script>
  import Vue from "vue";
  import Artifact from "./Artifact.vue";


  export default {
    components: {
      Artifact
    },
    data() {
      return {
        pagination: {
          currentPage: 1,
        },
        itemsPerPage: 20,
        selectedDate: "",
        name: "",
        inventionVOs: [],
        refresh: true
      }
    },

    async created() {
      await this.getInventions();
    },

    computed: {
      inventionsPerPage: function () {

        var totalData = Object.keys(this.inventionVOs).length;
        var total = 0;
        var currentPage = this.pagination.currentPage;
        var max = this.itemsPerPage
        var current = 0;
        var cp = 0;

        if (totalData > this.itemsPerPage) {
          max = this.itemsPerPage
          current = max * currentPage - max;
          total = current + max > totalData ? totalData : current + max;
        } else {
          total = totalData;
        }
        var inventionsPerPage = [];

        var primero = 1;
        for (cp = current; cp < total; cp++) {
          var inventionPerPage = {};
          inventionPerPage = this.inventionVOs[cp];
          var title = inventionPerPage.name;
          if (title.length > 25) {
            title = inventionPerPage.name.substring(0, 25) + ".."
          }
          inventionPerPage.title = `<span title="${inventionPerPage.name}"><img width="16px" src="${inventionPerPage.icon}"> ${title}</span>`;
          inventionsPerPage.push(inventionPerPage);

        }
        return inventionsPerPage;
      }

    },

    methods: {
      pageChange() {
        //console.log("this.pagination.currentPage >>" + this.pagination.currentPage);
      },

      async  getInventions() {
        await this.axios.get("/api/invention/all/").then(rs => {
          this.inventionVOs = rs.data;
        });
      },
      createInventionVO(name) {
        var inventionVO = {
          name: name,
          artifacts: [],
          id: this.inventionVOs.length + 1,
          title: "",
          subtitle: "",
          slots: name,
          icon: ""
        }

        return inventionVO
      },
      exist(name) {
        for (var i in this.inventionVOs) {
          var inv = this.inventionVOs[i];
          if (inv.name == name) {
            return true;
          }
        }
        return false;
      },
      async saveAllArtifact() {
        await this.axios.post("/api/invention/saveAll/", this.inventionVOs).then(rs => {
        });
        await this.getInventions();
      },
      notifySave() {
        //console.log("Notificar Guardado.");
      },
      createInvention() {
        if (!this.exist(this.name) && (this.name)) {
          this.refresh = false;
          setTimeout(() => {
            var invention = this.createInventionVO(this.name)
            this.inventionVOs.unshift(invention)
            this.name = "";
            this.refresh = true;
          }, 10);
        }
      }
    }
  };
</script>
<style>
  ul.pagination {}

  ul.pagination li.pagination-page {
    padding-left: 15px;
    padding-right: 15px;
  }

  li.pagination-page.ng-scope.active {
    background-color: #28a745;
  }

  li.pagination-page.ng-scope a {
    color: #999;

  }


  li.pagination-page.ng-scope.active a {
    color: #FFF;
    font-weight: bold;
  }

  ul.pagination li.pagination-prev {
    padding: 10px;
  }

  ul.pagination li.pagination-prev a {
    color: #4e1d00;
  }

  ul.pagination li.pagination-next {
    padding: 10px;
  }

  ul.pagination li.pagination-next a {
    color: #4e1d00;
  }

  ul.pagination li.pagination-first {
    padding: 10;
  }

  ul.pagination li.pagination-first a {
    color: #4e1d00;
  }


  ul.pagination li.pagination-last {
    padding: 10px;
    width: 5px;
  }

  ul.pagination li.pagination-last a {
    color: #4e1d00;
    width: 5px;
  }

  ul.pagination li {
    text-align: center;
    padding: 10px;
    width: 5px;
  }
</style>
<template>

  <div>
    <div class="card" style="min-height: 150px !important;">
      <div class="card-header">
        <div class="typo-line">
          <p class="category">
            <img width="128px" :src="project.avatar" />
            <br />

          </p>
          <blockquote>
            <h4>
              <img src="@/assets/img/SubDomain/d.epic.png" width="28px">  {{project.name}}</h4>
            <p v-html="project.resume"></p>
          </blockquote>
        </div>
      </div>
      <div class="card-body">

      </div>
      
      <div class="card-footer">
        <span class="click" @click="createNew()"><span class="ti-new-window"></span> Crear nuevo tipo de <strong>Epico</strong></span>
      </div>
    </div>

    <epic-design v-if="task==NEW" :epicType="epicForm" v-on:back="listEpic();" :inventions="inventions" :project="project"></epic-design>

    <div class="card-group" v-if="task==NONE">
      <div class="col-md-3" v-for="(epicType, index) in epicTypes" :key="index">
        <card style="background-color:khaki;">
          <drop-menu :title="tasktitle" style="float:left;top:-10px;" v-on:change="menuEpicType(option,epicType,index)" :options="epicOptions" :option.sync="option"></drop-menu>
          
          <template slot="header">
            <span class="card-title">
              <h5><strong> {{epicType.name}}</strong></h5>
            </span>
            <span class="row">
              <span class="col-md-12">
                {{epicType.definition}}
              </span>
            </span>
          </template>

        </card>
      </div>

    </div>
    <epic-design v-if="task==EDIT" :epicType="epicForm" :inventions="inventions" v-on:back="listEpic();" :project="project"></epic-design>
  </div>
  
</template>
<script>

  import EpicDesign from "./EpicDesign.vue";

  export default {
    name: "user-stories",
    computed: {
       filterInventions: function () {
        var filterInventions = [];
        for (var i in this.inventionVOs) {
          var inv = this.inventionVOs[i];
          if (
            inv.name.toUpperCase().indexOf(this.inventionName.toUpperCase()) >
            -1 ||
            this.inventionName == ""
          ) {
            var include = false;
            var invCount = 0;
            var artiNames = "";
            for (var a in inv.artifacts) {
              var artifact = inv.artifacts[a];
              if (!artifact.isEssential) {
                invCount++;
                include = true;
                artiNames += `\n${artifact.name} : ${artifact.nickName}`;
              }
            }
            if (include) {
              inv.resume = `${inv.name} : Depende de ${invCount} inventos para funcionar: ${artiNames}`;
            }
            filterInventions.push(inv);
          }
        }
        return filterInventions;
      }
    },

    components: {
      EpicDesign,
    },
    

    props: {
      project: {},
      inventions: {},
    },
    

    data() {
      return {
        inventionName: "",
        EDIT: 1,
        NEW: 0,
        NONE: -1,

        task: -1,

        epicForm: {
          name: "",
          definition: "",
          attributes: [],
          inventions: [],
        },
        tasktitle: `<small style="color:green"><i class="ti-angle-double-down"></i> Menu</small>`,
        option: -1,
        epicOptions: [
          { html: `<div><i style="color:green" class="ti-panel"></i> Editar</div>`, option: 1 }
          // { html: `<div><i style="color:red" class="ti-thumb-down"></i> Descativar</d>`, option: 0 },
        ],

        selectedDate: "",
        name: "",
        inventionVOs: [],
        epicTypes: [],
        refresh: true
      }
    },
    async mounted() {
      await this.getInventions();
      await this.getEpicTypes();
    },
    methods: {
      async listEpic() {
        await this.getEpicTypes();
        this.task = this.NONE;
      },
    setInvention(invention) {
        this.currentInvention = invention;
      },
      createNew() {
        this.epicForm = {
          name: "",
          definition: "",
          attributes: [],
          inventions: [],
        },
          this.task = this.NEW
      },

      menuEpicType(option, epicForm, index) {

        if (option == this.EDIT) {
          this.epicForm = epicForm
          this.task = this.EDIT;
          //console.log(JSON.stringify(epicForm));
        }
      },
      drag(ev) {
        if(ev.target.name){
          localStorage.setItem("text",ev.target.name);
        }else if(ev.target.textContent){
          localStorage.setItem("text",ev.target.textContent);
        }
      },
      drop(ev){
        ev.preventDefault();
        var data= ev.dataTransfer.getData("text");
        ev.target.innerHtml=document.getElementById(data).dataset.value;
      },
      async  getEpicTypes() {
        await this.axios.get(`/api/project/${this.project.code}/design/epic`).then(rs => {
          this.epicTypes = rs.data == null ? [] : rs.data;
        });
      },
      async  getInventions() {
        await this.axios.get("/api/invention/all/").then(rs => {
          this.inventionVOs = rs.data;
        });
      },
    }
  };
</script>
<style>
/* width */
  ::-webkit-scrollbar {
    width: 20px;
  }

  /* Track */
  /* Handle */
  ::-webkit-scrollbar-thumb {
    background: #fff;
    border-radius: 10px;
  }

  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background: #4e1d00;
    transform: translate3d(0, 0, 0);
    backface-visibility: hidden;
    perspective: 1000px;
  }

  .drag,
  .drop {
    font-family: sans-serif;
    display: inline-block;
    border-radius: 10px;
    background: transparent;
    color: #000;
    position: relative;
    padding: 2px;
    min-width: 80px;
    text-align: center;
    vertical-align: top;
  }

  .drag {
    cursor: grab !important;
    color: #000;
    font-size: 10px;
    text-transform: uppercase;
  }

  .drag:hover {
    animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
    transform: translate3d(0, 0, 0);
    backface-visibility: hidden;
    perspective: 1000px;
  }

  .drop {
    background: transparent;
    color: #000;
  }

  .scrolling {
    overflow-x: scroll;
    overflow-y: hidden;
    white-space: nowrap;
    min-height: 90px;
  }

  .close {
    cursor: pointer;
  }

  .labelPrj {
    font-size: 12px;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif;
    text-transform: uppercase;
  }

  .invention {
    margin-top: 3px;
    transition: 0.3s;
  }

  .invention:hover {
    font-size: 1em;
    margin-top: -5px;
  }

  .min-invention {
    cursor: default;
    margin-top: 3px;
    transition: 0.3s;
  }

  .min-invention:hover {
    font-size: 1em;
    margin-top: -5px;
  }

  .miniPrj {
    cursor: pointer;
    font-size: 12px;
    font-weight: bold;
    font-family: Arial, Helvetica, sans-serif;
    text-transform: capitalize;
    color: #000;
    display: inline-block;
    border-radius: 5px;
    background: #e4d6c0;
    position: relative;
    padding: 8px;
    min-width: 20px;
    border-right: 1px solid #906538;
    border-bottom: 1px solid #4e1d00;
    text-align: center;
    vertical-align: top;
  }
</style>
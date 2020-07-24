<template>

  <div>
    <slot name="task-area">
      <div v-for="(task,index) in tasks" track-by="id" :key="index">

        <div id="lastItem" v-if="index==(tasks.length-1)"></div>
        <epic v-if="task.isEpic && task.exist" :epic.sync="task.epic" v-on:remove="removeTask(index)" :form.sync="task.form" v-on:save="loadEpics()" :idUS.sync="idUS" v-on:sendToWorkSpace="takeUST"></epic>
        <user-story :task.sync="task" v-if="task.isUserStory  && task.exist" :userStory.sync="task.userStory" v-on:remove="removeTask(index)" :idUS.sync="idUS" v-on:sendToWorkSpace="takeUST"></user-story>
      </div>

    </slot>
  </div>

</template>
<script>
  import Epic from "./Epic/Epic.vue";
  import UserStory from "./UserStories/UserStory.vue";

  export default {
    name: "workspace",

    components: {
      Epic,
      UserStory
    },
    watch: {
      tasks(newValue, oldValue) {
        //TODO : Validar no tener tareas repetidas.
      }
    },
    props: {
      artifactForm: {},
      tasks: {}
    },

    data() {
      return {
        countEP: 1,
        countUS: 1,
        //projectCode: "",
        idUS: "",
        epics:{}
      }
    },

    async mounted() {
      //this.projectCode = await this.getCodeProject();
    },

    methods: {

      async findEpic(id) {
        var epic = {}
        var epics = await this.getEpics();
        this.epics=epics
        epics.forEach(epicache => {
          if (epicache.id == id) {
            epic = epicache;
          }
        });
        var itemType = {};
        var epicTypes = await this.getTypesOfEpics();

        epicTypes.forEach(typepic => {
          if (epic.type == typepic.Code) {
            itemType = typepic;
          }
        });


        var epicReference = "";
        var index = 0;
        epic.attributes.forEach(atr => {
          itemType.attributes[index].value = atr.value;
          index++;
        })
        /*
        console.log(`EPIC =  ${JSON.stringify(epic.attributes)}`);
        console.log(`itemType = ${JSON.stringify(itemType.attributes)}`);
        */
        var choiceEpic = {
          isEpic: true,
          exist: true,
          epic: {
            id: epic.id,
            formatDate: this.formatDays(epic.date),
            name: itemType.name,
            type: itemType.Code,
            definition: itemType.definition,
            title: epicReference,
            epicForm: {
              reference: itemType.reference,
              attributes: itemType.attributes,
              inventions: itemType.inventions,
              activities: epic.activities // <===Para llevar las actividades que vienen desde el Backend
            }
          }
        };
        this.countEP++;
        this.tasks.push(this.cloneObject(choiceEpic));
      },

      async takeEpic(id) {



        var goExit = false;
        this.tasks.forEach(task => {
          if (task.isEpic) {
            if (Object.keys(task.epic).indexOf("id") > -1) {
              if (id == task.epic.id) {
                goExit = true;
                return false;
              }
            }
          }
        })
        if (goExit) {
          return false;
        } else {
          await this.findEpic(id);
        }

        //this.$emit("update:tasks", this.tasks.slice().reverse());
        //this.task.reverse();
        setTimeout(() => {
          window.location.href = '#lastItem';
        }, 1000);
      },


      async takeUST() {
        //console.log(`WorkSpace recepcionando ${this.idUS}`);
        await this.takeUserStory(this.idUS)
        window.location.href = '#' + this.idUS;
      },
      async getUST(codeType) {
        var types = await this.getTypesOfUserStories();
        var typeUS = {}
        types.forEach(type => {
          if (type.code == codeType) {
            typeUS = type;
          }
        });
        return typeUS;
      },


      async takeUserStory(code) {
        var goExit = false;
        this.tasks.forEach(task => {
          if (task.isUserStory) {
            if (Object.keys(task.userStory).indexOf("us") > -1) {
              if (code == task.userStory.us.id) {
                //this.alertInfo("Esta Historia de Usuario (" + code + ") ya se encuentra abierta.");
                goExit = true;
                return false;
              }
            }
          }
        })
        if (goExit) {
          return false;
        }
        //console.log("buscando " + code);
        var pathname = window.location.pathname;
        var projectName = pathname.split("/")[2];
        var url = `/davinci/${projectName}/userStories/${code}`
        //console.log("buscando " + code + " en =" + url);

        var rsUS = {}

        await this.axios.get(url).then(rs => {
          //console.log(JSON.stringify(rs.data));
          if (rs.data.length > 0) {
            rsUS = rs.data[0]
          }
        });

        var itemType = await this.getUST(rsUS.type);

        //Rescato los datos de la ultima version de la historia de usuario.
        var last = rsUS.versions.length - 1;
        var lastVersionUS = rsUS.versions[last];

        var us = {
          isUserStory: true,
          exist: true,
          userStory: {
            title: rsUS.id,
            usType: itemType,
            us: rsUS,
            fields: lastVersionUS.fields,
            preConditions: lastVersionUS.preConditions,
            steps: lastVersionUS.steps
          }
        };
        this.countUS++;
        this.tasks.push(us);
        //this.$emit("update:tasks", this.tasks.slice().reverse());
        //this.task.reverse();
      },

      addNewUserStyory(itemType) {
        var emptyUS = {
          isUserStory: true,
          exist: true,
          userStory: {
            title: `Nuevo ${itemType.title} #${this.countUS} `,
            usType: itemType,
            fields: {
              description: "",
              expectedResult: ""
            },
            preConditions: itemType.preConditions,
            steps: {}
          }
        };
        this.countUS++;
        this.tasks.push(this.cloneObject(emptyUS));
        //this.$emit("update:tasks", this.tasks.slice().reverse());
        //this.task.reverse();
        setTimeout(() => {
          window.location.href = '#lastItem';
        }, 1000);


      },
      addNewEpic(itemType) {
        var emptyEpic = {
          isEpic: true,
          exist: true,
          epic: {
            name: itemType.name,
            type: itemType.Code,
            definition: itemType.definition,
            title: `Nuevo ${itemType.name} #${this.countEP} `,
            epicForm: {
              reference: itemType.reference,
              attributes: itemType.attributes,
              inventions: itemType.inventions
            }
          }
        };
        this.countEP++;
        this.tasks.push(this.cloneObject(emptyEpic));
        //this.$emit("update:tasks", this.tasks.slice().reverse());
        //this.task.reverse();
        setTimeout(() => {
          window.location.href = '#lastItem';
        }, 1000);

      },
      makeTask() {
        //////console.log("Crea Tarea");
      },
      removeTask(index) {
        //console.log("Index :" + index);
        if ((Object.keys(this.tasks[index]).indexOf("isEpic") > -1) && this.tasks[index].isEpic) {
          this.countEP--;
          this.tasks[index].isEpic = false;
          this.tasks[index].epic = null;
        }
        if ((Object.keys(this.tasks[index]).indexOf("isUserStory") > -1) && this.tasks[index].isUserStory) {
          this.countUS--;
          this.tasks[index].isUserStory = false;
          this.tasks[index].userStory = null;
        }
        //delete this.tasks[index];
        this.tasks[index].exist = false;
      },
    }
  };
</script>
<style></style>
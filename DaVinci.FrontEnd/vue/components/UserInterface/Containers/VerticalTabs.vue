<template>
  <div class="template">
    <div class="row" style="min-width: 100%">
      <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12 ftl-vertical-tab-container">
        <div class="col-lg-3 col-md-3 col-sm-3 col-xs-3 ftl-vertical-tab-menu">
          <div class="list-group">
            <span class="list-group-item text-center push" v-for="(tab, index) of mytabs" :key="index" v-bind:class="{'active': (tab.slots == activeTab.slots) }" v-on:click="activateTab(tab)">
              <span v-if="tab.slots == activeTab.slots" v-html="tab.title"></span>
              <i v-if="!(tab.slots == activeTab.slots)" v-html="tab.title"></i>
            </span>
          </div>
        </div>
        <div class="ftl-vertical-tab-content active" v-if="tabs.length>0">
          <slot :name="activeTab.slots">
          </slot>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  export default {
    name: "v-tabs",
    props: {
      tabs: {},
      activeIndex: Number,
      labelNext: String,
      value: {},
      labelBack: String
    },
    data: function () {
      return {
        currentIndex: 0,
        humanInteraction: false
      };
    },

    watch: {
      activeIndex(newValue, oldValue) {
        if (newValue != oldValue && this.humanInteraction) {
          this.humanInteraction = false;
        }
      }
    },

    computed: {
      activeTab: function () {
        var activeTAB = null;

        if (!this.humanInteraction) {
          //console.log(JSON.stringify(this.tabs));
          //console.log(JSON.stringify(this.activeIndex));
          activeTAB = (this.activeIndex == null) ? this.tabs[0] : this.tabs[this.activeIndex];
        } else {
          activeTAB =
            this.tabs[this.currentIndex] == null ?
              this.tabs[0] :
              this.tabs[this.currentIndex];
        }
        try {
          this.$emit("update:value", activeTAB.id);
        } catch (e) {
          this.$emit("update:value", 0);
        }
        this.$emit("loadTab");


        return activeTAB;
        //return activeTAB==nullthis.tabs[0];
      },
      Next: function () {
        return this.labelNext == null ? "Continuar" : this.labelNext;
      },
      Back: function () {
        return this.labelBack == null ? "Volver" : this.labelBack;
      },
      mytabs: function () {
        var myTabs = [];
        var max = Object.keys(this.tabs).length;
        var i = 1;
        var lastTab = {};
        for (var index in this.tabs) {
          var isLast = i == max;
          this.tabs[index]["last"] = isLast;
          this.tabs[index]["first"] = i == 1;
          this.tabs[index]["index"] = parseInt(index);
          myTabs.push(this.tabs[index]);
          i++;
        }

        return myTabs;
      }
    },
    methods: {
      activateTab: function (tab) {
        this.humanInteraction = true;
        this.currentIndex = tab.index;
      },
      activateTabByIndex: function (index) {
        if (this.currentIndex != index) {
          this.humanInteraction = true;
          this.currentIndex = index;
          this.$emit("changeTab");
        }
      }
    },
    /*
    created: function () {
      //console.log(JSON.stringify(this.mytabs));
      if (this.activeTab === undefined || this.activeTab == null) {
        this.activeTab = this.mytabs[0];
      }
    }
    */
  };
</script>
<style scoped type="scss">
</style>
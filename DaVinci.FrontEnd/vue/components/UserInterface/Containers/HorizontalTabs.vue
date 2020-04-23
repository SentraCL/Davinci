<template>
  <div>
    <ul class="nav nav-tabs">
      <li v-for="(tab, index) of tabs" :key="index">
        <a v-on:click="activateTab(tab)">
          <span class="tab titleTab" v-bind:class="{'active': (tab.id == activeTab.id)}" v-html="tab.title"> </span>
        </a>
      </li>
    </ul>

    <div class="tabContent" v-if="!viewScoped">
      <slot :name="activeTab.slot">
      </slot>
      <div class="clearfix"></div>
    </div>
    <div v-if="viewScoped" class="tabContent" v-for="(tab, index) of tabs" :key="index" v-bind:class="{'show': (tab.slot == activeTab.slot), 'hide': (tab.slot != activeTab.slot)}">
      <slot :name="tab.slot">
        <!-- activeTab.slots -->
      </slot>
      <div class="clearfix"></div>
    </div>
    <div class="clearfix"></div>
  </div>
</template>
<script>
  export default {
    name: "h-tabs",
    props: {
      tabs: {},
      viewScoped: Boolean
    },
    data: function () {
      return {
        activeTab: null
      }
    },
    watch: {
      tabs: function () {
        this.$emit("loadTab");
        this.activeTab = this.tabs[0];
      }

    },
    methods: {
      activateTab: function (tab) {
        this.activeTab = tab;
        this.$emit("activeTab");
        this.$emit("update");
      }
    },
    created: function () {
      this.$emit("loadTab");
      this.activeTab = this.tabs[0];
    }
  };
</script>

<style scoped>
  .nav {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .nav-tabs {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .titleTab {
    border-radius: 12px 12px 0px 0px;
    min-width: 100px;
  }

  .active {
    font-weight: bolder;
    font-size: 20px;
    position: relative;
    background-color: rgba(255, 255, 255, 0.9) !important;
    -webkit-box-shadow: 7px 3px 3px -4px rgba(0, 0, 0, 0.64);
    box-shadow: 7px 3px 3px -4px rgba(0, 0, 0, 0.64);
  }

  .tab {
    z-index: -10px;
    background-color: rgba(255, 255, 255, 0.3);
    display: inline-block;
    cursor: pointer;
    padding: 12px 12px;
    margin-bottom: 0;
    font-size: 16px;





    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  .tabContent {

    width: 100%;
    padding: 8px 10px;
    margin-bottom: 0;
    border: 0px;
    background-color: rgba(255, 255, 255, 0.9)
  }

  .show {
    display: block;
  }

  .hide {
    display: none;
  }
</style>
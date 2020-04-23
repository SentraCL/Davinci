<template>
  <span>
    <div class="input-group">
      <span class="input-group-addon">
        <strong v-bind:class="{'lock': inactive }">{{label}}</strong>
      </span>
      <div class="onoffswitch">
        <input v-if="!inactive" type="checkbox" name="onoffswitch" class="onoffswitch-checkbox" :id="id" v-model="optionSN" @click="change" >
        <input v-if="inactive" type="checkbox" name="onoffswitch" class="onoffswitch-checkbox" :id="id" v-model="optionSN" disabled>
        <label class="onoffswitch-label" :for="id">
          <span class="onoffswitch-inner"></span>
        </label>
      </div>
    </div>

  </span>
</template>
<script>
  export default {
    name: "switch-check",
    inheritAttrs: false,
    props: {
      label: String,
      inactive: Boolean,
      value: {}
    },
    data: function () {
      return {
        optionSN: this.value == 1
      };
    },
    computed: {
      id: function () {
        return this.label.replace(/\s+/g, "").replace(/[^a-zA-Z ]/g, "");
      }
    },
    mounted: function () {},
    created: function () {
      this.optionSN = this.value == 1;
    },
    updated: function () {
      this.optionSN = this.value == 1;
    },
    methods: {
      change() {
        this.optionSN = this.optionSN == 1 ? 0 : 1;
        //console.log("SN => " + this.optionSN);
        this.$emit("update:value", this.optionSN);
        this.$emit("click");
        this.$emit("change");
        this.$emit("update");
      }
    }
  };
</script>
<style scoped>
  .onoffswitch {
    position: relative;
    padding: 4px 5px;
    background-color: #fff;
    width: 60px;
    border: 1px solid #ccc;
    border-radius: 0px 4px 4px 0px;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
  }

  .onoffswitch-checkbox {
    display: none;
  }

  .onoffswitch-label {
    display: block;
    top: 2px;
    position: relative;
    overflow: hidden;
    cursor: pointer;
    border: 2px solid #999999;
    border-radius: 22px;
  }

  .onoffswitch-inner {
    display: block;
    width: 200%;
    margin-left: -100%;
    transition: margin 0.3s ease-in 0s;
  }

  .onoffswitch-inner:before,
  .onoffswitch-inner:after {
    display: block;
    float: left;
    width: 50%;
    height: 17px;
    padding: 0;
    line-height: 17px;
    font-size: 11px;
    color: white;
    font-family: Trebuchet, Arial, sans-serif;
    font-weight: bold;
    box-sizing: border-box;
  }

  .onoffswitch-inner:before {
    content: "Si";
    padding-left: 12px;
    background-color: #0766a5;
    color: #ffffff;
  }

  .onoffswitch-inner:after {
    content: "No";
    padding-right: 12px;
    background-color: #eeeeee;
    color: #999999;
    text-align: right;
  }

  .onoffswitch-switch {
    display: block;
    width: 13px;
    margin: 2px;
    background: #ffffff;
    position: absolute;
    top: 1;
    bottom: 0;
    right: 44px;
    border: 2px solid #999999;
    border-radius: 22px;
    transition: all 0.3s ease-in 0s;
  }

  .onoffswitch-checkbox:checked+.onoffswitch-label .onoffswitch-inner {
    margin-left: 0;
  }

  .onoffswitch-checkbox:checked+.onoffswitch-label .onoffswitch-switch {
    right: 0px;
  }
</style>
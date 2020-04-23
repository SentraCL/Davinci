<template>

  <component :is="tag" :type="nativeType" @click="clicked()" :disabled="disabled || loading" class="btn" :class="[
      { 'btn-round': round },
      { 'btn-block': block },
      { 'btn-sm': xs },
      { 'btn-just-icon': icon },
      { [`btn-${type}`]: type && to==value },
      { [`btn-outline-${type}`]: type && to!=value },
      { [`btn-${size}`]: size },
      { 'btn-link': simple }
    ]">
    <slot></slot>

    <input type="radio" :name="name" :value="value" style="display: none">

  </component>
</template>
<script>
  export default {
    name: "d-radio",

    data() {
      var init = this.value;
      return {
        isChecked: false
      };
    },


    methods: {
      clicked() {
        for (var i in document.getElementsByName(this.name)) {
          if (document.getElementsByName(this.name)[i].value == this.value) {
            document.getElementsByName(this.name)[i].click()
            this.$emit("click");
            this.isChecked = !this.isChecked;
          }
        }
      }
    },
    created() {

      this.isChecked = true;
    },
    props: {

      name: String,
      value: [String, Number],
      to: [String, Number],
      tag: {
        type: String,
        default: "button"
      },
      round: Boolean,
      icon: Boolean,
      outline: Boolean,
      block: Boolean,
      xs: Boolean,

      loading: Boolean,
      disabled: Boolean,
      type: {
        type: String,
        default: "default"
      },
      nativeType: {
        type: String,
        default: "button"
      },
      size: {
        type: String,
        default: ""
      },
      simple: Boolean
    }
  };
</script>
<style></style>
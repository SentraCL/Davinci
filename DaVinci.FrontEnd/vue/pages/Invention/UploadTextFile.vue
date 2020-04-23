<template>
  <label class="btn btn-success upload-text">
    {{label}}
    <input type="file" @change="loadTextFromFile" :accept="accept">
  </label>
</template>

<script>
  export default {
    name: "upload-text",
    props: {
      label: String,
      accept: String
    },
    methods: {
      loadTextFromFile(ev) {
        const file = ev.target.files[0];
        const reader = new FileReader();

        reader.onload = e => this.$emit("upload", e.target.result);
        reader.readAsText(file);
      }
    }
  };
</script>

<style>
  .upload-text {
    position: relative;
    overflow: hidden;
    display: inline-block;
    color: #FFF !important;
    /* Fancy button looking */
    cursor: pointer;
  }

  .upload-text input {
    color: #FFF !important;
    position: absolute;
    top: 0;
    left: 0;
    z-index: -1;
    opacity: 0;
  }
</style>
<template>
  <div class="card">
    <div class="card-image" v-if="$slots.image">
      <slot name="image"></slot>
    </div>
    <div class="card-header" v-if="$slots.header || title">
      <slot name="header">
        <h2 class="card-title" v-html="title"></h2>
        <p class="card-category" v-if="subTitle">{{ subTitle }}</p>
      </slot>
    </div>
    <div class="card-body custom-card" v-if="$slots.default">
      <slot></slot>
    </div>
    <slot name="raw-content"></slot>
    <div class="card-footer" v-if="$slots.footer">
      <hr />
      <slot name="footer"></slot>
    </div>
  </div>
</template>
<script>
  export default {
    name: "card",
    props: {
      title: String,
      subTitle: String
    }
  };
</script>
<style>
.custom-card{
    padding: 15px 15px 20px 15px !important;
}
</style>
<template>
  <component :is="tag" :type="nativeType" :disabled="disabled || loading" class="btn" :class="[
      { 'btn-round': round },
      { 'btn-block': block },
      { 'btn-sm': xs },
      { 'btn-just-icon': icon },
      { [`btn-${type}`]: type && !outline },
      { [`btn-outline-${type}`]: type && outline },
      { [`btn-${size}`]: size },
      { 'btn-link': simple }
    ]">
    <slot name="loading">
      <i v-if="loading" class="fa fa-spinner fa-spin"></i>
    </slot>
    <slot></slot>
  </component>
</template>
<script>
  export default {
    name: "d-button",
    props: {
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
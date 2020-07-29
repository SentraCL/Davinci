<template>
    <span :class="inputClass">
        <!--{{inputClass}} -->
        <input :maxlength="maxlength?maxlength:255" class="input_field input_field--davinci" v-if="!isInactive" :autocomplete="autocomplete" v-model="textbox" readonly="readonly" onfocus="javascript: this.removeAttribute('readonly')" @input="$emit('input', $event.target.value)" v-bind="$attrs" :type="type" @blur="save" />
        <span class="input_field input_field--davinci" v-if="isInactive" :class="isInactive?'disable':''">{{textbox}}</span>
        <label class="input_label input_label--davinci">
            <span class="input_label-content input_label-content--davinci" :data-content="label" v-html="label"></span>
        </label>
    </span>

</template>
<script>
    export default {
        inheritAttrs: false,
        name: "input-text",
        data() {
            var init = this.value;
            return {
                textbox: init,
                newValue: init
            };
        },
        watch: {
            textbox(newText, oldText) {
                this.textbox = this.format(newText)
                this.$emit("update:value", this.textbox);
                this.$emit("update");
            }
        },
        props: {
            label: String,
            maxlength: Number,
            type: String,
            alphabetic: Boolean,
            alphanumeric: Boolean,
            numeric: Boolean,
            removeSpace: Boolean,
            capitalize: Boolean,
            autocomplete: String,
            isInactive: Boolean,
            value: [String, Number]
        },

        computed: {
            typeDef: function () {
                return (this.type) ? this.type : "text";
            },

            inputClass: function () {
                var classInput = "input input--davinci";
                if (this.type == "datetime-local" || this.type == "date") {
                    classInput = "input input--davinci input--filled";
                } else {
                    if (this.textbox || this.value) {
                        classInput = "input input--davinci input--filled";
                    }
                }
                return classInput;
            }
        },
        created() {



        },
        updated: function () {
            this.$nextTick(function () {
                if (this.value != this.newValue) {
                    this.newValue = this.value;
                    this.textbox = this.value;
                }
            });
        },
        methods: {
            format: function (text) {
                if(text!==undefined){
                    if (this.capitalize) {
                        text = text.charAt(0).toUpperCase() + text.slice(1);
                    }
                    if (this.alphabetic) {
                        text = text.replace(/[^a-z-A-Z]/gi, '')
                    }
                    if (this.alphanumeric) {
                        text = text.replace(/[^a-z0-9]/gi, '')
                    }
                    if (this.numeric) {
                        text = text.replace(/[^0-9a-z-A-Z]/g, '')
                    }
                    if (this.removeSpace) {
                        text = text.replace(/\s/g, '')
                    }
                }
                return text;
            },
            /*
            keymonitor: function (event) {
                this.textbox = this.format(this.textbox)
                this.$emit("update:value", this.textbox);
                this.$emit("update");
            },
            */
            save: function () {
                this.textbox = this.format(this.textbox)
                this.$emit("update:value", this.textbox);
                this.$emit("blur");
                this.$emit("update");
            }

        }
    };
</script>
<style>
.disable{
    height: 27px;
}
</style>
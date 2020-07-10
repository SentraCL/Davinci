<template>
    <span :class="inputClass">
        <select :disabled="inactive" class="input_field input_field--davinci" :id="idValue" v-model="option" :title="placeholder" @blur="change" @change="change" style="cursor:pointer">
            <option class=".ui_option" v-for="item in list" :key="item[keyValue]" v-bind:value="item[keyValue]" :selected="item[keyValue] == value">{{showKey?item[keyValue] + ' : ':''}}{{item[keyLabel]}}</option>
        </select>
        <label class="input_label input_label--davinci">
            <span class="input_label-content input_label-content--davinci" :data-content="label">{{label}}</span>
        </label>

        <div class="input-group find-box">
        <div class="dropdown" v-if="list.length>0">
            <button class="dropbtn"><span style="float:left; position:relative">{{key}}</span> <span v-if="list.length>0" style="float:right; position:relative" class="ti-angle-double-down"></span></button>
            <div class="dropdown-content">
                <p class="selectItem" v-if="item!=key" v-for="item in list" @click="setContext(item)" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">{{item}}</p>
            </div>
        </div>
        <input tabindex="1" ref="search" :id="listID+ 'SS'" v-model="value" class="text-find" :list="listID" autocomplete="off" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">
        <datalist class="ui-autocomplete fixed-height" :id="listID">
            <option v-for="item in list" :key="item[keyValue]" v-bind:value="item[keyValue]" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">
                {{showKey?item[keyValue] + ' : ':''}}{{item[keyLabel]}}
            </option>
        </datalist>
    </div>
    </span>
    
</template>
<script>
    export default {
        name: "combo-data",
        inheritAttrs: false,
        props: {
            label: String,
            placeholder: String,
            value: {},
            showKey: Boolean,
            inactive: Boolean,
            keyLabel: String,
            keyValue: String,
            idValue:String,
            list: {}
        },
        watch: {
            value(n, o) {
                this.option = n;
            }
        },
        data: function () {
            return {
                title: "",
                option: this.value,
                initValue: this.value
            };
        },
        computed: {
            idOption: function () {
                return "#" + this.id;
            },
            defaultValue: function () {
                return this.inactive ? this.getTitle(this.value) : "";
            },
            inputClass: function () {
                var classInput = "input input--davinci";

                if (this.getTitle(this.option) != "") {
                    classInput = "input input--davinci input--filled";
                }
                /*
                if (this.option != this.initValue && this.initValue != null) {
                    classInput = "input input--davinci input--filled";
                }
                */
                return classInput;
            }
        },
        updated: function () {
            this.option = this.value;
        },
        methods: {
            change() {
                this.$emit("update:value", this.option);
                this.$emit("update");
                this.$emit("change");
                this.title = this.getTitle(this.option);
                //this.$refs.search.focus();
            },

            getTitle(key) {
                for (var index in this.list) {
                    if (this.list[index][this.keyValue] === key) {
                        return this.list[index][this.keyLabel];
                    }
                }
                return "";
            }
        }
    };
</script>
<style scoped>

</style>
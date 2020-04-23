<template>
    <div class="input-group find-box">
        <div class="dropdown" v-if="items.length>0">
            <button class="dropbtn"><span style="float:left; position:relative">{{key}}</span> <span v-if="items.length>0" style="float:right; position:relative" class="ti-angle-double-down"></span></button>
            <div class="dropdown-content">
                <p class="selectItem" v-if="item!=key" v-for="item in items" @click="setContext(item)" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">{{item}}</p>
            </div>
        </div>
        <input tabindex="1" ref="search" :id="listID+ 'SS'" v-model="value" class="text-find" :list="listID" autocomplete="off" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">
        <datalist class="ui-autocomplete fixed-height" :id="listID">
            <option v-for="entry,index in context[key]" :key="index" :value="entry.value?entry.value:entry" v-on:keyup="setValue" v-on:click="setValue" @focus="setValue" @blur="setValue">
                {{entry.label?entry.label:entry}}
            </option>
        </datalist>
    </div>

</template>
<script>
    export default {
        name: "find-box",
        inheritAttrs: false,

        data: function () {
            var _default = Object.keys(this.context)[0];
            if (this.default != null) {
                _default = this.default;
            }
            return {
                key: _default,
                value: "",
                localTarget: {}
            }
        },

        props: {
            target: {},
            context: {},
            default: String
        },

        computed: {
            items() {
                var items = Object.keys(this.context)
                return items;
            },

            listID: function () {
                return this.makeId(16) + "_list";
            }
        },
        created() {
            console.log(JSON.stringify(this.value));
            this.localTarget[this.key] = this.value;
            this.$emit("update:target", this.localTarget);
        },
        methods: {
            setContext(keyContext) {
                this.value = "";
                this.key = keyContext;
            },
            setValue() {
                this.localTarget = {};
                /*
                var jsonValue = this.value == "" ? {} : JSON.parse(this.value);
                if (Object.keys(jsonValue).indexOf("value") > -1) {
                    this.localTarget[this.key] = this.value.value;
                } else {
                    */
                this.localTarget[this.key] = this.value;
                //}
                this.$emit("update:target", this.localTarget)
            }
        }
    };
</script>
<style scoped>

</style>
<template>

    <div>
        <form @submit.prevent>
            <h4 v-if="title" v-html="title"></h4>
            <div v-for="input in inputs" class="input" :key="input.name">
                <div class="form-group" v-if="input.isMultitext">
                    <label>{{input.name}}</label>
                    <textarea rows="5" class="form-control border-input" v-model="form[input.name]">
                    </textarea>
                </div>

                <input-text :isInactive="isDisabled" v-if="input.isText" :label="input.name" v-model="form[input.name]"></input-text>
                <input-text :isInactive="isDisabled" v-if="input.isNumeric" type="number" :label="input.name" v-model="form[input.name]" numeric></input-text>
                <input-text :isInactive="isDisabled" v-if="input.isDate" type="datetime-local" :label="input.name" v-model="form[input.name]"></input-text>
                <!-- <combo-simple v-if="input.isBool" :label="input.name" :list="input.list" keyValue="id" keyLabel="name" :value.sync="form[input.name]"></combo-simple>-->
                <switch-box :isInactive="isDisabled" v-if="input.isBool" :label="input.name" :value.sync="form[input.name]"></switch-box>
                <combo-simple :inactive="isDisabled" v-if="input.isJson" :label="input.name" :list="input.list" keyValue="id" keyLabel="name" :value.sync="form[input.name]"></combo-simple>
            </div>
        </form>

    </div>

</template>
<script>
    /*
<input type="button">
<input type="checkbox">
<input type="color">
<input type="date">
<input type="datetime-local">
<input type="email">
<input type="file">
<input type="hidden">
<input type="image">
<input type="month">
<input type="number">
<input type="password">
<input type="radio">
<input type="range">
<input type="reset">
<input type="search">
<input type="submit">
<input type="tel">
<input type="text">
<input type="time">
<input type="url">
<input type="week">
    */
    export default {
        name: "artifact-form",
        computed: {

        },
        props: {
            inputs: {},
            values: {},
            title: String,
            isDisabled:false
        },
        data() {
            var _form = {}
            for (var i in this.inputs) {
                var input = this.inputs[i]
                var conditionalKey = "is" + input.typeName.charAt(0).toUpperCase() + input.typeName.slice(1);
                input[conditionalKey] = true

                if (input.isJson) {
                    input.isEssential = false
                    input['list'] = [
                        { id: 0, name: input.nickName + " 1" },
                        { id: 1, name: input.nickName + " 2" },
                        { id: 2, name: input.nickName + " 3" }
                    ]
                } else {
                    input.isEssential = true
                    if (input.isBool) {
                        input['list'] = [
                            { id: 0, name: "No" }, { id: 1, name: "Si" }
                        ]
                        input['value'] = 0
                    } else {
                        input['value'] = ""
                    }
                }
                _form[input.name] = ""
                this.$emit("update:values", _form);
            }
            return {
                form: _form
            }
        },
        watch: {
            async form(newForm, oldValue) {
                //console.log(JSON.stringify(newForm));
                this.$emit("update:values", newForm);
                this.$emit("update");
            }
        },
        methods: {

        }
    };
</script>
<style></style>
<template>
    <span :class="inputClass">
        <div class="labeltags">
            <ul class="labelul">
                <li class="labelli" v-for="item in list" :key="item[keyValue]" v-bind:value="item[keyValue]"><label :class="item.Disable ? 'disable' : ''" v-on:click="activate(item[keyValue])" :id="item[keyValue]" class="label btn btn-xs btn-round btn-success">{{item[keyLabel]}}</label></li>
            </ul>
        </div>
    </span>
</template>
<script>
    export default {
        name: "label-list",
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
                initValue: this.value,
                size:0
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
            activate(key){
                let listkey=[]
                if(this.size==0||!this.size){
                    this.size=this.list.length
                }
                for(let i=0;i<this.list.length;i++){
                    if(this.list[i].EnterpriseId==key&&this.list[i].Disable){
                        this.list[i].Disable=false;
                        this.size++;
                    }else 
                    if(this.list[i].EnterpriseId==key&&this.size>1){
                        this.list[i].Disable=true;
                        this.size--;
                    }
                    if(this.list[i].Disable){
                        listkey.push(this.list[i].EnterpriseId);
                    }
                }
                this.$emit('clicked', listkey)
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
    .labelli{
        padding: 4px 0;
        margin: 0 4px 4px 0;
        display: inline-block;
        color: #ffffff;
    }
    .labelul{
        list-style-type: none;
        padding: 0;
        margin: 0;
        display: inline-block;
    }
    .labeltags{
        vertical-align: top;
        padding: 0;
    }
    .label{    
        top:-6px!important;
    }
    .disable{
        background-color: #6c757d;
        border-color: #5f656a;
        border-right: 2px solid #4E1D00 !important;
        border-bottom: 2px solid #161b20 !important;
    }
</style>
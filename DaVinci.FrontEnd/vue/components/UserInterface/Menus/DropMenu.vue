<template>

    <div class="menudown">
       
        <span class="titleOption" v-html="title"></span></span>
        <div class="menudown-content card">
            <span v-for="menu,index in options" @click="setCurrent(index)" v-html="menu.html?menu.html:menu"></span>
        </div>
    </div>

</template>
<script>
    export default {
        name: "drop-menu",
        props: {
            option: {},
            title: String,
            options: {}
        },
        watch: {
            option(n, o) {
                this.$emit("update:option", n);
            }
        },
        methods: {
            setCurrent(index) {
                this.$emit("update:option", this.options[index].option);
                this.$emit("change");
            }
        }
    };
</script>
<style scoped>
    .titleOption {
        font-size: 16px;
        border: none;
        min-width: 120px;
        left: 0px;
        top: 5px;
        position: relative;
    }

    .menudown {
        position: relative;
        display: inline-block;
        -webkit-user-select: none;
        /* Safari 3.1+ */
        -moz-user-select: none;
        /* Firefox 2+ */
        -ms-user-select: none;
        /* IE 10+ */
        user-select: none;
        /* Standard syntax */
    }

    .menudown-content {
        z-index: 9000;
        display: none;
        position: absolute;
        top: 20px;
        background-color: #FFF;
        width: 250px;
        
        box-shadow: 0px 2px 4px 0px rgba(29, 48, 0, 0.8);
        z-index: 1;
        -webkit-user-select: none;
        /* Safari 3.1+ */
        -moz-user-select: none;
        /* Firefox 2+ */
        -ms-user-select: none;
        /* IE 10+ */
        user-select: none;
        /* Standard syntax */
    }

    .menudown-content span {
        z-index: 9000;
        color: black;
        
        padding: 12px 16px;
        text-decoration: none;
        display: block;
    }

    .menudown-content span:hover {
        z-index: 9000;
        background-color: azure;
        font-weight: bold;
    }

    .menudown:hover .menudown-content {
        z-index: 9000;
        display: block;
        cursor: pointer
    }

    .menudown:hover .titleOption {
        cursor: pointer
    }
</style>
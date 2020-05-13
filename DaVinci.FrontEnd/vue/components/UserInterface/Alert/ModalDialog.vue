<template>

    <span v-if="show">
        <!-- The Modal -->
        <div :id="id" class="fixedmodal" >

            <!-- Modal content -->
            <div class="fixedmodal-content" v-if="show">

                <span > 
                    <span>
                        <div class="md-title">
                            <span class="closemodal" v-if="doClose" @click="toClose()" >&times;</span>
                            <h4 v-html="title"></h4>
                        </div>
                    </span>
                    <slot name="dialog">
                    </slot>
                    <slot name="actions">
                    </slot>
                </span>




            </div>

        </div>
    </span>
</template>
<script>
    export default {
        name: "default",
        inheritAttrs: false,
        props: {
            id: String,
            text: String,
            time: Number,
            title: String,
            isClose: Boolean,
            show: Boolean
        },
        data() {
            var _isclose = this.isClose;
            return {
                doClose: _isclose
            };
        },
        computed: {},
        mounted: function () {
            if (this.show) {
                var fixedmodal = document.getElementById(this.id);
                fixedmodal.style.display = "block";
                if (this.time != null && this.time > 0) {
                    setTimeout(() => {
                        fixedmodal.style.display = "none";
                        this.$emit("update:time", 0);
                        this.$emit("update:show", false);
                        this.$emit("off");
                    }, this.time * 1000);
                }
            }
        },
        created: function () { },
        updated: function () {
            /*
            //console.log("time " + this.time);
            //console.log("closemodal " + this.closemodal);
            */
            var fixedmodal = document.getElementById(this.id);
            if (fixedmodal != null) {
                //fixedmodal.style.display = "block";
                this.screen("on");
            }
            if (this.time > 0) {
                if (fixedmodal == null) {
                    fixedmodal = document.getElementById(this.id);
                    this.screen("on");
                    //fixedmodal.style.display = "block";
                }
                setTimeout(() => {
                    fixedmodal.style.display = "none";
                    this.$emit("update:time", 0);
                    this.$emit("update:show", false);
                    this.$emit("off");
                    this.$emit("onClose");
                }, this.time * 1000);
            }
        },
        methods: {
            toClose: function () {
                var fixedmodal = document.getElementById(this.id);
                if (fixedmodal){
                    fixedmodal.style.display = "none";
                    this.$emit("update:time", 0);
                    this.$emit("update:show", false);
                    this.$emit("off");
                    this.$emit("onClose");
                }
            }
        }
    };
</script>

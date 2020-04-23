<template>
    <div>
        <m-dialog id="usDialog" :show.sync="maximize">
            <div slot="dialog">
                <slot name="content"></slot>
            </div>
        </m-dialog>

        <card style="background-color:khaki;" v-if="!maximize">

            <div class="card-title ">
                <h4 v-html="title"> </h4>
                <span @click="switchWindows()" class="fa-pull-right maximize-button"><i class="ti-new-window"></i></span>
            </div>
            <hr />
            <div class="row">
                <div class="col-md-12">
                    <slot name="resume"></slot>
                </div>
            </div>
        </card>
    </div>
</template>
<script>
    export default {
        name: "artifact",
        computed: {

        },
        props: {
            title: String
        },

        data() {
            return {
                maximize: false
            }
        },

        methods: {
            switchWindows() {
                this.maximize = !this.maximize;
            }
        }
    };
</script>
<style>
    .maximize-button {
        position: relative;
        top: -45px;
        cursor: pointer;
    }
</style>
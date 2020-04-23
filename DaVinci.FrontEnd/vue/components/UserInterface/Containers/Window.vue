<template>
    <form @submit.prevent v-if="showForm">
        <card style="min-height: 600px;background: var(--transparent);">
            <div class="topnav">
                <div>
                    <slot :name="buttons">
                    </slot>
                </div>
            </div>
            <div>
                <slot :name="content">
                </slot>
            </div>
        </card>
    </form>

</template>

<script>
    export default {
        name: "artifact",
        computed: {

        },
        props: {
            artifactForm: {}
        },

        data() {
            return {
                defaultArtifac: ["Texto", "Numerico", "Fecha y Hora", "Si/No"],
            }
        },

        methods: {
            maximizar: function (id) {
                this.maximize = true;

                $("#" + id).hide();
                $("#" + id)
                    .css("position", "fixed")
                    .css("top", "0px")
                    .css("height", "100%")
                    .css("width", "100%")
                    .css("z-index", "100001")
                    .css("overflow-y", "auto")
                    .css("overflow-x", "auto")
                    .css("left", "1px");
                $("#" + id).show("scale");
            },
            minimizar: function (id) {
                this.maximize = false;

                $("#" + id).hide();
                $("#" + id)
                    .css("position", "")
                    .css("top", "")
                    .css("height", "")
                    .css("z-index", "0")
                    .css("left", "");
                $("#" + id).show("puff");
            },
        }
    };
</script>
<style></style>
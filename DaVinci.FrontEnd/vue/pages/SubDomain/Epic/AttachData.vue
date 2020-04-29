<template>

    <div >
        <i>La data asociada debe tener los siguientes atributos:</i>
        <br/>
        
    </div>

</template>
<script>
    export default {
        name: "attach-data",
        computed: {

        },
        props: {
            epic: {},
            activity: {}
        },
        data() {
            return {
                defaultArtifac: ["Texto", "Numerico", "Fecha y Hora", "Si/No"],
            }
        },

        methods: {

        }
    };
</script>
<style></style>
<template>

    <div class="pagination" v-if="totalPages>1">
        <p :class="{ 'active': currentPage==p }" @click="clickPage(p)" v-for="p in totalPages">{{p}}</p>
    </div>

</template>
<script>
    export default {
        name: "pagination",


        props: {
            currentPage: Number,
            totalItems: Number,
            itemsPerPage: Number
        },

        computed: {
            totalPages() {
                var total = (this.totalItems / this.itemsPerPage) + (this.totalItems % this.itemsPerPage > 0 ? 1 : 0);


                var _totalPages = []
                for (var i = 1; i < total; i++) {
                    _totalPages.push(i)
                }
                return _totalPages
            }
        },

        data() {

            return {
            }
        },
        methods: {
            clickPage(p) {
                this.$emit("update:currentPage", p);
                this.$emit("clickPage");
            }
        }
    };
</script>
<style>

</style>
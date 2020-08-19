<template>
    <div class="row" v-if="enterprises">
        <div class="col-md-3" v-for="(enterprise,index) in enterprises" :key="index">
            <div class="drop">
                <div class="card card-user">
                    <div class="image">
                        <img src="@/assets/img/toolbar.jpg" alt="...">
                    </div>
                    <div class="content" :class="enterprise.Status==0?'isdisable':''">
                        <div class="author">
                            
                            <a href="#">
                                <img v-if="enterprise.Avatar" class="avatar border-white"  @click="avatarClick(enterprise)" :src="enterprise.Avatar" alt="...">
                                <img v-if="!enterprise.Avatar" class="avatar border-white"  @click="avatarClick(enterprise)" src="@/assets/img/davinci-logo.png" alt="...">
                            </a>
                            <h4 class="title" v-if="enterprise.Name">{{enterprise.Name}}<br></h4>
                        </div>
                        <div class="col-lg-12">
                            <p :tittle="enterprise.Description" class="prettySummary"> {{enterprise.Description}}</p>
                        </div>
                    </div>
                    <div class="text-center">

                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
export default {
        name: "enterprise-item",
        props:{
            enterprises:{},
            filteredEnterprise:""
        },
        data(){
            return{
                }
        },
        
        methods:{
            avatarClick(enterprise){
                this.$emit("selectedEnterprise",enterprise);
            },
            
        }
}
</script>
<style scoped>
.isdisable{
    background-color: #0000005e;
}
</style>
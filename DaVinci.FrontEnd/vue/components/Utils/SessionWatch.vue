<template>
    <span>

    </span>
</template>
<script>
    import io from 'socket.io-client'

    export default {
        name: "session-watch",
        data() {
            return {
                online: false,
                isSubDomain:false,
                socket: io(`${window.location.origin}`, {
                    forceNew: true
                }),
            }
        },
        watch: {
            online: function () {

            }
        },
        created: async function () {
                
                var pathname = window.location.pathname;
                let path=pathname.split("/");
                this.isSubDomain =path[1]=="davinci"?true:false;

                if(this.isSubDomain){
                    console.log("sub")
                var loginHash = sessionStorage.getItem("subHash")
                var projectName = path[2];
                var varCodeProject = projectName + "_CODE";
                var code=sessionStorage.getItem(varCodeProject);
                if (code == null) {
                    console.log("entro")
                    await this.axios.get(`/davinci/${projectName}/code`).then(rs => {
                    code = rs.data;
                    console.log("respuesta",code)
                    });
                }else{
                    code=code.replace('"','').replace('"','')
                }

                if (loginHash == null) {
                console.log("return")
                    this.$router.push('login');
                } else {
                    this.axios
                    .post(`/davinci/token/${code}`, {
                        "subHash": loginHash
                    }).then(rs => {
                        console.log("axios",rs)
                        if (rs.data) {
                            sessionStorage.setItem("subHash", rs.config.data.split('"')[4]);
                            console.log(sessionStorage.getItem("subHash"))
                            this.$router.push('Main');
                        } else {
                            this.$router.push('Main');
                            //descomenta
                            //this.$router.push('login');
                            sessionStorage.setItem("subHash", null);
                        }
                    });
                }

            }else{
            var loginHash = sessionStorage.getItem("loginHash")
            console.log("loginHash",loginHash)
                if (loginHash == null) {
                    console.log("not return")
                    this.$router.push('login');
                } else {
                    this.axios
                    .post("/api/token/", {
                        "loginHash": loginHash
                    }).then(rs => {
                        if (rs.data) {
                            this.$router.push('dashboard');
                        } else {
                            this.$router.push('login');
                            sessionStorage.setItem("loginHash", null);
                        }
                    });
                }
            }
        }
    }
</script>
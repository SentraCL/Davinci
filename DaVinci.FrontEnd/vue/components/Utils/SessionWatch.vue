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
            var loginHash = sessionStorage.getItem("loginHash")
            if (loginHash == null) {
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
</script>
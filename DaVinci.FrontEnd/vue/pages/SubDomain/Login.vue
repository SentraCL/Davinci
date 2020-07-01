<template>

    <div>
        <div class="row login" v-if="!isOnline">
            <div class="col-xl-4 col-lg-4 col-md-4">
            </div>
            <div class="col-xl-4 col-lg-4 col-md-4">
                <card class="card-user sub-background" :style="display">
                    <div slot="image" class="loginpaper">
                        <img src="@/assets/img/project-logo.png" style="height:100%;width:100%" />
                    </div>
                    <div>
                        <div class="author">
                            <img class="avatar border-white" ref="img" :src="urlLogo" />
                            <h4 class="title sub-text">
                                Comenzar
                                <br />
                            </h4>
                            <transition name="fade">
                                <small v-if="dijo" class="frase">
                                    <span v-html="frase"></span>
                                </small>
                            </transition>
                        </div>
                        <br />
                        <div class="panel">
                            <div class="col-md-12">
                                <input-text :label="title.alias" removeSpace @keyup.enter.native="doLogin" name="user" v-model="login.alias">
                                </input-text>
                            </div>
                            <div class="col-md-12">
                                <input-text :label="title.pass" name="password" v-model="login.password" @keyup.enter.native="doLogin" type="password" autocomplete="off">
                                </input-text>
                            </div>
                            <div class="col-md-12 ">
                                <button type="button" @click="doLogin" style="width: 100%" class="btn sub-button">
                                    Iniciar
                                </button>
                            </div>
                        </div>
                    </div>
                    <hr />
                    <div class="text-center">
                        <small><strong>DaVinci</strong> Diseñador Agíl. &copy; Creado por <a href="http://www.sentra.cl/" target="_blank">&nbsp;<strong>Sentra.</strong></a></small>
                    </div>
                </card>
            </div>
            <div class="col-xl-4 col-lg-4 col-md-4">
            </div>
        </div>
    </div>
</template>
<script>
    //https://codesandbox.io/s/vue-fixed-header-demo-jcm8p
    import io from 'socket.io-client'

    export default {
        data() {
            return {
                load: false,
                isOnline: false,
                urlLogo: "",
                display: "display:none",
                projectName: "",
                frase: "",
                dijo: false,
                title: {
                    alias: `<span class="ti-user" ></span> Usuario`,
                    pass: `<span class="ti-key" ></span> Contraseña`
                },
                socket: io(`${window.location.origin}`, {
                    forceNew: true
                }),
                login: {
                    alias: "",
                    password: ""
                }
            };
        },

        async mounted() {
            var me = this
            var pathname = window.location.pathname;
            this.projectName = pathname.split("/")[2]
            await this.getContextCSS(this.projectName);
            var isOnline = `/davinci/${this.projectName}/isOnline`;
            this.urlLogo = `/api/project/avatar/${this.projectName}.png`;
            await this.axios.get(isOnline).then(rs => {
                var status = rs.data.replace(/\n|\r/g, "");
                if (status == "OK") {
                    this.$router.push('home');
                }
                this.display = "display:block";
            });

            this.socket.emit("davinciSay");
            this.socket.on("DiloTuyo", function (frase) {
                me.dijo = (me.frase != frase);
                if (me.dijo) {
                    me.frase = frase;
                }
                setTimeout(() => {
                    me.socket.emit("davinciSay");
                }, 5000);
            })

        },



        methods: {

            doLogin: function () {
                var me = this
                this.axios
                    .post(`/davinci/${this.projectName}/logIn`, {
                        "user": this.login.alias,
                        "pass": this.login.password
                    }).then(rs => {

                        if (rs.data.Online) {
                            this.isOnline = true;
                            this.$router.push('main');
                        } else {
                            me.login = {
                                alias: "",
                                password: ""
                            }
                            me.dijo = true
                            me.frase = "<strong>Credenciales no válidas..</strong>"
                            setTimeout(() => {
                                me.frase = ""
                            }, 5000);
                        }
                    });

            }
        }
    };
</script>
<style>
    .login {
        position: absolute;
        width: 100%;
        top: 40px;
    }

    .fade-enter-active,
    .fade-leave-active {
        transition: opacity .5s;
    }

    .fade-enter,
    .fade-leave-to

    /* .fade-leave-active below version 2.1.8 */
        {
        opacity: 0;
    }

    .panel {
        height: 10em;
        width: 40%;
        min-width: 250px;
        display: block;
        margin-left: auto;
        margin-right: auto;
        align-items: center
    }

    .color {
        filter: invert(100%);
        background-color: rgba(255, 255, 255, 0.4);
    }

    .sub-background {
        background-color: var(--sub-background) !important;
    }

    .sub-card {
        background-color: var(--sub-card) !important;
    }

    .sub-wallpaper {
        background-color: var(--sub-wallpaper) !important;
    }

    .sub-toolbar {
        background-color: var(--sub-toolbar) !important;
    }

    .sub-text {
        color: var(--sub-text) !important;
        text-shadow: 1px 1px var(--sub-dark) !important;
    }

    .sub-button {
        background-color: rgb(10, 9, 65) !important;
        /*background: linear-gradient(var(--sub-button),var(--sub-wallpaper)) !important;*/
        color: #FFF !important;
        text-shadow: var(--sub-dark) 1px 1px 0 !important;
    }

    .sub-panel {
        background-color: var(--sub-panel) !important;
    }


    .loginpaper {
        background:
            linear-gradient(217deg, var(--sub-card), var(--sub-wallpaper) 70.71%),
            linear-gradient(127deg, var(--sub-background), var(--sub-toolbar) 70.71%),
            linear-gradient(336deg, var(--sub-text), var(--sub-button) 70.71%);
        height: 100%;
        width: 100%
    }
</style>
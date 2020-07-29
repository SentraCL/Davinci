<template>

    <span class="content">
        <div class="row login">
            <div class="col-xl-4 col-lg-4 col-md-4">
            </div>
            <div class="col-xl-4 col-lg-4 col-md-4">
                <card class="card-user">
                    <div slot="image">
                        <img src="@/assets/img/divina.jpg" height="280" width="540" />
                    </div>
                    <div>
                        <div class="author">
                            <img class="avatar border-white" src="@/assets/img/davinci-logo.png" alt=" ..." />
                            <h4 class="title">
                                Iniciar Sesión
                                <br />
                            </h4>
                            <transition name="fade">
                                <small v-if="dijo" class="frase">
                                    <span v-html="frase"></span>
                                </small>
                            </transition>
                            <small>
                                <br /></small>
                        </div>

                        <div class="panel">
                            <div class="col-md-12">
                                <input-text :label="title.alias" removeSpace name="user" @keyup.enter.native="doLogin" v-model="login.alias" autocomplete="off">
                                </input-text>
                            </div>
                            <div class="col-md-12">
                                <input-text :label="title.pass" name="password" v-model="login.password" @keyup.enter.native="doLogin" type="password" autocomplete="off">
                                </input-text>
                            </div>
                            <div class="col-md-12 ">
                                <button type="button" @click="doLogin" style="width: 100%" class="btn btn-xs btn-round btn-info">
                                    Iniciar
                                </button>
                            </div>
                        </div>
                    </div>
                    <hr />
                    <div class="text-center">

                        <small><strong>DaVinci</strong> Repositorio de Ideas y Diseños. &copy; Creado por <a href="http://www.sentra.cl/" target="_blank">&nbsp;<strong>Sentra.</strong></a></small>

                    </div>
                </card>
            </div>
            <div class="col-xl-4 col-lg-4 col-md-4">
            </div>
        </div>
    </span>
</template>
<script>
    import io from 'socket.io-client'
    export default {
        data() {
            return {
                frase: "",
                dijo: false,
                title: {
                    alias: `<span class="ti-user" ></span> Usuario`,
                    pass: `<span class="ti-key" ></span> Contraseña`
                },
                //127.0.0.01
                socket: io(`${window.location.origin}`, {
                    forceNew: true
                }),
                login: {
                    alias: "",
                    password: ""
                }
            };
        },
        mounted: async function () {
            var me = this
            this.login.alias = ""
            this.login.password = ""

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
                    .post("/api/login/", {
                        "user": this.login.alias,
                        "pass": this.login.password
                    }).then(rs => {
                        if (rs.data.Online) {
                            sessionStorage.setItem("username",this.login.alias);
                            sessionStorage.setItem("loginHash", rs.data.DavinciCode);
                            window.location.href = "/"; //+ rs.data.DavinciCode

                        } else {
                            me.login = {
                                alias: "",
                                password: ""
                            }
                            me.dijo = true
                            me.frase = "<strong>Credenciales no validas..</strong>"
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
</style>
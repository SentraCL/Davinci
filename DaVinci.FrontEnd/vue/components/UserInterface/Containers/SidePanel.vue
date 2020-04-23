<template>
    <span v-if="toShow">
        <div class="blockbackground" >

            <ul class="side-effect side-panel">
                <ul>
                    <div class="fa-pull-right">
                        <button class="btn btn-xs btn-success" v-if="close" @click="doClose()">x</button>
                    </div>
                    <slot name="content">

                    </slot>
                </ul>
            </ul>
        </div>

    </span>
</template>
<script>
    export default {
        name: "side-panel-rigth",
        watch: {
            show(newValue, oldValue) {
                if (newValue) {
                    this.showAnimated();
                } else {
                    this.hiddenAnimated();

                }
            }
        },
        data() {
            return {
                toShow: this.show
            }
        },
        props: {
            show: Boolean,
            close: Boolean
        },
        methods: {

            doClose() {
                this.hiddenAnimated();
            },

            showAnimated() {
                this.doShow = true;
                var x = document.getElementsByClassName("blockbackground");
                var i;
                for (i = 0; i < x.length; i++) {
                    // if (i == 0) {
                    x[i].style.display = "block";

                    //}
                }


                document.getElementById("SUBDOMAIN-MENU").classList.remove('main-menu');
                this.toShow = true;
                //console.log("Iniciando Animacion de Aparicion");
                var container = document.getElementById('WORKSPACE');
                container.className = 'st-content'; // clear
                this.toAddClass(container, "side-effect");
                var me = this;

                setTimeout(function () {
                    /*
                    var divHtml = `<div class="blockbackground"></div>`;
                    var e = document.createElement('div');
                    e.innerHTML = divHtml;
                    var element = document.body;
                    while (e.firstChild) {
                        element.appendChild(e.firstChild);
                    }
                    */
                    me.toAddClass(container, 'side-panel-open');
                }, 25);
            },
            hiddenAnimated() {
                if (this.doShow) {
                    //console.log("Iniciando Animacion de Ocultamiento");
                    var container = document.getElementById('WORKSPACE');
                    this.toRemoveClass(container, 'side-panel-open');
                    var me = this;
                    setTimeout(function () {
                        me.$emit("update:show", false);
                        me.toShow = false;
                        document.getElementById("SUBDOMAIN-MENU").classList.add('main-menu');
                        var x = document.getElementsByClassName("blockbackground");
                        var i;
                        for (i = 0; i < x.length; i++) {
                            x[i].style.display = "none";
                        }

                    }, 600);
                }
            }
            ,
            toAddClass(elem, c) {
                if (!this.hasClass(elem, c)) {
                    elem.className = elem.className + ' ' + c;
                }
            },
            toRemoveClass(elem, c) {
                elem.className = elem.className.replace(this.classReg(c), ' ');
            },
            hasClass(elem, c) {
                return elem.classList.contains(c);
            },
            classReg(className) {
                return new RegExp("(^|\\s+)" + className + "(\\s+|$)");
            }
        }
    };
</script>
<style>

</style>
<template>

    <div class="Expander">
        <div class="Expander__trigger" @click="open=!open" :class="open?'active':'beforeBorder'">
            <div class="row">
                <svg class="Expander__trigger-Icon" :class="{open:open}" width="40" height="12" stroke="cornflowerblue">
                    <polyline points="12,2 20,10 28,2" stroke-width="3" fill="none"></polyline>
                </svg>
                <span style="top:-5;position:relative" v-html="title"></span>
            </div>
        </div>
        <transition :name="animation">
            <div class="Expander__body" v-show="open">
                <slot></slot>
            </div>
        </transition>
    </div>

</template>
<script>
    export default {
        name: "accordion",
        computed: {

        },
        props: {
            title: {
                type: String,
                default: 'title'
            },
            animation: {
                type: String,
                default: 'rightToLeft'
                // validator: prop => ['leftToRight', 'bounceIn', 'bottomToTop'].includes(prop)
            }
        },
        data() {
            return {
                open: false
            }
        },

        methods: {

        }
    };
</script>
<style>
    .beforeBorder {
        position: relative;
    }

    .beforeBorder:before {
        transition: opacity 0.1s linear, transform 0.2s ease-in-out;
        position: absolute;
        border-bottom: 1px solid currentColor;
        content: '';
        width: 100%;
        left: 0;
        bottom: -1px;
    }

    .beforeBorder:not(:hover)::before {
        transform: scaleX(0);
        opacity: 0;
    }

    .Accordion {
        background: #fff;
        box-shadow: 0 1px 12px 1PX rgba(0, 0, 0, 0.25);
        overflow: hidden;
    }

    .Expander__trigger {
        cursor: pointer;
        padding: 0.7rem 0.5rem;
        border-bottom: 1px solid #efefef;
    }

    .Expander__trigger:hover {
        color: #477dca;
    }

    .Expander__trigger.active {
        border-bottom-color: #477dca;
    }

    .Expander__trigger-Icon {
        transition: transform 0.2s cubic-bezier(0.23, 1, 0.32, 1);
    }

    .Expander__trigger-Icon.open {
        stroke: #ff6347;
        transform: rotate(180deg);
    }

    .Expander__body {
        padding: 1rem 1.5rem;
        background: transparent;
    }

    .rightToLeft-enter-active {
        animation: rightToLeft 0.2s;
    }

    .rightToLeft-leave-active {
        animation: rightToLeft 0.2s reverse;
    }

    .leftToRight-enter-active {
        animation: leftToRight 0.2s;
    }

    .leftToRight-leave-active {
        animation: leftToRight 0.2s reverse;
    }

    .bounceIn-enter-active {
        animation: bounceIn 0.3s;
    }

    .bounceIn-leave-active {
        animation: bottomToTop 0.2s reverse;
    }

    .bottomToTop-enter-active {
        animation: bottomToTop 0.2s forwards;
    }

    .bottomToTop-leave-active {
        animation: bottomToTop 0.2s reverse;
    }

    @-moz-keyframes rightToLeft {
        0% {
            transform: translateX(100vw);
        }

        50% {
            transform: translateX(-2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-webkit-keyframes rightToLeft {
        0% {
            transform: translateX(100vw);
        }

        50% {
            transform: translateX(-2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-o-keyframes rightToLeft {
        0% {
            transform: translateX(100vw);
        }

        50% {
            transform: translateX(-2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @keyframes rightToLeft {
        0% {
            transform: translateX(100vw);
        }

        50% {
            transform: translateX(-2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-moz-keyframes leftToRight {
        0% {
            transform: translateX(-100vw);
        }

        50% {
            transform: translateX(2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-webkit-keyframes leftToRight {
        0% {
            transform: translateX(-100vw);
        }

        50% {
            transform: translateX(2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-o-keyframes leftToRight {
        0% {
            transform: translateX(-100vw);
        }

        50% {
            transform: translateX(2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @keyframes leftToRight {
        0% {
            transform: translateX(-100vw);
        }

        50% {
            transform: translateX(2em);
        }

        100% {
            transform: translateX(0);
        }
    }

    @-moz-keyframes bounceIn {

        from,
        20%,
        40%,
        60%,
        80%,
        to {
            animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
        }

        0% {
            opacity: 0;
            transform: scale3d(0.3, 0.3, 0.3);
        }

        20% {
            transform: scale3d(1.1, 1.1, 1.1);
        }

        40% {
            transform: scale3d(0.9, 0.9, 0.9);
        }

        60% {
            opacity: 1;
            transform: scale3d(1.03, 1.03, 1.03);
        }

        80% {
            transform: scale3d(0.97, 0.97, 0.97);
        }

        to {
            opacity: 1;
            transform: scale3d(1, 1, 1);
        }
    }

    @-webkit-keyframes bounceIn {

        from,
        20%,
        40%,
        60%,
        80%,
        to {
            animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
        }

        0% {
            opacity: 0;
            transform: scale3d(0.3, 0.3, 0.3);
        }

        20% {
            transform: scale3d(1.1, 1.1, 1.1);
        }

        40% {
            transform: scale3d(0.9, 0.9, 0.9);
        }

        60% {
            opacity: 1;
            transform: scale3d(1.03, 1.03, 1.03);
        }

        80% {
            transform: scale3d(0.97, 0.97, 0.97);
        }

        to {
            opacity: 1;
            transform: scale3d(1, 1, 1);
        }
    }

    @-o-keyframes bounceIn {

        from,
        20%,
        40%,
        60%,
        80%,
        to {
            animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
        }

        0% {
            opacity: 0;
            transform: scale3d(0.3, 0.3, 0.3);
        }

        20% {
            transform: scale3d(1.1, 1.1, 1.1);
        }

        40% {
            transform: scale3d(0.9, 0.9, 0.9);
        }

        60% {
            opacity: 1;
            transform: scale3d(1.03, 1.03, 1.03);
        }

        80% {
            transform: scale3d(0.97, 0.97, 0.97);
        }

        to {
            opacity: 1;
            transform: scale3d(1, 1, 1);
        }
    }

    @keyframes bounceIn {

        from,
        20%,
        40%,
        60%,
        80%,
        to {
            animation-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
        }

        0% {
            opacity: 0;
            transform: scale3d(0.3, 0.3, 0.3);
        }

        20% {
            transform: scale3d(1.1, 1.1, 1.1);
        }

        40% {
            transform: scale3d(0.9, 0.9, 0.9);
        }

        60% {
            opacity: 1;
            transform: scale3d(1.03, 1.03, 1.03);
        }

        80% {
            transform: scale3d(0.97, 0.97, 0.97);
        }

        to {
            opacity: 1;
            transform: scale3d(1, 1, 1);
        }
    }

    @-moz-keyframes bottomToTop {
        0% {
            opacity: 0;
            transform: translateY(100%);
        }

        100% {
            transform: translateY(0);
        }
    }

    @-webkit-keyframes bottomToTop {
        0% {
            opacity: 0;
            transform: translateY(100%);
        }

        100% {
            transform: translateY(0);
        }
    }

    @-o-keyframes bottomToTop {
        0% {
            opacity: 0;
            transform: translateY(100%);
        }

        100% {
            transform: translateY(0);
        }
    }

    @keyframes bottomToTop {
        0% {
            opacity: 0;
            transform: translateY(100%);
        }

        100% {
            transform: translateY(0);
        }
    }
</style>
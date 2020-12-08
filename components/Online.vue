<template lang="pug">
div.Online
</template>

<style>
.Online {
    background: #1414ff;
    z-index: 9999;
    width:8px;
    height:8px;
    position: absolute;
    top: 16px;
    right: 16px;
    border-radius: 50%;
}
.xOnline:before {
    content: '';
    position: absolute;
    top: -12px;
    left: -12px;
    width: 32px;
    height: 32px;
    border-radius: 100%;
    background: radial-gradient(transparent 10%,rgba(33,150,243,.6) 100%);
    -webkit-animation: aniTwo 1s ease .4s infinite;
    animation: aniTwo 1s ease .4s infinite;
}
@keyframes aniTwo {
    0% {
        transform: scale(.1);
        opacity: 1;
    }
    100% {
        transform: scale(1.2);
        opacity: .2;
    }
}
</style>
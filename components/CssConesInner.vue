<template>
    <ul class="cones-inner">
        <li class="bottom"></li>
        <li class="front"></li>
        <li class="right"></li>
        <li class="left"></li>
    </ul>
</template>

<style>
.cones-inner {
    padding: 0;
    list-style: none;
    display: block;
    position: absolute;
    top: 50%;
    left: 30%;
    margin: -25px 0 0 -25px;
    -webkit-transform-style: preserve-3d;
    transform-style: preserve-3d;
    -webkit-transform-origin: 25px 25px;
    transform-origin: 25px 25px;
    -webkit-transform: rotateY(35deg);
    transform: rotateY(35deg);
    -webkit-animation: slowRotate 6s ease-in-out infinite 0s;
    animation: slowRotate 6s ease-in-out infinite 0s;
}
.cones-inner>li {
    position: absolute;
    top: 0;
    z-index: 10;
    border-left: 28.87px solid transparent;
    border-right: 28.87px solid transparent;
}
.cones-inner .bottom {
    border-bottom: 50px solid rgba(200,200,200,.7);
}
.cones-inner .front {
    border-bottom: 50px solid rgba(80,80,80,.7);
    -webkit-transform: rotate3d(-1,1.732,0,70.5deg);
    transform: rotate3d(-1,1.732,0,70.5deg);
    -webkit-transform-origin: 25% 50%;
    transform-origin: 25% 50%;
}
.cones-inner .right {
    border-bottom: 50px solid rgba(160,160,160,.7);
    -webkit-transform: rotate3d(1,0,0,70.5deg);
    transform: rotate3d(1,0,0,70.5deg);
    -webkit-transform-origin: 50% 100%;
    transform-origin: 50% 100%;
}
.cones-inner .left {
    border-bottom: 50px solid rgba(120,120,120,.7);
    -webkit-transform: rotate3d(1,1.732,0,-70.5deg);
    transform: rotate3d(1,1.732,0,-70.5deg);
    -webkit-transform-origin: 75% 50%;
    transform-origin: 75% 50%;
}
@keyframes slowRotate {
    0% {
        -webkit-transform: rotateY(35deg);
        transform: rotateY(35deg);
    }
    100% {
        -webkit-transform: rotateY(395deg);
        transform: rotateY(395deg);
    }
}
</style>
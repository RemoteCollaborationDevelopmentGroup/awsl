<template lang="pug">
div.container_works
    vs-card(v-for="s in 8" :key="s")
        template(#title)
            h3 Pot with a plant
        template(#img)
            img(src="/foto5.png" alt="")
        template(#text)
            p Lorem ipsum dolor sit amet consectetur, adipisicing elit.
        template(#interactions)
            vs-button(danger icon)
                i.bx.bx-heart
            vs-button.btn-chat(shadow primary)
                i.bx.bx-chat
                span.span 54
</template>

<style>
.container_works {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  text-align: center;
}
.container_works .vs-card {
    text-align: left;
    margin: .5em;
}
</style>
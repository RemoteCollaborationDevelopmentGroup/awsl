<template lang="pug">
.container
  div
    header.content-logos
      logo
      span.plus +
      VuesaxLogo
    h1.title Nuxt.js + Vuesax
    h2.subtitle
      | <a href="https://vuesax.com/">Vuesax</a>
      | is a framework of ui components
      | for <a href="https://vuejs.org/">Vuejs</a>, It was created to make new
      | interfaces that have a new trend and are visually beautiful
    </h2>
    .links
      h3.h3 Vuesax
      a.button--green(
        href="https://vuesax.com/",
        target="_blank",
        rel="noopener noreferrer"
      ) Documentation
      a.button--grey(
        href="https://discordapp.com/invite/9dsKtvB",
        target="_blank",
        rel="noopener noreferrer"
      ) Discord
      a.button--grey(
        href="https://github.com/lusaxweb/vuesax",
        target="_blank",
        rel="noopener noreferrer"
      ) GitHub
    .links
      h3.h3 Nuxt.js
      a.button--green(
        href="https://nuxtjs.org/",
        target="_blank",
        rel="noopener noreferrer"
      ) Documentation
      a.button--grey(
        href="https://github.com/nuxt/nuxt.js",
        target="_blank",
        rel="noopener noreferrer"
      ) GitHub
  div
    p 关于社区

  h2 社区成员
  .member
    vs-table(v-model="selected")
      template(#thead)
        vs-tr
          vs-th name
          vs-th city
          vs-th roles
          vs-th skill
          vs-th skill_ex
          vs-th time
          vs-th mark
      template(#tbody)
        vs-tr(
          v-for="(user, i) in users",
          :data="user.id",
          :is-selected="selected == user.id",
          :key="i"
        )
          vs-td {{ user.name }}
          vs-td {{ user.city }}
          vs-td {{ user.roles }}
          vs-td {{ user.skill }}
          vs-td {{ user.skill_ex }}
          vs-td {{ user.time }}
          vs-td {{ user.mark }}
</template>

<script>
import Logo from "~/components/Logo.vue";
import VuesaxLogo from "~/components/VuesaxLogo.vue";
import users from "~/assets/users.json";
export default {
  components: {
    Logo,
    VuesaxLogo,
  },
  data: () => ({
    selected: null,
    users: [],
  }),
  mounted() {
    this.users = users;
  },
};
</script>

<style lang="sass" scoped>
.member
  max-width: 1280px
  margin: auto

.container
  margin: 0 auto
  min-height: calc(100vh - 120px - 64px)
  display: flex
  justify-content: center
  align-items: center
  text-align: center
  flex-direction: column

.title
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif
  display: block
  font-weight: 300
  font-size: 55px
  color: #35495e
  letter-spacing: 1px
  text-transform: capitalize
  margin: 25px 0

.subtitle
  font-weight: 300
  font-size: 1.1rem
  color: #526488
  word-spacing: 2px
  padding-bottom: 15px
  max-width: 600px

  & a
    font-weight: 500
    color: inherit

.links
  padding-top: 15px
  margin-bottom: 20px

.content-logos
  display: flex
  align-items: center
  justify-content: center
  min-width: 500px

.plus
  font-size: 2.5rem
  margin: 15px
  color: #35495e

.h3
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif
  font-weight: 400
  margin: 10px
</style>

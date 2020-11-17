<template lang="pug">
div.center.examplex
  vs-navbar(target-scroll="#padding-scroll-content" center-collapsed v-model="active" style="position:fixed;")
    template(#left)
      img(src="/logo2.png" alt="awsl")
    vs-navbar-item(to="/" :active="active == ''" id="") 开发组
    vs-navbar-item(to="/works" :active="active == 'works'" id="works") 作品集
    vs-navbar-item(to="/remote" :active="active == 'remote'" id="remote") 远程协作
    vs-navbar-item(to="/license" :active="active == 'license'" id="license") license
    template(#right)
      vs-button(flat) Login
      vs-button Get Started
  div(id="padding-scroll-content" class="square")
    Nuxt
    //div.child(style="height:18em;") child 1
    //div.child(style="height:18em;") child 2
    //div.child(style="height:18em;") child 3
    //padding-scroll
</template>

<script>
export default {
  data: ({ $route }) => ({
    active: $route.path.replace('/', '')
  }),
  mounted() {
    console.log(this.active);
  }
}
</script>

<style>
html {
  font-family:
    'Source Sans Pro',
    -apple-system,
    BlinkMacSystemFont,
    'Segoe UI',
    Roboto,
    'Helvetica Neue',
    Arial,
    sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
}

.button--green {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #3b8070;
  color: #3b8070;
  text-decoration: none;
  padding: 10px 30px;
}

.button--green:hover {
  color: #fff;
  background-color: #3b8070;
}

.button--grey {
  display: inline-block;
  border-radius: 4px;
  border: 1px solid #35495e;
  color: #35495e;
  text-decoration: none;
  padding: 10px 30px;
  margin-left: 15px;
}

.button--grey:hover {
  color: #fff;
  background-color: #35495e;
}
</style>

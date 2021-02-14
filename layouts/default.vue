<template lang="pug">
.center.examplex
  Online
  vs-navbar(center-collapsed, v-model="active")
    template(#left)
    vs-navbar-item(
      v-for="(item, i) in list",
      :active="active == item.path",
      :to="`/${item.path}`",
      :key="i"
    ) {{ item.name }}
    template(#right)
      vs-button(flat) Sign in
      vs-button Sign up
  #padding-scroll-content.square(
    style="padding: 120px 120px 64px 120px; min-height: 100vh"
  )
    Nuxt
</template>

<script>
export default {
  data: ({ $route }) => ({
    active: $route.path.replace("/", ""),
    list: [
      { path: "", name: "start" },
      //{ path: "member", name: "member" },
      { path: "task", name: "task" },
      { path: "project", name: "project" },
      { path: "idea", name: "idea" },
      //{ path: "guild", name: "guild" },
      { path: "team", name: "team" },
      { path: "laboratory", name: "laboratory" },
      //{ path: "works", name: "作品集" },
      //{ path: "remote", name: "远程协作" },
      { path: "show", name: "show" },
      { path: "about", name: "about" },
    ],
  }),
  watch: {
    $route: {
      handler: function (val, old) {
        this.active = val.path.replace(/\//, "").replace(/\/.*$/g, "");
      },
    },
  },
  mounted() {
    //console.log(this.active);
    this.init_websocket();
  },
  methods: {
    only_websocket() {},
    init_websocket() {
      if (!window["WebSocket"]) {
        console.log("Your browser does not support WebSockets.");
        this.append_log("Your browser does not support WebSockets.");
        return;
      }

      this.conn = new WebSocket("wss://" + document.location.host + "/ws");
      this.conn.onclose = function (e) {
        let message = e.stopImmediatePropagation.split("\n");
        console.log(message);
      };
      this.conn.onmessage = function (e) {
        let message = e.data.split("\n");
        console.log(message);
      };
      window.ws_send = (data) => {
        if (!this.conn) {
          console.log("Connection closed.");
          return;
        }
        this.conn.send(data);
      };
    },
  },
};
</script>

<style>
html {
  font-family: "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI",
    Roboto, "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}
body::-webkit-scrollbar {
  display: none;
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

export default {
  head: {
    title: 'awsl',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ],
    script: [
      { src: "/js/centcount.js" }
    ]
  },
  css: ['vuesax/dist/vuesax.css'],
  plugins: ['@/plugins/vuesax'],
  components: true,
  buildModules: [],
  modules: [],
  build: { extractCSS: true }
}

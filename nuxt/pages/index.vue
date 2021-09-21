<template>
  <div id="app">
    <v-toolbar>
      <img :src="require('@/assets/images/title_dot.png')" :height="logosize" />
    </v-toolbar>
    <v-container>
      <ShowNum style="padding-top: 1em; padding-bottom: 1em" />
      <ShowGraph style="padding-bottom: 2em" />
    </v-container>
  </div>
</template>

<script>
import ShowNum from '../components/ShowNum.vue'
import ShowGraph from '../components/ShowGraph.vue'
export default {
  components: {
    ShowNum,
    ShowGraph,
  },
  data: () => ({
    drawer: false,
    group: null,
    tab: 0,
    items: [
      { tab: 'mdi-human-greeting' },
      { tab: 'mdi-equalizer' },
      { tab: 'mdi-google-downasaur' },
    ],
  }),
  computed: {
    logosize: function () {
      switch (this.$vuetify.breakpoint.name) {
        case 'xs':
          return 40
        case 'sm':
          return 50
        case 'md':
          return 60
        case 'lg':
          return 60
        case 'xl':
          return 60
      }
    },
  },
}
</script>

<style scoped>
h1 {
  color: #000;
}
.bg {
  top: 0;
  left: 0;
  height: 110hv;
  background-size: cover;
  background-image: url('~@/assets/images/bg_bar.png');
}
</style>

<style lang="scss">
.v-application {
  font-family: 'DotsFont';
}
</style>

<template>
  <div id="app">
    <v-toolbar>
      <img :src="require('@/assets/images/title_dot.png')" :height="logosize" />
    </v-toolbar>
    <v-container>
      <ShowNum class="py-4" />
      <ShowGraph />
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

<style lang="scss">
.v-application {
  font-family: 'DotsFont';
}
</style>

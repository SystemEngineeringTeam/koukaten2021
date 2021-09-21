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
import ShowNum from '/components/shownum.vue'
import ShowGraph from '/components/showgraph.vue'
export default {
  components: {
    ShowNum,
    ShowGraph,
  },
  
  computed: {
    logosize: function () {
      switch (this.$vuetify.breakpoint.name) {
        case 'xs':
          return 40
        case 'sm':
          return 50
        default: 
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

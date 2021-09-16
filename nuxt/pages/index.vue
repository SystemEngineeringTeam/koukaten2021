<template>
  <div id="app">
    <!-- <v-card flat> -->
    <v-toolbar>
      <img :src="require('@/assets/images/title_dot.png')" :height="logosize" />
      <!-- <template v-slot:extension>
          <v-tabs
            color="white"
            centered
            fixed-tabs
            v-model="tab"
          >
            <v-tabs-slider color="yellow"></v-tabs-slider>
            <v-tab v-for="item in items" :key="item.tab">
              <v-icon>{{ item.tab }}</v-icon>
            </v-tab>
          </v-tabs>
        </template> -->
    </v-toolbar>
    <v-container>
      <row>
        <br />
        <!-- <v-tabs-items v-model="tab">
        <v-tab-item class="bg"> -->
        <ShowNum />
      </row>
      <!-- </v-tab-item>
        <v-tab-item class="bg"> -->
      <row>
        <br />
        <br />
        <ShowGraph />
        <br />
        <br />
      </row>
    </v-container>
    <!-- </v-tab-item>
        <v-tab-item  class="bg" style="height: 90vh">
          <h1>工事中！</h1>
        </v-tab-item>
      </v-tabs-items> -->
    <!-- </v-card> -->
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
  }
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
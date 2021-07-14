<template>
  <div id="app">
    <v-card flat>
      <v-toolbar color="blue" dark>
        <v-toolbar-title>シス研人数出る蔵</v-toolbar-title>
        <template v-slot:extension>
          <v-tabs
            color="white"
            background-color="blue"
            centered
            fixed-tabs
            v-model="tab"
          >
            <v-tabs-slider color="yellow"></v-tabs-slider>
            <v-tab v-for="item in items" :key="item.tab">
              <v-icon>{{ item.tab }}</v-icon>
            </v-tab>
          </v-tabs>
        </template>
      </v-toolbar>

      <v-tabs-items v-model="tab">
        <v-tab-item>
          <ShowNum />
        </v-tab-item>
        <v-tab-item>
          <h1>工事中！</h1>
        </v-tab-item>
        <v-tab-item>
          <h1>工事中！</h1>
        </v-tab-item>
      </v-tabs-items>
    </v-card>
  </div>
</template>

<script>
import ShowNum from '../components/ShowNum.vue'
export default {
  components: {
    ShowNum,
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
}
</script>

<style scoped>
h1 {
  color: #000;
}
</style>

<template>

  <v-app>
    <v-app-bar max-height="64px" color="blue" dark>
      <v-toolbar-title>シス研人数出る蔵</v-toolbar-title>
      <template v-slot:extension>
        <v-tabs　color="white" background-color="blue" centered fixed-tabs v-model="tab">
          <v-tabs-slider color="yellow"></v-tabs-slider>
            <v-tab v-for="item in items" :key="item.tab">
                <v-icon>{{ item.tab }}</v-icon>
            </v-tab>
        </v-tabs>
      </template>
      <v-tabs-items v-model="tab">
        <v-tab-item v-for="item in items" :key="item.tab">
          <v-card flat>
            <v-card-text>{{ item.content }}</v-card-text>
          </v-card>
        </v-tab-item>
      </v-tabs-items>
    </v-app-bar>
    <ShowNum />
  </v-app>
</template>

<script>
import ShowNum from '../components/ShowNum.vue'
export default {
  name: 'menubar',
  components: {
    ShowNum,
  },
  data: () => ({
    drawer: false,
    group: null,
  }),
  data() {
    return {
      tab: null,
      items: [
        { tab: 'mdi-human-greeting', content: 'Tab1' },
        { tab: 'mdi-equalizer', content: 'Tab2' },
        { tab: 'mdi-google-downasaur', content: 'Tab3' },
      ],
    }
  },
}
</script>
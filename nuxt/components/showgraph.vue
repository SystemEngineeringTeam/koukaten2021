<!--グラフ表示コンポーネント-->
<template>
  <v-row style="height: 500px" justify="center" align-content="center">
    <v-col cols="8">
      <v-card class="justify-center">
        <v-col cols="12" sm="3">
          <h3>人数予想</h3>
          <v-select
            v-model="selectedDoW"
            :items="dow"
            label="day of week"
            outlined
          ></v-select>
        </v-col>
        <v-row align="center">
          <v-col align="center">
            <line_chart
              v-if="loaded"
              v-bind:chart-value="computed.value"
            />
            </line_chart>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>

import Chart from './line_chart';

const gradients = [
  ['#222'],
  ['#42b3f4'],
  ['red', 'orange', 'yellow'],
  ['purple', 'violet'],
  ['#00c6ff', '#F0F', '#FF0'],
  ['#f72047', '#ffd200', '#1feaea'],
]

export default {
  name: 'ShowGraph',
  data: () => ({
    gradients,
    selectedDoW: '',
    dow: ['日曜日', '月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日'],
    graphKey: Math.random(),
    chartdata: {
    }
  }),
  methods: {
    indexToTime: function (index) {
      if (index % 3 === 0 && index !== 0) {
        return index + 'h'
      }
      return ''
    },
  },
  computed: {
    value: function () {
      switch (this.selectedDoW) {
        case '月曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 4, 6, 3, 3, 2, 3, 4, 2, 2, 1, 1, 0,
            0,
          ]
        // value = this.$store.state.monday
        case '火曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 2, 0, 0, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.tuesday
        case '水曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 2, 1, 2, 0, 0, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.wednesday
        case '木曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 3, 3, 2, 1, 1, 1, 3, 3, 2, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.thursday
        case '金曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 4, 3, 3, 1, 0, 0, 0, 0, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.friday
        case '土曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.saturday
        case '日曜日':
          return [
            0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            0,
          ]
        // value = this.$store.state.sunday
      }
    },
  },
  watch: {
    value: function () {
      this.graphKey = Math.random()
    },
  },
  created: function () {
    let date = new Date()
    this.selectedDoW = this.dow[date.getDay()]
  },
  components: {
    Chart,
  },
}
</script>

<style scoped>
h2 {
  display: inline;
  font-size: 32px;
}
h1 {
  display: inline;
  font-size: 64px;
}
h3 {
  font-size: 20px;
}
</style>
<!--グラフ表示コンポーネント-->
<template>
  <v-row justify="center">
    <v-col :cols="cardwidth">
      <v-card outlined>
        <v-container>
          <v-row>
            <v-card-text
              class="font-weight-bold mb-n2"
              style="font-size: 1.25rem"
              >【人数予想】</v-card-text
            >
          </v-row>
          <v-row>
            <v-col cols="12" sm="4">
              <v-select
                v-model="selectedDoW"
                :items="dow"
                label="day of week"
                outlined
                class="mb-n8"
              ></v-select>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <Chart
                :datasets="datasets"
                :labels="graphLabels"
                :options="graphOptions"
              />
            </v-col>
          </v-row>
        </v-container>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import Chart from './barchart.vue'

// グラフのラベル(定数)
const graphLabels = [...Array(24).keys()].map((element) => {
  return element + 'h'
})

// グラフの設定(定数)
const graphOptions = {
  scales: {
    y: {
      beginAtZero: true,
      max: 10,
      ticks: {
        stepSize: 5,
        callback: function (value) {
          return value + '人'
        },
      },
      grid: {
        color: '#151515',
      },
    },
    x: {
      beginAtZero: true,
    },
  },
  plugins: {
    legend: {
      display: false,
    },
  },
  responsive: true,
  maintainAspectRatio: false,
}

// 曜日(定数)
const dow = [
  '日曜日',
  '月曜日',
  '火曜日',
  '水曜日',
  '木曜日',
  '金曜日',
  '土曜日',
]

const graphBorderColor = [
  'rgba(51, 204, 204, 1)',
  'rgba(51, 204, 153, 1)',
  'rgba(255, 206, 86, 1)',
  'rgba(255, 159, 64, 1)',
  'rgba(255, 99, 132, 1)',
]

const graphBackGroundColor = [
  'rgba(51, 204, 204, 0.2)',
  'rgba(51, 204, 153, 0.2)',
  'rgba(255, 206, 86, 0.2)',
  'rgba(255, 159, 64, 0.2)',
  'rgba(255, 99, 132, 0.2)',
]

export default {
  name: 'ShowGraph',

  components: {
    Chart,
  },

  props: ['cardwidth'],

  data: () => ({
    selectedDoW: '',
  }),

  computed: {
    datasets: function () {
      let backGroundColors = this.graphValue.map((element) => {
        return this.graphBackGroundColor[
          Math.max(Math.ceil(element / 2) - 1, 0)
        ]
      })
      let borderColors = this.graphValue.map((element) => {
        return this.graphBorderColor[Math.max(Math.ceil(element / 2) - 1, 0)]
      })
      return [
        {
          label: '人数',
          data: this.graphValue,
          backgroundColor: backGroundColors,
          borderColor: borderColors,
          borderWidth: 1,
        },
      ]
    },
    graphValue: function () {
      switch (this.selectedDoW) {
        case '月曜日':
          return [
            10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 3, 4, 6, 3, 3, 2, 3, 4, 2, 2, 1, 1,
            0, 0,
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

  created: function () {
    // 定数をthisに定義
    this.graphLabels = graphLabels
    this.graphOptions = graphOptions
    this.dow = dow
    this.graphBorderColor = graphBorderColor
    this.graphBackGroundColor = graphBackGroundColor

    let date = new Date()
    this.selectedDoW = this.dow[date.getDay()]
  },
}
</script>

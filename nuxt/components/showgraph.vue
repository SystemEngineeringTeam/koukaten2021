<!--グラフ表示コンポーネント-->
<template>
  <v-row justify="center" align-content="center">
    <v-col cols="8">
      <v-card class="justify-center" outlined>
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
            <Chart
              :datasets="datasets"
              :labels="graphLabels"
              :options="graphOptions"
            />
          </v-col>
        </v-row>
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

export default {
  name: 'ShowGraph',
  components: {
    Chart,
  },
  data: () => ({
    selectedDoW: '',
  }),
  computed: {
    datasets: function () {
      return [
        {
          label: '人数',
          data: this.graphValue,
          backgroundColor: [
            'rgba(255, 99, 132, 0.2)',
            'rgba(54, 162, 235, 0.2)',
            'rgba(255, 206, 86, 0.2)',
            'rgba(75, 192, 192, 0.2)',
            'rgba(153, 102, 255, 0.2)',
            'rgba(255, 159, 64, 0.2)',
          ],
          borderColor: [
            'rgba(255,99,132,1)',
            'rgba(54, 162, 235, 1)',
            'rgba(255, 206, 86, 1)',
            'rgba(75, 192, 192, 1)',
            'rgba(153, 102, 255, 1)',
            'rgba(255, 159, 64, 1)',
          ],
          borderWidth: 1,
        },
      ]
    },
    graphValue: function () {
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

  created: function () {
    // 定数をthisに定義
    this.graphLabels = graphLabels
    this.graphOptions = graphOptions
    this.dow = dow

    let date = new Date()
    this.selectedDoW = this.dow[date.getDay()]
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

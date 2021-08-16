<!--グラフ表示コンポーネント-->
<template>
  <v-row style="height: 500px" justify="center" align-content="center">
    <v-col cols="8">
      <v-card class="justify-center">
        <v-col cols="12" sm="3">
          <h3>人数予想</h3>
          <v-select
            v-model="selectedWoD"
            :items="wod"
            label="day of week"
            outlined
          ></v-select>
        </v-col>
        <v-row align="center">
          <v-col align="center">
            <v-sparkline
              :value="value"
              :gradient="gradient"
              :smooth="radius || false"
              :padding="padding"
              :line-width="width"
              :stroke-linecap="lineCap"
              :gradient-direction="gradientDirection"
              :fill="fill"
              :type="type"
              :auto-line-width="autoLineWidth"
              :show-labels="true"
              :key="graphKey"
              auto-draw
            >
              <template v-slot:label="item">
                {{ indexToTime(item.index) }}
              </template>
            </v-sparkline>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
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
    width: 5,
    radius: 0,
    padding: 8,
    lineCap: 'round',
    gradient: gradients[5],
    // value: this.$store.state.monday,
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'bar',
    autoLineWidth: false,
    selectedWoD: '',
    wod: ['月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日', '日曜日'],
    graphKey: Math.random(),
  }),
  methods: {
    indexToTime: function (index) {
      if (index % 3 === 0) {
        return index
      }
      return ''
    },
  },
  computed: {
    value: function () {
      switch (this.selectedWoD) {
        case '月曜日':
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
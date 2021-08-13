<!--グラフ表示コンポーネント-->
<template>
  <v-row style="height: 500px" justify="center" align-content="center">
    <v-col cols="8">
      <v-card class="justify-center">
        <v-col cols="12" sm="3">
            <h3>人数予想</h3>
          <v-select :items="items" label=" day of week" outlined></v-select>
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
              auto-draw
            ></v-sparkline>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import monday from "../store/index.js"
import tuesday from "../store/index.js"
import wednesday from "../store/index.js"
import thursday from "../store/index.js"
import friday from "../store/index.js"
import saturday from "../store/index.js"
import sunday from "../store/index.js"
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
    width: 2,
    radius: 0,
    padding: 8,
    lineCap: 'round',
    gradient: gradients[5],
    value: monday,
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'bar',
    autoLineWidth: false,
    items: [
      '月曜日',
      '火曜日',
      '水曜日',
      '木曜日',
      '金曜日',
      '土曜日',
      '日曜日',
    ],
  }),
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
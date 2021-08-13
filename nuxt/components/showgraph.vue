<!--グラフ表示コンポーネント-->
<template>
  <v-row style="height: 450px" justify="center" align-content="center">
    <v-col cols="6">
      <v-card elecaiton="50" class="justify-center">
        <v-row align="center">
          <v-col cols="12">
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
    padding: 5,
    lineCap: 'round',
    gradient: gradients[5],
    value: [0, 2, 5, 9, 5, 20, 3, 5, 0, 0, 1, 8, 2, 9, 0],
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'trend',
    autoLineWidth: false,
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
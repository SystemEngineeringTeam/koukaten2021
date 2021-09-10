<!--グラフ表示コンポーネント-->
<template>
  <v-row justify="center" align-content="center">
    <v-col :cols="width">
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
            <v-sparkline
              :value="value"
              :gradient="gradients[5]"
              padding="8"
              line-width="5"
              stroke-linecap="round"
              gradient-direction="top"
              fill="false"
              type="bar"
              auto-line-width="false"
              show-labels="true"
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
    gradients,
    selectedDoW: '',
    dow: ['日曜日', '月曜日', '火曜日', '水曜日', '木曜日', '金曜日', '土曜日'],
    graphKey: Math.random(),
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
    width: function () {
      switch (this.$vuetify.breakpoint.name) {
        case 'xs':
          return 11
        case 'sm':
          return 8
        case 'md':
          return 8
        case 'lg':
          return 8
        case 'xl':
          return 8
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
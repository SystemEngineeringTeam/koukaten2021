<!--人数表示コンポーネント-->
<template>
  <v-row justify="center">
    <v-col :cols="cardwidth">
      <v-card outlined>
        <v-container>
          <v-row justify="center">
            <v-card-title class="font-weight-bold" style="font-size: 1.25rem"
              >最終更新時刻：{{ nowTime }}</v-card-title
            >
          </v-row>
          <v-row justify="center">
            <!--人数表示部分-->
            <v-card-text
              class="text-center font-weight-bold"
              style="font-size: 2rem"
              >現在：
              <span style="font-size: 4rem">{{
                this.$store.state.people
              }}</span>
              人</v-card-text
            >
          </v-row>
          <v-row justify="center">
            <!-- 画像表示部分 -->
            <v-card-title>
              <v-img v-bind:src="meterImage" max-width="250px"></v-img>
            </v-card-title>
          </v-row>
        </v-container>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'ShowNum',

  props: ['cardwidth'],

  data: () => ({
    //画像用配列
    meterArray: [
      require('@/assets/images/meter_0.png'),
      require('@/assets/images/meter_01.png'),
      require('@/assets/images/meter_02.png'),
      require('@/assets/images/meter_03.png'),
      require('@/assets/images/meter_04.png'),
      require('@/assets/images/meter_05.png'),
      require('@/assets/images/meter_06.png'),
      require('@/assets/images/meter_07.png'),
      require('@/assets/images/meter_08.png'),
      require('@/assets/images/meter_09.png'),
      require('@/assets/images/meter_10.png'),
    ],
  }),

  computed: {
    meterImage: function () {
      if (this.$store.state.people > 10) {
        return this.meterArray[10]
      }
      return this.meterArray[this.$store.state.people]
    },

    nowTime: function () {
      return this.$store.state.date.slice(-8)
    },
  },
}
</script>

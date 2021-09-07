<!--人数表示コンポーネント-->
<template>
  <v-row
    justify="center"
    align-content="center"
  >
    <v-col cols="6">
      <v-card elecaiton="50" class="justify-center" outlined>
        <v-row align="center">
          <v-col cols="12">
            <v-card-text class="text-center">
              <h3>最終更新時刻：{{ this.$store.state.date }}</h3>
              <br /><br />
              <!--人数表示部分-->
              <h2>現在：</h2>
              <h1>{{ this.$store.state.people }}</h1>
              <h2>人</h2>
            </v-card-text>
            <br />
            <!-- 画像表示部分 -->
            <v-row justify="center">
              <v-img v-bind:src="meterImage" max-width="250px"></v-img>
            </v-row>
            <br />
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'ShowNum',
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
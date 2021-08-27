<!--人数表示コンポーネント-->
<template>
  <v-row style="height: 700px" justify="center" align-content="center" class="bg">
    <v-col cols="6">
      <v-card elecaiton="50" class="justify-center">
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
              <div v-if="this.$store.state.people　=== 0">
                <v-img v-bind:src="meter_0" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 1">
                <v-img v-bind:src="meter_01" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 2">
                <v-img v-bind:src="meter_02" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 3">
                <v-img v-bind:src="meter_03" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 4">
                <v-img v-bind:src="meter_04" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 5">
                <v-img v-bind:src="meter_05" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 6">
                <v-img v-bind:src="meter_06" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 7">
                <v-img v-bind:src="meter_07" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 8">
                <v-img v-bind:src="meter_08" max-width="250px"></v-img>
              </div>
              <div v-else-if="this.$store.state.people === 9">
                <v-img v-bind:src="meter_09" max-width="250px"></v-img>
              </div>
              <div v-else>
                <v-img v-bind:src="meter_10" max-width="250px"></v-img>
              </div>
            </v-row>
            <br />
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</v-content>
</template>

<script>
export default {
  name: 'ShowNum',
  data() {
    return {
      //画像用変数
      meter_0: require('@/assets/images/meter_0.png'),
      meter_01: require('@/assets/images/meter_01.png'),
      meter_02: require('@/assets/images/meter_02.png'),
      meter_03: require('@/assets/images/meter_03.png'),
      meter_04: require('@/assets/images/meter_04.png'),
      meter_05: require('@/assets/images/meter_05.png'),
      meter_06: require('@/assets/images/meter_06.png'),
      meter_07: require('@/assets/images/meter_07.png'),
      meter_08: require('@/assets/images/meter_08.png'),
      meter_09: require('@/assets/images/meter_09.png'),
      meter_10: require('@/assets/images/meter_10.png'),
    }
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
.bg {
  top: 0;
  left: 0;
  background-size: cover;
  background-image: url("~@/assets/images/sysken.png");
}
</style>
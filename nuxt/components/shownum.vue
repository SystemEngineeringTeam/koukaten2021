<!--人数表示コンポーネント-->
<template>
  <v-row style="height: 450px" justify="center" align-content="center">
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
              <v-img v-bind:src="image_src" max-width="250px"></v-img>
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
  data() {
    return {
      //画像用変数
      image_src: require('@/assets/meter.png'),
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
</style>
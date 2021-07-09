<!--人数表示コンポーネント-->
<template>
  <v-row style="height: 450px" justify="center" align-content="center">
    <v-col cols="6">
      <v-card elecaiton="50" class="justify-center">
        <v-row align="center">
          <v-col cols="12">
            <v-card-text class="text-center">
              <h3>最終更新時刻：ｎ月ｎ日ｎ時ｎ分</h3><br><br>
              <!--人数表示部分-->
              <h2>現在：</h2>
              <h1>X</h1>
              <h1>{{ this.$store.state.tasks }}</h1>
              <h2>人</h2>
            </v-card-text>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'ShowNum',
}
</script>

<style scoped>
h2{
  display: inline;
  font-size: 32px;
}
h1{
  display: inline;
  font-size: 64px;
}
h3{
  font-size: 20px;
}
</style>
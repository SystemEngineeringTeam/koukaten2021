<!--人数表示コンポーネント-->
<template>
  <div>
    <v-row
      style="height: 450px"
      justify="center"
      align-content="center"
    >
      <v-card
        elecaiton="50"
        width="500"
        height="250"
      >
        <v-responsive :aspect-raito="16 / 9">
          <div class="main">
            <!--人数表示部分-->
            <h2>現在：</h2>
            <h1>{{ this.$store.state.tasks }}</h1>
            <h2>人</h2>
          </div>
        </v-responsive>
      </v-card>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'ShowNum',
}
</script>

<style scoped>
h1 {
  display: inline;
}
h2 {
  display: inline;
}
.main {
  text-align: center;
}
</style>
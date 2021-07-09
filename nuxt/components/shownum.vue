<!--人数表示コンポーネント-->
<template>
  <v-row style="height: 450px" justify="center" align-content="center">
    <v-col cols="6">
      <v-card elecaiton="50" width="500" class="justify-center">
        <v-row align="center">
          <v-col cols="12">
            <v-card-text class="text-center">
              <!--人数表示部分-->
              <h2>現在：</h2>
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
h1 {
  display: inline;
}
h2 {
  display: inline;
}
</style>
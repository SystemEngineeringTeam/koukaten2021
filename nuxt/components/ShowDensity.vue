<template>
  <v-container fluid>
    <v-row justify="space-around">
      <v-col cols="7">
        <v-img v-bind:src="image_src"></v-img>
      </v-col>
    </v-row>
  </v-container>
  </div>
</template>

<script>
export default {
  name: 'ShowDensity',
  data() {
    return {
      image_src: require('@/assets/sysken.png'),
    }
  },
}
</script>

<style scoped>
</style>
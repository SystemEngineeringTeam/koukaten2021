<!--グラフデータコンポーネント-->
<template>
  <canvas id="chart"></canvas>
</template>

<script>
import Chart from 'chart.js/auto'

export default {
  props: ['labels', 'datasets', 'options'],
  data: () => ({
    myChart: null,
  }),
  methods: {
    renderChart: function () {
      // グラフを生成
      let ctx = document.getElementById('chart')

      this.myChart = new Chart(ctx, {
        type: 'bar',
        data: {
          labels: this.labels, // propsでもらったものをぶち込む1
          datasets: this.datasets, // propsでもらったものをぶち込む2
        },
        options: this.options, // propsでもらったものをぶち込む3
      })
    },
  },
  watch: {
    datasets: {
      handler() {
        this.myChart.destroy() // 再描画の前にdestroy
        this.renderChart()
      },
      deep: true,
    },
  },
  mounted() {
    this.renderChart() // 描画
  },
}
</script>

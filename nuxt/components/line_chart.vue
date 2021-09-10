
<!--グラフデータコンポーネント-->
<script>
import { Line,mixins } from 'vue-chartjs';
const { reactiveProp } = mixins

export default {
  extends: Line,
  name: 'chart',
  mixins: [reactiveProp],
  props: ['chartvalue']
  data:() => ({
    chartdata: {
      labels: ['1', '2', '3', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13', '14', '15', '16', '17', '18', '19', '20', '21', '22', '23', '24'],
      datasets: [
        {
          label: '平均人数',
          data: this.chartvalue,
          borderColor: 'rgba(255, 146, 51, 1)',
          backgroundColor: 'rgba(255, 146, 51, 0.2)',
        }
      ]
    },
    options: {
     scales: {
         yAxes: [{
             ticks: {
                 beginAtZero: true,
                 callback: function(value, index, values) {
                     return setDegreeCelsius(value, index, values);
                 }
             }
         }],
          xAxes: [{
            ticks:{
              max: 10
            }
         }]
     },
     responsive: true,
     maintainAspectRatio: false
   }
  }),
  mounted () {
    this.renderChart(this.chartdata, this.options)
  }
}
</script>
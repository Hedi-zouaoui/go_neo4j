import * as echarts from 'echarts/dist/echarts.js';
var dom = document.getElementById('chart-container');
var myChart = echarts.init(dom, null, { renderer: 'canvas', useDirtyRect: false
});
var app = {};
var ROOT_PATH = 'https://echarts.apache.org/examples';
var option;

myChart.showLoading();
$.get(ROOT_PATH + 'records.json', function (data) {
  myChart.hideLoading();
  data.children.forEach(function (datum, index) {
    index % 2 === 0 && (datum.collapsed = true);
  });
  myChart.setOption(
    (option = {
      tooltip: {
        trigger: 'item',
        triggerOn: 'mousemove'
      },
      series: [
        {
          type: 'tree',
          data: [data],
          top: '1%',
          left: '7%',
          bottom: '1%',
          right: '20%',
          symbolSize: 7,
          label: {
            position: 'left',
            verticalAlign: 'middle',
            align: 'right',
            fontSize: 9
          },
          leaves: {
            label: {
              position: 'right',
              verticalAlign: 'middle',
              align: 'left'
            }
          },
          emphasis: {
            focus: 'descendant'
          },
          expandAndCollapse: true,
          animationDuration: 550,
          animationDurationUpdate: 750
        }
      ]
    })
  );
});

if (option && typeof option === 'object') { myChart.setOption(option);
}

window.addEventListener('resize', myChart.resize);
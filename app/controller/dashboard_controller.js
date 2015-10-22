var dashboard = dashboard || {};

(function() {
  'use strict';

  dashboard.controller = function() {
    this.options = m.prop({
      chart: {
        backgroundColor: 'transparent',
        renderTo: 'chart-block',
        type: 'line'
      },
      credits: {
        enabled: false
      },
      legend: {
        borderWidth: 0,
        itemStyle: {
          color: '#818a91'
        },
        layout: 'horizontal'
      },
      plotOptions: {
        series: {
          marker: {
            enabled: false
          },
          states: {
            hover: {
              enabled: false
            }
          }
        }
      },
      subtitle: {},
      title: {
        text: null
      },
      tooltip: {
        enabled: false
      },
      xAxis: {
        labels: {
          style: {
            color: '#818a91'
          }
        }
      },
      yAxis: {
        labels: {
          style: {
            color: '#818a91'
          },
          format: '{value:.1f}'
        },
        max: 1.0,
        min: 0.0,
        title: ''
      }
    });
  };

  dashboard.chart = m.prop({});
  
  dashboard.plotter = function(ctrl) {
    return function(element, isInitialized) {
      if (!isInitialized) {
        m.startComputation();
        dashboard.chart(new Highcharts.Chart(ctrl.options()));
        m.endComputation();
      }
    };
  };
})();

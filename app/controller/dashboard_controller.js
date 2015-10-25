var dashboard = dashboard || {};

(function() {
  'use strict';

  dashboard.controller = function() {
    this.hypervisors = m.request({
      method: 'GET',
      url: '/api/cluster/hypervisors',
      initialValue: []
    });

    this.options = {
      chart: {
        backgroundColor: 'transparent',
        renderTo: 'chart-block',
        type: 'line'
      },
      colors: ['#003744', '#00596e', '#00728c', '#0090a3', '#00adb5'],
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
      series: [],
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
    };
  };

  dashboard.chart = {};

  dashboard.plotter = function(ctrl) {
    return function(element, isInitialized) {
      if (!isInitialized) {
        m.startComputation();

        var arr = Array.apply(null, Array(100)).map(function() {
          return 0.0;
        });
        ctrl.hypervisors().map(function(hypervisor) {
          hypervisor.virtual_machines.map(function(machine) {
            ctrl.options.series.push({
              name: machine.name,
              data: arr
            });
          });
        });
        dashboard.chart = new Highcharts.Chart(ctrl.options);

        m.endComputation();
      }
    };
  };
})();

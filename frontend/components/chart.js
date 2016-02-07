var chart = chart || {};

(function() {
  'use strict';

  chart.highcharts = {};

  chart.width = 100;

  chart.options = {
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
      tickInterval: 0.1,
      title: ''
    }
  };

  chart.initialize = function(ctrl) {
    var arr = Array.apply(null, Array(chart.width)).map(function() {
      return 0.0;
    });

    ctrl.hypervisors().map(function(hypervisor) {
      hypervisor.virtual_machines.map(function(machine) {
        chart.options.series.push({
          name: machine.name,
          data: arr
        });
      });
    });

    chart.highcharts = new Highcharts.Chart(chart.options);
  };

  chart.update = function(str) {
    var loads = str.split(',');
    var diff = chart.highcharts.series.length - loads.length;

    if (diff > 0) {
      loads = loads.concat(Array.apply(null, Array(diff)).map(function() {
        return '0.0';
      }));
    }

    loads.map(function(load, i) {
      chart.highcharts.series[i].addPoint(Number(load), true, true);
    });
    // chart.highcharts.redraw();
  };
})();

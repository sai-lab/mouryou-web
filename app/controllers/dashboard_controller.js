var chart = chart || {};
var realtime = realtime || {};
var dashboard = dashboard || {};

(function() {
  'use strict';

  dashboard.controller = function() {
    this.hypervisors = m.request({
      method: 'GET',
      url: '/api/cluster/hypervisors',
      initialValue: []
    });
  };

  dashboard.plotter = function(ctrl) {
    return function(element, isInitialized) {
      if (!isInitialized) {
        m.startComputation();
        chart.initialize(ctrl);
        realtime.initialize(new WebSocket(wsURL))
        m.endComputation();
      }
    };
  };
})();

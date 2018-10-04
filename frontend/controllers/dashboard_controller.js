var chart = chart || {};
var dashboard = dashboard || {};
var realtime = realtime || {};

(function() {
  'use strict';

  dashboard.controller = function() {
    this.vendors = m.request({
      method: 'GET',
      url: '/api/cluster/vendors',
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

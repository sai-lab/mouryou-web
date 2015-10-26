var chart = chart || {};
var realtime = realtime || {};

(function() {
  'use strict';

  realtime.initialize = function(ws) {
    ws.onopen = function(e) {
      console.log('Open: ' + ws.url);
    };

    ws.onclose = function(e) {
      console.log('Close: ' + ws.url);
    };

    ws.onerror = function(e) {
      console.log('Error: ' + e);
    };

    ws.onmessage = function(e) {
      console.log('Message: ' + e.data);

      if (e.data.match(/^Loads:\s/g)) {
        chart.update(e.data.replace(/^Loads:\s/g, ''));
      }
    };
  };
})();

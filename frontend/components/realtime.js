var chart = chart || {};
var power = power || {};
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
      } else if (e.data.match(/^Booting\sup:\s/g)) {
        power.bootingUp(e.data.replace(/^Booting\sup:\s/g, ''));
      } else if (e.data.match(/^Booted\sup:\s/g)) {
        power.bootedUp(e.data.replace(/^Booted\sup:\s/g, ''));
      } else if (e.data.match(/^Shutting\sdown:\s/g)) {
        power.shuttingDown(e.data.replace(/^Shutting\sdown:\s/g, ''));
      } else if (e.data.match(/^Shutted\sdown:\s/g)) {
        power.shuttedDown(e.data.replace(/^Shutted\sdown:\s/g, ''));
      }
    };
  };
})();

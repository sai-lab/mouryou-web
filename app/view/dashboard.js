var dashboard = dashboard || {};

(function() {
  'use strict';

  dashboard.view = function(ctrl) {
    document.title = 'mouryou';

    return [
      m('.chart', [
        m('#chart-block', {
          config: dashboard.plotter(ctrl)
        })
      ])
    ];
  };
})();

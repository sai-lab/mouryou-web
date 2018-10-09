var anime = anime || {};
var dashboard = dashboard || {};

const operationType = new Map([
  ['booting up', 'bg-booting-up'],
  ['booted up', 'bg-mouryou'],
  ['shutting down', 'bg-shutting-down'],
  ['shutted down', 'bg-default'],
]);

(function() {
  'use strict';

  dashboard.view = function(ctrl) {
    document.title = 'mouryou';
    var i = 0;

    return m('.row', [
      m('.col-md-9', [
        m('.chart', [
          m('#chart-block.m-a-05', {
            config: dashboard.plotter(ctrl)
          })
        ])
      ]),
      m('.col-md-3', [
        m('ul.list-group.small', [
          ctrl.vendors().map(function(vendor) {
            return [
              m('li.list-group-item.bg-mouryou.text-white.border-mouryou.p-a-05.animated.fadeInRight', [
                m('i.fa.fa-fw.fa-server.m-r-05'),
                m('span.bold', vendor.name)
              ]),
              Object.entries(vendor.virtual_machines).map(function (machine) {
                var view =  m('li.list-group-item.p-a-05.animated.fadeInRight', {
                  id: 'machine-list-' + machine[1].name,
                  style: animationDelay(i),
                  config: anime.initialize
                }, [
                  m('i.fa.fa-fw.fa-hdd-o.m-r-05'),
                  m('span.bold', machine[1].name),
                  m('span.label.label-pill.bg-default.text-white.pull-right', {
                    id: 'machine-label-' + machine[1].name,
                    class: operationType.get(machine[1].operation)
                  }, [
                    m('i.fa.fa-fw.fa-power-off')
                  ])
                ]);
                i += 1;
                return view;
              })
            ];
          })
        ])
      ])
    ]);
  };
})();

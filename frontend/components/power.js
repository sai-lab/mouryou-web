var anime = anime || {};
var power = power || {};

(function() {
  'use strict';

  power.classes = [
    'bg-default',
    'bg-mouryou',
    'bg-booting-up',
    'bg-shutting-down',
  ]

  power.bootingUp= function(str) {
    anime.shake(document.getElementById('machine-list-' + str));

    var label = document.getElementById('machine-label-' + str);
    power.reset(label);
    label.classList.add('bg-booting-up');
  };

  power.bootedUp = function(str) {
    anime.shake(document.getElementById('machine-list-' + str));

    var label = document.getElementById('machine-label-' + str);
    power.reset(label);
    label.classList.add('bg-mouryou');
  };

  power.shuttingDown = function(str) {
    anime.shake(document.getElementById('machine-list-' + str));

    var label = document.getElementById('machine-label-' + str);
    power.reset(label);
    label.classList.add('bg-shutting-down');
  };

  power.shuttedDown = function(str) {
    anime.shake(document.getElementById('machine-list-' + str));

    var label = document.getElementById('machine-label-' + str);
    power.reset(label);
    label.classList.add('bg-default');
  };

  power.reset = function(element) {
    power.classes.map(function(c) {
      element.classList.remove(c);
    });
  };
})();

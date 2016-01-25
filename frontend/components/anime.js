var anime = anime || {};

(function() {
  'use strict';

  anime.events = [
    'webkitAnimationEnd',
    'mozAnimationEnd',
    'MSAnimationEnd',
    'oanimationend',
    'animationend'
  ];

  anime.classes = [
    'animated',
    'fadeInRight',
    'shake'
  ];

  anime.initialize = function(element, isInitialized) {
    if (!isInitialized) {
      m.startComputation();
      anime.reset(element)
      m.endComputation();
    }
  };

  anime.reset = function(element) {
    anime.events.map(function(event) {
      element.addEventListener(event, function() {
        anime.classes.map(function(c) {
          element.classList.remove(c);
        });
      });
    });
  };

  anime.shake = function(element) {
    element.classList.add('animated');
    element.classList.add('shake');
  };
})();

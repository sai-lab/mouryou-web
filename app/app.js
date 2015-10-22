var dashboard = dashboard || {};

var extractStatus = function(xhr, xhrOptions) {
  return xhr.status !== 200 ? xhr.status : xhr.responseText;
};

(function(window) {
  'use strict';

  m.route(document.getElementById('app-body'), '/', {
    '/': dashboard
  });

  document.getElementById('app-title').onclick = function(e) {
    e.preventDefault();
    m.route('/');
  };
})(window);

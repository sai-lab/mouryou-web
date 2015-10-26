## mouryou-web

![screen.gif](https://raw.githubusercontent.com/sai-lab/mouryou-web/master/screen.gif)

#### Requirements

  - [Golang](https://golang.org/) : 1.5
  - [Mithril.js](https://lhorie.github.io/mithril/) : 0.2.0
  - [Bootstrap](http://v4-alpha.getbootstrap.com/) : v4.0.0-alpha
  - [Font Awesome](http://fontawesome.io/) : 4.4.0
  - [Animate.css](http://daneden.github.io/animate.css/) : 3.4.0
  - [Highcharts](http://www.highcharts.com/) : 4.1.9

#### Installation

    $ git clone git://github.com/sai-lab/mouryou-web.git
    $ cd mouryou-web
    $ make gom link
    $ make build
    $ bin/mouryou-web

### Highcharts

`/assets/javascripts/highcharts.js` is downloaded from [download builder](http://www.highcharts.com/download).  
The enabled options are shown in below.

  - Compile code
  - Adapters : Standalone framework
  - Core : Core
  - Dynamics and interaction : Dynamics
  - Chart and serie types : Line

#### License

mouryou-web is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/mouryou-web/master/LICENSE).

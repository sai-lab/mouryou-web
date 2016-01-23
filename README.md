## mouryou-web

![screen.gif](https://raw.githubusercontent.com/sai-lab/mouryou-web/master/screen.gif)

#### Requirements

  - [Golang](https://golang.org/) : 1.5.3
  - [Mithril.js](https://lhorie.github.io/mithril/) : 0.2.2-rc.1
  - [Bootstrap](http://v4-alpha.getbootstrap.com/) : 4.0.0-alpha.2
  - [Font Awesome](http://fontawesome.io/) : 4.5.0
  - [Animate.css](http://daneden.github.io/animate.css/) : 3.5.1
  - [Highcharts](http://www.highcharts.com/) : 4.2.1

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
  - Core : Core
  - Dynamics and interaction : Dynamics
  - Chart and serie types : Line

#### License

mouryou-web is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/mouryou-web/master/LICENSE).

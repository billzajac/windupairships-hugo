Hugo powered static site generator content for https://windupairships.com
Repo: https://github.com/billzajac/windupairships-hugo

[![Netlify Status](https://api.netlify.com/api/v1/badges/b1e4b805-fa98-42e8-8e8b-337236abbdae/deploy-status)](https://app.netlify.com/sites/happy-hodgkin-40cf1d/deploys)

Primary Tech
---------------
* https://gohugo.io/
* https://app.netlify.com/
* https://github.com/
* https://analytics.google.com/

The SVG logo was made quickly with
* https://inkscape.org/

The favicon with
* https://favicon.io/favicon-converter/

Theme
* https://github.com/CaiJimmy/hugo-theme-stack

    git clone https://github.com/CaiJimmy/hugo-theme-stack/ themes/hugo-theme-stack
    git submodule add https://github.com/CaiJimmy/hugo-theme-stack/ themes/hugo-theme-stack

### Hugo commands

    git submodule init
    git submodule update

    hugo new post/hello-beautiful/index.md
    hugo server -D

### Images

* https://scop.io/search


### Netlify Function

* https://docs.netlify.com/functions/build/?fn-language=go

Create with

    ntl functions:create --name tao

Get deps

    go mod tidy

Test service with

    netlify functions:serve

Serves from /.netlify/functions/tao

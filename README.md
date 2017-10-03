# Readable 

A bookmarklet that makes pages readable :link: :book:.

This is like a self-hosted version of [Pocket](https://getpocket.com/), or [Firefox Reader View](https://support.mozilla.org/en-US/kb/firefox-reader-view-clutter-free-web-pages), or any other extension that helps you read an article on the web when using a Desktop browser. But, instead of a browser extension that you have to install, its just a bookmark you can keep on your toolbar.

Websites are parsed with either [the free Mercury Web Parser API](https://mercury.postlight.com/web-parser/) or a self-hosted version of [Mozilla's Readability program](https://github.com/mozilla/readability).

Demo
=====

Try it out at [readable.schollz.com](https://readable.schollz.com). 


<center>
<img src="http://i.imgur.com/k5ArA0A.gif" alt="Example of parsing a website">
</center>


Getting Started
===============

## Install

If you have Go installed:

```
go get github.com/schollz/readable
```

## Run (using Mercury Web Parser API)

First [get a Mercury Web Parser API Key](https://mercury.postlight.com/web-parser/).

Then use

```shell
$ readable -key YOUR_API_KEY
```

## Run (using self-hosted Web Parser)

Instead of using the Mercury Web Parser API, you can use [Mozilla's Readability](https://github.com/mozilla/readability) running on your own machine.

You will need to [install Docker](https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/). Then to build the Docker image, use

```
$ cd $GOPATH/src/github.com/schollz/readable
$ docker build -t readable .
```

Then you can run `readable` without specifying a `key` to use the self-hosted Docker image for parsing websites:

```
$ readable
```

You can also just use the image pretty easily:

```
$ docker run --rm -v `pwd`:/data -t readable URL
```

which will result in a file `some_hash.json` which contains the results.


License
=======

MIT

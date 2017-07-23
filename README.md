# Readable 

A bookmarklet that makes pages readable :link: :book:.

This is like a self-hosted version of [Pocket](https://getpocket.com/), or [Firefox Reader View](https://support.mozilla.org/en-US/kb/firefox-reader-view-clutter-free-web-pages), or any other extension that helps you read an article on the web when using a Desktop browser.

![Example of parsing website](http://i.imgur.com/k5ArA0A.gif)

Demo
=====

Try it out at [readable.schollz.com](https://readable.schollz.com). 



Getting Started
===============

## Install

If you have Go installed:

```
go get github.com/schollz/readable
```

## Run

First [get a Mercury Web Parser API Key](https://mercury.postlight.com/web-parser/).

Then use

```shell
$ readable -key YOUR_API_KEY
```

## Self-hosted Readability

Instead of using the Mercury Web Parser API, you can use [Mozilla's Readability](https://github.com/mozilla/readability) running on your own machine.

You will need to install Docker and build the image.

```
$ docker build -t readable .
```

Then you can run `readable` without the key to use the self-hosted Docker image:

```
$ readable
```


License
=======

MIT
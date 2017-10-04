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


# Getting Started


## Install

First install *readable*. Download Go and run:

```
go get github.com/schollz/readable
```

To run you can use either Docker (to self-host the web parser that uses [Mozilla's Readability](https://github.com/mozilla/readability)) or use Mercury API (which is free).

## Run (Docker)

First pull the latest Docker image.

```
$ docker pull schollz/readable
```

_Note_: Alternatively you can build it yourself using `docker build -t readable .` in the main source). 

Then to run, just do

```
$ readable
```

_Note_: If you built yourself, add `-docker readable` to specify your own image.

## Run (using Mercury Web Parser API)

First [get a Mercury Web Parser API Key](https://mercury.postlight.com/web-parser/).

Then use

```shell
$ readable -key YOUR_API_KEY
```

# Advanced usage

## Read articles from the command line

You can use the Docker image to directly read articles from the command line:

```
$ docker run --rm -t schollz/readable URL | more
```

where `URL` is the URL of some article that you want to read.

## Download readable data to computer

You can use the Docker image to download the parsed contents into a json file:

```
$ docker run --rm -v `pwd`:/data -t schollz/readable URL data.json
```

where `URL` is the URL of some article that you want to read. This will result in a file `data.json` which contains the results.

License
=======

MIT

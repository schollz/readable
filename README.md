# Readable 

Making web pages readable in a browser and in the command line :link: :book:.

This is like a self-hosted version of [Pocket](https://getpocket.com/), or [Firefox Reader View](https://support.mozilla.org/en-US/kb/firefox-reader-view-clutter-free-web-pages), or any other extension that helps you read an article on the web when using a Desktop browser. But, instead of a browser extension that you have to install, its just a bookmark you can keep on your toolbar or a single line of bash you can run at the terminal.

Websites are parsed with either [the free Mercury Web Parser API](https://mercury.postlight.com/web-parser/) or a self-hosted version of [Mozilla's *readability* package](https://github.com/mozilla/readability). The *readability* package was trasnformed into a Docker image that automatically performs some UTF-8 conversions and tidying.

# Demo

Try it out at [readable.schollz.com](https://readable.schollz.com). 

<center>
<img src="http://i.imgur.com/k5ArA0A.gif" alt="Example of parsing a website">
</center>

# Quickstart

First [download the latest release of *readable* for your OS](https://github.com/schollz/readable/releases/latest). Alternatively, if you have Go installed you can do `go get github.com/schollz/readable`.

You can run *readable* with or without Docker.

## with Docker

```shell
$ docker pull schollz/readable
$ ./readable
```

## without Docker 

Get `YOUR_API_KEY` Mercury Web Parser API Key [from here (its free)](https://mercury.postlight.com/web-parser/).

```shell
$ readable -key YOUR_API_KEY
```

# Advanced usage

The Docker image in this repo allows you to manipulate websites into readable ones. You can do some neat things like the following:

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

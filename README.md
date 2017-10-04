# Readable 

Making web pages readable in a browser and in the command line :link: :book:.

This is like a self-hosted version of [Pocket](https://getpocket.com/), or [Firefox Reader View](https://support.mozilla.org/en-US/kb/firefox-reader-view-clutter-free-web-pages), or any other extension that helps you read an article on the web when using a Desktop browser. But, instead of a browser extension that you have to install, its just a bookmark you can keep on your toolbar or a single line of bash you can run at the terminal.

Websites are parsed with either [the free Mercury Web Parser API](https://mercury.postlight.com/web-parser/) or a self-hosted version of [Mozilla's *readability* package](https://github.com/mozilla/readability). The *readability* package was trasnformed into a Docker image that automatically performs some UTF-8 conversions and tidying.

# Demo

Try it out at [readable.schollz.com](https://readable.schollz.com). 

[![Readable example](https://user-images.githubusercontent.com/6550035/31201819-59d78922-a91d-11e7-8bc5-b9b2668d0123.png)](https://readable.schollz.com)

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
$ URL=http://www.cnn.com/2017/10/03/world/nobel-physics-prize-2017/index.html
$ docker run --rm -t schollz/readable $URL | more

----------------------------------------------------------
Nobel Prize in Physics goes to 'black hole telescope' trio
----------------------------------------------------------

Story highlights

-   The development proves Einstein's prediction of gravitational waves
-   More than 1,000 people worked on the technology over four decades

(CNN)The 2017 Nobel Prize in Physics has been awarded to Rainer Weiss,
Barry C. Barish and Kip S. Thorne for their detection of gravitational
waves, a development scientists believe could give vital clues to the
origins of the universe.
...
```

## Download readable data to computer

You can use the Docker image to download the parsed contents into a json file:

```shell
$ URL=http://www.cnn.com/2017/10/03/world/nobel-physics-prize-2017/index.html
$ docker run --rm -v `pwd`:/data -t schollz/readable $URL data.json
$ cat data.json | jq .title
"Nobel Prize in Physics goes to 'black hole telescope' trio"
```

where `URL` is the URL of some article that you want to read. This will result in a file `data.json` which contains the results.

License
=======

MIT

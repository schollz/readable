FROM ubuntu:17.04

RUN apt-get update
RUN apt-get install -y wget
RUN wget https://github.com/htacg/tidy-html5/releases/download/5.4.0/tidy-5.4.0-64bit.deb
RUN dpkg --install tidy-5.4.0-64bit.deb
RUN rm -rf tidy-*

RUN apt-get install -y gcc
RUN wget http://www.it.uc3m.es/jaf/html2xhtml/downloads/html2xhtml-1.3.tar.gz
RUN tar -xvzf html2xhtml-1.3.tar.gz
WORKDIR html2xhtml-1.3
RUN ./configure 
RUN apt-get install -y make
RUN make
RUN make install

WORKDIR /root
RUN rm -rf html2xhtml*

RUN apt-get install -y pandoc
RUN apt-get install -y python3
RUN apt-get install -y nodejs
RUN apt-get install -y npm

RUN apt-get install -y git
RUN git clone https://github.com/mozilla/readability
WORKDIR /root/readability
RUN npm install
WORKDIR /root/readability/test

RUN apt-get install nodejs-legacy 

RUN echo "OK"
COPY run.py /root/readability/test/run.py

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENTRYPOINT ["/usr/bin/python3","run.py"]

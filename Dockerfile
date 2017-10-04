FROM node:slim

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

# Install Pandoc
RUN wget https://github.com/jgm/pandoc/releases/download/1.19.2.1/pandoc-1.19.2.1-1-amd64.deb
RUN dpkg --install pandoc*deb

# Install readability
RUN apt-get install -y zip
RUN wget https://github.com/mozilla/readability/archive/master.zip
RUN unzip master.zip
RUN rm master.zip
WORKDIR /root/readability-master
RUN npm install
WORKDIR /root/readability-master/test

COPY run.sh /root/readability-master/test/run.sh
RUN chmod +x /root/readability-master/test/run.sh

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

RUN apt-get purge -y make gcc
RUN apt-get autoremove -y

ENTRYPOINT ["./run.sh"]

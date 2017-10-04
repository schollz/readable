#!/bin/bash

rm -rf test-pages/foo
mkdir test-pages/foo

wget -q "$1" -O test-pages/foo/1.html
cat test-pages/foo/1.html | pandoc -f html -t html > test-pages/foo/2.html
tidy -q -utf8 -o test-pages/foo/3.html test-pages/foo/2.html
cat test-pages/foo/3.html | html2xhtml > test-pages/foo/source.html
node generate-testcase.js foo
chmod -R 0755 /data
cat test-pages/foo/expected-metadata.json  | grep title | sed 's/"//g' | sed 's/title://g' | sed 's/   //g' | sed 's/,//g' >> '{"title":"' >> /data/`echo -n "$1" | md5sum | awk '{print $1}'`.title
cat test-pages/foo/expected.html >> /data/`echo -n "$1" | md5sum | awk '{print $1}'`.html

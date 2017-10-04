#!/bin/bash

rm -rf test-pages/foo
mkdir test-pages/foo

wget -q "$1" -O test-pages/foo/1.html
cat test-pages/foo/1.html | pandoc -f html -t html > test-pages/foo/2.html
tidy -q -utf8 -o test-pages/foo/3.html test-pages/foo/2.html  > /dev/null 2>&1
cat test-pages/foo/3.html | html2xhtml > test-pages/foo/source.html
node generate-testcase.js foo 
V=`cat test-pages/foo/expected-metadata.json | jq .title | sed 's/^.\(.*\).$/\1/'`
s=$(printf "%-${#V}s" "-")
if [ "${#V}" -gt "70" ]; then 
	s=$(printf "%-70s" "-");
fi
if [ -z "$2" ]
then
	echo ""
	echo "${s// /-}"	
	echo $V | fold -w 70 -s
	echo "${s// /-}"	
	echo ""
	pandoc -f html -t plain test-pages/foo/expected.html
else
	chmod -R 0755 /data
	cp test-pages/foo/expected-metadata.json temp.json
	sed -i 's/title/content":"","title/g' temp.json
	echo '. as $file' > content.jq
	echo '| $json' >> content.jq
	echo '| (.content = $file + "\n" + .content)' >> content.jq
	jq -R -s --argfile json temp.json -f content.jq test-pages/foo/expected.html > temp2.json
	mv temp2.json /data/$2
fi



#!/bin/bash

# sudo apt-get install curl jq pandoc
# https://mercury.postlight.com/web-parser/ for API key
# ./html2text.sh APIKEY URL
curl -s -H "x-api-key: $1" "https://mercury.postlight.com/parser?url=$2" | jq '.content' | pandoc -f html -t plain | sed 's/\\n//g' | sed 's/_//g' | sed 's/\[\]//g' | sed ':1;s/([^)]*)//g;/(/{N;b1};' | tail -n +3 | head -n -2

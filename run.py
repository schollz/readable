# -*- coding: utf-8 -*-

import json
import os
import sys
import hashlib
import shutil

url = sys.argv[1]
print('-' * 30)
print(url)
print('-' * 30)
url_hash = hashlib.md5(url.encode('utf-8')).hexdigest()
print(url_hash)
os.mkdir('test-pages/{}'.format(url_hash))
os.system('wget -q "{}" -O test-pages/{}/1.html'.format(url, url_hash))
os.system(
    'cat test-pages/{url_hash}/1.html | pandoc -f html -t html > test-pages/{url_hash}/2.html'.format(url_hash=url_hash))
os.system(
    'tidy -q -utf8 -o test-pages/{url_hash}/3.html test-pages/{url_hash}/2.html'.format(url_hash=url_hash))
os.system(
    'cat test-pages/{url_hash}/3.html | html2xhtml > test-pages/{url_hash}/source.html'.format(url_hash=url_hash))
os.system('node generate-testcase.js {url_hash}'.format(url_hash=url_hash))

with open('test-pages/{}/expected-metadata.json'.format(url_hash), 'r') as f:
    data = json.loads(f.read())

payload = {}
payload['title'] = data['title']
payload['content'] = open(
    'test-pages/{}/expected.html'.format(url_hash), 'r').read()
with open('/data/' + url_hash + '.json', 'w') as f:
    f.write(json.dumps(payload))


os.system('chmod -R 0755 /data')

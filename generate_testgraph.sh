#!/bin/bash

if [[ -d "./db" ]]; then
  rm -r "./db"
fi

mkdir "./db"

cd "./db"

touch nodes.jf
touch edges.jf
touch links.jf
touch graphs.jf

echo '{"name": "a"}' >> nodes.jf
echo '{"name": "b"}' >> nodes.jf
echo '{"name": "c"}' >> nodes.jf
echo '{"name": "d"}' >> nodes.jf

echo '{"name": "ab", "directed": 0}' >> edges.jf
echo '{"name": "bc", "directed": 1}' >> edges.jf
echo '{"name": "cd", "directed": 0}' >> edges.jf
echo '{"name": "bd", "directed": 0}' >> edges.jf

echo '{"start": 0, "end": 1}' >> links.jf
echo '{"start": 1, "end": 2}' >> links.jf
echo '{"start": 2, "end": 3}' >> links.jf
echo '{"start": 1, "end": 3}' >> links.jf

echo '{"name": "test", "links": [0,1,2,3]}' >> graphs.jf

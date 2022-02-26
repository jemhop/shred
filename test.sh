#!/bin/bash

go build .

touch test.txt

./shredder -l 
./shredder -d test.txt 
./shredder -l
./shredder -r test.txt 
./shredder -l
cp test.txt test2.txt 
./shredder -s test.txt 
./shredder -d test2.txt 
./shredder -e test2.txt



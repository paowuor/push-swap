#! /bin/bash

for i in {1..100}
do
    ARG=$(shuf -i 1-100 -n 100)

    RESULT=$(./push-swap $ARG | ./checker $ARG)

    if [ "RESULT" != "OK" ]
    then
        echo "Test Failed"
        echo $ARG
        exit 1
    fi
done

echo "All tests passed"
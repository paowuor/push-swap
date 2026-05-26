#!/bin/bash

for i in {1..100}
do
    ARG=$(shuf -i 1-100 -n 100)

    OPS=$(./push-swap $ARG)

    RESULT=$(echo "$OPS" | ./checker $ARG)

    if [ "$RESULT" != "OK" ]
    then
        echo "Test Failed"
        echo
        echo "INPUT:"
        echo $ARG
        echo
        echo "OPERATIONS:"
        echo "$OPS"
        echo
        echo "CHECKER RESULT:"
        echo "$RESULT"
        exit 1
    fi
done

echo "All tests passed"
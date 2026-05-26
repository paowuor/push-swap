#!/bin/bash

ARG=$(shuf -i 1-100 -n 100)

COUNT=$(./push-swap $ARG | wc -l)

echo "Operations: $COUNT"

./push-swap $ARG | ./checker $ARG
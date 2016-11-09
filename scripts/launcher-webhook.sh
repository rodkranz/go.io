#!/bin/bash

while : ;
do
    curl "http://localhost:5000/webhook?txt=Joao-$(((RANDOM % 100) + 1))";
    sleep $((( RANDOM % 5 ) +1));
done

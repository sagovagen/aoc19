#!/bin/bash
pfile=$1
shift

echo 0 | ./p7.1 $pfile $1 | ./p7.1 $pfile $2 | ./p7.1 $pfile $3 | ./p7.1 $pfile $4 | ./p7.1 $pfile $5

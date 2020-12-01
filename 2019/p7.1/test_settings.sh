#!/bin/bash

maxthrust=0
params="00000"
for i1 in 0 1 2 3 4 ; do
  for i2 in 0 1 2 3 4 ; do
    [ $i1 -eq $i2 ] && continue
    for i3 in 0 1 2 3 4 ; do
      [ $i1 -eq $i3 -o $i2 -eq $i3 ] && continue
      for i4 in 0 1 2 3 4 ; do
        [ $i1 -eq $i4 -o $i2 -eq $i4 -o $i3 -eq $i4 ] && continue
        for i5 in 0 1 2 3 4 ; do
          [ $i1 -eq $i5 -o $i2 -eq $i5 -o $i3 -eq $i5 -o $i4 -eq $i5 ] && continue
          echo "Testing: $i1 $i2 $i3 $i4 $i5"
          thrust=$(./run_amplifiers.sh input.txt $i1 $i2 $i3 $i4 $i5)
          echo "Thrust: $thrust"
          if [ $thrust -gt $maxthrust ]; then
            maxthrust=$thrust
            params="$i1$i2$i3$i4$i5"
          fi
        done
      done
    done
  done
done
echo "Max thrust: $maxthrust"
echo "Max params: $params"

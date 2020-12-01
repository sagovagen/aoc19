package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
)

func printMap(amap [][]bool) {
  for r:=0;r<len(amap);r++ {
    for c:=0;c<len(amap[r]);c++ {
      if amap[r][c] {
        fmt.Printf("*")
      } else {
        fmt.Printf(".")
      }
    }
    fmt.Println()
  }
}

func dirString(dx, dy int) string {
  maxdelta := math.Max(math.Abs(float64(dx)), math.Abs(float64(dy)))
  if (maxdelta == 0) {
    fmt.Println("Error: maxdelta is zero: ", dx, dy)
  }
  return fmt.Sprintf("%f:%f",float64(dx)/maxdelta,float64(dy)/maxdelta)
}

func countVisible(amap [][]bool, x, y int) int {
  count := 0
  ny := len(amap)
  nx := len(amap[0])
  hmap := make(map[string]bool)

  for i:=0;i<ny;i++ {
    for j:=0;j<nx;j++ {
      if i==y && j==x {continue} // do not count me
      if !amap[i][j] {continue} // only count asteroids
      dir := dirString(i-y, j-x)
      //if (y==2 && x==3) { fmt.Println(dir) }
      if !hmap[dir] {
        count++
        hmap[dir]=true
      }
    }
  }
  return count
}

func main() {
  amap := [][]bool{}
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    str := scanner.Text()
    row := make([]bool, len(str))
    for i:=0;i<len(str);i++ {
      row[i] = (str[i] == '#')
    }
    amap = append(amap, row)
  }
  ny := len(amap)
  nx := len(amap[0])
  //fmt.Printf("Rows: %v, Cols: %v\n", ny, nx)

  //printMap(amap)

  maxcount := 0
  maxx := -1
  maxy := -1
  for y:=0;y<ny;y++ {
    for x:=0;x<nx;x++ {
      if amap[y][x] {
        c := countVisible(amap,x,y)
        if (c > maxcount) {
          maxx = x
          maxy = y
          maxcount = c
        }
      }
    }
  }
  fmt.Printf("Max count: %v for position: (%v,%v)\n",maxcount,maxx,maxy)
}

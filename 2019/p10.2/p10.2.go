package main

import (
    "fmt"
    "bufio"
    "os"
    "math"
    "sort"
)


type Point struct {
  x,y int
}
type Vector struct {
  p Point
  angle float64
  dist int
}
func angle(p1, p2 Point) float64 {
  var a float64
  d := delta(p1, p2)
  lower := (d.y > 0)
  left := (d.x < 0)
  if (d.y == 0) {
    if (d.x < 0) {
      a = math.Pi * 1.5
    } else {
      a = math.Pi * 0.5
    }
  } else if (d.x == 0) {
    if (d.y < 0) {
      a = 0
    } else {
      a = math.Pi
    }
  } else {
    a = -math.Atan(float64(d.x)/float64(d.y))

    if (lower) {
      a = a + math.Pi
    }
    if (a < 0) {
      a = a + math.Pi * 2
    }
    
    fmt.Println("left and lower:", left, lower)
  }
  a = 180 * (a / math.Pi)
  fmt.Printf("p: %v, dx: %v, dy: %v, angle: %v\n",p1,d.x,d.y,a)
  return a
}

func delta(p1, p2 Point) Point {
  return Point{x: p1.x-p2.x, y: p1.y-p2.y}
}
func dist(p1, p2 Point) int {
  d := delta(p1, p2)
  return int(math.Abs(float64(d.x))+math.Abs(float64(d.y)))
}
func newVector(origin, p Point) Vector {
  a := angle(p, origin)
  d := dist(p, origin)
  return Vector{p:p, angle: a, dist: d}
}

func printMap(amap [][]bool, station Point) {
  for y:=0;y<len(amap);y++ {
    for x:=0;x<len(amap[0]);x++ {
      if amap[y][x] {
        fmt.Printf("*")
      } else if y==station.y && x==station.x {
        fmt.Printf("X")
      } else {
        fmt.Printf(".")
      }
    }
    fmt.Println()
  }
}

func main() {
  var station Point
  amap := [][]bool{}
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    str := scanner.Text()
    row := make([]bool, len(str))
    for i:=0;i<len(str);i++ {
      row[i] = (str[i] == '#')
      if str[i]=='X' {
          station = Point{x: i, y: len(amap)}
      }
    }
    amap = append(amap, row)
  }
  ny := len(amap)
  nx := len(amap[0])
  fmt.Printf("Rows: %v, Cols: %v\n", ny, nx)
  fmt.Printf("Station: %v\n", station)

  asteroids := []Vector{}

  for y:=0;y<ny;y++ {
    for x:=0;x<nx;x++ {
      if amap[y][x] {
        p := Point{x: x, y: y}
        if p != station {
          asteroids = append(asteroids, newVector(station, p))
          //fmt.Println("Appending new asteroid: ", asteroids[len(asteroids)-1])
        }
      }
    }
  }

  // Sort for angle and then distance
  sort.Slice(asteroids, func(i,j int) bool {
    if asteroids[i].angle == asteroids[j].angle {
      return asteroids[i].dist < asteroids[j].dist
    }
    return asteroids[i].angle < asteroids[j].angle
  })

  // Set rotation order
  numHidden := 0
  for i:=0;i<len(asteroids); {
    angle := asteroids[i].angle
    order := 0
    asteroids[i].dist = order
    //fmt.Println("Checking asteroid angle: ",asteroids[i])
    for i++; i<len(asteroids) && angle==asteroids[i].angle; i++ {
      //fmt.Println("Found identical angle: ", asteroids[i])
      order++
      numHidden++
      asteroids[i].dist = order
    }
  }

  // Sort for rotation order then angle
  sort.Slice(asteroids, func(i,j int) bool {
    if asteroids[i].dist == asteroids[j].dist {
      return asteroids[i].angle < asteroids[j].angle
    }
    return asteroids[i].dist < asteroids[j].dist
  })

  for i:=0;i<len(asteroids);i++ {
    fmt.Println(i+1, asteroids[i])
  }

  printMap(amap, station)
}

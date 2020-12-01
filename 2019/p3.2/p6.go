package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)
type Position struct {
    x, y, dist int
}
func (p Position) move(dir byte) Position {
  //fmt.Println(p)
  switch(dir) {
  case 'R': return Position{x: p.x+1, y: p.y, dist: p.dist+1}
  case 'L': return Position{x: p.x-1, y: p.y, dist: p.dist+1}
  case 'U': return Position{x: p.x, y: p.y+1, dist: p.dist+1}
  case 'D': return Position{x: p.x, y: p.y-1, dist: p.dist+1}
  default: fmt.Println("Error parsing move direction: ", dir)
  }
  return p
}
func (p Position) ident() string {
  return fmt.Sprintf("%v:%v", p.x, p.y)
}
type Wire struct {
    name, dist int
}
var m = make(map[string]Wire)
var xing [100]Position
var numx int = 0

func doMove(wire int, p Position, move string) Position {
  d := move[0]
  n, err := strconv.ParseInt(move[1:], 10, 32)
  if (err != nil) {
    fmt.Println("Error parsing move: ", move)
  }
  //fmt.Printf("move: %c,%d\n",d,n)
  for i:=0;i<int(n);i++ {
    p=p.move(d)
    w, ok := m[p.ident()]
    if (ok && w.name != wire) {
      fmt.Println("Found xing: ",p, w.dist+p.dist)
      m[p.ident()] = Wire{name: wire, dist: w.dist + p.dist}
      xing[numx]=p; numx++
    } else {
      m[p.ident()] = Wire{name: wire, dist: p.dist}
    }
  }
  //fmt.Println(p)
  return p
}


func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  var wire int = 0
  for scanner.Scan() {
    wire++
    str := scanner.Text()
    moves := strings.Split(str, ",")
    num := len(moves)
    var p = Position{x:0, y:0, dist:0}
    for i:=0; i<num; i++ {
      p = doMove(wire, p, moves[i])
    }
    //fmt.Println("final position: ", p.ident())
  }
  if (numx == 0) {
    fmt.Println("No crossing found")
    os.Exit(1)
  }

  var p Position = xing[0]
  for i:=1;i<numx;i++ {
    if (m[xing[i].ident()].dist < m[p.ident()].dist) {
      p = xing[i]
    }
    //fmt.Printf("Crossing: %v, dist: %v\n",xing[i],m[xing[i].ident()])
  }
  fmt.Printf("Shortest crossing path: %v for position: %v\n",m[p.ident()].dist, p)
}

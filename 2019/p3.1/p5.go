package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "math"
)
type Position struct {
    x, y int
}
func (p Position) move(dir byte) Position {
  switch(dir) {
  case 'R': return Position{x: p.x+1, y: p.y}
  case 'L': return Position{x: p.x-1, y: p.y}
  case 'U': return Position{x: p.x, y: p.y+1}
  case 'D': return Position{x: p.x, y: p.y-1}
  default: fmt.Println("Error parsing move direction: ", dir)
  }
  return p
}
func (p Position) ident() string {
  return fmt.Sprintf("%v:%v", p.x, p.y)
}
func (p Position) dist() int {
  return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

var m = make(map[string]int)
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
    pw, ok := m[p.ident()]
    if (ok && pw != wire) {
      //fmt.Println("Found xing: ",p)
      xing[numx] = p; numx++
    } else {
      m[p.ident()] = wire
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
    var p = Position{x:0, y:0}
    for i:=0; i<num; i++ {
      p = doMove(wire, p, moves[i])
    }
    //fmt.Println("final position: ", p.ident())
  }
  if (numx == 0) {
    fmt.Println("No crossing found")
    os.Exit(1)
  }

  var p = xing[0]
  for i:=0;i<numx;i++ {
      if (xing[i].dist() < p.dist()) {
        p = xing[i]
      }
  }
  fmt.Printf("Crossing: %v, distance: %v\n", p, p.dist())
}

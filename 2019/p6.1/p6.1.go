package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)


const MAXARR int = 1500
var branches = [][]int{}
var orbits = []string{}
var added = []bool{}

func printBranches() {
  for i:=0;i<len(branches);i++ {
    for j:=0;j<len(branches[i]);j++ {
      fmt.Printf("%v ",orbits[branches[i][j]])
    }
    fmt.Println()
  }
}


func linkBranches() bool {
  return false
}


func countOrbits() int {
  counted := make([]bool, len(orbits))
  var count int = 0
  for i:=0;i<len(branches);i++ {
    for j:=0;j<len(branches[i]);j++ {
      orb := branches[i][j]
      if (!counted[orb]) {
        count += j+1
        counted[orb] = true
      }
    }
    fmt.Printf("count after branch %v: %v\n",i,count)
  }
  return count
}

func findOrbitsFrom(from string) []int {
  fmt.Println("Searching for orbit around ", from)
  matches := []int{}
  pref := from + ")"
  for i:=0;i<len(orbits);i++ {
    if (strings.HasPrefix(orbits[i], pref)) {
      fmt.Println("Found match: ", i, orbits[i])
      matches = append(matches, i)
    }
  }
  return matches
}

func orbitOf(orb string) string {
  s := strings.Split(orb, ")")
  if (len(s)>1) {
    return s[1]
  }
  return ""
}

func copyBranch(branch []int, orb int) {
  last := len(branches)
  branches = append(branches, []int{})
  for i:=0;i<len(branch);i++ {
    branches[last] = append(branches[last], branch[i])
  }
  branches[last] = append(branches[last], orb)
}

func createBranches() {
  m := findOrbitsFrom("COM")
  if (len(m)!=1) {
    fmt.Println("ERROR: more than one COM")
    os.Exit(2)
  }
  added = make([]bool, len(orbits))
  branches = append(branches, []int{m[0]})
  added[m[0]] = true
  for i:=0;i<len(branches);i++ {
    o := orbitOf(orbits[branches[i][len(branches[i])-1]])
    m = findOrbitsFrom(o)
    for len(m)>0 {
      for j:=1;j<len(m);j++ {
        copyBranch(branches[i],m[j])
        added[m[j]] = true
      }
      nextOrb := m[0]
      branches[i] = append(branches[i], nextOrb)
      added[nextOrb] = true
      o := orbitOf(orbits[nextOrb])
      m = findOrbitsFrom(o)
    }
  }
}

func readInput() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    orbits = append(orbits, scanner.Text())
  }
}

func main() {
  readInput()

  createBranches()
  for i:=0;i<len(orbits);i++ {
    if (!added[i]) {
      fmt.Println("Orbit not added to any branch: ", orbits[i])
    }
  }

  printBranches()

  fmt.Println("Number of branches: ", len(branches))
  fmt.Println("Number of orbits: ", len(orbits))
  fmt.Println("Number of total orbits: ", countOrbits())
}

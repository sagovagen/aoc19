package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "log"
)

var inputs = []string{}

type Int int64
type Codes []Int
type Stream []Int

func newStream() Stream {
  return Stream{}
}
func streamRead(stream *Stream) Int {
  value := (*stream)[0]
  *stream = (*stream)[1:]
  log.Printf("input stream %p: %v\n",stream, value)
  return value
}

func streamWrite(stream *Stream, value Int) {
  log.Printf("output stream %p: %v\n",stream, value)
  *stream = append(*stream, value)
}
func streamPending(stream *Stream) bool {
  return len(*stream)>0
}

func oper(opcode Int) Int {
  return opcode % Int(10)
}

func numParams(oper Int) Int {
  switch(oper) {
  case 1,2,7,8: return 3
  case 3,4,5,6: return 2
  case 9: return 1
  default: fmt.Println("Error, unknown operation: ", oper)
  }
  return 0
}

func opParams(mem Codes, ip Int, offset Int) Codes {
  opcode := mem[ip]
  op := oper(opcode)
  c := opcode - op
  n := numParams(op)
  modes := make([]Int,n)
  p := make(Codes, n)
  d := Int(100)
  for i:=Int(0);i<n;i++ {
    m := c / d
    modes[i] = m % 10
    switch(modes[i]) {
    case 0: p[i] = mem[ip+i+1]
    case 1: p[i] = ip+i+1
    case 2: p[i] = mem[ip+i+1]+offset
    }
    d = d * 10
  }
  fmt.Println(modes)
  return p
}

func run(mem Codes, ip Int, input *Stream, output *Stream) Int {
  offset := Int(0)
  for ip<Int(len(mem)) {
    log.Printf("ip: %v, op: %v\n",ip,mem[ip])
    if (mem[ip] == Int(99)) { // Halt
      return -1
    }
    op := oper(mem[ip])
    p := opParams(mem, ip, offset)
    //fmt.Printf("op: %v, params: %v\n",mem[ip],p)
    //ip += Int(len(p))+1
    switch(op) {
    case 1: // add
      mem[p[2]] = mem[p[0]] + mem[p[1]]
      ip+=4
    case 2: // mul
      mem[p[2]] = mem[p[0]] * mem[p[1]]
      ip+=4
    case 3: // inp
      if !streamPending(input) { return ip } else {
        mem[p[0]]=streamRead(input)
        ip+=2;
      }
    case 4: // out
      streamWrite(output, mem[p[0]])
      ip+=2;
    case 5: // jit
      if (mem[p[0]] != 0) { ip=mem[p[1]] } else { ip+=3 }
    case 6: // jif
      if (mem[p[0]] == 0) { ip=mem[p[1]] } else { ip+=3 }
    case 7: // lt
      if (mem[p[0]] < mem[p[1]]) { mem[p[2]] = 1 } else { mem[p[2]] = 0 }
      ip+=4;
    case 8: // eq
      if (mem[p[0]] == mem[p[1]]) { mem[p[2]] = 1 } else { mem[p[2]] = 0 }
      ip+=4;
    case 9: // ofs
      offset += mem[p[0]]
      ip+=2;
    default: // unknown
      log.Printf("ERROR: Unknown opcode %v at ip=%v\n", mem[ip], ip)
    }

    /*

    switch(mem[ip]) {
    case 1: //fmt.Printf("ip: %v, add [%v] + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]+mem[mem[ip+2]];
      ip+=4;
    case 101: //fmt.Printf("ip: %v, add %v + [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]+mem[mem[ip+2]]
      ip+=4;
    case 1001: //fmt.Printf("ip: %v, add [%v] + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]+mem[ip+2]
      ip+=4;
    case 1101: //fmt.Printf("ip: %v, add %v + %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]+mem[ip+2]
      ip+=4;
    case 201: //log.Printf("ip: %v, add [%v+%v] + [%v] -> [%v]\n", ip, mem[ip+1], offset, mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]+offset]+mem[mem[ip+2]]
      ip+=4;
    case 1201: //log.Printf("ip: %v, add [%v+%v] + %v -> [%v]\n", ip, mem[ip+1], offset, mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]+offset]+mem[ip+2]
      ip+=4;
    case 2101: //log.Printf("ip: %v, add %v + [%v+%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], offset, mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]+mem[mem[ip+2]+offset]
      ip+=4;
    case 21101: //log.Printf("ip: %v, add %v + %v -> [%v+%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3], offset);
      mem[mem[ip+3]+offset]=mem[ip+1]+mem[mem[ip+2]+offset]
      ip+=4;
    case 22101: //log.Printf("ip: %v, add %v + [%v+%v] -> [%v+%v]\n", ip, mem[ip+1], mem[ip+2], offset, mem[ip+3], offset);
      mem[mem[ip+3]+offset]=mem[ip+1]+mem[mem[ip+2]+offset]
      ip+=4;
    case 2201: //log.Printf("ip: %v, add [%v+%v] + [%v+%v] -> [%v]\n", ip, mem[ip+1], offset, mem[ip+2], offset, mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]+offset]+mem[mem[ip+2]+offset]
      ip+=4;
    case 2: //fmt.Printf("ip: %v, mul [%v] * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]*mem[mem[ip+2]];
      ip+=4;
    case 102: //fmt.Printf("ip: %v, mul %v * [%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[mem[ip+2]]
      ip+=4;
    case 1002: //fmt.Printf("ip: %v, mul [%v] * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[mem[ip+1]]*mem[ip+2]
      ip+=4;
    case 1102: //fmt.Printf("ip: %v, mul %v * %v -> [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[ip+2]
      ip+=4;
    case 21102: //log.Printf("ip: %v, mul %v * %v -> [%v+%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3], offset);
      mem[mem[ip+3]+offset]=mem[ip+1]*mem[ip+2]
      ip+=4;
    case 2102: //log.Printf("ip: %v, mul %v * [%v+%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], offset, mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[mem[ip+2]+offset]
      ip+=4;
    case 2202: //log.Printf("ip: %v, mul [%v+%v] * [%v+%v] -> [%v]\n", ip, mem[ip+1], offset, mem[ip+2], offset, mem[ip+3]);
      mem[mem[ip+3]]=mem[ip+1]*mem[mem[ip+2]+offset]
      ip+=4;
    case 3: //fmt.Printf("ip: %v, inp -> [%v]\n", ip, mem[ip+1]);
      if !streamPending(input) { return ip } else {
        mem[mem[ip+1]]=streamRead(input)
        ip+=2;
      }
    case 203: log.Printf("ip: %v, inp -> [%v+%v]\n", ip, mem[ip+1],offset);
      if !streamPending(input) { return ip } else {
        mem[mem[ip+1]+offset]=streamRead(input)
        ip+=2;
      }
    case 4: //fmt.Printf("ip: %v, out [%v]\n", ip, mem[ip+1]);
      streamWrite(output, mem[mem[ip+1]])
      ip+=2;
    case 104: //fmt.Printf("ip: %v, out %v\n", ip, mem[ip+1]);
      streamWrite(output, mem[ip+1])
      ip+=2;
    case 204: //fmt.Printf("ip: %v, out [%v+%v]\n", ip, mem[ip+1],offset);
      streamWrite(output, mem[mem[ip+1]+offset])
      ip+=2;
    case 5: //fmt.Printf("ip: %v, jit [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] != 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 105: //fmt.Printf("ip: %v, jit %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] != 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 2105: log.Printf("ip: %v, jit %v [%v+%v]\n", ip, mem[ip+1], mem[ip+2], offset);
      if (mem[ip+1] != 0) { ip=mem[mem[ip+2]+offset] } else { ip+=3 }
    case 1005: //fmt.Printf("ip: %v, jit [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] != 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1205: log.Printf("ip: %v, jit [%v+%v] %v\n", ip, mem[ip+1], offset, mem[ip+2]);
      if (mem[mem[ip+1]] != 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1105: //fmt.Printf("ip: %v, jit %v %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] != 0) { ip=mem[ip+2] } else { ip+=3 }
    case 6: //fmt.Printf("ip: %v, jif [%v] [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] == 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 106: //fmt.Printf("ip: %v, jif %v [%v]\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] == 0) { ip=mem[mem[ip+2]] } else { ip+=3 }
    case 2106: log.Printf("ip: %v, jif %v [%v+%v]\n", ip, mem[ip+1], mem[ip+2], offset);
      if (mem[ip+1] == 0) { ip=mem[mem[ip+2]+offset] } else { ip+=3 }
    case 1006: //fmt.Printf("ip: %v, jif [%v] %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[mem[ip+1]] == 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1206: log.Printf("ip: %v, jif [%v+%v] %v\n", ip, mem[ip+1], offset, mem[ip+2]);
      if (mem[mem[ip+1]+offset] == 0) { ip=mem[ip+2] } else { ip+=3 }
    case 1106: //fmt.Printf("ip: %v, jif %v %v\n", ip, mem[ip+1], mem[ip+2]);
      if (mem[ip+1] == 0) { ip=mem[ip+2] } else { ip+=3 }
    case 7: //fmt.Printf("ip: %v, lt [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 107: //fmt.Printf("ip: %v, lt %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] < mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 2107: log.Printf("ip: %v, lt %v [%v+%v] [%v]\n", ip, mem[ip+1], mem[ip+2], offset, mem[ip+3]);
      if (mem[ip+1] < mem[mem[ip+2]+offset]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1007: //fmt.Printf("ip: %v, lt [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1207: log.Printf("ip: %v, lt [%v+%v] %v [%v]\n", ip, mem[ip+1], offset, mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]+offset] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1107: //fmt.Printf("ip: %v, lt %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] < mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 8: //fmt.Printf("ip: %v, eq [%v] [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 108: //fmt.Printf("ip: %v, eq %v [%v] [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] == mem[mem[ip+2]]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 2108: log.Printf("ip: %v, eq %v [%v+%v] -> [%v]\n", ip, mem[ip+1], mem[ip+2], offset, mem[ip+3]);
      if (mem[ip+1] == mem[mem[ip+2]+offset]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1008: //fmt.Printf("ip: %v, eq [%v] %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1208: log.Printf("ip: %v, eq [%v+%v] %v -> [%v]\n", ip, mem[ip+1], offset, mem[ip+2], mem[ip+3]);
      if (mem[mem[ip+1]+offset] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 1108: //fmt.Printf("ip: %v, eq %v %v [%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3]);
      if (mem[ip+1] == mem[ip+2]) { mem[mem[ip+3]] = 1 } else { mem[mem[ip+3]] = 0 }
      ip+=4;
    case 21108: log.Printf("ip: %v, eq %v %v -> [%v+%v]\n", ip, mem[ip+1], mem[ip+2], mem[ip+3],offset);
      if (mem[ip+1] == mem[ip+2]) { mem[mem[ip+3]+offset] = 1 } else { mem[mem[ip+3]+offset] = 0 }
      ip+=4;
    case 9: log.Printf("ip: %v, ofs [%v]=%v\n",ip,mem[ip+1],mem[mem[ip+1]])
      offset += mem[mem[ip+1]]
      ip+=2;
    case 109: log.Printf("ip: %v, ofs %v\n",ip,mem[ip+1])
      offset += mem[ip+1]
      ip+=2;
    case 209: log.Printf("ip: %v, ofs [%v+%v]=%v\n",ip,mem[ip+1],offset,mem[mem[ip+1]+offset])
      offset += mem[mem[ip+1]+offset]
      ip+=2;
    case 99:
      //fmt.Printf("ip: %v, Halt\n", ip);
      return -1
    default: fmt.Printf("ERROR: Unknown opcode %v at ip=%v\n", mem[ip], ip)
    }
    */
  }
  return -1
}

func copyIntcode(intcode Codes) Codes {
  copy := make(Codes, len(intcode))
  for i:=0;i<len(intcode);i++ {
    copy[i] = intcode[i]
  }
  return copy
}

const EXTRAMEM int = 10000

func main() {
  if (len(os.Args) != 3) {
    fmt.Printf("usage: %v intcodefile input\n", os.Args[0])
    os.Exit(2)
  }
  program := os.Args[1]
  input := os.Args[2]

  pfile, err := os.Open(program)
  if err != nil {
    fmt.Println("Error reading program file")
    os.Exit(2)
  }
  defer pfile.Close()

  // Read program
  pscanner := bufio.NewScanner(pfile)
  pscanner.Split(bufio.ScanLines)
  if (!pscanner.Scan()) {
    fmt.Println("Error reading intcode program")
    os.Exit(2)
  }
  str := pscanner.Text()
  //fmt.Println("input: ",str)
  strarr := strings.Split(str,",")
  mem := make(Codes, len(strarr) + EXTRAMEM) // Extra memory storage
  codelen := len(strarr)
  fmt.Println("intcode length: ", codelen)
  for i := 0; i<codelen; i++ {
    c, err := strconv.ParseInt(strarr[i], 10, 64)
    if err != nil {
      fmt.Println("Oops. Error reading intcode: ", strarr[i])
    } else {
      mem[i] = Int(c)
    }
  }

/*
  opcodes := []Int{1,101,1001,1101,1201,201,2101,2201,21101,12101,0,109,209}
  for i:=0;i<len(opcodes);i++ {
    fmt.Printf("Opcode %v param modes: %v\n",opcodes[i],opParams(opcodes[i],mem,0))
  }
  os.Exit(2)
*/

  // Create input/output streams
  instream := newStream()
  fmt.Println("Parsing input: ", input)
  v,err := strconv.ParseInt(input, 10, 64)
  if err == nil {
    streamWrite(&instream, Int(v))
  } else {
    fmt.Println("Error parsing input: ", err)
  }
  outstream := newStream()

  ip := run(mem,0,&instream, &outstream)

  fmt.Println("Program finished at ip: ", ip)
  fmt.Println("Program output: ", outstream)
}

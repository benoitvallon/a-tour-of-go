package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
  return fmt.Sprintf("\"%d.%d.%d.%d\"", p[0], p[1], p[2], p[3])
}

func main() {
  addrs := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for n, a := range addrs {
    fmt.Printf("%v: %v\n", n, a)
  }
}

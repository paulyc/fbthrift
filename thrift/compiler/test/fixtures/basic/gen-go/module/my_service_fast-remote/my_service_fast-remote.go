// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        thrift "github.com/facebook/fbthrift-go"
        "../../module"
)

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void ping()")
  fmt.Fprintln(os.Stderr, "  string getRandomData()")
  fmt.Fprintln(os.Stderr, "  bool hasDataById(i64 id)")
  fmt.Fprintln(os.Stderr, "  string getDataById(i64 id)")
  fmt.Fprintln(os.Stderr, "  void putDataById(i64 id, string data)")
  fmt.Fprintln(os.Stderr, "  void lobDataById(i64 id, string data)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.Transport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewHTTPPostClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewSocket(thrift.SocketAddr(net.JoinHostPort(host, portStr)))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.ProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := module.NewMyServiceFastClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "ping":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Ping requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Ping())
    fmt.Print("\n")
    break
  case "getRandomData":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetRandomData requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetRandomData())
    fmt.Print("\n")
    break
  case "hasDataById":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "HasDataById requires 1 args")
      flag.Usage()
    }
    argvalue0, err30 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err30 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.HasDataById(value0))
    fmt.Print("\n")
    break
  case "getDataById":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDataById requires 1 args")
      flag.Usage()
    }
    argvalue0, err31 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err31 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDataById(value0))
    fmt.Print("\n")
    break
  case "putDataById":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "PutDataById requires 2 args")
      flag.Usage()
    }
    argvalue0, err32 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err32 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.PutDataById(value0, value1))
    fmt.Print("\n")
    break
  case "lobDataById":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "LobDataById requires 2 args")
      flag.Usage()
    }
    argvalue0, err34 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err34 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.LobDataById(value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}

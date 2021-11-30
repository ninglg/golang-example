package server

import ( 
    "log"
    "net"
    "os"
)

type Server struct {
    Address   string 
}

func NewServer(address string) *Server{
    return &Server{
        Address: address,
    }
}

// 开启TCP服务
func (this *Server) Start(){
    
    l, err := net.Listen("tcp", this.Address);
    checkError(err)
    defer l.Close()

    log.Printf("Server start: %s\n", this.Address)

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Println(err.Error())
            continue
        }

        go handleConnection(conn);
    }
}

// 连接接收
func handleConnection(conn net.Conn){
    
    rAddr := conn.RemoteAddr()
    log.Printf("Connect from client: %s\n", rAddr)
 
  	defer conn.Close()
  
  	_, err := conn.Write([]byte("Welcome client"))
    if err != nil {
      return
    }
  	
  	input := bufio.NewScanner(conn)
    for input.Scan() {
       msg := input.Text()
       log.Println("Receive from client", rAddr.String(), msg)
    }
}

func checkError(err error) {
    if err != nil {
        log.Printf("Fatal error %s", err.Error())
        os.Exit(1)
    }
}

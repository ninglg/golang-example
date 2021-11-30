package main

import(
    "server"
)

func main()  {  
      server := server.NewServer(":8999");
      server.Start();
}

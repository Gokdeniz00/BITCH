package backdoor

import (
	"fmt"
	"log"
	"net"
)

func main(){
	h:=DNSLookUp("",0)//enter your host here in form protocol://ipv4(xxx.xxx.xxx.xxx):Port_Number and use mode 0 or domain name such as http://example.com and use mode 1
}
type Host struct{
	Protocol  string
	IPAddress string
	Port	  uint
}
func DNSLookUp(url string,mode int) Host{
	ipaddr,err:=net.LookupHost(url)
	if err != nil {
		log.Fatal(err)
	}
	var protocol string
	if url[0:5]=="https"{
		protocol="https"
	}else{
		protocol="http"
	}
	return Host{Protocol: protocol,IPAddress: ipaddr[0],Port: 443}
}
func ExecuteCommands(commands string){
	
}

func (h Host) Address() string{
	address:=fmt.Sprintf("%s://%s:%d",h.Protocol,h.IPAddress,h.Port)
	return address
}
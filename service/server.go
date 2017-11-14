package service 
 
import ( 
	// use martini framework
	"github.com/go-martini/martini" 
) 

 
func NewServer(port string) { 
	m := martini.Classic()  
	// get from client
	m.Get("/hello/:name", func(params martini.Params) string { 
		return "Hello " + params["name"] 
	}) 
	// change the port 
	m.RunOnAddr(":"+port)	 
} 

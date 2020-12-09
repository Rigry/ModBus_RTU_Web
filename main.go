package main

import "modbus_client"
import "fmt"
import "time"
import "net/http"
import "html/template"
import "encoding/binary"

type Contain struct {
	Var uint16
	Click bool
}

var (
	tmpl = template.Must(template.ParseFiles("web.html"))
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 10
	handler.Timeout = 1 * time.Second

	client := modbus.NewClient(handler)
	
	// for {
		cont := Contain {}
		// if cont.Click {
			results, _ := client.ReadHoldingRegisters(1, 1)
			cont.Var = binary.BigEndian.Uint16(results)
			fmt.Println(cont.Var, cont.Click)
			tmpl.Execute(w, cont)
			cont.Click = false;
			// time.Sleep(1 * time.Second)
			// cont := Contain {}
			// cont.Var = 23
			// tmpl.Execute(w, cont)
		// }
	// }
	
	
}

func main() {
	
	// err := handler.Connect()
	// defer handler.Close()
	
	handleRequest()
	
	

}

func handleRequest() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs)) 
	
	http.HandleFunc("/", Handler)
   http.ListenAndServe(":8080", nil)
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	t:= template.Must(template.ParseFiles("web.html"))
// 	// t.ExecuteTemplate(os.Stdout, "web")
//    t.Execute(w, nil)
// }

package nettools

import (
	f "customs/filetools"
	"log"
	"net"
)

func Send_file(filename string, network, address string) error {
	f := f.File_read(filename, "none")
	connstat, _ := net.Dial(network, address)
	_, err := connstat.Write(f.File)
	if err != nil {
		log.Fatal()
		return err
	}
	connstat.Close()
	return err
}

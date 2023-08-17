package client

import (
	"cameraClient/entities"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/use-go/onvif"
)

func (s *cameraClient) GetCameraInterface() ([]string, error) {
	// TODO: Scan possible interface and return
	log.Println("---------------------Start-Scan---------------------")
	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ints := []string{}
	for _, it := range itfs {
		flags := strings.Split(it.Flags.String(), "|")
		if len(flags) == 4 {
			ints = append(ints, it.Name)
			log.Printf("---------------------Find-%s---------------------\n", it.Name)
		}
	}
	log.Println("---------------------Complete-Scan---------------------")
	return ints, nil
}

func (c *cameraClient) ScanCamera(interfaces ...string) ([]entities.Camera, error) {
	log.Println("---------------------Start-Scan---------------------")
	onvifDevices := []onvif.Device{}
	cameras := []entities.Camera{}
	for _, i := range interfaces {
		log.Printf("---------------------Interface-%s--------------------\n", i)
		devices, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface(i)
		if err != nil {
			return nil, err
		}

		for _, device := range devices {
			// info := device.GetDeviceInfo()
			log.Printf("Info: %v", device)
			endpoints := device.GetServices()
			for _, endpoint := range endpoints {
				log.Printf("address: %v", device.GetEndpoint(endpoint))
				camera := entities.Camera{
					SourceUrls: []entities.SourceUrl{},
					Comment:    fmt.Sprintf("%s:%s", i, endpoint),
				}
				cameras = append(cameras, camera)
			}
		}

		onvifDevices = append(onvifDevices, devices...)
		log.Println("---------------------Done-Interface---------------------")
	}
	log.Println("---------------------Complete-Scan---------------------")
	return cameras, nil
}

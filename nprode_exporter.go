package main

import (
	"fmt"
	"net"
	"time"
)

// global constants mostly later handle using yaml
// configuration ...
const (
	prome_gw = "http://10.100.254.6:9091"
	job_name = "nProde_exporter"
)

// check ip: string at port: int and
// return bool
func checkPort(ep string, port int) bool {
	address := fmt.Sprintf("%s:%d", ep, port)
	fmt.Println("Address to probe is : ", address)

	con, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		fmt.Println("Connection failed to ", address)
		return false
	}

	// close the tcp socket back
	defer con.Close()
	fmt.Println("Connection successful to : ", address)
	return true
}

// push the actual metric to pushGateway
func pushToGw(ep string, port int, status float64) {
	fmt.Printf("Pushing to GW for endpoint: %s, port: %d, status: %.2f \n", ep, port, status)
}

func main() {
	endPoint2PortsMap := map[string][]int{
		"google.com":     {443},
		"34.117.118.44":  {443},
		"myrepublic.net": {443},
	}

	for {
		for ep, ports := range endPoint2PortsMap {
			for _, port := range ports {
				// check port and get the status
				// status 1 means it is open 0 means
				// either closed or unreachable
				portStatus := checkPort(ep, port)
				var status float64
				if portStatus {
					status = 1
				} else {
					status = 0
				}

				// now that status is here time to push
				pushToGw(ep, port, status)
			}
		}

		// sleep a while before doing the whole thing again
		// will handle from configuration later
		time.Sleep(10 * time.Second) // check every 10 second
	}
}

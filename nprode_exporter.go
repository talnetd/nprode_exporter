package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"gopkg.in/yaml.v2"
)

// global constants mostly later handle using yaml
// configuration ...
var (
	configPath = kingpin.Flag("config", "Path to configuration YAML file usually this should be /etc/nprode_config.yaml").Required().String()
)

// parse the config into a struct
type Config struct {
	PGW       string           `yaml:"pgw"`
	EndPoints map[string][]int `yaml:"endpoints"`
	JobName   string           `yaml:"job_name"`
}

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
	kingpin.Parse()

	configFile, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	// Log the Prometheus Push Gateway and Job Name
	fmt.Printf("Using Prometheus Push Gateway: %s\n", config.PGW)
	fmt.Printf("Using Job Name: %s\n", config.JobName)

	for {
		for ep, ports := range config.EndPoints {
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

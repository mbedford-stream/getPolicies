package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/Juniper/go-netconf/netconf"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// SystemInformation provides a representation of the system-information container
type SystemInformation struct {
	HardwareModel string `xml:"system-information>hardware-model"`
	OsName        string `xml:"system-information>os-name"`
	OsVersion     string `xml:"system-information>os-version"`
	SerialNumber  string `xml:"system-information>serial-number"`
	HostName      string `xml:"system-information>host-name"`
}

var (
	host         = flag.String("host", "192.168.86.3", "Hostname")
	username     = flag.String("username", "", "Username")
	key          = flag.String("key", os.Getenv("HOME")+"/.ssh/id_rsa", "SSH private key file")
	passphrase   = flag.String("passphrase", "", "SSH private key passphrase (cleartext)")
	nopassphrase = flag.Bool("nopassphrase", false, "SSH private key does not contain a passphrase")
	pubkey       = flag.Bool("pubkey", false, "Use SSH public key authentication")
	agent        = flag.Bool("agent", false, "Use SSH agent for public key authentication")
)

func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileIsADirectory(file string) bool {
	if stat, err := os.Stat(file); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func FileExistsAndIsADirectory(file string) bool {
	if FileExists(file) && FileIsADirectory(file) {
		return true
	}
	return false
}

func getRPC(devIP string, sshConfig *ssh.ClientConfig, rpcCommand string) *netconf.RPCReply {

	fmt.Printf("Connecting: %s\n", devIP)
	s, err := netconf.DialSSH(devIP, sshConfig)
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
	}

	defer s.Close()

	// Sends raw XML
	res, err2 := s.Exec(netconf.RawMethod(rpcCommand))
	if err2 != nil {
		panic(err2)
	}
	return res
}

func BuildConfig() *ssh.ClientConfig {
	var config *ssh.ClientConfig
	var pass string
	if *pubkey {
		if *agent {
			var err error
			config, err = netconf.SSHConfigPubKeyAgent(*username)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if *nopassphrase {
				pass = "\n"
			} else {
				if *passphrase != "" {
					pass = *passphrase
				} else {
					var readpass []byte
					var err error
					fmt.Printf("Enter Passphrase for %s: ", *key)
					readpass, err = term.ReadPassword(int(syscall.Stdin))
					if err != nil {
						log.Fatal(err)
					}
					pass = string(readpass)
					fmt.Println()
				}
			}
			var err error
			config, err = netconf.SSHConfigPubKeyFile(*username, *key, pass)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Printf("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()

		config = netconf.SSHConfigPassword(*username, string(bytePassword))
	}
	return config
}

func main() {
	flag.Parse()

	sshConfig := BuildConfig()

	fmt.Println("\nConnecting....")
	s, err := netconf.DialSSH(*host, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	policiesCmd := "<get-firewall-policies/>"
	// policiesCmdGlobal := "<get-global-firewall-policies/>"

	var secPolicies MultiRoutingEngineResults

	fmt.Println("Calling getData function")
	RPCreply := getRPC(*host, sshConfig, policiesCmd)
	fmt.Println("Finished getting data.... ")

	err = xml.Unmarshal([]byte(RPCreply.Data), &secPolicies)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range secPolicies.MultiRoutingEngineItem.SecurityPolicies.SecurityContext {
		fmt.Printf("\n\nFrom Zone: %s\nTo Zone: %s\n\t============================\n", v.ContextInformation.SourceZoneName, v.ContextInformation.DestinationZoneName)
		for _, v2 := range v.Policies {
			appList := ""
			for _, app := range v2.PolicyInformation.Applications.Application {
				appList = fmt.Sprintf("%s, %s", appList, app.ApplicationName)
			}
			srcList := ""
			for _, src := range v2.PolicyInformation.SourceAddresses.SourceAddress {
				srcList = fmt.Sprintf("%s, %s", srcList, src.AddressName)
			}
			dstList := ""
			for _, dst := range v2.PolicyInformation.DestinationAddresses.DestinationAddress {
				dstList = fmt.Sprintf("%s, %s", dstList, dst.AddressName)
			}
			if len(v2.PolicyInformation.MultipleDestinationZones.DestinationZone) != 0 || len(v2.PolicyInformation.MultipleSourceZones.SourceZone) != 0 {
				var srcZoneList []string
				var dstZoneList []string
				for _, srcZone := range v2.PolicyInformation.MultipleSourceZones.SourceZone {
					srcZoneList = append(srcZoneList, srcZone.SourceZoneName)
				}
				for _, dstZone := range v2.PolicyInformation.MultipleDestinationZones.DestinationZone {
					dstZoneList = append(dstZoneList, dstZone.DestinationZoneName)
				}
				fmt.Printf("\t---GLOBAL POLICY---\n\tSource Zone/s: %s\n\tDestination Zone/s: %s\n", srcZoneList, dstZoneList)
			}
			fmt.Printf("\tPolicy: %s(%s)\n\tSrc Addr: %s\n\tDst Addr: %s\n\tApp/Port: %s\n\tAction: %s\n\t============================\n", v2.PolicyInformation.PolicyName, v2.PolicyInformation.PolicyIdentifier, srcList, dstList, appList, v2.PolicyInformation.PolicyAction.ActionType)
		}
	}
	writeOut := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s\n", "Source Zone", "Destination Zone", "Policy(ID)", "Source Addr", "Destination Addr", "Applications", "Action")
	var writePol string
	// fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", "Source Zone", "Destination Zone", "Policy(ID)", "Source Addr", "Destination Addr", "Applications", "Action")
	for _, v := range secPolicies.MultiRoutingEngineItem.SecurityPolicies.SecurityContext {
		var srcZoneList []string
		var dstZoneList []string
		for _, v2 := range v.Policies {
			if len(v2.PolicyInformation.MultipleDestinationZones.DestinationZone) != 0 || len(v2.PolicyInformation.MultipleSourceZones.SourceZone) != 0 {
				// var srcZoneList []string
				// var dstZoneList []string
				for _, srcZone := range v2.PolicyInformation.MultipleSourceZones.SourceZone {
					srcZoneList = append(srcZoneList, srcZone.SourceZoneName)
				}
				for _, dstZone := range v2.PolicyInformation.MultipleDestinationZones.DestinationZone {
					dstZoneList = append(dstZoneList, dstZone.DestinationZoneName)
				}
			} else {
				srcZoneList = nil
				dstZoneList = nil
				srcZoneList = append(srcZoneList, v.ContextInformation.SourceZoneName)
				dstZoneList = append(dstZoneList, v.ContextInformation.DestinationZoneName)
			}

			// appList := ""
			var appList []string
			for _, app := range v2.PolicyInformation.Applications.Application {
				appList = append(appList, app.ApplicationName)
			}
			var srcList []string
			for _, src := range v2.PolicyInformation.SourceAddresses.SourceAddress {
				srcList = append(srcList, src.AddressName)
			}
			var dstList []string
			for _, dst := range v2.PolicyInformation.DestinationAddresses.DestinationAddress {
				dstList = append(dstList, dst.AddressName)
			}
			writePol = fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s", srcZoneList, dstZoneList, v2.PolicyInformation.PolicyName+"("+v2.PolicyInformation.PolicyIdentifier+")", srcList, dstList, appList, v2.PolicyInformation.PolicyAction.ActionType)
			if len(v2.PolicyInformation.MultipleDestinationZones.DestinationZone) != 0 || len(v2.PolicyInformation.MultipleSourceZones.SourceZone) != 0 {
				writePol = writePol + ", GLOBAL\n"
			} else {
				writePol = writePol + "\n"
			}
			writeOut = writeOut + writePol
			// fmt.Println(writeOut)
		}
	}

	todayDate := time.Now().Format("02-01-2006")
	// todayFormatted := todayDate.Format("02-01-2006")

	outputFile := fmt.Sprintf("%s--%s.csv", todayDate, strings.Replace(*host, ".", "_", -1))

	if !FileExists(outputFile) || FileExistsAndIsADirectory(outputFile) {
		fNew, err := os.Create(outputFile)
		if err != nil {
			log.Fatal("Could not create output file:\n", err)
		}
		defer fNew.Close()
	}

	f, err := os.OpenFile(outputFile, os.O_WRONLY, 0755)
	// f, err := os.Open(outputFile)
	if err != nil {
		color.Red("Problems opening output file: %s", outputFile)
		log.Fatal(err)
	}

	_, err = f.WriteString(writeOut)
	if err != nil {
		color.Red("problems writing file")
		fmt.Println(err)
	}
	defer f.Close()

}

# Setup

*Instructions written assuming Linux machine and Belkin wireless router. Commands and dialogue may differ for other systems.*

## Basic Configuration

Install `go`.

Run with `go run MAIN.go`.

Access the chat room by going to `localhost:8080` in a browser.

## External Connection Configuration

### Firewall Access

Create a rule to allow access to the port `sudo iptables -I INPUT 1 -i eth0 -p tcp --dport 8080 -j ACCEPT`.

Save the new rule `sudo iptables-save`.

### Router Port Forwarding

Find the private IP for the router `ip route | grep default.`

Access the private IP through a browser & go to Virtual Servers (Port Forwarding).

Find the private IP for the machine running the chat service `ifconfig`.

Add and enable a TCP rule to map Inbound Port 8080 to the machine's private IP on port 8080.


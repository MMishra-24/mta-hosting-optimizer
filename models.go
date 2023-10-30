package main

type IPConfig struct {
	IP       string
	Hostname string
	Active   bool
}

var IpConfigData = []IPConfig{
	{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
	{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
	{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
	{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
	{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: false},
	{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
}

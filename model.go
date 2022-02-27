package main

type HTTPModel struct {
	HTTP HTTP `json:"http"`
}

type (
	Server struct {
		URL string `json:"url"`
	}
	LB struct {
		HealtCheck *HealtCheck `json:"healthCheck,omitempty"`
		Servers    []Server    `json:"servers"`
	}
	Service struct {
		LB LB `json:"loadBalancer"`
	}

	HealtCheck struct {
		Path     string `json:"path,omitempty"`
		Interval string `json:"interval,omitempty"`
		Port     string `json:"port,omitempty"`
		Scheme   string `json:"scheme,omitempty"`
		Timeout  string `json:"timeout,omitempty"`
	}
)

type (
	Router struct {
		Rule    string `json:"rule"`
		Service string `json:"service"`
	}
)
type HTTP struct {
	Routers  map[string]Router  `json:"routers"`
	Services map[string]Service `json:"services,omitempty"`
}

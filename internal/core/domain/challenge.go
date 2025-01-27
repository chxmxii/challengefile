package domain

type Challenge struct {
	Name       string      `yaml:"name" json:"name"`
	Metadata   *Metadata   `yaml:"Metadata" json:"metadata"`
	Deployment *Deployment `yaml:"Deployment" json:"deployment"`
	Service    *Service    `yaml:"Service" json:"service"`
}

type Metadata struct {
	Namespace string `yaml:"namespace" json:"namespace"`
	Category  string `yaml:"category" json:"category"`
}

type Deployment struct {
	Name        string `yaml:"name" json:"name"`
	Image       string `yaml:"image" json:"image"`
	Replicas    int    `yaml:"replicas" json:"replicas"`
	HPA         bool   `yaml:"hpa" json:"hpa"`
	HealthCheck bool   `yaml:"healthCheck" json:"healthCheck"`
}

type Service struct {
	Name        string   `yaml:"name" json:"name"`
	Port        int32    `yaml:"port" json:"port"`
	DnsEndpoint string   `yaml:"dnsEndpoint" json:"dnsEndpoint"`
	Protocol    Protocol `yaml:"protocol" json:"protocol"`
}

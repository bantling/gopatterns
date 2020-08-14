package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Product is the data returned by a Service
type Product struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

// Service is an interface describing a Service that returns Products.
type Service interface {
	Call() []Product
}

// LocalService is a service that runs locally, no tnet connection reqiured
type LocalService struct {
	products []Product
}

// WithProducts sets the products to return on future invocations of Call()
func (ls *LocalService) WithProducts(products ...Product) {
	ls.products = products
}

// configure ensures the service has default products if none were provided
func (ls *LocalService) configure() {
	if len(ls.products) == 0 {
		ls.products = []Product{
			{Name: "Oranges", Price: "2.50/lb"},
			{Name: "Apples", Price: "5.00/lb"},
		}
	}
}

// Call returns Products
func (ls LocalService) Call() []Product {
	// Ensure we use a configured service
	(&ls).configure()

	return ls.products
}

// RenoteService is a service that acquires info from a remote RESTful service.
// The zero value is ready to use.
type RemoteService struct {
	method string
	theURL string
}

// WithMethod override the default method of GET
func (rs *RemoteService) WithMethod(method string) *RemoteService {
	rs.method = method
	return rs
}

// WithRemote overrides the default remote of http://localhost:80
func (rs *RemoteService) WithRemote(URL string) *RemoteService {
	rs.theURL = URL

	return rs
}

// configure ensures the service has non-empty values for method and remote
func (rs *RemoteService) configure() {
	if rs.method == "" {
		rs.method = http.MethodGet
	}

	if rs.theURL == "" {
		rs.theURL = "http://localhost:80"
	}
}

// Call fetches data from the remote service and returns it
// Params only have to be provided if the service requires it
func (rs RemoteService) Call() []Product {
	// Ensure we use a configured service
	(&rs).configure()

	req, err := http.NewRequest(rs.method, rs.theURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	var products []Product
	if err = json.NewDecoder(resp.Body).Decode(&products); err != nil {
		panic(err)
	}

	return products
}

// ServiceFactory creates services definitions from a JSON file
type ServiceFactory struct {
	filename string
	services map[string]Service
}

// WithFilename provides a filename to load (default is services.json)
func (f *ServiceFactory) WithFilename(filename string) *ServiceFactory {
	f.filename = filename
	return f
}

// configure ensures the factory has values
func (f *ServiceFactory) configure() {
	if f.filename == "" {
		f.filename = "services.json"
	}

	if len(f.services) == 0 {
		f.services = map[string]Service{}
	}
}

// Load loads the services based on a config file
func (f *ServiceFactory) Load() {
	f.configure()

	file, err := os.Open(f.filename)
	if err != nil {
		panic(err)
	}

	var svcs []struct {
		Name     string    `json:"name"`
		Products []Product `json:"products"`
		Method   string    `json:"method"`
		URL      string    `json:"url"`
	}

	if err := json.NewDecoder(file).Decode(&svcs); err != nil {
		panic(err)
	}

	for _, svc := range svcs {
		if svc.Method == "" {
			f.services[svc.Name] = LocalService{products: svc.Products}
		} else {
			f.services[svc.Name] = RemoteService{method: svc.Method, theURL: svc.URL}
		}
	}
}

func main() {
	// Set up a factory
	var f ServiceFactory
	f.Load()

	// Dump the services
	fmt.Printf("%s:\n", f.filename)

	for name, svc := range f.services {
		if ls, isa := svc.(LocalService); isa {
			(&ls).configure()
			fmt.Printf("- configured = %+v\n", ls)
		} else {
			rs := svc.(RemoteService)
			(&rs).configure()
			fmt.Printf("- configured = %+v\n", rs)
		}

		fmt.Printf("Name: %s, Type: %T, %+v\n", name, svc, svc)
	}
}

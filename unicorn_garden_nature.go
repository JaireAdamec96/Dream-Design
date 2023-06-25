package main

import (
	"fmt"
	"log"
)

// EventPlanningDesignService is an event planning and design service.
type EventPlanningDesignService struct {
	Clients []Client
	Events  []Event
}

// Client represents a client we will be working with.
type Client struct {
	Name  string
	Email string
	Phone string
}

// Event represents an event we will be organizing and designing for a client.
type Event struct {
	Client      Client
	Name        string
	Type        string
	Description string
	StartTime   string
	EndTime     string
	Location    string
	Guests      []Guest
	Vendors     []Vendor
}

// Guest represents a guest of the event.
type Guest struct {
	Name  string
	Email string
	Phone string
}

// Vendor represents a vendor providing services to the event.
type Vendor struct {
	Name      string
	Service   string
	Cost      float64
	StartTime string
	EndTime   string
}

// CreateClient creates a new Client record in our database.
func (eps *EventPlanningDesignService) CreateClient(c Client) {
	eps.Clients = append(eps.Clients, c)
}

// GetClient retrieves a client record from our database.
func (eps *EventPlanningDesignService) GetClient(name string) (Client, error) {
	for _, c := range eps.Clients {
		if c.Name == name {
			return c, nil
		}
	}

	return Client{}, fmt.Errorf("no such client: %s", name)
}

// CreateEvent creates a new Event record in our database.
func (eps *EventPlanningDesignService) CreateEvent(e Event) {
	eps.Events = append(eps.Events, e)
}

// GetEvent returns an Event record from our database based on the given name.
func (eps *EventPlanningDesignService) GetEvent(name string) (Event, error) {
	for _, e := range eps.Events {
		if e.Name == name {
			return e, nil
		}
	}

	return Event{}, fmt.Errorf("no such event: %s", name)
}

// CreateGuest creates a new Guest record in our database.
func (e *Event) CreateGuest(g Guest) {
	e.Guests = append(e.Guests, g)
}

// GetGuest retrives a Guest record from our database based on the given name.
func (e *Event) GetGuest(name string) (Guest, error) {
	for _, g := range e.Guests {
		if g.Name == name {
			return g, nil
		}
	}

	return Guest{}, fmt.Errorf("no such guest: %s", name)
}

// CreateVendor creates a new Vendor record in our database.
func (e *Event) CreateVendor(v Vendor) {
	e.Vendors = append(e.Vendors, v)
}

// GetVendor retrieves a Vendor record from our database based on the given name.
func (e *Event) GetVendor(name string) (Vendor, error) {
	for _, v := range e.Vendors {
		if v.Name == name {
			return v, nil
		}
	}

	return Vendor{}, fmt.Errorf("no such vendor: %s", name)
}

func main() {
	eps := &EventPlanningDesignService{
		Clients: []Client{},
		Events:  []Event{},
	}

	c := Client{
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phone: "1234567890",
	}

	eps.CreateClient(c)

	// Check that the client was successfully created.
	client, err := eps.GetClient("John Doe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", client)

	// Create a new event.
	event := Event{
		Client:      c,
		Name:        "Company Holiday Party",
		Type:        "Corporate",
		Description: "Annual holiday party for the employees of Acme Inc.",
		StartTime:   "December 24th, 2019 8pm",
		EndTime:     "December 25th, 2019 12am",
		Location:    "Acme Inc. HQ",
		Guests:      []Guest{},
		Vendors:     []Vendor{},
	}

	eps.CreateEvent(event)

	// Check that the event was successfully created.
	getEvent, err := eps.GetEvent("Company Holiday Party")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", getEvent)
}
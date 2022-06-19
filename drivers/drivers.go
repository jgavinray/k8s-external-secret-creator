/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package drivers

// Interfaces for External Provider service
// A Provider should be thought of in this case as the system that the driver will interact with.
// The assumption for the provider is that it will acces a key/value pair.
type Provider interface {
	PutParameter(key string, value string) error
}

// Function to return a Provider interface
type ProviderLoad func() Provider

// List of available drivers - map to be populated by RegisterDriver
var drivers = make(map[string]*Driver)

// Driver definition
type Driver struct {
	name string
	new  ProviderLoad
}

// Create a new Provider interface when loading a Driver
func (d *Driver) New() Provider {
	return d.new()
}

// Can the requested Driver be found and loaded
func (d *Driver) Available() bool {
	return d.new != nil
}

// Pass the string name of the driver should be registere
func RegisterDriver(name string, p ProviderLoad) {
	drivers[name] = &Driver{
		name: name,
		new:  p,
	}
}

//Lookup to see if a Driver is Registered
func Lookup(name string) *Driver {
	driver := drivers[name]
	if driver == nil {
		return nil
	}
	return driver
}

// Create a new driver object that is based on the Provider interface, if driver exists
func New(provider string) Provider {
	driver := Lookup(provider)
	if driver == nil {
		return nil
	}

	return driver.New()
}

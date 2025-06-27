package common

import (
	"fmt"
	"sync"
)

// GeneratorRegistry manages registered generators
type GeneratorRegistry struct {
	generators map[string]Generator
	mutex      sync.RWMutex
}

// NewGeneratorRegistry creates a new generator registry
func NewGeneratorRegistry() *GeneratorRegistry {
	return &GeneratorRegistry{
		generators: make(map[string]Generator),
	}
}

// Register registers a new generator
func (gr *GeneratorRegistry) Register(generator Generator) {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	gr.generators[generator.GetName()] = generator
}

// Get retrieves a generator by name
func (gr *GeneratorRegistry) Get(name string) (Generator, error) {
	gr.mutex.RLock()
	defer gr.mutex.RUnlock()
	
	generator, exists := gr.generators[name]
	if !exists {
		return nil, fmt.Errorf("generator '%s' not found", name)
	}
	
	return generator, nil
}

// List returns all registered generator names
func (gr *GeneratorRegistry) List() []string {
	gr.mutex.RLock()
	defer gr.mutex.RUnlock()
	
	names := make([]string, 0, len(gr.generators))
	for name := range gr.generators {
		names = append(names, name)
	}
	
	return names
}

// Global registry instance
var Registry = NewGeneratorRegistry()

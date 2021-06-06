package template

import "fmt"

// Add template
func Add(template *TemplateConfig) error {
	fmt.Printf("add template %s with weight %v\n", template.Name, template.Weight)
	fmt.Printf("tags: %v\n", template.Tags)
	return nil
}

package template

import "fmt"

// Remove template
func Remove(template *TemplateConfig) error {
	fmt.Printf("remove template %s\n", template.Name)
	return nil
}

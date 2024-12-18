package nfs

import (
	"strings"
	"testing"
)

// TestReplaceTemplate_AllPlaceholdersReplaced verifies that all placeholders
// in the template are correctly replaced when no tags are blocked.
func TestReplaceTemplate_AllPlaceholdersReplaced(t *testing.T) {
	template := `<Test>
	<Placeholder1>{%Placeholder1%}</Placeholder1>
	<Placeholder2>{%Placeholder2%}</Placeholder2>
	</Test>`

	result, err := ReplaceTemplate(template)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Check that placeholders are replaced
	if strings.Contains(resultStr, "{%Placeholder1%}") || strings.Contains(resultStr, "{%Placeholder2%}") {
		t.Errorf("Expected all placeholders to be replaced")
	}

	// Check that replaced values are non-empty and tags exist
	if !strings.Contains(resultStr, "<Placeholder1>") || !strings.Contains(resultStr, "<Placeholder2>") {
		t.Errorf("Expected placeholders to have replaced values")
	}

	// Optionally, verify that the replaced values are non-empty
	if !strings.Contains(resultStr, "<Placeholder1>") || !strings.Contains(resultStr, "<Placeholder2>") {
		t.Errorf("Expected placeholders to have non-empty replaced values")
	}
}

// TestReplaceTemplate_BlockSinglePlaceholder verifies that blocking a single placeholder
// removes its corresponding XML tags from the generated output.
func TestReplaceTemplate_BlockSinglePlaceholder(t *testing.T) {
	template := `<Test>
	<Placeholder1>{%Placeholder1%}</Placeholder1>
	<Placeholder2>{%Placeholder2%}</Placeholder2>
	</Test>`

	blocked := []string{"Placeholder1"}

	// Create options using WithBlockedPlaceholders
	options := []Option{
		WithBlockedPlaceholders(blocked...),
	}

	result, err := ReplaceTemplate(template, options...)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Check that Placeholder1 tags are removed
	if strings.Contains(resultStr, "<Placeholder1>") || strings.Contains(resultStr, "</Placeholder1>") {
		t.Errorf("Expected <Placeholder1> tags to be removed")
	}

	// Check that Placeholder2 is still replaced
	if strings.Contains(resultStr, "{%Placeholder2%}") {
		t.Errorf("Expected Placeholder2 to be replaced")
	}

	// Ensure Placeholder2 has been replaced with a non-empty value
	if !strings.Contains(resultStr, "<Placeholder2>") {
		t.Errorf("Expected Placeholder2 to have a replaced value")
	}
}

// TestReplaceTemplate_BlockMultiplePlaceholders verifies that blocking multiple placeholders
// removes their corresponding XML tags from the generated output.
func TestReplaceTemplate_BlockMultiplePlaceholders(t *testing.T) {
	template := `<Test>
	<Placeholder1>{%Placeholder1%}</Placeholder1>
	<Placeholder2>{%Placeholder2%}</Placeholder2>
	<Placeholder3>{%Placeholder3%}</Placeholder3>
	</Test>`

	blocked := []string{"Placeholder1", "Placeholder3"}

	// Create options using WithBlockedPlaceholders
	options := []Option{
		WithBlockedPlaceholders(blocked...),
	}

	result, err := ReplaceTemplate(template, options...)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Check that Placeholder1 and Placeholder3 tags are removed
	if strings.Contains(resultStr, "<Placeholder1>") || strings.Contains(resultStr, "</Placeholder1>") {
		t.Errorf("Expected <Placeholder1> tags to be removed")
	}
	if strings.Contains(resultStr, "<Placeholder3>") || strings.Contains(resultStr, "</Placeholder3>") {
		t.Errorf("Expected <Placeholder3> tags to be removed")
	}

	// Check that Placeholder2 is still replaced
	if strings.Contains(resultStr, "{%Placeholder2%}") {
		t.Errorf("Expected Placeholder2 to be replaced")
	}

	// Ensure Placeholder2 has been replaced with a non-empty value
	if !strings.Contains(resultStr, "<Placeholder2>") {
		t.Errorf("Expected Placeholder2 to have a replaced value")
	}

	// Ensure no blank lines are left after tag removal
	lines := strings.Split(resultStr, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

// TestReplaceTemplate_NestedPlaceholders verifies that blocking a parent placeholder
// removes its corresponding XML tag along with any nested child tags.
func TestReplaceTemplate_NestedPlaceholders(t *testing.T) {
	template := `<Parent>
	<Child>
		<GrandChild>{%GrandChild%}</GrandChild>
	</Child>
	<Child2>{%Child2%}</Child2>
	</Parent>`

	blocked := []string{"Child"}

	// Create options using WithBlockedPlaceholders
	options := []Option{
		WithBlockedPlaceholders(blocked...),
	}

	result, err := ReplaceTemplate(template, options...)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Check that Child tag is removed, including GrandChild
	if strings.Contains(resultStr, "<Child>") || strings.Contains(resultStr, "</Child>") || strings.Contains(resultStr, "<GrandChild>") {
		t.Errorf("Expected <Child> and nested <GrandChild> tags to be removed")
	}

	// Check that Child2 is still replaced
	if strings.Contains(resultStr, "{%Child2%}") {
		t.Errorf("Expected Child2 to be replaced")
	}

	// Ensure Child2 has been replaced with a non-empty value
	if !strings.Contains(resultStr, "<Child2>") {
		t.Errorf("Expected Child2 to have a replaced value")
	}

	// Ensure no blank lines are left after tag removal
	lines := strings.Split(resultStr, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

// TestReplaceTemplate_NoPlaceholders verifies that a template without any placeholders
// remains unchanged after processing.
func TestReplaceTemplate_NoPlaceholders(t *testing.T) {
	template := `<Test>
	<NoPlaceholder>Static Content</NoPlaceholder>
	</Test>`

	result, err := ReplaceTemplate(template)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Ensure the template remains unchanged
	if resultStr != template {
		t.Errorf("Expected template to remain unchanged")
	}
}

// TestReplaceTemplate_AllPlaceholdersBlocked verifies that blocking all placeholders
// results in all corresponding XML tags being removed from the generated output.
func TestReplaceTemplate_AllPlaceholdersBlocked(t *testing.T) {
	template := `<Test>
	<Placeholder1>{%Placeholder1%}</Placeholder1>
	<Placeholder2>{%Placeholder2%}</Placeholder2>
	</Test>`

	blocked := []string{"Placeholder1", "Placeholder2"}

	// Create options using WithBlockedPlaceholders
	options := []Option{
		WithBlockedPlaceholders(blocked...),
	}

	result, err := ReplaceTemplate(template, options...)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resultStr := string(result)

	// Check that all Placeholder tags are removed
	if strings.Contains(resultStr, "<Placeholder1>") || strings.Contains(resultStr, "</Placeholder1>") ||
		strings.Contains(resultStr, "<Placeholder2>") || strings.Contains(resultStr, "</Placeholder2>") {
		t.Errorf("Expected all Placeholder tags to be removed")
	}

	// Ensure no blank lines are left after tag removal
	lines := strings.Split(resultStr, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

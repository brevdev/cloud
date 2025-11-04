package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// ReadInputFileAsYaml reads a YAML file and returns a pointer to the root node
func ReadInputFileAsYaml(filename string) (*yaml.Node, error) {
	iFile, err := os.ReadFile(filename) //nolint:gosec // internal tool
	if err != nil {
		fmt.Println("Error reading swagger file:", err)
		return nil, err //nolint:wrapcheck // script
	}

	var iYaml yaml.Node
	if err := yaml.NewDecoder(bytes.NewReader(iFile)).Decode(&iYaml); err != nil {
		fmt.Println("Error decoding swagger file:", err)
		return nil, err
	}

	return &iYaml, nil
}

// WriteYamlToFile writes a YAML node to a file
func WriteYamlToFile(filename string, yamlNode *yaml.Node) error {
	oFile, err := os.Create(filename) //nolint:gosec // internal tool
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err //nolint:wrapcheck // script
	}
	defer func() {
		if err := oFile.Close(); err != nil {
			fmt.Println("Error closing output file:", err)
		}
	}()

	yamlEncoder := yaml.NewEncoder(oFile)
	yamlEncoder.SetIndent(2)

	if err := yamlEncoder.Encode(yamlNode); err != nil {
		fmt.Println("Error encoding swagger file:", err)
		return err //nolint:wrapcheck // script
	}

	return nil
}

// GetYamlNode gets a node from a YAML file by dot-separated path. For example, in the following YAML:
//
//	components:
//	  schemas:
//	    Foo:
//	      properties:
//	        bar:
//
// The path "components.schemas.Foo.properties.bar" would return the node for the "bar" property.
func GetYamlNode(yamlNode *yaml.Node, path string) *yaml.Node {
	// Split the path into pathParts
	pathParts := strings.Split(path, ".")

	var currentNode *yaml.Node
	// As a special case, if the node is a document (root) node, we need to get its first child node to begin
	// operating on the contents of the actual data.
	if yamlNode.Kind == yaml.DocumentNode {
		currentNode = yamlNode.Content[0]
	} else {
		currentNode = yamlNode
	}

	// For each path part, find the child node with the same value
	for _, pathPart := range pathParts {
		// If the node is not a mapping node, return early as we can't find the child node. In the future we may
		// want to support other node types (as in sequences), but for now we only support mapping nodes.
		if currentNode.Kind != yaml.MappingNode {
			return nil
		}

		var found *yaml.Node

		// In the go YAML library, mapping nodes are represented as a slice of key-value pairs. In this slice, entries
		// with even indices are keys, and entries with odd indices are values.
		for i := 0; i < len(currentNode.Content); i += 2 {
			key := currentNode.Content[i]
			val := currentNode.Content[i+1]

			// If the key matches the path part, we found the node we're looking for. Break out of the loop to allow
			// for the next path part to be searched.
			if key.Value == pathPart {
				found = val
				break
			}
		}

		// At this point if the "found" node is nil, then we never found a node with a key that matches the path part.
		// If the child node is not found, return early as we can't find the child node
		if found == nil {
			return nil
		}

		// We found the node matching the path part, continue with the next part
		currentNode = found
	}

	return currentNode
}

// GetYamlChildrenMap is a convenience function that returns a map of the children of a mapping node.
// In the go YAML library, mapping nodes are represented as a slice of key-value pairs. In this slice, entries
// with even indices are keys, and entries with odd indices are values. This function returns a map of the keys to
// the values.
func GetYamlChildrenMap(yamlNode *yaml.Node) map[string]*yaml.Node {
	if yamlNode.Kind != yaml.MappingNode {
		return nil
	}

	// We know that the number of children is half the length of the content slice, so we can pre-allocate the map
	// to the correct size.
	children := make(map[string]*yaml.Node, len(yamlNode.Content)/2)

	for i := 0; i < len(yamlNode.Content); i += 2 {
		children[yamlNode.Content[i].Value] = yamlNode.Content[i+1]
	}

	return children
}

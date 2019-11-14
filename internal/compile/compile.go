package compile

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/payfazz/fazz-swagger/internal/compile/model"
	"github.com/payfazz/fazz-swagger/internal/value"
	"gopkg.in/yaml.v2"
)

func walkMatch(root string, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func updateSwaggerFile(filePaths string) (*model.Swagger, error) {
	var swagger model.Swagger

	paths, err := walkMatch(filePaths, "swagger.yaml")
	if nil != err {
		return nil, err
	}

	for _, path := range paths {
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(yamlFile, &swagger)
		if err != nil {
			return nil, err
		}
	}

	return &swagger, err
}

func createPaths(filePaths string) (map[string]interface{}, error) {
	summary := make(map[string]interface{})

	paths, err := walkMatch(filePaths, "*.yaml")
	if nil != err {
		return nil, err
	}

	for _, path := range paths {
		if !strings.Contains(path, "path") {
			continue
		}

		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		var data map[string]interface{}
		err = yaml.Unmarshal(yamlFile, &data)
		if err != nil {
			return nil, err
		}

		for k, v := range data {
			if _, ok := summary[k]; !ok {
				summary[k] = v
			} else {
				child := v.(map[interface{}]interface{})
				childSummary := summary[k].(map[interface{}]interface{})
				for ck, cv := range child {
					childSummary[ck] = cv
				}
			}
		}
	}

	return summary, nil
}

func createComponentItem(filePaths string, componentType value.ComponentType) (map[string]interface{}, error) {
	component := make(map[string]interface{})

	paths, err := walkMatch(filePaths, string(componentType)+".yaml")
	if nil != err {
		return nil, err
	}

	for _, path := range paths {
		yamlFile, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(yamlFile, &component)
		if err != nil {
			return nil, err
		}
	}

	return component, nil
}

func Compile(filePaths string) {
	swagger, err := updateSwaggerFile(filePaths)
	if nil != err {
		log.Println("error fetching swagger file:", err)
	}

	var component model.Component

	paths, err := createPaths(filePaths)
	if nil != err {
		log.Println("creating path error:", err)
	}

	component.Schemas, err = createComponentItem(filePaths, value.COMPONENT_SCHEMAS)
	if nil != err {
		log.Println("creating Schemas error:", err)
	}
	component.Responses, err = createComponentItem(filePaths, value.COMPONENT_RESPONSES)
	if nil != err {
		log.Println("creating Responses error:", err)
	}
	component.Parameters, err = createComponentItem(filePaths, value.COMPONENT_PARAMETERS)
	if nil != err {
		log.Println("creating Parameters error:", err)
	}
	component.Examples, err = createComponentItem(filePaths, value.COMPONENT_EXAMPLES)
	if nil != err {
		log.Println("creating Examples error:", err)
	}
	component.RequestBodies, err = createComponentItem(filePaths, value.COMPONENT_REQUESTBODIES)
	if nil != err {
		log.Println("creating RequestBodies error:", err)
	}
	component.Headers, err = createComponentItem(filePaths, value.COMPONENT_HEADERS)
	if nil != err {
		log.Println("creating Headers error:", err)
	}
	component.SecuritySchemes, err = createComponentItem(filePaths, value.COMPONENT_SECURITYSCHEMES)
	if nil != err {
		log.Println("creating SecuritySchemes error:", err)
	}
	component.Links, err = createComponentItem(filePaths, value.COMPONENT_LINKS)
	if nil != err {
		log.Println("creating Links error:", err)
	}
	component.Callbacks, err = createComponentItem(filePaths, value.COMPONENT_CALLBACKS)
	if nil != err {
		log.Println("creating Callbacks error:", err)
	}

	swagger.Paths = &paths
	swagger.Components = &component

	swaggerData, err := yaml.Marshal(swagger)
	if nil != err {
		log.Println(err)
	}

	f, err := os.Create(filePaths + "fazz-swagger.yaml")
	if nil != err {
		log.Println(err)
	}

	_, err = f.WriteString(string(swaggerData))
	if nil != err {
		log.Println(err)
	}
}

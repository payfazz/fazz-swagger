
# fazz-swagger 3.0
Documentation tools to help creates API documentations with Swagger 3.0 ([OpenAPI 3.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)) easier.

## Features
`fazz-swagger` lets you split OAS3 Standard Yaml File to smaller components for easier documentations and readability. 

## How It Works
`fazz-swagger` compiles multiple YAML files into a single OAS3 Standard File.

## Example
```
ProjectName
    docs/
        api/
            user/
                paths/
                    all.yaml
                    create.yaml
                    update.yaml
                    delete.yaml
                schemas.yaml {optional}
            pet/
                paths/
                    list.yaml
        components/ {optional}
            schemas.yaml {optional}
            responses.yaml {optional}
            parameters.yaml {optional}
            examples.yaml {optional}
            requestBodies.yaml {optional}
            headers.yaml {optional}
            securitySchemes.yaml {optional}
            links.yaml {optional}
            callbacks.yaml {optional}
        swagger.yaml { See Example swagger.yaml file }
```

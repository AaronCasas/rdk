package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"
	"go.viam.com/utils"

	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
)

// UpdateActionType help hint the reconfigure process on whether one should reconfigure a resource or rebuild it.
type UpdateActionType int

const (
	// None is returned when the new configuration doesn't change the resource.
	None UpdateActionType = iota
	// Reconfigure is returned when the resource should be updated without recreating its proxies.
	// Note that two instances (old&new) will coexist, all dependencies will be destroyed and recreated.
	Reconfigure
	// Rebuild is returned when the resource and it's proxies should be destroyed and recreated,
	// all dependencies will be destroyed and recreated.
	Rebuild
)

// ComponentUpdate interface that a component can optionally implement.
// This interface helps the reconfiguration process.
type ComponentUpdate interface {
	UpdateAction(config *Component) UpdateActionType
}

type dependencyValidator interface {
	Validate(path string) ([]string, error)
}

type validator interface {
	Validate(path string) error
}

// A ResourceConfig represents an implmentation of a config for any type of resource.
type ResourceConfig interface {
	validator
	String() string
	ResourceName() resource.Name
	Get() interface{}
	Set(val string) error
}

// A ResourceLevelServiceConfig describes component or remote configuration for a service.
type ResourceLevelServiceConfig struct {
	Type                resource.SubtypeName `json:"type"`
	Attributes          AttributeMap         `json:"attributes"`
	ConvertedAttributes interface{}          `json:"-"`
}

// A Component describes the configuration of a component.
type Component struct {
	Name string `json:"name"`

	Namespace resource.Namespace   `json:"namespace"`
	Type      resource.SubtypeName `json:"type"`
	// TODO(PRODUCT-266): API replaces Type and Namespace when Service/Component merge, so json needs to be enabled.
	API           resource.Subtype             `json:"-"`
	Model         resource.Model               `json:"model"`
	Frame         *referenceframe.LinkConfig   `json:"frame,omitempty"`
	DependsOn     []string                     `json:"depends_on"`
	ServiceConfig []ResourceLevelServiceConfig `json:"service_config"`

	Attributes          AttributeMap `json:"attributes"`
	ConvertedAttributes interface{}  `json:"-"`
	ImplicitDependsOn   []string     `json:"-"`
}

// Dependencies returns the deduplicated union of user-defined and implicit dependencies.
func (config *Component) Dependencies() []string {
	result := make([]string, 0, len(config.DependsOn)+len(config.ImplicitDependsOn))
	seen := make(map[string]struct{})
	appendUniq := func(dep string) {
		if _, ok := seen[dep]; !ok {
			seen[dep] = struct{}{}
			result = append(result, dep)
		}
	}
	for _, dep := range config.DependsOn {
		appendUniq(dep)
	}
	for _, dep := range config.ImplicitDependsOn {
		appendUniq(dep)
	}
	return result
}

// Ensure Component conforms to flag.Value.
var _ = flag.Value(&Component{})

// String returns a verbose representation of the config.
func (config *Component) String() string {
	return fmt.Sprintf("%#v", config)
}

// ResourceName returns the  ResourceName for the component.
func (config *Component) ResourceName() resource.Name {
	if err := config.fixAPI(); err != nil {
		panic(err)
	}
	remotes := strings.Split(config.Name, ":")
	if len(remotes) > 1 {
		rName := resource.NameFromSubtype(config.API, remotes[len(remotes)-1])
		return rName.PrependRemote(resource.RemoteName(strings.Join(remotes[:len(remotes)-1], ":")))
	}
	return resource.NameFromSubtype(config.API, config.Name)
}

// Validate ensures all parts of the config are valid and returns dependencies.
func (config *Component) Validate(path string) ([]string, error) {
	var deps []string
	if err := config.fixAPI(); err != nil {
		return nil, err
	}

	if config.Namespace == "" {
		// NOTE: This should never be removed in order to ensure RDK is the
		// default namespace.
		config.Namespace = resource.ResourceNamespaceRDK
	}

	if config.Model.Namespace == "" {
		config.Model.Namespace = resource.ResourceNamespaceRDK
		config.Model.ModelFamily = resource.DefaultModelFamily
	}

	if config.Name == "" {
		return nil, utils.NewConfigValidationFieldRequiredError(path, "name")
	}
	if !ValidNameRegex.MatchString(config.Name) {
		return nil, ErrInvalidName(config.Name)
	}
	if err := resource.ContainsReservedCharacter(config.Name); err != nil {
		return nil, err
	}

	if err := config.Model.Validate(); err != nil {
		return nil, err
	}

	// this effectively checks reserved characters and the rest for namespace and type
	if err := config.API.Validate(); err != nil {
		return nil, err
	}

	for key, value := range config.Attributes {
		fieldPath := fmt.Sprintf("%s.%s", path, key)
		switch v := value.(type) {
		case validator:
			if err := v.Validate(fieldPath); err != nil {
				return nil, err
			}
		case dependencyValidator:
			validatedDeps, err := v.Validate(fieldPath)
			if err != nil {
				return nil, err
			}
			deps = append(deps, validatedDeps...)
		default:
			continue
		}
	}
	switch v := config.ConvertedAttributes.(type) {
	case validator:
		if err := v.Validate(path); err != nil {
			return nil, err
		}
	case dependencyValidator:
		validatedDeps, err := v.Validate(path)
		if err != nil {
			return nil, err
		}
		deps = append(deps, validatedDeps...)
	}
	return deps, nil
}

// Set hydrates a config based on a flag like value.
func (config *Component) Set(val string) error {
	parsed, err := ParseComponentFlag(val)
	if err != nil {
		return err
	}
	*config = parsed
	return nil
}

// Get gets the config itself.
func (config *Component) Get() interface{} {
	return config
}

// fixAPI updates Namespace and Type into the API (subtype field), or vice versa, depending on what fields are empty.
func (config *Component) fixAPI() error {
	//nolint:gocritic
	if config.API.Namespace == "" && config.Namespace == "" {
		config.Namespace = resource.ResourceNamespaceRDK
		config.API.Namespace = config.Namespace
	} else if config.API.Namespace == "" {
		config.API.Namespace = config.Namespace
	} else {
		config.Namespace = config.API.Namespace
	}

	if config.API.ResourceType == "" {
		config.API.ResourceType = resource.ResourceTypeComponent
	}

	if config.API.ResourceSubtype == "" {
		config.API.ResourceSubtype = config.Type
	} else if config.Type == "" {
		config.Type = config.API.ResourceSubtype
	}

	// this shouldn't be able to happen except with directly instantiated config structs
	if config.API.Namespace != config.Namespace || config.API.ResourceSubtype != config.Type {
		return errors.New("component namespace and/or type do not match component api field")
	}
	return nil
}

// ParseComponentFlag parses a component flag from command line arguments.
//
//nolint:dupl
func ParseComponentFlag(flag string) (Component, error) {
	// TODO(PRODUCT-266): Needs triplet support for model and api
	cmp := Component{}
	componentParts := strings.Split(flag, ",")
	for _, part := range componentParts {
		keyVal := strings.SplitN(part, "=", 2)
		if len(keyVal) != 2 {
			return Component{}, errors.New("wrong component format; use type=name,host=host,depends_on=name1|name2,attr=key:value")
		}
		switch keyVal[0] {
		case "name":
			cmp.Name = keyVal[1]
		case "type":
			cmp.Type = resource.SubtypeName(keyVal[1])
		case "model":
			m, err := resource.NewModelFromString(keyVal[1])
			if err != nil {
				return cmp, err
			}
			cmp.Model = m
		case "depends_on":
			split := strings.Split(keyVal[1], "|")

			var dependsOn []string
			for _, s := range split {
				if s != "" {
					dependsOn = append(dependsOn, s)
				}
			}
			cmp.DependsOn = dependsOn
		case "attr":
			if cmp.Attributes == nil {
				cmp.Attributes = AttributeMap{}
			}
			attrKeyVal := strings.SplitN(keyVal[1], ":", 2)
			if len(attrKeyVal) != 2 {
				return Component{}, errors.New("wrong attribute format; use attr=key:value")
			}
			cmp.Attributes[attrKeyVal[0]] = attrKeyVal[1]
		}
	}
	if string(cmp.Type) == "" {
		return Component{}, errors.New("component type is required")
	}
	return cmp, nil
}

// A Service describes the configuration of a service.
type Service struct {
	Name string `json:"name"`

	Namespace resource.Namespace   `json:"namespace"`
	Type      resource.SubtypeName `json:"type"`
	Model     resource.Model       `json:"model"`
	DependsOn []string             `json:"depends_on"`

	Attributes          AttributeMap `json:"attributes"`
	ConvertedAttributes interface{}  `json:"-"`
	ImplicitDependsOn   []string     `json:"-"`
}

// Ensure Service conforms to flag.Value.
var _ = flag.Value(&Service{})

// String returns a verbose representation of the config.
func (config *Service) String() string {
	return fmt.Sprintf("%#v", config)
}

// ResourceName returns the  ResourceName for the component.
func (config *Service) ResourceName() resource.Name {
	remotes := strings.Split(config.Name, ":")
	cType := string(config.Type)
	rName := resource.NewName(config.Namespace,
		resource.ResourceTypeService,
		resource.SubtypeName(cType),
		config.Name)
	if len(remotes) > 1 {
		return rName.PrependRemote(resource.RemoteName(strings.Join(remotes[:len(remotes)-1], ":")))
	}
	return rName
}

// ResourceName returns the  ResourceName for the component within a service_config.
func (config *ResourceLevelServiceConfig) ResourceName() resource.Name {
	cType := string(config.Type)
	// TODO(PRODUCT-266): needs triplet support here
	return resource.NewName(
		resource.ResourceNamespaceRDK,
		resource.ResourceTypeService,
		resource.SubtypeName(cType),
		cType,
	)
}

// Set hydrates a config based on a flag like value.
func (config *Service) Set(val string) error {
	parsed, err := ParseServiceFlag(val)
	if err != nil {
		return err
	}
	*config = parsed
	return nil
}

// Get gets the config itself.
func (config *Service) Get() interface{} {
	return config
}

// ParseServiceFlag parses a service flag from command line arguments.
//
//nolint:dupl
func ParseServiceFlag(flag string) (Service, error) {
	svc := Service{}
	serviceParts := strings.Split(flag, ",")
	for _, part := range serviceParts {
		keyVal := strings.SplitN(part, "=", 2)
		if len(keyVal) != 2 {
			return Service{}, errors.New("wrong service format; use type=name,depends_on=name1|name2,attr=key:value")
		}
		switch keyVal[0] {
		case "name":
			svc.Name = keyVal[1]
		case "type":
			svc.Type = resource.SubtypeName(keyVal[1])
		case "model":
			m, err := resource.NewModelFromString(keyVal[1])
			if err != nil {
				return svc, err
			}
			svc.Model = m
		case "depends_on":
			split := strings.Split(keyVal[1], "|")

			var dependsOn []string
			for _, s := range split {
				if s != "" {
					dependsOn = append(dependsOn, s)
				}
			}
			svc.DependsOn = dependsOn
		case "attr":
			if svc.Attributes == nil {
				svc.Attributes = AttributeMap{}
			}
			attrKeyVal := strings.SplitN(keyVal[1], ":", 2)
			if len(attrKeyVal) != 2 {
				return Service{}, errors.New("wrong attribute format; use attr=key:value")
			}
			svc.Attributes[attrKeyVal[0]] = attrKeyVal[1]
		}
	}
	if string(svc.Type) == "" {
		return Service{}, errors.New("service type is required")
	}
	return svc, nil
}

// Validate ensures all parts of the config are valid.
func (config *Service) Validate(path string) ([]string, error) {
	if config.Type == "" {
		return nil, utils.NewConfigValidationFieldRequiredError(path, "type")
	}
	// If services do not have a name use the name builtin
	if config.Name == "" {
		golog.Global().Debugw("no name given, defaulting name to builtin")
		config.Name = resource.DefaultServiceName
	}
	if !ValidNameRegex.MatchString(config.Name) {
		return nil, ErrInvalidName(config.Name)
	}
	if config.Model.Name == "" {
		golog.Global().Debugw("no model given; using default")
		config.Model = resource.DefaultServiceModel
	}

	if config.Model.Namespace == "" {
		config.Model.Namespace = resource.ResourceNamespaceRDK
		config.Model.ModelFamily = resource.DefaultModelFamily
	}

	if config.Namespace == "" {
		// NOTE: This should never be removed in order to ensure RDK is the
		// default namespace.
		config.Namespace = resource.ResourceNamespaceRDK
	}
	if err := resource.ContainsReservedCharacter(string(config.Namespace)); err != nil {
		return nil, err
	}
	if err := resource.ContainsReservedCharacter(config.Name); err != nil {
		return nil, err
	}

	if err := config.Model.Validate(); err != nil {
		return nil, err
	}

	var deps []string
	for key, value := range config.Attributes {
		switch v := value.(type) {
		case validator:
			if err := v.Validate(fmt.Sprintf("%s.%s", path, key)); err != nil {
				return nil, err
			}
		case dependencyValidator:
			validatedDeps, err := v.Validate(fmt.Sprintf("%s.%s", path, key))
			if err != nil {
				return nil, err
			}
			deps = append(deps, validatedDeps...)
		default:
			continue
		}
	}
	switch v := config.ConvertedAttributes.(type) {
	case validator:
		if err := v.Validate(path); err != nil {
			return nil, err
		}
	case dependencyValidator:
		validatedDeps, err := v.Validate(path)
		if err != nil {
			return nil, err
		}
		deps = append(deps, validatedDeps...)
	}
	return deps, nil
}

// Dependencies returns the deduplicated union of user-defined and implicit dependencies.
func (config *Service) Dependencies() []string {
	result := make([]string, 0, len(config.DependsOn)+len(config.ImplicitDependsOn))
	seen := make(map[string]struct{})
	appendUniq := func(dep string) {
		if _, ok := seen[dep]; !ok {
			seen[dep] = struct{}{}
			result = append(result, dep)
		}
	}
	for _, dep := range config.DependsOn {
		appendUniq(dep)
	}
	for _, dep := range config.ImplicitDependsOn {
		appendUniq(dep)
	}
	return result
}

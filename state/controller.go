// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"github.com/juju/errors"

	"github.com/juju/juju/environs/config"
)

// controllerSettingsGlobalKey is the key for the controller and its settings.
const controllerSettingsGlobalKey = "controllerSettings"

func controllerOnlyAttribute(attr string) bool {
	for _, a := range config.ControllerOnlyConfigAttributes {
		if attr == a {
			return true
		}
	}
	return false
}

// controllerConfig returns the controller config attributes that result when we have
// have a current config and want to save a new config, possible overwriting some current values.
func controllerConfig(currentControllerCfg, cfg map[string]interface{}) map[string]interface{} {
	controllerCfg := make(map[string]interface{})

	if len(currentControllerCfg) == 0 {
		// No controller config yet, so we are setting up a
		// new controller. We'll grab the controller config
		// attributes from the passed in config.
		for _, attr := range config.ControllerOnlyConfigAttributes {
			controllerCfg[attr] = cfg[attr]
		}
	} else {
		// Copy across attributes only valid for the controller config.
		for _, attr := range config.ControllerOnlyConfigAttributes {
			if v, ok := currentControllerCfg[attr]; ok {
				controllerCfg[attr] = v
			}
		}
	}
	return controllerCfg
}

// modelConfig returns the model config attributes that result when we
// have a current controller config and want to save a new model config.
// currentControllerCfg is not currently used - it will be when we support inheritance.
func modelConfig(currentControllerCfg, cfg map[string]interface{}) map[string]interface{} {
	modelCfg := make(map[string]interface{})
	// The model config contains any attributes not controller only.
	for attr, value := range cfg {
		if controllerOnlyAttribute(attr) {
			continue
		}
		modelCfg[attr] = value
	}
	return modelCfg
}

// ControllerConfig returns the config values for the controller.
func (st *State) ControllerConfig() (map[string]interface{}, error) {
	settings, err := readSettings(st, controllersC, controllerSettingsGlobalKey)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return settings.Map(), nil
}

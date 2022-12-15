package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobInput = MLTableJobInput{}

type MLTableJobInput struct {
	Mode *InputDeliveryMode `json:"mode,omitempty"`
	Uri  string             `json:"uri"`

	// Fields inherited from JobInput
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = MLTableJobInput{}

func (s MLTableJobInput) MarshalJSON() ([]byte, error) {
	type wrapper MLTableJobInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MLTableJobInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MLTableJobInput: %+v", err)
	}
	decoded["jobInputType"] = "mltable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MLTableJobInput: %+v", err)
	}

	return encoded, nil
}
// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

// MasterRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/cat/master/types.ts#L20-L39
type MasterRecord struct {
	// Host host name
	Host *string `json:"host,omitempty"`
	// Id node id
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
}

// MasterRecordBuilder holds MasterRecord struct and provides a builder API.
type MasterRecordBuilder struct {
	v *MasterRecord
}

// NewMasterRecord provides a builder for the MasterRecord struct.
func NewMasterRecordBuilder() *MasterRecordBuilder {
	r := MasterRecordBuilder{
		&MasterRecord{},
	}

	return &r
}

// Build finalize the chain and returns the MasterRecord struct
func (rb *MasterRecordBuilder) Build() MasterRecord {
	return *rb.v
}

// Host host name

func (rb *MasterRecordBuilder) Host(host string) *MasterRecordBuilder {
	rb.v.Host = &host
	return rb
}

// Id node id

func (rb *MasterRecordBuilder) Id(id string) *MasterRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Ip ip address

func (rb *MasterRecordBuilder) Ip(ip string) *MasterRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Node node name

func (rb *MasterRecordBuilder) Node(node string) *MasterRecordBuilder {
	rb.v.Node = &node
	return rb
}
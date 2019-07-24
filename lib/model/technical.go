/*
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

type Content struct {
	Id              string          `json:"id"`
	ContentVariable ContentVariable `json:"content_variable"`
}

type ValueType string

const (
	String  ValueType = "https://schema.org/Text"
	Integer ValueType = "https://schema.org/Integer"
	Float   ValueType = "https://schema.org/Float"
	Boolean ValueType = "https://schema.org/Boolean"

	List      ValueType = "https://schema.org/ItemList"
	Structure ValueType = "https://schema.org/StructuredValue"
)

type ContentVariable struct {
	Id                  string            `json:"id"`
	Type                ValueType         `json:"type"`
	SubContentVariables []ContentVariable `json:"sub_content_variables"`
	ExactMatch          string            `json:"exact_match"`
}
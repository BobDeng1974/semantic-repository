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

package lib

import (
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/controller"
	"github.com/SENERGY-Platform/semantic-repository/lib/database"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/SENERGY-Platform/semantic-repository/lib/source/producer"
	"testing"
)

func TestProduceValidDeviceType(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:device1"
	devicetype.Name = "Device1"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:device-1-1",
		"Device1.localId1",
		"Device1.service1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e-1", Name: "brightnessAdjustment1", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:device-1-2",
		"Device1.localId2",
		"Device1.service2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e-2", Name: "brightnessAdjustment2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")

	devicetype1 := model.DeviceType{}
	devicetype1.Id = "urn:infai:ses:devicetype:device2"
	devicetype1.Name = "Device2"
	devicetype1.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2e2e",
		Name: "Lamp",
	}
	devicetype1.Description = "description"
	devicetype1.Image = "image"
	devicetype1.Services = []model.Service{}
	devicetype1.Services = append(devicetype1.Services, model.Service{
		"urn:infai:ses:service:device-2-1",
		"Device2.localId1",
		"Device2.service1",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e-1", Name: "brightnessAdjustment1", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype1.Services = append(devicetype1.Services, model.Service{
		"urn:infai:ses:service:device-2-2",
		"Device2.localId2",
		"Device2.service2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4e4e", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5e5e-2", Name: "brightnessAdjustment2", ConceptId: "urn:ses:infai:concept:1a1a1a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype1, "sdfdsfsf")
}

func TestReadDeviceTypesWithDeviceClassIdAndOneFunctionId(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceTypesFiltered("urn:infai:ses:deviceclass:2e2e", []string{"urn:infai:ses:function:5e5e-1"}, []string{})

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if deviceType[0].Id != "urn:infai:ses:devicetype:device1" {
		t.Fatal("error id")
	}

	if deviceType[0].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[0].Name != "Device1" {
		t.Fatal("error name")
	}

	if deviceType[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[0].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[0].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[0].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[0].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType[0].Services[0].Id != "urn:infai:ses:service:device-1-1" {
		t.Fatal("error service -> 0 -> id", deviceType[0].Services[0].Id)
	}
	if deviceType[0].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[0].Name != "Device1.service1" {
		t.Log(deviceType[0].Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}

	if deviceType[1].Id != "urn:infai:ses:devicetype:device2" {
		t.Fatal("error id")
	}

	if deviceType[1].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[1].Name != "Device2" {
		t.Fatal("error name", deviceType[1].Name)
	}

	if deviceType[1].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[1].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[1].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[1].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[1].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType[1].Services[0].Id != "urn:infai:ses:service:device-2-1" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[0].Name != "Device2.service1" {
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e-1" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Functions[0].Name != "brightnessAdjustment1" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		b, err := json.Marshal(deviceType)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Log(string(b))
	}
}

func TestReadDeviceTypesWithDeviceClassIdAndTwoFunctionIds(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceTypesFiltered("urn:infai:ses:deviceclass:2e2e", []string{"urn:infai:ses:function:5e5e-1", "urn:infai:ses:function:5e5e-2"}, []string{})

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if deviceType[0].Id != "urn:infai:ses:devicetype:device1" {
		t.Fatal("error id")
	}

	if deviceType[0].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[0].Name != "Device1" {
		t.Fatal("error name")
	}

	if deviceType[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[0].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[0].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[0].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[0].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service 0
	if deviceType[0].Services[0].Id != "urn:infai:ses:service:device-1-2" {
		t.Fatal("error service -> 0 -> id", deviceType[0].Services[0].Id)
	}
	if deviceType[0].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[0].Name != "Device1.service2" {
		t.Log(deviceType[0].Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}

	// Service 0
	if deviceType[0].Services[1].Id != "urn:infai:ses:service:device-1-1" {
		t.Fatal("error service -> 0 -> id", deviceType[0].Services[0].Id)
	}
	if deviceType[0].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[1].Name != "Device1.service1" {
		t.Log(deviceType[0].Services[1].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[1].Description != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}

	// deviceType 1

	if deviceType[1].Id != "urn:infai:ses:devicetype:device2" {
		t.Fatal("error id")
	}

	if deviceType[1].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[1].Name != "Device2" {
		t.Fatal("error name", deviceType[1].Name)
	}

	if deviceType[1].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[1].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[1].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[1].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[1].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service 0
	if deviceType[1].Services[0].Id != "urn:infai:ses:service:device-2-1" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[0].Name != "Device2.service1" {
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e-1" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Functions[0].Name != "brightnessAdjustment1" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}

	// Service 1
	if deviceType[1].Services[1].Id != "urn:infai:ses:service:device-2-2" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[1].Name != "Device2.service2" {
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[1].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e-2" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[1].Functions[0].Name != "brightnessAdjustment2" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		b, err := json.Marshal(deviceType)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Log(string(b))
	}
}

func TestReadDeviceTypesWithAspect(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceTypesFiltered("", []string{}, []string{"urn:infai:ses:aspect:4e4e"})

	b, err := json.Marshal(deviceType)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Log(string(b))

	if deviceType[0].Id != "urn:infai:ses:devicetype:device1" {
		t.Fatal("error id")
	}

	if deviceType[0].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[0].Name != "Device1" {
		t.Fatal("error name")
	}

	if deviceType[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[0].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[0].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[0].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[0].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service 0
	if deviceType[0].Services[0].Id != "urn:infai:ses:service:device-1-2" {
		t.Fatal("error service -> 0 -> id", deviceType[0].Services[0].Id)
	}
	if deviceType[0].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[0].Name != "Device1.service2" {
		t.Log(deviceType[0].Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[0].Description != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}

	// Service 0
	if deviceType[0].Services[1].Id != "urn:infai:ses:service:device-1-1" {
		t.Fatal("error service -> 0 -> id", deviceType[0].Services[0].Id)
	}
	if deviceType[0].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[0].Services[1].Name != "Device1.service1" {
		t.Log(deviceType[0].Services[1].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[0].Services[1].Description != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[0].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[0].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[0].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[0].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}

	// deviceType 1

	if deviceType[1].Id != "urn:infai:ses:devicetype:device2" {
		t.Fatal("error id")
	}

	if deviceType[1].RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType[1].Name != "Device2" {
		t.Fatal("error name", deviceType[1].Name)
	}

	if deviceType[1].Description != "" { // not stored as TRIPLE
		t.Fatal("error description")
	}

	if deviceType[1].Image != "" { // not stored as TRIPLE
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType[1].DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType[1].DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType[1].DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service 0
	if deviceType[1].Services[0].Id != "urn:infai:ses:service:device-2-1" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[0].Name != "Device2.service1" {
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e-1" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[0].Functions[0].Name != "brightnessAdjustment1" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}

	// Service 1
	if deviceType[1].Services[1].Id != "urn:infai:ses:service:device-2-2" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType[1].Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType[1].Services[1].Name != "Device2.service2" {
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType[1].Services[1].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType[1].Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType[1].Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType[1].Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType[0].Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType[1].Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e-2" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType[1].Services[1].Functions[0].Name != "brightnessAdjustment2" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType[1].Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType[1].Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		b, err := json.Marshal(deviceType)
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Log(string(b))
	}
}

func TestReadDeviceTypeWithId1(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:device1")

	if deviceType.Id != "urn:infai:ses:devicetype:device1" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Device1" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:device-1-1" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "Device1.service1" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e-1" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment1" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType.Services[1].Id != "urn:infai:ses:service:device-1-2" {
		t.Fatal("error service -> 1 -> id")
	}
	if deviceType.Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType.Services[1].Name != "Device1.service2" {
		t.Log(deviceType.Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType.Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType.Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType.Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType.Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType.Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e-2" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType.Services[1].Functions[0].Name != "brightnessAdjustment2" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType.Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func TestReadDeviceTypeWithId2(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:device2")

	if deviceType.Id != "urn:infai:ses:devicetype:device2" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Device2" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2e2e" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:device-2-1" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "Device2.service1" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5e5e-1" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment1" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds")
	}
	/// service 2
	if deviceType.Services[1].Id != "urn:infai:ses:service:device-2-2" {
		t.Fatal("error service -> 1 -> id")
	}
	if deviceType.Services[1].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 1 -> RdfType")
	}
	if deviceType.Services[1].Name != "Device2.service2" {
		t.Log(deviceType.Services[1].Name)
		t.Fatal("error service -> 1 -> name")
	}
	if deviceType.Services[1].Description != "" {
		t.Fatal("error service -> 1 -> description")
	}
	if deviceType.Services[1].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 1 -> LocalId")
	}
	if deviceType.Services[1].Aspects[0].Id != "urn:infai:ses:aspect:4e4e" {
		t.Fatal("error aspect -> 1/0 -> id")
	}
	if deviceType.Services[1].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[1].Aspects[0].Name)
		t.Fatal("error aspect -> 1/0 -> Name")
	}
	if deviceType.Services[1].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].Id != "urn:infai:ses:function:5e5e-2" {
		t.Fatal("error function -> 1/0 -> id")
	}
	if deviceType.Services[1].Functions[0].Name != "brightnessAdjustment2" {
		t.Fatal("error function -> 1/0 -> Name")
	}
	if deviceType.Services[1].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 1/0 -> RdfType")
	}
	if deviceType.Services[1].Functions[0].ConceptId != "urn:ses:infai:concept:1a1a1a" {
		t.Fatal("error function -> 1/0/0 -> ConceptIds")
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func TestCreateAndDeleteDeviceTypePart1(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3a",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4a", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5a", Name: "brightnessAdjustment", ConceptId: "urn:infai:ses:concept:6a", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3b",
		"localId",
		"setBrightness",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4b", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5b", Name: "brightnessAdjustment", ConceptId: "urn:infai:ses:concept:6b", RdfType: model.SES_ONTOLOGY_MEASURING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
}

func TestCreateAndDeleteDeviceTypePart2(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	err = producer.PublishDeviceTypeDelete("urn:infai:ses:devicetype:1", "sdfdsfsf")
	if err != nil {
		t.Fatal(err)
	}
}

func TestProduceValidDeviceTypeWithoutConceptId(t *testing.T) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	producer, _ := producer.New(conf)
	devicetype := model.DeviceType{}
	devicetype.Id = "urn:infai:ses:devicetype:1_4-12-2019"
	devicetype.Name = "Philips Hue Color"
	devicetype.DeviceClass = model.DeviceClass{
		Id:   "urn:infai:ses:deviceclass:2_4-12-2019",
		Name: "Lamp",
	}
	devicetype.Description = "description"
	devicetype.Image = "image"
	devicetype.Services = []model.Service{}
	devicetype.Services = append(devicetype.Services, model.Service{
		"urn:infai:ses:service:3_4-12-2019",
		"localId",
		"setBrightness2",
		"",
		[]model.Aspect{{Id: "urn:infai:ses:aspect:4_4-12-2019", Name: "Lighting", RdfType: "asasasdsadas"}},
		"asdasda",
		[]model.Content{},
		[]model.Content{},
		[]model.Function{{Id: "urn:infai:ses:function:5_4-12-2019", Name: "brightnessAdjustment", RdfType: model.SES_ONTOLOGY_CONTROLLING_FUNCTION}},
		"asdasdsadsadasd",
	})

	producer.PublishDeviceType(devicetype, "sdfdsfsf")
}

func TestReadDeviceTypeWithoutConceptId(t *testing.T) {
	err, con, _ := StartUpScript(t)
	deviceType, err, code := con.GetDeviceType("urn:infai:ses:devicetype:1_4-12-2019")

	t.Log(deviceType)

	if deviceType.Id != "urn:infai:ses:devicetype:1_4-12-2019" {
		t.Fatal("error id")
	}

	if deviceType.RdfType != model.SES_ONTOLOGY_DEVICE_TYPE {
		t.Fatal("error model")
	}

	if deviceType.Name != "Philips Hue Color" {
		t.Fatal("error name")
	}

	if deviceType.Description != "" {
		t.Fatal("error description")
	}

	if deviceType.Image != "" {
		t.Fatal("error image")
	}
	// DeviceClass
	if deviceType.DeviceClass.Id != "urn:infai:ses:deviceclass:2_4-12-2019" {
		t.Fatal("error deviceclass id")
	}
	if deviceType.DeviceClass.Name != "Lamp" {
		t.Fatal("error deviceclass name")
	}
	if deviceType.DeviceClass.RdfType != model.SES_ONTOLOGY_DEVICE_CLASS {
		t.Fatal("error deviceclass rdf type")
	}
	// Service
	if deviceType.Services[0].Id != "urn:infai:ses:service:3_4-12-2019" {
		t.Fatal("error service -> 0 -> id")
	}
	if deviceType.Services[0].RdfType != model.SES_ONTOLOGY_SERVICE {
		t.Fatal("error service -> 0 -> RdfType")
	}
	if deviceType.Services[0].Name != "setBrightness2" {
		t.Log(deviceType.Services[0].Name)
		t.Fatal("error service -> 0 -> name")
	}
	if deviceType.Services[0].Description != "" {
		t.Fatal("error service -> 0 -> description")
	}
	if deviceType.Services[0].LocalId != "" { // not stored as TRIPLE
		t.Fatal("error service -> 0 -> LocalId")
	}
	if deviceType.Services[0].Aspects[0].Id != "urn:infai:ses:aspect:4_4-12-2019" {
		t.Fatal("error aspect -> 0/0 -> id")
	}
	if deviceType.Services[0].Aspects[0].Name != "Lighting" {
		t.Log(deviceType.Services[0].Aspects[0].Name)
		t.Fatal("error aspect -> 0/0 -> Name")
	}
	if deviceType.Services[0].Aspects[0].RdfType != model.SES_ONTOLOGY_ASPECT {
		t.Fatal("error aspect -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].Id != "urn:infai:ses:function:5_4-12-2019" {
		t.Fatal("error function -> 0/0 -> id")
	}
	if deviceType.Services[0].Functions[0].Name != "brightnessAdjustment" {
		t.Fatal("error function -> 0/0 -> Name")
	}
	if deviceType.Services[0].Functions[0].RdfType != model.SES_ONTOLOGY_CONTROLLING_FUNCTION {
		t.Fatal("error function -> 0/0 -> RdfType")
	}
	if deviceType.Services[0].Functions[0].ConceptId != "" {
		t.Fatal("error function -> 0/0/0 -> ConceptIds", deviceType.Services[0].Functions[0].ConceptId)
	}

	if err != nil {
		t.Fatal(deviceType, err, code)
	} else {
		t.Log(deviceType)
	}
}

func StartUpScript(t *testing.T) (error, *controller.Controller, database.Database) {
	conf, err := config.Load("../config.json")
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = producer.New(conf)
	if err != nil {
		t.Fatal(err)
	}
	con, err := controller.New(conf, db)
	if err != nil {
		t.Fatal(err)
	}
	return err, con, db
}

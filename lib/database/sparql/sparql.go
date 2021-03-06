/*
 *
 * Copyright 2019 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *
 */

package sparql

import (
	"bytes"
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"github.com/SENERGY-Platform/semantic-repository/lib/model"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Database struct {
	conf config.Config
}

func New(conf config.Config) (db *Database, err error) {
	return &Database{conf: conf}, nil
}

func (*Database) Disconnect() {
	panic("implement me")
}

func (*Database) GetDevice(ctx context.Context, id string) (device model.Device, exists bool, err error) {
	panic("implement me")
}

func (*Database) SetDevice(ctx context.Context, device model.Device) error {
	panic("implement me")
}

func (*Database) RemoveDevice(ctx context.Context, id string) error {
	panic("implement me")
}

func (*Database) SetDeviceType(ctx context.Context, deviceType model.DeviceType) error {
	panic("implement me")
}

func (*Database) RemoveDeviceType(ctx context.Context, id string) error {
	panic("implement me")
}

func (this *Database) InsertData(triples string) (err error) {
	requestBody := []byte(triples)
	resp, err := http.Post(this.conf.RyaUrl+"/web.rya/loadrdf?format=N-Triples", "text/plain", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("ERROR: Statuscode " + string(resp.StatusCode))
	}
}

func (this *Database) DeleteDeviceType(s string) (err error) {
	query := this.getDeleteDeviceTypeQuery(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("ERROR: Statuscode " + string(resp.StatusCode))
	}
}

func (this *Database) GetDeviceType(deviceTypeId string, deviceClassId string, functionIds []string, aspectIds []string) (rdfxml string, err error) {

	query := this.getDeviceTypeQuery(deviceTypeId, deviceClassId, functionIds, aspectIds)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetDeviceClassesFunctions(s string) (rdfxml string, err error) {
	query := this.getDeviceClassesFunctions(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetDeviceClassesControllingFunctions(s string) (rdfxml string, err error) {
	query := this.getDeviceClassesControllingFunctions(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetAspectsMeasuringFunctions(s string) (rdfxml string, err error) {
	query := this.getAspectsMeasuringFunctions(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) DeleteConcept(s string, deleteNested bool) (err error) {
	query := ""
	if deleteNested {
		query = this.getDeleteConceptWithNestedQuery(s)
	} else {
		query = this.getDeleteConceptWithouthNestedQuery(s)
	}
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("ERROR: Statuscode " + string(resp.StatusCode))
	}
}

func (this *Database) DeleteCharacteristic(s string) (err error) {
	query := this.getDeleteCharacteristicQuery(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("ERROR: Statuscode " + string(resp.StatusCode))
	}
}

func (this *Database) GetListWithoutSubProperties(p string, o string) (rdfxml string, err error) {
	query := this.getConstructListWithoutSubProperties(p, o)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetWithAllSubProperties(subject string) (rdfxml string, err error) {
	query := this.getConstructWithAllSubProperties(subject)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetFunctions", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func (this *Database) GetWithoutSubProperties(s string) (rdfxml string, err error) {
	query := this.getConstructWithoutSubProperties(s)
	resp, err := http.Get(this.conf.RyaUrl + "/web.rya/queryrdf?query=" + query)
	if err != nil {
		log.Println("ERROR: GetConcept", err)
		return "", err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	return string(byteArray), nil
}

func getTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

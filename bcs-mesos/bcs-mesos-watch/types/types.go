/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

//BcsSyncData holder for sync data
type BcsSyncData struct {
	DataType string      //data type: reflect.TypeOf(Item).Name()
	Action   string      //operation, like Add, Delete, Update
	Item     interface{} //SyncData, data is Endpoint, Service, Pod
}

//CmdConfig hold all command line config item
type CmdConfig struct {
	ClusterID   string
	ClusterInfo string
	CAFile      string
	CertFile    string
	KeyFile     string
	PassWord    string

	RegDiscvSvr            string
	Address                string
	ApplicationThreadNum   int
	TaskgroupThreadNum     int
	ExportserviceThreadNum int

	MetricPort uint

	ServerCAFile   string
	ServerCertFile string
	ServerKeyFile  string
	ServerPassWord string
	ServerSchem    string
}

const (
	ApplicationChannelPrefix   = "Application_"
	TaskgroupChannelPrefix     = "TaskGroup_"
	ExportserviceChannelPrefix = "Exportservice_"
)
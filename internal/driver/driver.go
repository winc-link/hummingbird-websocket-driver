/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package driver

import (
	"context"
	"github.com/winc-link/hummingbird-sdk-go/commons"
	"github.com/winc-link/hummingbird-sdk-go/model"
	"github.com/winc-link/hummingbird-sdk-go/service"
	"github.com/winc-link/hummingbird-websocket-driver/internal/device"
	"github.com/winc-link/hummingbird-websocket-driver/internal/server"
)

type WebsocketProtocolDriver struct {
	sd *service.DriverService
}

// CloudPluginNotify 云插件启动/停止通知
func (dr WebsocketProtocolDriver) CloudPluginNotify(ctx context.Context, t commons.CloudPluginNotifyType, name string) error {
	//TODO implement me
	panic("implement me")
}

// DeviceNotify 设备添加/修改/删除通知
func (dr WebsocketProtocolDriver) DeviceNotify(ctx context.Context, t commons.DeviceNotifyType, deviceId string, device model.Device) error {
	//TODO implement me
	panic("implement me")
}

// ProductNotify 产品添加/修改/删除通知
func (dr WebsocketProtocolDriver) ProductNotify(ctx context.Context, t commons.ProductNotifyType, productId string, product model.Product) error {
	//TODO implement me
	panic("implement me")
}

// Stop 驱动退出通知。
func (dr WebsocketProtocolDriver) Stop(ctx context.Context) error {
	for _, dev := range device.GetAllDevice() {
		dr.sd.Offline(dev.GetDeviceId())
	}
	return nil
}

// HandlePropertySet 设备属性设置
func (dr WebsocketProtocolDriver) HandlePropertySet(ctx context.Context, deviceId string, data model.PropertySet) error {
	return nil
}

// HandlePropertyGet 设备属性查询
func (dr WebsocketProtocolDriver) HandlePropertyGet(ctx context.Context, deviceId string, data model.PropertyGet) error {
	//TODO implement me
	panic("implement me")
}

// HandleServiceExecute 设备服务调用
func (dr WebsocketProtocolDriver) HandleServiceExecute(ctx context.Context, deviceId string, data model.ServiceExecuteRequest) error {
	//TODO implement me
	panic("implement me")
}

// NewWebsocketProtocolDriver Websocket协议驱动
func NewWebsocketProtocolDriver(sd *service.DriverService) *WebsocketProtocolDriver {
	loadDevices(sd)
	go server.NewWebsocketService(sd).Start()
	return &WebsocketProtocolDriver{
		sd: sd,
	}
}

// loadDevices 获取所有已经创建成功的设备，保存在内存中。
func loadDevices(sd *service.DriverService) {
	for _, dev := range sd.GetDeviceList() {
		device.PutDevice(dev.DeviceSn, device.NewDevice(dev.Id, dev.DeviceSn, dev.ProductId, dev.Status == commons.DeviceOnline))
	}
}

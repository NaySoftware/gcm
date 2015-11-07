
package gcm

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
  // "errors"
  "encoding/json"
)

const (
  server_url = "https://gcm-http.googleapis.com/gcm/send"
  MAX_TTL = 2419200
  Priority_HIGH   = "high"
  Priority_NORMAL = "normal"
)

type GcmClient struct {
  ApiKey string
  Message GcmMsg
}

type GcmMsg struct {
  Data map[string]string            `json:"data,omitempty"`
  To string                         `json:"to,omitempty"`
  RegistrationIds []string          `json:"registration_ids,omitempty"`
  CollapseKey   string              `json:"collapse_key,omitempty"`
  Priority string                   `json:"priority,omitempty"`
  Notification NotificationPayload  `json:"notification,omitempty"`
  ContentAvailable bool             `json:"content_available,omitempty"`
  DelayWhileIdle  bool              `json:"delay_while_idle,omitempty"`
  TimeToLive    int                 `json:"time_to_live,omitempty"`
  RestrictedPackageName string      `json:"restricted_package_name,omitempty"`
  DryRun      bool                  `json:"dry_run,omitempty"`

}

type GcmResponseStatus struct {
  Ok bool   
  StatusCode    int
  MulticastId int                   `json:"multicast_id"`
  Success int                       `json:"success"`
  Fail int                          `json:"failure"`
  Canonical_ids int                 `json:"canonical_ids"`
  Results []map[string]string       `json:"results,omitempty"`
} 


type NotificationPayload struct {
  Title string   `json:"title,omitempty"`
  Body  string   `json:"body,omitempty"`
  Icon  string   `json:"icon,omitempty"`
  Sound string   `json:"sound,omitempty"`
  Badge string   `json:"badge,omitempty"`
  Tag   string   `json:"tag,omitempty"`
  Color string   `json:"color,omitempty"`
  ClickAction string  `json:"click_action,omitempty"`
  BodyLocKey  string  `json:"body_loc_key,omitempty"`
  BodyLocArgs string  `json:"body_loc_args,omitempty"`
  TitleLocKey string  `json:"title_loc_key,omitempty"`
  TitleLocArgs string `json:"title_loc_args,omitempty"`

}

func NewGcmClient(apiKey string) (*GcmClient) {
  gcmc := new(GcmClient)
  gcmc.ApiKey = apiKey

  return gcmc
}


func (this *GcmClient) NewGcmMsgTo(to string, body map[string]string) (*GcmClient) {
  this.Message.To = to
  this.Message.Data = body

  return this
}

func (this *GcmClient) SetMsgData(body map[string]string) (*GcmClient) {
  
  this.Message.Data = body

  return this

}



func (this *GcmClient) NewGcmRegIdsMsg(list []string, body map[string]string) (*GcmClient) {
  this.NewDevicesList(list)
  this.Message.Data = body

  return this

}


func (this *GcmClient) NewDevicesList(list []string) (*GcmClient) {
  this.Message.RegistrationIds = make([]string, len(list))
  copy(this.Message.RegistrationIds, list)

  return this

}

func (this *GcmClient) AppendDevices(list []string) (*GcmClient)  {

  this.Message.RegistrationIds = append(this.Message.RegistrationIds, list...)

  return this
}


func (this *GcmClient) Send() (*GcmResponseStatus, error) {

  gcmRespStatus := new(GcmResponseStatus)


  jsonByte , err := this.Message.toJsonByte()
  if err != nil {
    gcmRespStatus.Ok = false
    return gcmRespStatus, err 
  }

  fmt.Println(string(jsonByte))

  request, err := http.NewRequest("POST", server_url , bytes.NewBuffer(jsonByte))
  request.Header.Set("Authorization", this.apiKeyHeader())
  request.Header.Set("Content-Type", "application/json")


  client := &http.Client{}
  response, err := client.Do(request)

  if err != nil {
    gcmRespStatus.Ok = false
    return gcmRespStatus, err
  }
  defer response.Body.Close()


  body, err := ioutil.ReadAll(response.Body)

  gcmRespStatus.StatusCode = response.StatusCode

  if response.StatusCode == 200 && err == nil {


  gcmRespStatus.Ok = true

  // fmt.Println(response.Status)
  eror := gcmRespStatus.parseStatusBody(body)
  if eror != nil {
    return gcmRespStatus, eror
  }

  return gcmRespStatus, nil

  } else  {
    gcmRespStatus.Ok = false

    eror := gcmRespStatus.parseStatusBody(body)
  if eror != nil {
    return gcmRespStatus, eror
  }

    return gcmRespStatus, err 

  }

}

func (this *GcmMsg) toJsonByte() ([]byte, error) {

  return json.Marshal(this)

}


func (this *GcmResponseStatus) parseStatusBody(body []byte ) (error)  {

  if err := json.Unmarshal([]byte(body), &this); err != nil {
    return err
  } 
  return nil

}

func (this *GcmClient) apiKeyHeader() string {
  return fmt.Sprintf("key=%v", this.ApiKey)
}

func (this *GcmClient) SetPriorety(p string) {

  if p == Priority_HIGH {
  this.Message.Priority = Priority_HIGH
  } else {
    this.Message.Priority = Priority_NORMAL
  }
}

func (this *GcmClient) SetCollapseKey(val string) (*GcmClient) {

  this.Message.CollapseKey = val

  return this
}

func (this *GcmClient) SetNotificationPayload(payload *NotificationPayload) (*GcmClient) {

  this.Message.Notification = *payload

  return this
}

func (this *GcmClient) SetContentAvailable(isContentAvailable bool) (*GcmClient) {
  
  this.Message.ContentAvailable = isContentAvailable

  return this
}

func (this *GcmClient) SetDelayWhileIdle(isDelayWhileIdle bool) (*GcmClient) {
  
  this.Message.DelayWhileIdle = isDelayWhileIdle

  return this
}
func (this *GcmClient) SetTimeToLive(ttl int) (*GcmClient) {
  
  if ttl > MAX_TTL {
    
    this.Message.TimeToLive = MAX_TTL

   } else {

    this.Message.TimeToLive = ttl

   }
  return this
}

func (this *GcmClient) SetRestrictedPackageName(pkg string) (*GcmClient) {
  
  this.Message.RestrictedPackageName = pkg

  return this
}


func (this *GcmClient) SetDryRun(drun bool ) (*GcmClient) {
  
  this.Message.DryRun = drun

  return this
}






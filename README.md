# gcm
Go - GCM library ( Google Cloud Messaging )

#Example

```golang
package main


import (
  "gcm"
  "fmt"
)

const (
  GCM_API_KEY = "GCM_API_KEY"

)

func main() {

  p := &gcm.NotificationPayload{
    Title : "Android",
    Body: "Hello Android",
    Icon: "ic_stat_ic_notification",
    ClickAction: "OPEN_ACTIVITY_1",
    Sound: "default",
    Tag: "ab",
  }


  ids := []string{
    "GCM_TOKEN1",
    "GCM_TOKEN2",
  }

  // data := map[string]string{
  //   "message": "Hello World",
  // }

  client := gcm.NewGcmClient(GCM_API_KEY)

  // data message exmaple
  // status, err := client.NewDevicesList(ids).SetMsgData(data).Send()

  status, err := client.NewDevicesList(ids).SetNotificationPayload(p).Send()

  if err == nil {

    // fmt.Println(status.MulticastId)
    fmt.Println("Success:", status.Success)
    fmt.Println("Fail:", status.Fail)
    fmt.Println("Canonical_ids:",status.Canonical_ids)
    // fmt.Println(status.Results)
  } else {
    fmt.Println(err)
  }


}


```

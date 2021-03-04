package main

import (
    "fmt"
    "log"

    "github.com/golang/protobuf/proto"
)

func main() {
    /////////////////////////////////////////////////////
    // GET OBJECT DATA
    /////////////////////////////////////////////////////
    elliot := Person{
        Name: "Elliot",
        Age:  24,
        SocialFollowers: &SocialFollowers{
            Youtube: 2500,
            Twitter: 1400,
        },
    }
    
    ////////////////////////////////////////////////////
    // CREATE PROTOBUF OBJECT FROM JSON DATA
    ////////////////////////////////////////////////////
    data, err := proto.Marshal(&elliot)//now a protobuf object
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    
    ///////////////////////////////////////////////////
    // CREATE JSON DATA OUT OF PROTOBUF DATA
    //////////////////////////////////////////////////
    newElliot := &Person{}//object to put data in
    err = proto.Unmarshal(data, newElliot)//return to object format
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    //print out fields
    fmt.Println(newElliot.GetName())
    fmt.Println(newElliot.GetAge())
    fmt.Println(newElliot.SocialFollowers.GetTwitter())
    fmt.Println(newElliot.SocialFollowers.GetYoutube())
}

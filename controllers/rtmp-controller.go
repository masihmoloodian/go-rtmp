package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type NginxConfigStruct struct {
	AppName string
	UserId string
	SecretKey string
}

var nginxRtmpDir string = "./rtmp/include"
var nginxHttpDir string = "./rtmp/http-include"

func CreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n NginxConfigStruct
	err :=json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		log.Fatalf("Failed to decode: %v", err)
	}

	ok := WriteToFile(n.UserId, n.SecretKey)
	if !ok {
		json.NewEncoder(w).Encode("Failed to add configuration")
		return
	}
	json.NewEncoder(w).Encode("Configuration added!")
	return
}

func WriteToFile(userId string, secretKey string) bool {
	appName := fmt.Sprintf("live-"+userId)

	rtmpConfig := fmt.Sprintf("application %s {\n" + 
		"live on;\n"+
		"exec ffmpeg -i rtmp://localhost:1935/%s/$name\n"+
		  "-c:a libfdk_aac -b:a 128k -c:v libx264 -b:v 2500k -f flv -g 30 -r 30 -s 1280x720 -preset superfast -profile:v baseline rtmp://localhost:1935/hls/$name_720p2628kbs\n"+
		  "-c:a libfdk_aac -b:a 128k -c:v libx264 -b:v 1000k -f flv -g 30 -r 30 -s 854x480 -preset superfast -profile:v baseline rtmp://localhost:1935/hls/$name_480p1128kbs\n"+
		  "-c:a libfdk_aac -b:a 128k -c:v libx264 -b:v 750k -f flv -g 30 -r 30 -s 640x360 -preset superfast -profile:v baseline rtmp://localhost:1935/hls/$name_360p878kbs\n"+
		  "-c:a libfdk_aac -b:a 128k -c:v libx264 -b:v 400k -f flv -g 30 -r 30 -s 426x240 -preset superfast -profile:v baseline rtmp://localhost:1935/hls/$name_240p528kbs\n"+
		  "-c:a libfdk_aac -b:a 64k -c:v libx264 -b:v 200k -f flv -g 15 -r 15 -s 426x240 -preset superfast -profile:v baseline rtmp://localhost:1935/hls/$name_240p264kbs;\n"+
	      "}\n", appName, appName)

	httpConfig := fmt.Sprintf("location /callback-%s {\n"+
		"if ($arg_name = '%s') {\n"+
		"return 200;\n"+
		"}\n"+
		"return 403;\n"+
	  	"}\n",appName, secretKey)  

	d1 := []byte(rtmpConfig)
	rtmpFileName := fmt.Sprintf(nginxRtmpDir+"/rtmp-%s.conf", userId)
    errRtmp := os.WriteFile(rtmpFileName, d1, 0644)


	d2 := []byte(httpConfig)
	httpFileName := fmt.Sprintf(nginxHttpDir+"/rtmp-%s.conf", userId)
    errHttp := os.WriteFile(httpFileName, d2, 0644)

	if errRtmp != nil || errHttp != nil{
		return false
	}

	return true
}
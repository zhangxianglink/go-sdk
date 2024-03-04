package main

import (
	"asr/nls"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

func fileAsr() {
	url := "http://192.168.6.102:9997/vad/asr"
	// 以下的参数必填
	data := nls.NewVadParam("John Doe",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJndWFuZ2RvbmciLCJpYXQiOjE2ODE5NzU3MzQsImV4cCI6MjAzMDgwMzIwMH0.wn1FMgemnqj5_jaBZ6nPrKpKGsva-UBUnXbO2-MDgCQ",
		"http://192.168.6.55:10000/data/model/2024-02-20-17-32-40_30300013023096150_18500194588.mp3",
		"http://192.168.6.102:10101/sdk/vad",
		true)

	// 将数据编码为 JSON 格式
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	// 设置请求头部
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request sending error:", err)
		return
	}
	defer resp.Body.Close()

	// 解析响应
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Response decoding error:", err)
		return
	}

	// 处理响应数据
	fmt.Println("Response:", result)
}

func onMessage(msg []byte) {
	log.Printf("recv: %s", msg)
}
func onClose(err error) {
	log.Println("连接断开:", err)
}

func streamAsr() {
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	// 创建一个`os.Signal`类型的通道，用于接收操作系统的中断信号（如Ctrl+C）。
	taskId := "自定义"
	token := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJndWFuZ2RvbmciLCJpYXQiOjE2ODE5NzU3MzQsImV4cCI6MjAzMDgwMzIwMH0.wn1FMgemnqj5_jaBZ6nPrKpKGsva-UBUnXbO2-MDgCQ"
	u := url.URL{Scheme: "ws", Host: "192.168.6.102:9997", Path: "/asr/" + taskId}
	// 构造WebSocket服务的URL
	headers := http.Header{
		// 密钥
		"token": []string{token},
		// 模式
		"appKey": []string{"default"},
		// 采样率 8000/16000
		"sampleRate": []string{"8000"},
		// 8000采样率升级16000   1 启用 2 关闭
		"resample": []string{"1"},
		// 以下非必填 中间结果返回
		"enableIntermediateResult": []string{"true"},
		//是否添加标点符号
		"punctuation": []string{"1"},
		// 热词替换
		"hotRule": []string{"1"},
		// 中文转阿拉伯数字
		"numRule": []string{"1"},
		// 热词转换前后对比
		"cnRule": []string{"1"},
		// 分贝
		"dbRule": []string{"2"},
		//语速
		"speedRule": []string{"2"},
	}
	c := nls.Connection(u, headers, 5*time.Second)
	defer c.Close()
	// 使用Gorilla WebSocket的Dial函数连接到服务器。如果连接失败，程序将终止。defer c.Close()确保在函数返回前关闭连接。
	done := nls.Run(c, onMessage, onClose)

	// 从这里开始，你可以直接读取PCM数据流进行推送
	file, err := os.Open("D:\\data\\8kt\\20088.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // 文件结束
		}
		if err != nil {
			log.Fatal(err) // 处理其他可能的错误
		}
		// 发送读取的数据
		c.WriteMessage(websocket.BinaryMessage, buffer[:n])
	}
	log.Println("PCM数据流推送结束")

	// 等待WebSocket服务器断开连接
	<-done
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			streamAsr()
			fmt.Printf("Thread %d finished\n", i)
		}(i)
	}
	wg.Wait()
}

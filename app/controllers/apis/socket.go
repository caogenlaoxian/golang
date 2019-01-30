package apis

import (
	. "../../../app/models"
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	PATH = "http://quote.eastmoney.com/stocklist.html"
	SHPATH = "http://hq.sinajs.cn/rn=1457346056960&list=sh%s"
	SZPATH = "http://hq.sinajs.cn/rn=1457346056960&list=sz%s"
)
func GetSocketsDetail(c *gin.Context){
	//total := SocketsTotal()
	limit := 4000
	//num := int(math.Ceil(float64(total) / float64(limit)))
	list,_ := SocketList(0,limit)
	for _,val := range(list){
		var path string
		if strings.Index(val.Code,"0") == 0 || strings.Index(val.Code,"3") == 0{
			path = fmt.Sprintf(SZPATH,val.Code)
		}else{
			path = fmt.Sprintf(SHPATH,val.Code)
		}

		resp := Get(path)
		array := strings.Split(resp,",")
		if len(array)-3 > 0{
			index := (len(array)-3)
			if len(array[index]) > 0 {
				_date := strings.Replace(array[index],"-","",0)
				s := SocketsDetail{SocketCode:val.Code,Kai:array[1],Shou:array[2],Height:array[3],Low:array[4],Liang:array[5],Jia:array[6],SocketDate:_date}
				go s.SocketDetailAdd()
			}
		}
	}
}

//采集股票代码
func GetSockets(c *gin.Context){

	//获取到内容
	resp := Get(PATH)
	//fmt.Println(resp)
	result := _reg(resp)
	//写入日志
	//fileName := "f:/socket.log"
	//logFile,_  := os.Create(fileName)
	//defer logFile.Close()
	_s := Sockets{}
	_s.SocketTruncate()
	for i := 0 ; i < len(result) ; i++{
		//fmt.Println(result[i])
		_r1 := strings.Split(result[i],">")
		_r2 := strings.Split(_r1[1],")<")

		if !strings.Contains(_r2[0],"("){
			continue
		}
		_r3 := strings.Split(_r2[0],"(")

		if strings.Index(_r3[1],"6") != 0 && strings.Index(_r3[1],"0") != 0 && strings.Index(_r3[1],"3") != 0{
			continue
		}
		//debugLog := log.New(logFile,"[Info]",log.Llongfile)
		//debugLog.Println(_r3[0],_r3[1])
		name := UseNewEncoder(_r3[0],"gbk","utf-8")
		s := Sockets{Name:string(name) , Code: _r3[1]}
		////SocketsAdd这个方法必须大写
		id,err := s.SocketsAdd()
		if err != nil {
			//debugLog := log.New(logFile,"[Info]",log.Llongfile)
			//debugLog.Println(err)
		}
		if id > 0 {
			fmt.Println(id)
		}
		//_r1 := strings.Replace(result[i],"","",0)

	}
}

//字符集转换
func UseNewEncoder(src string,oldEncoder string,newEncoder string) string{
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr:= srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder .Translate([]byte(resStr), true)
	return string(resBytes)
}

//提取内容正则表达式
func _reg(content string) []string {
	reg :=regexp.MustCompile(`\">.*<\/a>`)
	if reg == nil{
		fmt.Println("hello,world")
	}

	return reg.FindAllString(content,-1)
}

func _reg2(content string) []string {
	reg :=regexp.MustCompile(`^\">.*<\/a>$`)
	if reg == nil{
		fmt.Println("hello,world")
	}

	return reg.FindAllString(content,-1)
}

func Get(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	response = result.String()
	return
}
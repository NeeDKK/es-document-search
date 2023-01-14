package controller

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/NeeDKK/esDocumentSearch/config"
	"github.com/NeeDKK/esDocumentSearch/entity"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"math/rand"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var nameForbidenMap = map[string]string{
	"路原理":   "name",
	"简历":    "name",
	"经历":    "name",
	"简介":    "name",
	"通信协议":  "name",
	"爱动漫":   "name",
	"程师":    "name",
	"公司":    "name",
	"简 历":   "name",
	"经验":    "name",
	"习经历":   "name",
	"练掌握":   "name",
	"方式":    "name",
	"年龄":    "name",
	"程序编程":  "name",
	"相信自己":  "name",
	"公司 技术": "name",
	"通信工":   "name",
	"公司 协议": "name",
	"项目维护":  "name",
	"简历编号":  "name",
	"查询":    "name",
}

type UploadController struct {
}

func UploadFile(c *gin.Context) {
	_, file, err := c.Request.FormFile("file")
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名
	name := strings.TrimSuffix(file.Filename, ext)
	// 读取文件
	f, openError := file.Open()
	if openError != nil {
		fmt.Println("function file.Open() Filed", openError.Error())
		entity.FailWithDetailed(openError, "上传文件"+name+"失败", c)
		return
	}
	defer f.Close() // 创建文件 defer 关闭
	if err != nil {
		fmt.Println("文件上传失败:", err.Error())
	}
	//创建文件大小的字节数组
	bytes := make([]byte, file.Size)
	buffer := bufio.NewReader(f)
	//将文件读取到字节数组中
	buffer.Read(bytes)
	//将字节数组转换为base64的字符串
	doc := base64.StdEncoding.EncodeToString(bytes)
	do := FileToEs(c, doc)
	entity.OkWithDetailed(do, "上传成功", c)
}

func FileToEs(c *gin.Context, doc string) *elastic.IndexResponse {
	var resume entity.Resume
	rand.Seed(time.Now().UnixNano())
	//正常从数据库读取,这里随机生成
	resume.ID = uint(rand.Intn(9999))
	//录入简历数据时，从用户输入读取,这里随机生成
	resume.Name = "name:" + strconv.Itoa(int(resume.ID))
	resume.Content = doc
	//定义索引和pipline写入es
	do, err := config.EsClient.Index().Index(config.RESUMEINDEX).BodyJson(resume).Pipeline(config.RESUMEPIPLINE).Do(c)
	if err != nil {
		fmt.Println(err.Error())
		entity.FailWithDetailed(err, "上传失败", c)
		return nil
	}
	//通过携程异步的方式执行回查赋值
	go searchPiplineHandleContent(do.Id, c)
	return do
}

//通过完成上传后的简历反查，填充简历的姓名和学校
func searchPiplineHandleContent(id string, c *gin.Context) {
	time.Sleep(3 * time.Second)
	query := elastic.NewBoolQuery().Must(elastic.NewMatchQuery("_id", id))
	do, err := config.EsClient.Search(config.RESUMEINDEX).Query(query).Do(c)
	if err != nil {
		fmt.Println("查询失败", err.Error())
		return
	}
	var School string
	var Name string
	if len(do.Hits.Hits) != 0 {
		var Source entity.Source
		err := json.Unmarshal(do.Hits.Hits[0].Source, &Source)
		if err != nil {
			fmt.Println("转换失败", err.Error())
		}
		content := Source.Attachment.Content
		for k, _ := range config.UnversityMap {
			if strings.Contains(content, k) {
				if School == "" {
					School += k
				} else {
					School += " & " + k
				}
			}
		}
		rexFirst := "(姓)[\u4E00-\u9FA5|\u00A0|\u0020|\u3000]{1,3}"
		compile, err := regexp.Compile(rexFirst)
		if err != nil {
			log.Fatalf("正则匹配失败:", err.Error())
		}
		matchStringFirstName := compile.FindAllString(content, -1)
		if len(matchStringFirstName) != 0 {
			content = strings.Split(content, matchStringFirstName[0])[1]
		}
		regEx := "(王|李|张|刘|陈|杨|黄|赵|吴|周|徐|孙|马|朱|胡|郭|何|高|林|罗|郑|梁|谢|宋|唐|许|韩|冯|邓|曹|彭|曾" +
			"|肖|田|董|袁|潘|于|蒋|蔡|余|杜|叶|程|苏|魏|吕|丁|任|沈|姚|卢|姜|崔|钟|谭|陆|汪|范|金|石|廖|贾|夏|韦|傅" +
			"|方|白|邹|孟|熊|秦|邱|江|尹|薛|闫|段|雷|侯|龙|史|黎|贺|顾|毛|郝|龚|邵|万|钱|覃|武|戴|孔|汤|庞|樊|兰|殷" +
			"|施|陶|洪|翟|安|颜|倪|严|牛|温|芦|季|俞|章|鲁|葛|伍|申|尤|毕|聂|柴|焦|向|柳|邢|岳|齐|沿|梅|莫|庄|辛|管" +
			"|祝|左|涂|谷|祁|时|舒|耿|牟|卜|路|詹|关|苗|凌|费|纪|靳|盛|童|欧|甄|项|曲|成|游|阳|裴|席|卫|查|屈|鲍|位" +
			"|覃|霍|翁|隋|植|甘|景|薄|单|包|司|柏|宁|柯|阮|桂|闵|欧阳|解|强|丛|华|车|冉|房|边|辜|吉|饶|刁|瞿|戚|丘" +
			"|古|米|池|滕|晋|苑|邬|臧|畅|宫|来|嵺|苟|全|褚|廉|简|娄|盖|符|奚|木|穆|党|燕|郎|邸|冀|谈|姬|屠|连|郜|晏" +
			"|栾|郁|商|蒙|计|喻|揭|窦|迟|宇|敖|糜|鄢|冷|卓|花|艾|蓝|都|巩|稽|井|练|仲|乐|虞|卞|封|竺|冼|原|官|衣|楚" +
			"|佟|栗|匡|宗|应|台|巫|鞠|僧|桑|荆|谌|银|扬|明|沙|薄|伏|岑|习|胥|保|和|蔺|水|云|昌|凤|酆|常|皮|康|元|平" +
			"|萧|湛|禹|无|贝|茅|麻|危|骆|支|咎|经|裘|缪|干|宣|贲|杭|诸|钮|嵇|滑|荣|荀|羊|於|惠|家|芮|羿|储|汲|邴|松" +
			"|富|乌|巴|弓|牧|隗|山|宓|蓬|郗|班|仰|秋|伊|仇|暴|钭|厉|戎|祖|束|幸|韶|蓟|印|宿|怀|蒲|鄂|索|咸|籍|赖|乔" +
			"|阴|能|苍|双|闻|莘|贡|逢|扶|堵|宰|郦|雍|却|璩|濮|寿|通|扈|郏|浦|尚|农|别|阎|充|慕|茹|宦|鱼|容|易|慎|戈" +
			"|庚|终|暨|居|衡|步|满|弘|国|文|寇|广|禄|阙|东|殴|殳|沃|利|蔚|越|夔|隆|师|厍|晃|勾|融|訾|阚|那|空|毋|乜" +
			"|养|须|丰|巢|蒯|相|后|红|权逯|盖益|桓|公|万俟|司马|上官|夏侯|诸葛|闻人|东方|赫连|皇甫|尉迟|公羊|澹台" +
			"|公冶|宗政|濮阳|淳于|单于|太叔|申屠|公孙|仲孙|轩辕|令狐|钟离|宇文|长孙|慕容|鲜于|闾丘|司徒|司空|亓官" +
			"|司寇|仉|督|子车|颛孙|端木|巫马|公西|漆雕|乐正|壤驷|公良|拓跋|夹谷|宰父|谷粱|法|汝|钦|段干|百里|东郭" +
			"|南门|呼延|归海|羊舌|微生|帅|缑|亢|况|郈|琴|梁丘|左丘|东门|西门|佘|佴|伯|赏|南宫|墨|哈|谯" +
			"|笪|年|爱|仝|代)[\u4E00-\u9FA5|\u00A0|\u0020|\u3000]{1,4}"
		compile, err = regexp.Compile(regEx)
		if err != nil {
			log.Fatalf("正则匹配失败:", err.Error())
		}
		matchString := compile.FindAllString(content, -1)
	OuterLoop:
		for _, v := range matchString {
			space := strings.Replace(v, " ", "", -1)
			if len([]rune(space)) > 1 && len([]rune(space)) <= 4 {
				//截取的姓名必须不包含简历中可能存在的关键字
				if s := nameForbidenMap[space]; s == "name" {
					continue OuterLoop
				}
				//姓名一般在简历开头。所以首次匹配到就返回
				Name = v
				break
			}
		}

		_, err = config.EsClient.Update().Index(config.RESUMEINDEX).Id(id).Doc(map[string]interface{}{"school": School, "name": Name}).Do(c)
		if err != nil {
			log.Fatalf("update es failure:" + err.Error())
		}
	}

}

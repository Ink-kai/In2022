package RoleManage

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 请求体
type (
	req_meta struct {
		Url          string
		Method       string
		Url_form     *strings.Reader
		Content_Type string
		Cookie       string
	}
)

func Log_err(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Read_excel(path, role_name, rolecode_name string, roleNameNum, roleCodeNum, userNameNum, userIdNum int) string {
	f, err := excelize.OpenFile(path)
	Log_err(err)
	active := f.GetSheetName(f.GetActiveSheetIndex())
	cols, err := f.GetRows(active)
	Log_err(err)
	var (
		data string
	)
	for _, col := range cols {
		if col[roleNameNum] == role_name && col[roleCodeNum] == rolecode_name {
			user_name := strings.Split(col[userNameNum], ",")
			user_id := strings.Split(col[userIdNum], ";")

			if len(user_id) > 1 {
				for i := 0; i < len(user_name); i++ {
					name := strconv.QuoteToASCII(user_name[i])
					uid := strconv.QuoteToASCII(user_id[i])
					if i == (len(user_name) - 1) {
						data += "{\"userUid\":" + uid + ",userName:" + name + "}"
					} else {
						data += "{\"userUid\":" + uid + ",userName:" + name + "},"
					}
				}
			} else {
				data += "{\"userUid\":" + col[userIdNum] + ",\"userName\":" + col[userNameNum] + "}"
			}
		}
	}
	var user_data string = "[" + data + "]"
	defer f.Close()
	return user_data
}

func Request_url(r Req_meta) {
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, r.Url, r.Url_form)
	Log_err(err)
	req.Header.Set("Content-Type", r.Content_Type)
	req.Header.Set("Cookie", r.Cookie)
	resp, err := client.Do(req)
	Log_err(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	Log_err(err)
	log.Println(string(body))
}

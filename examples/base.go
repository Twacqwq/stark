package main

import (
	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/app"
	starkConf "github.com/FarmerChillax/stark/config"
)

func main() {
	app.New(&stark.Application{
		Name:   "base-demo",
		Host:   "127.0.0.1",
		Port:   6000,
		Config: &starkConf.Config{},
		LoadConfig: func() error {
			return nil
		},
		SetupVars: func() error {
			return nil
		},
		RegisterCallback: make(map[stark.CallbackPosition]stark.CallbackFunc),
	})

}

// func groupAnagrams(strs []string) [][]string {
// 	mp := map[[26]int][]string{}
// 	for _, str := range strs {
// 		cnt := [26]int{}
// 		for _, b := range str {
// 			cnt[b-'a']++
// 		}
// 		mp[cnt] = append(mp[cnt], str)
// 	}
// 	ans := make([][]string, 0, len(mp))
// 	for _, v := range mp {
// 		ans = append(ans, v)
// 	}
// 	return ans
// }

// func ConcurrencyWrtie(src io.Reader, dest [2]io.Writer) (err error) {
// 	errCh := make(chan error, 1)

// 	// 管道，主要是用来写、读流转化
// 	pr, pw := io.Pipe()
// 	// teeReader ，主要是用来 IO 流分叉
// 	wr := io.TeeReader(src, pw)

// 	// 并发写入
// 	go func() {
// 		var _err error
// 		defer func() {
// 			pr.CloseWithError(_err)
// 			errCh <- _err
// 		}()
// 		_, _err = io.Copy(dest[1], pr)
// 	}()

// 	defer func() {
// 		// TODO：异常处理
// 		pw.Close()
// 		_err := <-errCh
// 		_ = _err
// 	}()

// 	// 数据写入
// 	_, err = io.Copy(dest[0], wr)

// 	return err
// }

package qike

import "fmt"
import "bufio"
import "strings"
import "os"
import "time"

func About() {
	fmt.Println("输入m查看作者信息")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Text()

	if strings.ToLower(input.Text()) == "m" {
		arr := [...]string{"六", "壬", "通", "神", "暄", "未", "然", "\n", "因", "缘", "造", "化", "系", "自", "身", "\n", "阴", "阳", "共", "判", "劝", "君", "明", "\n", "虚", "实", "明", "辨", "壬", "中", "求", "\n"}

		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i])
			time.Sleep(500 * time.Millisecond)
			//		fmt.Println(i)
			if i == 7 || i == 15 || i == 23 || i == 31 {
				fmt.Println(" ")
			}
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("						/by liangzi")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("						mail: bGlhbmd6aTEyMTZAb3V0bG9vay5jb20K")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	} else {
		fmt.Println("输入错误 退出...")
		os.Exit(0)
	}
}

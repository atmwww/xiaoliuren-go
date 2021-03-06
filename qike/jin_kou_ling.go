package qike

import (
	"fmt"
)

type 金口定信息列表 struct {
	大安 string
	留连 string
	速喜 string
	赤口 string
	小吉 string
	空亡 string

	大安留连 string
	大安速喜 string
	大安重位 string
	大安白虎 string
	大安小吉 string
	大安同宫 string
	三个大安 string

	留连空亡 string

	速喜白虎 string
	速喜重位 string
	速喜大安 string

	白虎重位   string
	白虎留连   string
	白虎留连空亡 string
	双虎     string
	白虎速喜   string
	速喜空亡   string
	白虎小吉   string

	小吉赤口 string
	小吉留连 string
}

func jin_kou_ding() {

	金口定 := 金口定信息列表{
		大安: "大安忠厚能当官　大安命占妻倔强，多是姐妹二呈祥，老人定是体不康，" +
			"岳家必有婚败梁",
		留连: "流连操劳手艺端　流连走关心不闲，赚的都是远方钱 ，白虎流连命两立",
		速喜: "速喜靠嘴走江川",
		赤口: "白虎刚烈婚不安　白虎占位金神关　不伤财宝难多端　白虎金神细推行，" +
			"不撞南墙头不停，脾气火爆劝不听，遇事较真必弄清　白虎落在空亡星，" +
			"行走强运莫动情，命中必定伤娄病，兄弟二三不多名 ，白虎流连命两立",
		小吉: "小吉精明犯桃欢　小吉命占人最灵，世间万事样样行，可惜命欠阴间债，" +
			"若不积善早归冥 小吉最怕重重来，须记寿促遇三灾，阳世虽得满宿财，" +
			"阴府却欠无数债 小吉最怕虎口张，一生婚姻暗情伤，福运最忌私情落，" +
			"不是藏娇便添房",
		空亡: "空亡财源心路宽　时落空亡牢狱关，须防牢狱身不安　空亡倔强豆腐心，" +
			"孝敬堂前敬双亲　空亡占命克弟兄，取得贤妇是英雄，流连重位离亲远，" +
			"外方动名富贵圈",
		大安留连: "大安流连也同讲，必与生肖按分祥，虽然婚后不顺当，命必富贵性情强",
		大安速喜: "大安速喜两星全，定是修仙有佛缘 大安速喜命中主，阳相生来运受楚，" +
			"早年必走坎坷路，晚年定保享安福",
		大安重位: "大安重位在命间，老坟风水远不圆，必要土年动坟迁，下辈儿孙免私奸",
		大安白虎: "大安白虎如风飘，阴阳风水细推敲，东北洼来西南高，必得龙凤命中娇" +
			"最忌东北坑门开，破财伤灾犯鬼妖，正北占火更不吉，运行猎狗坏运交",
		大安小吉: "大安小吉西南路，运见金猴不可住，正南高坡还有树，儿女定有不吉目" +
			"大安小吉对头皮，必有一个桃花迷，最怕空亡婚不强，克漏命中财宝箱",
		大安同宫: "    ",
		三个大安: "三个大安共并连，一生最忌本命年，三个大安阳相生，兄弟一个犯孤单",
		留连空亡: "流连空亡在命间，须是留意车马前　流连空亡铁扫盆，家有少亡无影坟，" +
			"女人命占要私奔，男人逢之动三婚",
		速喜白虎: "速喜白虎火与金，十五之前克双亲，若想破解免克星，认拜干亲或征兵",
		速喜重位: "速喜重位婚必离",
		速喜大安: "速喜大安定命全，老坟必得动过玄，家有少亡寿不添，也有空坟引葬圈",
		白虎重位: "白虎重位见空亡，必因官司进牢房，木土之年运不多，就连兄弟运也连" +
			"白虎重位姻缘伤，必有破败外情张",
		白虎留连:   "    ",
		白虎留连空亡: "白虎流连空亡落，运强也要走车祸，虽然一生财宝多，大忌藏娇风情波",
		双虎: "双虎怕逢甲运年，破祥灾星破财连，就连家人也不宁，定是口舌疾病添" +
			"丁年不宜远行间，若不破财官司牵",
		白虎速喜: "白虎速喜好姻缘，必是女人掌财权",
		速喜空亡: "速喜空亡婚下格，姻缘忽忙过路客",
		白虎小吉: "白虎小吉两分张，六月天里下冷霜，大安同宫婚百祥，是男必带桃花香",
		小吉赤口: "    ",
		小吉留连: "小吉流连人人亏，细风扬柳应桃花，姻缘和合定阴阳，流连重位必见伤",
	}

	msg1, msg2, msg3 := <-yi_gong, <-er_gong, <-san_gong

	//单宫定位
	if msg1 == "大安" || msg2 == "大安" || msg3 == "大安" {
		fmt.Println(金口定.大安)
	}
	if msg1 == "留连" || msg2 == "留连" || msg3 == "留连" {
		fmt.Println(金口定.留连)
	}
	if msg1 == "速喜" || msg2 == "速喜" || msg3 == "速喜" {
		fmt.Println(金口定.速喜)
	}
	if msg1 == "赤口" || msg2 == "赤口" || msg3 == "赤口" {
		fmt.Println(金口定.赤口)
	}
	if msg1 == "小吉" || msg2 == "小吉" || msg3 == "小吉" {
		fmt.Println(金口定.小吉)
	}
	if msg1 == "空亡" || msg2 == "空亡" || msg3 == "空亡" {
		fmt.Println(金口定.空亡)
	}
	//双工定位 大安留连
	if (msg1 == "大安" && (msg2 == "留连" || msg3 == "留连")) || (msg2 == "大安" && (msg1 == "留连" || msg3 == "留连")) || (msg3 == "大安" && (msg1 == "留连" || msg2 == "留连")) {
		fmt.Println(金口定.大安留连)
	}
	//大安速喜
	if (msg1 == "大安" && (msg2 == "速喜" || msg3 == "速喜")) || (msg2 == "大安" && (msg1 == "速喜" || msg3 == "速喜")) || (msg3 == "大安" && (msg1 == "速喜" || msg2 == "速喜")) {
		fmt.Println(金口定.大安速喜)
	}
	//大安重位
	if (msg1 == "大安" && (msg2 == "大安" || msg3 == "大安")) || (msg2 == "大安" && (msg1 == "大安" || msg3 == "大安")) || (msg3 == "大安" && (msg1 == "大安" || msg2 == "大安")) {
		fmt.Println(金口定.大安重位)
	}
	//大安白虎
	if (msg1 == "大安" && (msg2 == "赤口" || msg3 == "赤口")) || (msg2 == "大安" && (msg1 == "赤口" || msg3 == "赤口")) || (msg3 == "大安" && (msg1 == "赤口" || msg2 == "赤口")) {
		fmt.Println(金口定.大安白虎)
	}
	//大安小吉
	if (msg1 == "大安" && (msg2 == "小吉" || msg3 == "小吉")) || (msg2 == "大安" && (msg1 == "小吉" || msg3 == "小吉")) || (msg3 == "大安" && (msg1 == "小吉" || msg2 == "小吉")) {
		fmt.Println(金口定.大安小吉)
	}
	//大安同宫 有待进一步确认
	if msg1 == "大安" || msg2 == "大安" || msg3 == "大安" {
		fmt.Println(金口定.大安同宫)
	}
	//三个大安
	if msg1 == "大安" && msg2 == "大安" && msg3 == "大安" {
		fmt.Println(金口定.三个大安)
	}
	//留连空亡
	if (msg1 == "留连" && (msg2 == "空亡" || msg3 == "空亡")) || (msg2 == "留连" && (msg1 == "空亡" || msg3 == "空亡")) || (msg3 == "留连" && (msg1 == "空亡" || msg2 == "空亡")) {
		fmt.Println(金口定.留连空亡)
	}
	//速喜白虎
	if (msg1 == "速喜" && (msg2 == "赤口" || msg3 == "赤口")) || (msg2 == "速喜" && (msg1 == "赤口" || msg3 == "赤口")) || (msg3 == "速喜" && (msg1 == "赤口" || msg2 == "赤口")) {
		fmt.Println(金口定.速喜白虎)
	}
	//速喜重位
	if (msg1 == "速喜" && (msg2 == "速喜" || msg3 == "速喜")) || (msg2 == "速喜" && (msg1 == "速喜" || msg3 == "速喜")) || (msg3 == "速喜" && (msg1 == "速喜" || msg2 == "速喜")) {
		fmt.Println(金口定.速喜重位)
	}
	//速喜大安
	if (msg1 == "速喜" && (msg2 == "大安" || msg3 == "大安")) || (msg2 == "速喜" && (msg1 == "大安" || msg3 == "大安")) || (msg3 == "速喜" && (msg1 == "大安" || msg2 == "大安")) {
		fmt.Println(金口定.速喜大安)
	}
	//赤口重位

	if (msg1 == "赤口" && msg2 == "赤口" && msg3 == "空亡") || (msg2 == "空亡" && msg1 == "赤口" && msg3 == "赤口") || (msg3 == "空亡" && msg1 == "赤口" && msg2 == "赤口") {
		fmt.Println(金口定.白虎重位)
	}
	//双虎
	if (msg1 == "赤口" && (msg2 == "赤口" || msg3 == "赤口")) || (msg2 == "赤口" && (msg1 == "赤口" || msg3 == "赤口")) || (msg3 == "赤口" && (msg1 == "赤口" || msg2 == "赤口")) {
		fmt.Println(金口定.双虎)
	}
	//赤口留连
	if (msg1 == "赤口" && (msg2 == "留连" || msg3 == "留连")) || (msg2 == "赤口" && (msg1 == "留连" || msg3 == "留连")) || (msg3 == "赤口" && (msg1 == "留连" || msg2 == "留连")) {
		fmt.Println(金口定.白虎留连)
	}
	//赤口留连空亡
	if (msg1 == "赤口" && msg2 == "留连" && msg3 == "空亡") || (msg2 == "赤口" && msg1 == "留连" && msg3 == "空亡") || (msg3 == "赤口" && msg1 == "空亡" && msg2 == "留连") || (msg2 == "赤口" && msg3 == "留连" && msg1 == "空亡") || (msg1 == "赤口" && msg3 == "留连" && msg2 == "空亡") || (msg3 == "赤口" && msg2 == "留连" && msg1 == "空亡") {
		fmt.Println(金口定.白虎留连空亡)
	}
	//赤口速喜
	if (msg1 == "赤口" && (msg2 == "速喜" || msg3 == "速喜")) || (msg2 == "赤口" && (msg1 == "速喜" || msg3 == "速喜")) || (msg3 == "赤口" && (msg1 == "速喜" || msg2 == "速喜")) {
		fmt.Println(金口定.白虎速喜)
	}
	//速喜空亡
	if (msg1 == "速喜" && (msg2 == "空亡" || msg3 == "空亡")) || (msg2 == "速喜" && (msg1 == "空亡" || msg3 == "空亡")) || (msg3 == "速喜" && (msg1 == "空亡" || msg2 == "空亡")) {
		fmt.Println(金口定.白虎速喜)
	}
	//赤口小吉
	if (msg1 == "赤口" && (msg2 == "小吉" || msg3 == "小吉")) || (msg2 == "赤口" && (msg1 == "小吉" || msg3 == "小吉")) || (msg3 == "赤口" && (msg1 == "小吉" || msg2 == "小吉")) {
		fmt.Println(金口定.白虎小吉)
	}
	//小吉赤口
	if (msg1 == "赤口" && (msg2 == "小吉" || msg3 == "小吉")) || (msg2 == "赤口" && (msg1 == "小吉" || msg3 == "小吉")) || (msg3 == "赤口" && (msg1 == "小吉" || msg2 == "小吉")) {
		fmt.Println(金口定.小吉赤口)
	}
	//小吉留连
	if (msg1 == "小吉" && (msg2 == "留连" || msg3 == "留连")) || (msg2 == "小吉" && (msg1 == "留连" || msg3 == "留连")) || (msg3 == "小吉" && (msg1 == "留连" || msg2 == "留连")) {
		fmt.Println(金口定.小吉留连)
	}
}

package quick

import (
	"cyhd/common/utils"
	"fmt"
	"strconv"
	"strings"
)

/**
QuickSDK游戏同步加解密算法描述


解密方法
strEncode 密文
keys 解密密钥 为游戏接入时分配的 callback_key
*/
func Decode(str string, keys string) string {

	strs := strings.Split(str, "@")
	strs = strs[1:]
	//fmt.Printf("strs %#v %#v\n", strs, len(strs))

	keysNum := GetBytes(keys)

	//fmt.Printf("keysNum %#v\n", keysNum)

	_data := []int{}
	_len := len(keysNum)

	//fmt.Printf("_len  %#v \n", _len)

	for i, v := range strs {
		keyVar := keysNum[i%_len]
		kn, _ := strconv.Atoi(v)
		//fmt.Printf("keyVar  %#v %#v %#v - %#v\n", i, keyVar, kn, 0xff&keyVar)
		_data = append(_data, kn-0xff&keyVar)
	}

	//fmt.Printf("_data  %#v \n", _data)

	return ToStr(_data)

}

/**
计算游戏同步签名
*/
func GetSign(nt_data string, sign string, callback_key string) string {
	str := fmt.Sprintf("%s%s%s", nt_data, sign, callback_key)
	md5sign := utils.MD5(str)
	return md5sign
}

/**
MD5签名替换
*/
func replaceMD5(md5 string) string {

	keysNum := GetBytes(md5)

	_len := len(keysNum)

	if _len >= 23 {
		change := keysNum[1]
		keysNum[1] = keysNum[13]
		keysNum[13] = change

		change2 := keysNum[5]
		keysNum[5] = keysNum[17]
		keysNum[17] = change2

		change3 := keysNum[7]
		keysNum[7] = keysNum[23]
		keysNum[23] = change3
	} else {
		return md5
	}

	return ToStr(keysNum)

}

/**
 * 转成字符数据
 */
func GetBytes(strs string) []int {
	_keys := []byte(strs)
	//fmt.Printf("keys %#v\n", _keys)

	keysNum := []int{}
	for _, _n := range _keys {
		num := int(_n)
		keysNum = append(keysNum, num)
	}
	return keysNum
}

/**
 * 转化字符串
 */
func ToStr(keysNum []int) string {
	_b := []string{}
	for _, v := range keysNum {
		_b = append(_b, string(v))
	}
	return strings.Join(_b, "")
}

/**
<?php


 * QuickSDK游戏同步加解密算法描述
 * @copyright quicksdk 2015
 * @author quicksdk
 * @version quicksdk v 0.0.1 2014/9/2


class quickAsy{


 * 解密方法
 * $strEncode 密文
 * $keys 解密密钥 为游戏接入时分配的 callback_key

	public function decode($strEncode, $keys) {
		if(empty($strEncode)){
			return $strEncode;
		}
		preg_match_all('(\d+)', $strEncode, $list);
		$list = $list[0];
		if (count($list) > 0) {
			$keys = self::getBytes($keys);
			for ($i = 0; $i < count($list); $i++) {
				$keyVar = $keys[$i % count($keys)];
				$data[$i] =  $list[$i] - (0xff & $keyVar);
			}
			return self::toStr($data);
		} else {
			return $strEncode;
		}
	}


 * 计算游戏同步签名

public static function getSign($params,$callbackkey){

return md5($params['nt_data'].$params['sign'].$callbackkey);
}


 * MD5签名替换

static private function replaceMD5($md5){

strtolower($md5);
$bytes = self::getBytes($md5);

$len = count($bytes);

if ($len >= 23){
$change = $bytes[1];
$bytes[1] = $bytes[13];
$bytes[13] = $change;

$change2 = $bytes[5];
$bytes[5] = $bytes[17];
$bytes[17] = $change2;

$change3 = $bytes[7];
$bytes[7] = $bytes[23];
$bytes[23] = $change3;
}else{
return $md5;
}

return self::toStr($bytes);
}


 * 转成字符数据

    private static function getBytes($string) {
        $bytes = array();
        for($i = 0; $i < strlen($string); $i++){
             $bytes[] = ord($string[$i]);
        }
        return $bytes;
    }

 * 转化字符串

private static function toStr($bytes) {
$str = '';
foreach($bytes as $ch) {
$str .= chr($ch);
}
return $str;
}
}

?>

*/

package tools

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

//random
// A default and returns the first one does not start with numbers
func RandomString (l int,c int,pa ... string) []string {
	if c <=0{c=1}
	var result []string = make([]string,c,c)
	var randStringLow = "abcdefghijklmnopqrstuvwxyz"
//	var randStringCap = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var randStringLN ="abcdefghijklmnopqrstuvwxyz0123456789"
//	var randStringCN ="ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//	var randStringAll="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var sam int
if len(pa[0])>0{
	switch pa[0]{
	case "A":
		if c!=1{
			var same bool
			for i:=0;i<c;i++{
				same = false
				temp:=a(l,randStringLN,randStringLow)
				if len(result)==0{
					result[i] = temp
					continue
				}
				for j:=0;j<len(result);j++{
					if temp ==result[j]{
						i--
						sam++
						break
					}
					if !same{
						same = true
					}
				}
				if same{
					result[i] = temp
				}
			}
		}else{
			result[0] =a(l,randStringLN,randStringLow)
		}
	case "B":
		if c!=1{
			var same bool
			for i:=0;i<c;i++{
				same = false
				temp:=b(l,randStringLN)
				if len(result)==0{
					result[i] = temp
					continue
				}
				for j:=0;j<len(result);j++{
					if temp ==result[j]{
						i--
						break
					}
					if !same{
						same = true
					}
				}
				if same{
					result[i] = temp
				}
			}
		}else{
			result[0] =b(l,randStringLN)
		}
	}
}
	fmt.Println("sam:",sam)
return result
}

type randNum struct {
	NumPort []int
}
func (rn *randNum)randInt(min int , max int)int{
	if(len(rn.NumPort)<1 || rn.NumPort[0]==0){
		if len(rn.NumPort)>1{
			if rn.NumPort[1]==0 {
				rn.NumPort = randInts(min,max)
			}
		}else{
			rn.NumPort  = randInts(min,max)
		}
	}
	re := rn.NumPort[0]
	if re<min || re>max{
		re = rn.randInt(min,max)
	}
	rn.NumPort = rn.NumPort[1:len(rn.NumPort)]
	return  re
}

func randInts(min int , max int) []int{
	//time.Sleep(1*time.Nanosecond)
	timeseed :=time.Now().UTC().UnixNano()
	timeseedstr := strconv.FormatInt(timeseed,10)
	var seed string
	var length =len(timeseedstr)
	for i:=0;i<length;i++{
		ija,_ := strconv.Atoi(string(timeseedstr[length-3]))
		ij,_ := strconv.Atoi(string(timeseedstr[length-4]))
		ijs,_ := strconv.Atoi(string(timeseedstr[length-5]))
		j := ((i+ijs)*ij)*ija
		t :=j%len(timeseedstr)
		seed = seed+string(timeseedstr[t])
	}
	seedInt,_ :=strconv.Atoi(seed)
	rand.Seed(int64(seedInt))
	res := make([]int,16,16);
	for x:=0;x<16;x++{
		re := min + rand.Intn(max-min)
		res[x] =re
	}
	return  res
}

func randInt(min int , max int) int {
	time.Sleep(1*time.Nanosecond)
	timeseed :=time.Now().UTC().UnixNano()
	timeseedstr := strconv.FormatInt(timeseed,10)
	var seed string
	var length =len(timeseedstr)
	for i:=0;i<length;i++{
		ij,_ := strconv.Atoi(string(timeseedstr[length-4]))
		ijs,_ := strconv.Atoi(string(timeseedstr[length-5]))
		j := ((i+ijs)*ij)
		t :=j%len(timeseedstr)
		seed = seed+string(timeseedstr[t])
	}
	seedInt,_ :=strconv.Atoi(seed)
	rand.Seed(int64(seedInt))
	return min + rand.Intn(max-min)
}

func a(l int ,seedString string,fseedString string)string{
	var rs string
	var t int
	var randnum randNum
	randnum.NumPort = make([]int,16,16);
	for i:=0 ; i<l ;i++ {
		if i== 0{
			 j:=len(fseedString)
					t =randnum.randInt(0,j)
			rs = string(fseedString[t])
		}else{
			if i==1{
				randnum.NumPort = make([]int,16,16);
			}
			j:=len(seedString)
					t =randnum.randInt(0,j)
			rs = rs+string(seedString[t])
		}
	}
	return rs
}

func b(l int ,seedString string,)string{
	var rs string
	var t int
	var randnum randNum
	randnum.NumPort = make([]int,16,16);
	for i:=0 ; i<l ;i++ {
			j:=len(seedString)
			t =randnum.randInt(0,j)
			rs = rs+string(seedString[t])
	}
	return rs
}


package main

import (
    "fmt"
    "log"

    "os"
    "strconv"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

func frev(slice []byte)[]byte{
	for i, j := 0, len(slice)-1; i< j; i, j = i+1,j-1 {
		slice[i],slice[j]=slice[j],slice[i]
	}
	return slice
}

func main() {
    if len(os.Args) <2{
		fmt.Println("args num < 2")
		return
    }

    if len(os.Args[1]) <3{
		fmt.Println("subPubKey len <2")
		return
    }

    priKey:=os.Args[1]
    subPubKey:=[]byte(os.Args[2][2:])

    privateKey, err := crypto.HexToECDSA(priKey)
    if err != nil {
        log.Fatal(err)
		return
    }

    data:=[]byte("\x19Ethereum Signed Message:\n")

    data2 := []byte("evm:")
    data2=append(data2,subPubKey...)
    var rev []byte;
    l:=len(data2)
    for {
	if l>0{
	  e:=strconv.Itoa(l%10)
	  rev=append(rev,byte(int(e[0])));
	  l /= 10;
	 }else{
	  break;
	}
    }
    rev=frev(rev)

    data3 :=[]byte{}
    data3=append( data3,data...)
    data3=append( data3,rev...)
    data3=append(data3,data2...)
    fmt.Println("sig data",string(data3))
    hash := crypto.Keccak256Hash(data3)
    fmt.Println(hash.Hex())

    signature, err := crypto.Sign(hash.Bytes(), privateKey)
    if err != nil {
        log.Fatal(err)
		return
    }

    fmt.Println("signature:", hexutil.Encode(signature), ", len:", len(signature))
}

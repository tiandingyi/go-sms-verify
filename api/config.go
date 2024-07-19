package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// godotenv.Unmarshal() 函数的作用是将 .env 文件中的环境变量读取到内存中，
// 以便程序后续可以方便地获取和使用这些环境变量的值。

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		// 会终止程序
		log.Fatalln(err)
		log.Fatal("Error loading .env")
	}
	return os.Getenv("TWILTO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env")
	}
	return os.Getenv("TWILTO_AUTHTOKEN")
}

func envSERVICESID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env")
	}
	return os.Getenv("TWILTO_SERVICES_ID")
}

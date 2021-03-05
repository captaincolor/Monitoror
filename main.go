package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"github.com/shirou/gopsutil/v3/cpu"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "mysql991004"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "Serverinfo"
)

type cpuInfo struct {
	cpuid int32
	pcore int
	lcore int
	occupancy float64
	Mhz float64
	CacheSize int32
}

func main() {
	cpuCurrentStat := visorCpu()
	dSN := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db, err := sql.Open("mysql", dSN)
	if err != nil {
		log.Println("open mysql failed, ", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("ping failed, ", err)
		return
	}
	insertData(db, cpuCurrentStat)
	defer db.Close()
}

func visorCpu() cpuInfo {
	pcore, _ := cpu.Counts(false)
	lcore, _ := cpu.Counts(true)
	//fmt.Printf("物理核数: %v, 逻辑核数: %v \n", physicCore, logicCore)

	seconds := 5
	occupancy, _ := cpu.Percent(time.Duration(seconds)*time.Second, false) // false, 总cpu使用率
	//fmt.Printf("cpu总占用率: %v \n", cpuOccupancy[0])

	cpuStat, _ := cpu.Info()
	//fmt.Printf("%v, %v \n", cpuStat[0].Mhz, cpuStat[0].CacheSize)

	return cpuInfo{
		cpuid: cpuStat[0].CPU,
		pcore: pcore,
		lcore: lcore,
		occupancy: occupancy[0],
		Mhz: cpuStat[0].Mhz,
		CacheSize: cpuStat[0].CacheSize,
	}
}

func insertData(db *sql.DB, cpuinfo cpuInfo) {
	result, err := db.Exec("insert into cpuinfo(physical_core, logical_core, occupancy, mhz, cachesize) values(?,?,?,?,?)", cpuinfo.pcore, cpuinfo.lcore, cpuinfo.occupancy, cpuinfo.Mhz, cpuinfo.CacheSize)
	if err != nil {
		log.Println("exec failed, ", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("exec failed, ", err)
		return
	}
	fmt.Println("insert succ: ", id)
}
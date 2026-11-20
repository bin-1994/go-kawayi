package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"time"
)

func GetSystemInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"result": gin.H{
			"cpu": getCpuPercent(),
			"memory": getMemPercent(),
			"disk": getDiskPercent(),
		},
	})
}

func getCpuPercent() float64 {
	value, _ := cpu.Percent(time.Second, false)
	return value[0]
}

func getMemPercent() float64 {
	value, _ := mem.VirtualMemory()
	return value.UsedPercent
}

func getDiskPercent() float64 {
	value, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(value[0].Mountpoint)
	return diskInfo.UsedPercent
}

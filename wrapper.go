package qprometheus

import (
	"strings"
	"strconv"
	"github.com/pkg/errors"
)

// QPS 统计
type QPSRecord struct {
	Idc    string  // 机房
	Times  float64 // 统计次数
	Api    string  // 统计路径
	Module string  // 所属模块
	Method string  // 请求方法
	Code   int     // 状态码
}

// 延迟数据 统计
type LatencyRecord struct {
	Idc    string  // 机房
	Time   float64 // 花费时间
	Api    string  // 统计路径
	Module string  // 所属模块
	Method string  // 请求方法
}

// 获取wrapper句柄
func GetWrapper() *prom {
	return Wrapper
}

// QPS记录
func (p *prom) QpsCountLog(r QPSRecord) (ret bool, err error) {
	if strings.TrimSpace(r.Api) == "" {
		return ret, errors.New("QPSRecord.Api Can't Be Empty")
	}

	if r.Times <= 0 {
		r.Times = 1
	}

	if strings.TrimSpace(r.Module) == "" {
		r.Module = "self"
	}

	if strings.TrimSpace(r.Method) == "" {
		r.Method = "GET"
	}

	if r.Code == 0 {
		r.Code = 200
	}

	if r.Idc == "" {
		r.Idc = p.Idc
	}

	p.counter.WithLabelValues(p.Appname, r.Module, r.Api, r.Method, strconv.Itoa(r.Code), r.Idc).Add(r.Times)

	return true, nil
}

// 延迟记录
func (p *prom) LatencyLog(r LatencyRecord) (ret bool, err error) {
	if r.Time <= 0 {
		return ret, errors.New("LatencyRecord.Time Must Greater Than 0")
	}

	if strings.TrimSpace(r.Module) == "" {
		r.Module = "self"
	}

	if strings.TrimSpace(r.Method) == "" {
		r.Method = "GET"
	}

	if r.Idc == "" {
		r.Idc = p.Idc
	}

	p.histogram.WithLabelValues(p.Appname, r.Module, r.Api, r.Method, r.Idc).Observe(r.Time)

	return true, nil
}

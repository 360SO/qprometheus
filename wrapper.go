package qprometheus

import (
	"strings"
	"strconv"
	"github.com/pkg/errors"
)

type QPSRecord struct {
	Times float64
	Api string
	Module string
	Method string
	Code int
}

type LatencyRecord struct {
	Time float64
	Api string
	Module string
	Method string
}

func GetWrapper() *prom {
	return Wrapper
}

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

	p.counter.WithLabelValues(p.Appname, r.Module, r.Api, r.Method, strconv.Itoa(r.Code), p.Idc).Add(r.Times)

	return true, nil
}

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

	p.histogram.WithLabelValues(p.Appname, r.Module, r.Api, r.Method, p.Idc).Observe(r.Time)

	return true, nil
}

package iplibrary

type WriterInterface interface {
	WriteMeta() error
	Write(ipFrom string, ipTo string, countryId int64, provinceId int64, cityId int64, townId int64, providerId int64) error
	Sum() string
}

package service

import (
	"encoding/json"
	"errors"
	"majooTest/domain"
	"majooTest/dto"
	"majooTest/repository"
)

type ReportService interface {
	GetMerchantReport(filterData dto.PageDTO, userId int) (resultPagingData dto.PageDTO, err error)
	GetOutletReport(filterData dto.PageDTO, userId int) (resultPagingData dto.PageDTO, err error)
}

type reportService struct {
	userRepository        repository.UsersRepository
	merchantRepository    repository.MerchantRepository
	outletRepository      repository.OutletRepository
	transactionRepository repository.TransactionRepository
}

func NewReportService(
	userRepository repository.UsersRepository,
	merchantRepository repository.MerchantRepository,
	outletRepository repository.OutletRepository,
	transactionRepository repository.TransactionRepository,
) ReportService {
	return &reportService{
		userRepository:        userRepository,
		merchantRepository:    merchantRepository,
		outletRepository:      outletRepository,
		transactionRepository: transactionRepository,
	}
}

func (service *reportService) GetMerchantReport(filterData dto.PageDTO, userId int) (resultPagingData dto.PageDTO, err error) {
	var merchantDetailData = dto.MerchantOmzetDTO{}
	var resultListDTO = []dto.MerchantOmzetDTO{}
	dataByte, _ := json.Marshal(filterData.Data)
	json.Unmarshal(dataByte, &merchantDetailData)
	merchantData, _ := service.merchantRepository.GetDetailMerchantByUserId(userId)

	if merchantDetailData.MerchantId != 0 {
		if merchantDetailData.MerchantId != int(merchantData.Id) {
			err = errors.New("UnAuthorized Merchant")
			return resultPagingData, err
		}
	}

	dataList, total, page, lastPage, err := service.transactionRepository.FindAllMerchantTransactionWithFilterPage(
		userId, int64(merchantDetailData.MerchantId), 0, filterData.StartDate,
		filterData.EndDate, filterData.Page, filterData.Limit)

	if len(dataList) > 0 {
		for _, data := range dataList {
			dtoData := MappingMerchantReportDomainToDTO(data)
			resultListDTO = append(resultListDTO, dtoData)
		}
	} else {
		dtoData := dto.MerchantOmzetDTO{}
		dtoData.MerchantId = int(merchantData.Id)
		dtoData.MerchantName = merchantData.MerchantName
		dtoData.Omzet = 0

		resultListDTO = append(resultListDTO, dtoData)
	}

	resultPagingData.Data = resultListDTO
	resultPagingData.Page = page
	resultPagingData.TotalData = total
	resultPagingData.LastPage = lastPage
	resultPagingData.StartDate = filterData.StartDate
	resultPagingData.EndDate = filterData.EndDate

	if err != nil {
		return resultPagingData, err
	}
	return resultPagingData, nil
}

func (service *reportService) GetOutletReport(filterData dto.PageDTO, userId int) (resultPagingData dto.PageDTO, err error) {
	var outletDetailData = dto.OutletOmzetDTO{}
	var resultListDTO = []dto.OutletOmzetDTO{}
	dataByte, _ := json.Marshal(filterData.Data)
	json.Unmarshal(dataByte, &outletDetailData)
	merchantData, _ := service.merchantRepository.GetDetailMerchantByUserId(userId)

	if outletDetailData.MerchantId != 0 {
		if outletDetailData.MerchantId != int(merchantData.Id) {
			err = errors.New("UnAuthorized Merchant")
			return resultPagingData, err
		}
	}

	outletDataList, err := service.outletRepository.GetOutletByMerchantId(int(merchantData.Id))
	if err != nil {
		return resultPagingData, err
	}

	if outletDetailData.OutletId != 0 {
		var isAllowedOutlet = false
		for _, outletData := range outletDataList {
			if outletDetailData.OutletId == int(outletData.Id) {
				isAllowedOutlet = true
			}
		}

		if !isAllowedOutlet {
			err = errors.New("UnAuthorized Outlet")
			return resultPagingData, err
		}
	}

	dataList, total, page, lastPage, err := service.transactionRepository.FindAllOutletTransactionWithFilterPage(
		userId, int64(outletDetailData.MerchantId), int64(outletDetailData.OutletId), filterData.StartDate,
		filterData.EndDate, filterData.Page, filterData.Limit)

	if len(dataList) > 0 {
		for _, data := range dataList {
			dtoData := MappingOutletReportDomainToDTO(data)
			resultListDTO = append(resultListDTO, dtoData)
		}
	} else {
		for _, outletData := range outletDataList {
			dtoData := dto.OutletOmzetDTO{}
			dtoData.MerchantId = int(merchantData.Id)
			dtoData.MerchantName = merchantData.MerchantName
			dtoData.OutletId = int(outletData.Id)
			dtoData.OutletName = outletData.OutletName
			dtoData.Omzet = 0

			resultListDTO = append(resultListDTO, dtoData)
		}
	}

	resultPagingData.Data = resultListDTO
	resultPagingData.Page = page
	resultPagingData.TotalData = total
	resultPagingData.LastPage = lastPage
	resultPagingData.StartDate = filterData.StartDate
	resultPagingData.EndDate = filterData.EndDate

	if err != nil {
		return resultPagingData, err
	}
	return resultPagingData, nil
}

func MappingMerchantReportDomainToDTO(domain domain.DataReportTemplate) dto.MerchantOmzetDTO {
	result := dto.MerchantOmzetDTO{}

	if domain.MerchantId != 0 {
		result.MerchantId = int(domain.MerchantId)
	}

	if domain.MerchantName != "" {
		result.MerchantName = domain.MerchantName
	}

	if domain.Omzet != 0 {
		result.Omzet = domain.Omzet
	}

	return result
}

func MappingOutletReportDomainToDTO(domain domain.DataReportTemplate) dto.OutletOmzetDTO {
	result := dto.OutletOmzetDTO{}

	if domain.MerchantId != 0 {
		result.MerchantId = int(domain.MerchantId)
	}

	if domain.MerchantName != "" {
		result.MerchantName = domain.MerchantName
	}

	if domain.OutletId != 0 {
		result.OutletId = int(domain.OutletId)
	}

	if domain.OutletName != "" {
		result.OutletName = domain.OutletName
	}

	if domain.Omzet != 0 {
		result.Omzet = domain.Omzet
	}

	return result
}

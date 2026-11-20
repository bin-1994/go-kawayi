package service

import (
	"fmt"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IBangumiService interface {
	FindById(int) model.BangumiInfoModel
	FindAll(model.BangumiQueryModel) ([]model.BangumiListModel, int64)
	FindFilter() model.FilterModel
	EpToId(string) int
}

type BangumiService struct {
	Repository        repository.IBangumiRepository
	StyleRepository   repository.IStyleRepository
	EpisodeRepository repository.IEpisodeRepository
}

func NewBangumiService(repository repository.IBangumiRepository, styleRepository repository.IStyleRepository, episodeRepository repository.IEpisodeRepository) IBangumiService {
	return &BangumiService{Repository: repository, StyleRepository: styleRepository, EpisodeRepository: episodeRepository}
}

func (service *BangumiService) EpToId(episode string) int {
	id := service.Repository.EpToId(episode)
	return id
}

func (service *BangumiService) FindById(id int) model.BangumiInfoModel {
	bangumi := service.Repository.FindById(id)

	var info model.BangumiInfoModel
	info.MediaId = bangumi.MediaId
	info.Title = bangumi.Title
	info.Subtitle = bangumi.Subtitle
	info.Keyword = bangumi.Keyword
	info.Level = bangumi.Level
	info.IsLimit = bangumi.IsLimit
	info.Cover = bangumi.Cover
	info.VideoCover = bangumi.VideoCover
	info.Staff = bangumi.Staff
	info.Cast = bangumi.Cast
	info.Story = bangumi.Story

	info.Publish.EpId = bangumi.EpId
	info.Publish.EpIndex = bangumi.EpIndex
	info.Publish.EpTotal = bangumi.EpTotal
	info.Publish.IsFinish = bangumi.IsFinish
	info.Publish.OnAir = bangumi.OnAir
	info.Publish.PubString = bangumi.PubString
	info.Publish.PubTime = bangumi.PubTime

	info.Count.Follows = bangumi.Follows
	info.Count.Views = bangumi.Views

	info.Styles = service.StyleRepository.FindByIds(bangumi.MediaId)

	info.Seasons = service.Repository.FindSeasonById(bangumi.SeasonId)

	info.Episode = service.EpisodeRepository.FindAll(bangumi.MediaId)
	return info
}

func (service *BangumiService) FindAll(query model.BangumiQueryModel) ([]model.BangumiListModel, int64) {

	orderBy := []string{"pub_time", "update_time", "score", "follows", "views"}
	orderSort := []string{"DESC", "ASC"}
	order := fmt.Sprintf("%s %s, media_id DESC", orderBy[query.Order], orderSort[query.Sort])

	conditions := []string{}
	params := []any{}

	conditions = append(conditions, "is_limit <= ?")
	params = append(params, 1)

	conditions = append(conditions, "level < ?")
	params = append(params, 18)

	if category := query.Category; category != 0 {
		conditions = append(conditions, "type = ?")
		params = append(params, category)
	}

	if version := query.Version; version != 0 {
		if version == 3 {
			conditions = append(conditions, "version NOT IN (?)")
			params = append(params, []int{1, 2})
		} else {
			conditions = append(conditions, "version = ?")
			params = append(params, version)
		}
	}

	if area := query.Area; area != 0 {
		switch area {
		case 1:
			conditions = append(conditions, "area IN (?)")
			params = append(params, []int{201, 101201, 201301})
		case 2:
			conditions = append(conditions, "area IN (?)")
			params = append(params, []int{101, 102, 103, 101201, 101301})
		case 3:
			conditions = append(conditions, "area IN (?)")
			params = append(params, []int{301, 101301, 201301})
		default:
			conditions = append(conditions, "area NOT IN (?)")
			params = append(params, []int{101, 102, 103, 201, 301, 101201, 101301})
		}
	}

	if year := query.Year; year != "0" && year != "" {
		switch year {
		case "更早":
			conditions = append(conditions, "FROM_UNIXTIME(pub_time, '%Y') < ?")
			params = append(params, 1980)
		case "80年代":
			conditions = append(conditions, "FROM_UNIXTIME(pub_time, '%Y') BETWEEN 1980 AND 1989")
		case "90年代":
			conditions = append(conditions, "FROM_UNIXTIME(pub_time, '%Y') BETWEEN 1990 AND 1999")
		default:
			if matched, _ := regexp.MatchString(`[\d+]-[\d+]`, query.Year); matched {
				arrYear := strings.Split(query.Year, "-")
				conditions = append(conditions, "FROM_UNIXTIME(pub_time, '%Y') BETWEEN ? AND ?")
				params = append(params, arrYear[1], arrYear[0])
			} else {
				conditions = append(conditions, "FROM_UNIXTIME(pub_time, '%Y') = ?")
				params = append(params, query.Year)
			}
		}
	}

	if month := query.Month; month != 0 {
		conditions = append(conditions, "quarter(FROM_UNIXTIME(pub_time)) = ?")
		params = append(params, month)
	}

	if style := query.Style; style != 0 {
		ids := service.StyleRepository.FindByBangumiIds(style)
		conditions = append(conditions, "media_id IN (%s)")
		params = append(params, ids)
	}

	if state := query.State; state != 0 {
		conditions = append(conditions, "is_finish = ?")
		params = append(params, state)
	}

	if isNew := query.IsNew; isNew != 0 {
		conditions = append(conditions, "is_new = ?")
		params = append(params, isNew)
	}

	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 {
		query.Size = 200
	}
	offset := (query.Page - 1) * query.Size
	list := service.Repository.FindAll(conditions, params, order, offset, query.Size)

	total := service.Repository.Count(conditions, params)

	return list, total
}

func (service *BangumiService) FindFilter() model.FilterModel {
	category := model.FilterList{
		Key:   "category",
		Title: "分类",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
			{Value: "1", Name: "动画"},
			{Value: "2", Name: "电视剧"},
		},
	}

	version := model.FilterList{
		Key:   "version",
		Title: "版本",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
			{Value: "1", Name: "正片"},
			{Value: "2", Name: "剧场版/电影"},
			{Value: "3", Name: "其他"},
		},
	}

	area := model.FilterList{
		Key:   "area",
		Title: "地区",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
			{Value: "1", Name: "日本"},
			{Value: "2", Name: "中国"},
			{Value: "3", Name: "欧美"},
			{Value: "4", Name: "其他"},
		},
	}

	state := model.FilterList{
		Key:   "state",
		Title: "状态",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
			{Value: "1", Name: "连载"},
			{Value: "2", Name: "完结"},
		},
	}

	month := model.FilterList{
		Key:   "month",
		Title: "季度",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
			{Value: "1", Name: "1月"},
			{Value: "2", Name: "4月"},
			{Value: "3", Name: "7月"},
			{Value: "4", Name: "10月"},
		},
	}

	nowYear, nowMonth, _ := time.Now().Date()
	if nowMonth >= 10 {
		nowYear++
	}
	yearItem := []model.FilterItem{{Value: "0", Name: "全部"}}
	index := 0
	for i := nowYear; i >= 2000; i-- {
		if index < 2 {
			yearItem = append(yearItem, model.FilterItem{Value: strconv.Itoa(i), Name: strconv.Itoa(i)})
		} else {
			if index >= 2 && i%5 == 0 {
				str := fmt.Sprintf("%s-%s", strconv.Itoa(i+4), strconv.Itoa(i))
				yearItem = append(yearItem, model.FilterItem{Value: str, Name: str})
			}
		}
		if i%5 == 0 {
			index++
		}
	}

	yearItem = append(yearItem, model.FilterItem{Value: "90年代", Name: "90年代"})
	yearItem = append(yearItem, model.FilterItem{Value: "80年代", Name: "80年代"})
	yearItem = append(yearItem, model.FilterItem{Value: "更早", Name: "更早"})

	year := model.FilterList{
		Key:   "year",
		Title: "年份",
		Items: yearItem,
	}

	style := model.FilterList{
		Title: "风格",
		Key:   "style",
		Items: []model.FilterItem{
			{Value: "0", Name: "全部"},
		},
	}

	styles := service.StyleRepository.FindAll()
	for _, item := range styles {
		style.Items = append(style.Items, model.FilterItem{Value: strconv.Itoa(item.StyleId), Name: item.Title})
	}

	var order model.FilterList
	order.Title = "排序"
	order.Items = []model.FilterItem{
		{Value: "0", Name: "开播时间", Sort: "0,1"},
		{Value: "1", Name: "最近更新", Sort: "0,1"},
		//{Value: "2", Name: "最高评分", Sort: "0"},
		{Value: "3", Name: "追番人数", Sort: "0,1"},
		{Value: "4", Name: "人气指数", Sort: "0,1"},
	}

	var list model.FilterModel
	list.Filter = append(list.Filter, category, version, area, state, month, year, style)
	list.Order = order
	return list
}

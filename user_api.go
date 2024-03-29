package qingflowapi

import (
	"fmt"
	"strconv"
)

type UserApi struct {
	client Client
}

func NewUserApi(client Client) UserApi {
	return UserApi{client: client}
}

type User struct {
	AreaCode         string `json:"areaCode"`
	BeingActive      bool   `json:"beingAtive"`
	BeingDisabled    bool   `json:"beingDisabled"`
	CustomDepartment []ID   `json:"customDepartment"`
	CustomRole       []ID   `json:"customRole"`
	Department       []ID   `json:"department"`
	Email            string `json:"email"`
	HeadImg          string `json:"headImg"`
	MobileNum        string `json:"mobileNum"`
	Name             string `json:"name"`
	Role             []ID   `json:"role"`
	UserId           SID    `json:"userId"`
	OptionId         ID     `json:"optionId"`
	SuperiorId       ID     `json:"superiorId"`
}

/*
仅专有轻流支持；
「超管accessToken」和「权限组accessToken」都可以调用，「权限组accessToken」须具有通讯录权限。
*/
func (api UserApi) Delete(userId SID) error {
	endpoint := fmt.Sprintf("user/%s", userId)
	var result ApiResponse[User]
	err := api.client.deleteRequest(endpoint, nil, &result)
	if err != nil {
		return err
	}
	return nil
}

/*
仅支持专有轻流；
「超管accessToken」和「权限组accessToken」都可以调用，「权限组accessToken」须具有通讯录权限。
*/
func (api UserApi) Update(user User) error {
	endpoint := fmt.Sprintf("user/%s", user.UserId)
	var result ApiResponse[User]
	err := api.client.post(endpoint, user, &result)
	if err != nil {
		return err
	}
	return nil
}

func (api UserApi) Create(user User) (SID, error) {
	endpoint := "user"
	var result ApiResponse[User]
	err := api.client.post(endpoint, user, &result)
	if err != nil {
		return "", err
	}
	return result.Result.UserId, nil
}

/*
仅专有轻流支持；
「超管accessToken」和「权限组accessToken」都可以调用，「权限组accessToken」须具有通讯录权限。
*/
func (api UserApi) CreateBatch(users []User) (RequestID, error) {
	endpoint := "user/batch"
	request := map[string]any{
		"members": users,
	}
	var result ApiResponse[struct {
		RequestId RequestID `json:"requestId"`
	}]
	err := api.client.post(endpoint, request, &result)
	if err != nil {
		return "", err
	}
	return result.Result.RequestId, nil
}

/*
「超管accessToken」和「权限组accessToken」都可以调用，「权限组accessToken」须具有通讯录权限。
*/
func (api UserApi) Get(userId SID) (User, error) {
	endpoint := fmt.Sprintf("user/%s", userId)
	var result ApiResponse[User]
	err := api.client.get(endpoint, nil, &result)
	if err != nil {
		return User{}, err
	}
	return result.Result, nil
}

func (api UserApi) Page(page int, size int) (PageResult[User], error) {
	endpoint := "user"
	params := map[string]string{
		"pageSize": strconv.Itoa(page),
		"pageNum":  strconv.Itoa(size),
	}
	var result ApiResponse[PageResult[User]]
	err := api.client.get(endpoint, params, &result)
	if err != nil {
		return PageResult[User]{}, err
	}
	return result.Result, nil
}

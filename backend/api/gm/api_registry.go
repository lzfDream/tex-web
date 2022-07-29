package gm

import (
    "strconv"
    "strings"
    "github.com/labstack/echo"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-go/sdp/rpc"
    "github.com/yellia1989/tex-web/backend/common"
)

func RegistryList(c echo.Context) error {
    ctx := c.(*mid.Context)
    page, _ := strconv.Atoi(ctx.QueryParam("page"))
    limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

    comm := common.GetLocator()

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    var vObj []rpc.ObjEndpoint
    ret, err := queryPrx.GetAllEndpoints(&vObj)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    vPage := common.GetPage(vObj, page, limit)
    return ctx.SendArray(vPage, len(vObj))
}

func registryAdd(sObj string, sDivision string, sEp string) error {
    comm := common.GetLocator()

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    ret, err := queryPrx.AddEndpoint(sObj, sDivision, sEp)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    return nil
}

func registryDel(sObj string, sDivision string, sEp string) error {
    comm := common.GetLocator()

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    ret, err := queryPrx.RemoveEndpoint(sObj, sDivision, sEp)
    if err := checkRet(ret, err); err != nil {
        return err
    }

    if err := checkRet(ret, err); err != nil {
        return err
    }

    return nil
}

func RegistryAdd(c echo.Context) error {
    ctx := c.(*mid.Context)
    sObj := ctx.FormValue("sObj")
    sDivision := ctx.FormValue("sDivision")
    sEp := ctx.FormValue("sEp")

    if sObj == "" || sDivision == "" || sEp == "" {
        return ctx.SendError(-1, "参数非法")
    }

    if err := registryAdd(sObj, sDivision, sEp); err != nil {
        return err
    }

    return ctx.SendResponse("添加registry成功")
}

func RegistryDel(c echo.Context) error {
    ctx := c.(*mid.Context)

    ids := strings.Split(ctx.FormValue("idsStr"), ",")
    if len(ids) == 0 {
        return ctx.SendError(-1, "registry不存在")
    }

    comm := common.GetLocator()

    queryPrx := new(rpc.Query)
    comm.StringToProxy("tex.mfwregistry.QueryObj", queryPrx)

    for _, id := range ids {
        tmp := strings.Split(id, "$")
        if len(tmp) != 3 {
            return ctx.SendError(-1, "参数非法")
        }
        ret, err := queryPrx.RemoveEndpoint(tmp[0], tmp[1], tmp[2])
        if err := checkRet(ret, err); err != nil {
            return err
        }
    }

    return ctx.SendResponse("删除registry成功")
}

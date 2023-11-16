package log

// EditRoles
//	@Tags			系统管理
//	@Summary		编辑角色
//	@Description	这是一段接口描述
//	@Produce		json
//	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
//	@Router			/api/v1/sys/roles/:id [put]
//func EditRoles(c *gin.Context) {
//	code := e.SUCCESS
//	data := make(map[string]interface{})
//
//	ID := c.Param("id")
//	Name := c.Query("name")
//	AssignedPerson := c.Query("assigned_person")
//
//	data["name"] = Name
//	data["assigned_person"] = AssignedPerson
//
//	id, _ := strconv.Atoi(ID)
//	sys.EditRoles(id, data)
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
//// DeleteErrorLogs
////	@Tags			系统管理
////	@Summary		删除角色
////	@Description	这是一段接口描述
////	@Produce		json
////	@Success		200	{string}	json	"{"code":200,"data":{},"msg":"ok"}"
////	@Router			/api/v1/sys/roles/:id [delete]
//func DeleteErrorLogs(c *gin.Context) {
//	code := e.SUCCESS
//	ID := c.Param("id")
//	id, _ := strconv.Atoi(ID)
//
//	sys.DeleteRoles(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}

package routes

import (
	v1 "YamadaKESHE/api/v1"
	"YamadaKESHE/middleware"
	"YamadaKESHE/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//r.HTMLRender = createMyRender()
	//r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//教师管理
	Teacher := r.Group("api/v1")
	Teacher.POST("Login", v1.Login)
	Teacher.Use(middleware.JwtTokenTeacher())
	{
		//查询班级成员
		Teacher.GET("admin/users", v1.GetUsers)
		//签到
		Teacher.GET("admin/sigincheckclass", v1.SIgnInCheckClass)
		Teacher.GET("admin/sigincheckall", v1.SIgnInCheckAll)
		Teacher.GET("admin/siginchecksignday", v1.SIgnInCheckSignDay)
		//请假
		Teacher.GET("admin/leavecheckclass", v1.LeaveCheckClass)
		Teacher.GET("admin/leavecheckall", v1.LeaveCheckAll)
		Teacher.GET("admin/leavechecksignday", v1.LeaveCheckSignDay)
		Teacher.POST("admin/leaveupdate", v1.LeaveUpdateTeacher)
		//实习报告
		Teacher.GET("admin/reportcheckclass", v1.ReportCheckClass)
		Teacher.GET("admin/reportcheckuserid", v1.ReportCheckuserid)
		Teacher.GET("admin/reportcheckday", v1.ReportCheckDay)
		Teacher.POST("admin/reportdelete", v1.ReportDelete)
		//工作日志
		Teacher.POST("admin/worklogpost", v1.WorkLogWrite)
		Teacher.POST("admin/worklogdupdate", v1.WorkLogUpdate)
		Teacher.GET("admin/worklogcheckclass", v1.WorkLogCheckClass)
		Teacher.GET("admin/worklogcheckuserid", v1.WorkLogCheckuserid)
		Teacher.GET("admin/worklogcheckday", v1.WorkLogCheckDay)
		Teacher.GET("admin/worklogcheckitself", v1.WorkLogCheckItself)
		//实习通知
		Teacher.POST("admin/noticepost", v1.NoticeWrite)
		Teacher.POST("admin/noticeupdate", v1.NoticeUpdate)
		Teacher.GET("admin/noticecheckclass", v1.NoticeCheckClass)
		Teacher.GET("admin/noticecheckday", v1.NoticeCheckDay)
		Teacher.GET("admin/noticecheckitself", v1.NoticeCheckItselfTeacher)
		Teacher.POST("admin/noticedelete", v1.NoticeDelete)
		//企业需求
		Teacher.POST("admin/demandpost", v1.DemandWrite)
		Teacher.POST("admin/demandupdate", v1.DemandUpdate)
		Teacher.POST("admin/demanddelete", v1.DemandDelete)
		Teacher.GET("admin/demandcheckclass", v1.DemandCheckClass)
		Teacher.GET("admin/demandcheckall", v1.DemandCheckAll)
		Teacher.GET("admin/demandcheckday", v1.DemandCheckDay)
		Teacher.GET("admin/demandcheckitself", v1.DemandCheckItself)
		//实习能力考核
		Teacher.POST("admin/assessmentcreate", v1.AssessmentWriteAll)
		Teacher.POST("admin/assessmentupdate", v1.AssessmentUpdate)
		Teacher.GET("admin/assessmentcheckclass", v1.AssessmentCheckClass)
		Teacher.GET("admin/assessmentcheckuserid", v1.AssessmentCheckuserid)
		Teacher.GET("admin/assessmentcheckall", v1.AssessmentCheckAll)
		Teacher.POST("admin/assessmentdelete", v1.AssessmentDelete)
		//学业成绩
		Teacher.POST("admin/gpacreate", v1.GpaWriteAll)
		Teacher.POST("admin/gpaupdate", v1.GpaUpdate)
		Teacher.POST("admin/gpacheckupdate", v1.GpaCheckUpdate)
		Teacher.GET("admin/gpacheckclass", v1.GpaCheckClass)
		Teacher.GET("admin/gpacheckuserid", v1.GpaCheckuserid)
		Teacher.GET("admin/gpacheckall", v1.GpaCheckAll)
		Teacher.POST("admin/gpaupdelete", v1.GpaDelete)
		//实时定位
		Teacher.POST("admin/rtslcheck", v1.RstlCheck)
	}

	//学生管理
	studentV1 := r.Group("api/v1")
	//登录模块
	studentV1.POST("Loginfront", v1.LoginFront)
	studentV1.Use(middleware.JwtTokenStudent())
	{
		//签到
		studentV1.POST("student/sigin", v1.SignInCreat)
		studentV1.GET("student/signcheckitself", v1.SignInCheckItself)
		//请假
		studentV1.POST("student/leavepost", v1.LeavePost)
		studentV1.GET("student/leavecheckitself", v1.LeaveCheckItself)
		studentV1.POST("student/leaveupdate", v1.LeaveUpdateStudent)
		studentV1.POST("student/leavedelete", v1.LeaveDelete)
		//实习报告
		studentV1.POST("student/reportpost", v1.ReportWrite)
		studentV1.POST("student/reportdelete", v1.ReportDelete)
		studentV1.POST("student/reportupdate", v1.ReportUpdate)
		studentV1.GET("student/reportcheckitself", v1.ReportCheckItself)
		//工作日志
		studentV1.POST("student/worklogcheckpost", v1.WorkLogWrite)
		studentV1.GET("student/worklogcheckitself", v1.WorkLogCheckItself)
		studentV1.POST("student/worklogdupdate", v1.WorkLogUpdate)
		//实习通知
		studentV1.GET("student/noticecheckclass", v1.NoticeCheckClass)
		studentV1.GET("student/noticecheckday", v1.NoticeCheckDay)
		//企业需求
		studentV1.GET("student/demandcheckall", v1.DemandCheckAll)
		studentV1.GET("student/demandcheckday", v1.DemandCheckDay)
		//实习能力考核
		studentV1.GET("student/assessmentcheckitself", v1.AssessmentCheckItself)
		//学业成绩
		studentV1.GET("student/gpacheckitself", v1.GpaCheckItself)
		//实时定位
		studentV1.POST("student/rtslupdate", v1.RtslInCreat)
	}

	//本地模块
	LocalManager := r.Group("api/v1")
	LocalManager.POST("Locallogin", v1.LoginLocalManager)
	LocalManager.Use(middleware.JwtTokenLocalManager())
	{
		//注册
		LocalManager.POST("user/add", v1.AddUser)
		//赋予权限
		LocalManager.POST("user/update", v1.EditUser)
	}
	r.Run(utils.HttpPort)
}

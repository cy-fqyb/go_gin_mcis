SELECT 
a.id as "rowId"
,'WM2002' "uploadType"
,a.ID "uploadRowId"
,org_shoucha.license "__license" --授权码
,nvl(doctor_info.doctor_account,org_shoucha.xj_czyid) "__user" --授权User
,R."outId" as "recordId" --档案ID
,R_BASIC."outId" as "recordIdReg" --孕产登记ID
,a.M_NAME as "husbandName" --丈夫姓名
,(case when a.M_BIRTH_DATE is not null then to_char(a.M_BIRTH_DATE,'YYYY-MM-DD') when LENGTH(a.M_IDNO)=18 then substr(a.M_IDNO,7,4) || '-' || substr(a.M_IDNO,11,2) || '-' || substr(a.M_IDNO,13,2) else '' end) as "husbandBirthday" --丈夫出生日期
,to_char(a.M_AGE) as "husbandAge" --丈夫年龄
,a.M_PHONE as "husbandPhoneNo" --丈夫电话
,(case when a.M_JOB_TYPE_NAME is null then '0'
when a.M_JOB_TYPE_NAME='国家机关、党群组织、企业、事业单位负责人' then '1'
when a.M_JOB_TYPE_NAME='专业技术人员' then '2'
when a.M_JOB_TYPE_NAME='办事人员和有关人员' then '3'
when a.M_JOB_TYPE_NAME='商业、服务业人员' then '4'
when a.M_JOB_TYPE_NAME='农、林、牧、渔、水利业生产人员' then '5'
when a.M_JOB_TYPE_NAME='生产、运输设备操作人员及有关人员' then '6'
when a.M_JOB_TYPE_NAME='军人' then '7'
when a.M_JOB_TYPE_NAME='不便分类的其他从业人员' then '8'
when a.M_JOB_TYPE_NAME='无业或待业' then '9'
when a.M_JOB_TYPE_NAME='学生' then '10'
when a.M_JOB_TYPE_NAME='其他' then '11'
 else '' end) as "occupation" --丈夫职业
,a.M_NATION_CODE as "nation" --丈夫民族
,a.M_IDNO as "idcardNo" --丈夫身份证号
,a.M_COMPANY as "husbandWorkplace" --丈夫工作单位
,a.WM_YUNCI as "pregnantCount" --孕次
,a.wm_shunchan_count as "deliverCountV" --顺产
,a.wm_operate_count as "deliverCountC" --剖宫产
,to_char(a.WM_YUCHANQI_DATE,'YYYY-MM-DD') as "edcDate" --预产期
,(case when b.yun_week is not null then TO_CHAR(b.yun_week)
WHEN CALEYUNWEEK_WEEK(a.WM_MENSTRUAL_LAST,b.check_date) >= '13' then '12'
else CALEYUNWEEK_WEEK(a.WM_MENSTRUAL_LAST,b.check_date) end) as "gestWeeks" --随访孕周-周
,(case when b.yun_day is not null then TO_CHAR(b.yun_day)
else CALEYUNWEEK_DAY(a.WM_MENSTRUAL_LAST,b.check_date) end)  as "gestDays" --随访孕周-天
,SYFS_CODE "conceptionMode" -- 受孕方式
,a.TAISHU as "fetalnum" --胎数
,case when sc.zysc_code = '0' then '1' else '2' end as "voluntaryDelivery" --自愿顺产
,(CASE WHEN a.JMJKTJ_NAME = '是' THEN '1' WHEN a.JMJKTJ_NAME = '否' THEN '2' ELSE '2' END) as "healthExam" --全民健康检查：是 否
,(CASE WHEN a.JBGGWS_NAME is null THEN '1' WHEN a.JBGGWS_NAME = '是' THEN '1' WHEN a.JBGGWS_NAME = '否' THEN '2' ELSE '1' END) as "basicPubhealth" --基本公共卫生：是 否
,'select DISTINCT platform_code item from ZX_DIC_MAP_TAB where status=1 and TYPE_CODE=''pastHistory'' and instr('''||a.wm_sick_history_qt_name||''',local_name)>0' as "SArr:pastHistory" --既往史
,a.history_sick_other as "pastOther" --既往史-其他
,(case when (a.wm_family_history_name is null or a.wm_family_history_name='无特殊')
 and (a.m_family_history_name is null or a.m_family_history_name='无特殊')
 and (a.family_father_name is null or a.family_father_name='无特殊')
 and (a.family_mother_name is null or a.family_mother_name='无特殊')
 and (a.family_brother_name is null or a.family_brother_name='无特殊')
 and (a.family_children_name is null or a.family_children_name='无特殊')
 then ''
 else '' end) as "fh" --家族史
,'' as "fhOther" --家族史-其他
,'select DISTINCT platform_code item from ZX_DIC_MAP_TAB where status=1 and TYPE_CODE=''self'' and instr('''||
(case when a.wm_xiyan_name='有' then ',吸烟' else '' end)||(case when a.wm_yinjiu_name='有' then ',饮酒' else '' end)||(case when a.wm_fuyao_name='有' then ',服用药物' else '' end)
||(case when a.wm_contact_harmful_factor_name='有' then ',接触有毒有害物质' else '' end)||(case when a.wm_jiechu_fangshexian_name='有' then ',接触放射线' else '' end)||(case when a.grs_qt_name='有' then ',其他' else '' end)
||''',local_name)>0'  as "SArr:self" --个人史
,a.GRS_QT_DESC as "selfOther" --个人史-其他
,a.SSS_SM as "gyn" --妇科手术史
,'' as "gynOther" --妇科手术史-其他
,a.WM_RENGONG_LIUCHAN_COUNT as "mhAbortion" --孕产史-流产
,a.WM_SITAI_COUNT as "mhDeathFetus" --孕产史-死胎
,a.WM_SICHAN_COUNT as "mhDeathDelivery" --孕产史-死产
,a.WM_XINSHENER_DEAD_COUNT as "mhNewbornDeath" --孕产史-新生儿死亡
,a.WM_BIRTH_DEFECT_COUNT as "mhBirthDefect" --孕产史-出生缺陷儿
,'' as "mhLiveBirth" --孕产史-活产
,a.WM_RENSHEN_SICK as "mhPastPncomp" --孕产史-既往妊娠合并症并发症
,(CASE WHEN a.YQBJ_YW_NAME = '是' or a.YQBJ_YW_NAME = '有' THEN '1' WHEN a.YQBJ_YW_NAME = '否' or a.YQBJ_YW_NAME = '无' THEN '0' ELSE '' END) as "pregCare" --孕前保健-有无
,(CASE WHEN a.YQBJ_YSZX_NAME = '是' or a.YQBJ_YSZX_NAME = '有' THEN '1' WHEN a.YQBJ_YSZX_NAME = '否' or a.YQBJ_YSZX_NAME = '无' THEN '0' ELSE '' END) as "pgcHealthyBabies" --孕前保健-优生咨询
,(CASE WHEN a.YQBJ_YQJC_NAME = '是' or a.YQBJ_YQJC_NAME = '有' THEN '1' WHEN a.YQBJ_YQJC_NAME = '否' or a.YQBJ_YQJC_NAME = '无' THEN '0' ELSE '' END) as "pgcPreCheck" --孕前保健-孕前检查
,(CASE WHEN a.YQBJ_YQZB_NAME = '是' or a.YQBJ_YQZB_NAME = '有' THEN '1' WHEN a.YQBJ_YQZB_NAME = '否' or a.YQBJ_YQZB_NAME = '无' THEN '0' ELSE '' END) as "pgcPrePrepare" --孕前保健-孕前准备
,a.YQBJ_QT as "pgcOther" --孕前保健-其他
,(CASE WHEN a.WM_IS_TAKE_YESUAN_NAME = '是' or a.WM_IS_TAKE_YESUAN_NAME = '有' THEN '1' WHEN a.WM_IS_TAKE_YESUAN_NAME = '否' or a.WM_IS_TAKE_YESUAN_NAME = '无' THEN '0' ELSE '' END) as "pgcFolicAcid" --孕前保健-服用叶酸
,a.ys_lqr as "pgcFaReceiveName" --孕前保健-叶酸-领取人姓名
,a.YS_SL as "pgcFaSendQuantity" --孕前保健-叶酸-发放数量
,to_char(a.YS_FFSJ,'YYYY-MM-DD') as "pgcFaSendDate" --孕前保健-叶酸-发放时间
,a.YS_FFDW as "pgcFaSendUnit" --孕前保健-叶酸-发放单位
,(CASE WHEN a.YC_YJDZHX_NAME = '是' or a.YC_YJDZHX_NAME = '有' THEN '1' WHEN  a.YC_YJDZHX_NAME = '否' or a.YC_YJDZHX_NAME = '无' THEN '0' ELSE '' END) as "gestAbnSitnight" --妊娠期异常-夜间端坐呼吸
,(CASE WHEN a.YC_KS_NAME = '是' or a.YC_KS_NAME = '有' THEN '1' WHEN  a.YC_KS_NAME = '否' or a.YC_KS_NAME = '无' THEN '0' ELSE '' END) as "gestCough" --妊娠期异常-咳嗽
,(CASE WHEN a.YC_FG_NAME = '是' or a.YC_FG_NAME = '有' THEN '1' WHEN  a.YC_FG_NAME = '否' or a.YC_FG_NAME = '无' THEN '0' ELSE '' END) as "gestCyanosis" --妊娠期异常-发绀
,(CASE WHEN a.YC_FL_NAME = '是' or a.YC_FL_NAME = '有' THEN '1' WHEN  a.YC_FL_NAME = '否' or a.YC_FL_NAME = '无' THEN '0' ELSE '' END) as "gestFatigue" --妊娠期异常-乏力
,(CASE WHEN a.YC_EX_NAME = '是' or a.YC_EX_NAME = '有' THEN '1' WHEN  a.YC_EX_NAME = '否' or a.YC_EX_NAME = '无' THEN '0' ELSE '' END) as "gestFeelSick" --妊娠期异常-恶心
,(CASE WHEN a.YC_SFBBS_NAME = '是' or a.YC_SFBBS_NAME = '有' THEN '1' WHEN  a.YC_SFBBS_NAME = '否' or a.YC_SFBBS_NAME = '无' THEN '0' ELSE '' END) as "gestDiscomfort" --妊娠期异常-上腹部不适
,(CASE WHEN a.YC_OT_NAME = '是' or a.YC_OT_NAME = '有' THEN '1' WHEN  a.YC_OT_NAME = '否' or a.YC_OT_NAME = '无' THEN '0' ELSE '' END) as "gestVomit" --妊娠期异常-呕吐
,(CASE WHEN a.YC_DXT_NAME = '是' or a.YC_DXT_NAME = '有' THEN '1' WHEN  a.YC_DXT_NAME = '否' or a.YC_DXT_NAME = '无' THEN '0' ELSE '' END) as "gestHypoglycemia" --妊娠期异常-低血糖
,(CASE WHEN a.YC_GZMXYC_NAME = '是' or a.YC_GZMXYC_NAME = '有' THEN '1' WHEN  a.YC_GZMXYC_NAME = '否' or a.YC_GZMXYC_NAME = '无' THEN '0' ELSE '' END) as "gestUnliver" --妊娠期异常-肝脏酶学异常
,(CASE WHEN a.YC_XS_NAME = '是' or a.YC_XS_NAME = '有' THEN '1' WHEN  a.YC_XS_NAME = '否' or a.YC_XS_NAME = '无' THEN '0' ELSE '' END) as "gestLoseWeight" --妊娠期异常-消瘦
,(CASE WHEN a.YC_CH_NAME = '是' or a.YC_CH_NAME = '有' THEN '1' WHEN  a.YC_CH_NAME = '否' or a.YC_CH_NAME = '无' THEN '0' ELSE '' END) as "gestSweat" --妊娠期异常-出汗
,(CASE WHEN a.YC_NXGNYC_NAME = '是' or a.YC_NXGNYC_NAME = '有' THEN '1' WHEN  a.YC_NXGNYC_NAME = '否' or a.YC_NXGNYC_NAME = '无' THEN '0' ELSE '' END) as "gestDyblood" --妊娠期异常-凝血功能异常
,(CASE WHEN a.YC_SSZC_NAME = '是' or a.YC_SSZC_NAME = '有' THEN '1' WHEN  a.YC_SSZC_NAME = '否' or a.YC_SSZC_NAME = '无' THEN '0' ELSE '' END) as "gestTremorHands" --妊娠期异常-双手震颤
,(CASE WHEN a.YC_YT_NAME = '是' or a.YC_YT_NAME = '有' THEN '1' WHEN  a.YC_YT_NAME = '否' or a.YC_YT_NAME = '无' THEN '0' ELSE '' END) as "gestOcularProcess" --妊娠期异常-眼突
,(CASE WHEN a.YC_XXBJS_NAME = '是' or a.YC_XXBJS_NAME = '有' THEN '1' WHEN  a.YC_XXBJS_NAME = '否' or a.YC_XXBJS_NAME = '无' THEN '0' ELSE '' END) as "gestThrombocytopenia" --妊娠期异常-血小板减少
,a.YBJC_HX_CODE as "gestRespiration" --妊娠期异常-呼吸(次/分)
,(case when a.ZSZZ_XM_NAME='是' or a.ZSZZ_XM_NAME='有' then '1'
 when a.ZSZZ_XM_NAME='否' or a.ZSZZ_XM_NAME='无' then '2' else '' end) as "gestChestDistress" --妊娠期异常-胸闷
,(case when a.WM_FUZHONG_NAME='是' or a.WM_FUZHONG_NAME='有' then '1'
 when a.WM_FUZHONG_NAME='否' or a.WM_FUZHONG_NAME='无' then '2' else '' end) as "gestOedema" --妊娠期异常-浮肿
,a.YBJC_XL_CODE as "gestHeartRate" --妊娠期异常-心率
,(case when a.ZSZZ_HD_NAME='是' or a.ZSZZ_HD_NAME='有' then '1'
 when a.ZSZZ_HD_NAME='否' or a.ZSZZ_HD_NAME='无' then '2' else '' end) as "gestJaundice" --妊娠期异常-黄疸
,(case when a.ZSZZ_XJ_NAME='是' or a.ZSZZ_XJ_NAME='有' then '1'
 when a.ZSZZ_XJ_NAME='否' or a.ZSZZ_XJ_NAME='无' then '2' else '' end) as "gestPalpitate" --妊娠期异常-心悸
,a.wm_height as "giHeight" --一般检查-身高(CM)
,(CASE WHEN sc.wm_weight IS NULL THEN to_char(a.wm_weight) ELSE sc.wm_weight END) as "giWeight" --一般检查-体重(kg)
,(CASE when sc.WM_WEIGHT IS NOT NULL AND a.WM_HEIGHT IS NOT NULL THEN round(sc.WM_WEIGHT/(a.WM_HEIGHT*0.01*a.WM_HEIGHT*0.01),2)
 ELSE a.WM_PRE_BMI END) as "giBodyMass" --一般检查-体质指数
,(CASE WHEN sc.wm_blood_high IS NOT NULL THEN sc.wm_blood_high ELSE a.wm_blood_high END) as "giSystolicPressure" --一般检查-血压-收缩压(mmHg)
,(CASE WHEN sc.wm_blood_low IS NOT NULL THEN sc.wm_blood_low ELSE a.wm_blood_low END) as "giDiastolicPressure" --一般检查-血压-舒张压(mmHg)
,(case when sc.XINZANG_NAME IS NULL OR  sc.XINZANG_NAME = '未见异常' or sc.XINZANG_NAME='正常' or sc.XINZANG_NAME='其他' then '1' else '2' end) as "atHeart" --听诊-心脏
,(case when sc.FEIBU_NAME IS NULL OR  sc.FEIBU_NAME = '未见异常' or sc.FEIBU_NAME='正常' or sc.FEIBU_NAME='其他'  then '1' else '2' end) as "atLung" --听诊-肺部
,(case when sc.WAIYIN_NAME IS NULL OR  sc.WAIYIN_NAME = '未见异常' or sc.WAIYIN_NAME='正常' or sc.WAIYIN_NAME='其他' then '1' else '2' end) as "gyneVulva" --妇科检查-外阴
,(case when sc.YINDAO_NAME IS NULL OR  sc.YINDAO_NAME = '未见异常' or sc.YINDAO_NAME='正常' OR sc.YINDAO_NAME = '通畅' or sc.YINDAO_NAME='其他' then '1' else '2' end) as "gyneVagina" --妇科检查-阴道
,(case when sc.GONGJING_NAME IS NULL OR  sc.GONGJING_NAME = '未见异常' OR sc.GONGJING_NAME = '正常' or sc.GONGJING_NAME='其他' then '1' else '2' end) as "gyneCervix" --妇科检查-宫颈
,(case when sc.GONGTI_NAME IS NULL OR  sc.GONGTI_NAME = '未见异常' or sc.GONGTI_NAME='其他' OR sc.GONGTI_NAME LIKE '%正常%' then '1' else '2' end) as "gyneWomb" --妇科检查-子宫
,(case when sc.FUJIAN_NAME IS NULL OR  sc.FUJIAN_NAME = '未见异常' or sc.FUJIAN_NAME='其他' OR sc.FUJIAN_NAME LIKE '%正常%' then '1' else '2' end) as "gyneAttachment" --妇科检查-附件
,'' as "gyneOther" --妇科检查-其他
,(CASE WHEN B.XUEHONG_DANBAI IS NULL OR INSTR(B.XUEHONG_DANBAI, '未')>0 or B.XUEHONG_DANBAI='/' or B.XUEHONG_DANBAI='-' THEN ''
 when regexp_like (trim(B.XUEHONG_DANBAI),'([^.0-9])+') then null 
ELSE to_char(to_number(TRIM(B.XUEHONG_DANBAI))) END) as "axRbHemoglobinValue" --辅助检查-血常规-血红蛋白值(g/L)
,(CASE WHEN B.XUEHONG_DANBAI IS NULL OR INSTR(B.XUEHONG_DANBAI, '未')>0 or B.XUEHONG_DANBAI='/' or B.XUEHONG_DANBAI='-' THEN '1'
 when regexp_like (trim(B.XUEHONG_DANBAI),'([^.0-9])+') then '1' 
ELSE '0' END) as "axRbHemoglobin" --辅助检查-血常规-血红蛋白值-未查
,b.BAI_XIBAO as "axRbWbc" --辅助检查-血常规-白细胞计数(X10^9/L)
,b.WM_XUEXIAOBAN_COUNT as "axRbPlatelet" --辅助检查-血常规-血小板总数(X10^9/L)
,'' as "axRbAnemiaDiagnosis" --辅助检查-血常规-贫血诊断
,b.XCG_QT as "axRbOther" --辅助检查-血常规-其他
,'' as "axRbTips" --辅助检查-血常规-贫血提示
,(case when B.WM_NIAO_DANBAI_NAME is null or B.WM_NIAO_DANBAI_NAME='未查' then '0' else B.WM_NIAO_DANBAI_NAME end) as "axUrUrinaryAlbumin" --辅助检查-尿常规-尿蛋白
,(case when B.WM_NIAO_TANG_NAME is null or B.WM_NIAO_TANG_NAME='未查' then '0' else B.WM_NIAO_TANG_NAME end) as "axUrUrineSugar" --辅助检查-尿常规-尿糖
,(case when B.WM_NIAO_TONGTI_NAME is null or B.WM_NIAO_TONGTI_NAME='未查' then '0' else B.WM_NIAO_TONGTI_NAME end) as "axUrUrineKetone" --辅助检查-尿常规-尿酮体
,(case when b.NCG_NQX_NAME is null or B.NCG_NQX_NAME='未查' then '0' else b.NCG_NQX_NAME end) as "axUrUrineOccult" --辅助检查-尿常规-尿潜血
,(CASE WHEN B.WM_NIAO_CHANGGUI_OTHER IS NULL or B.WM_NIAO_CHANGGUI_OTHER='2' THEN '0'
 WHEN B.WM_NIAO_CHANGGUI_OTHER='无特殊情况' OR B.WM_NIAO_CHANGGUI_OTHER='未见明显异常' OR B.WM_NIAO_CHANGGUI_OTHER='无' OR B.WM_NIAO_CHANGGUI_OTHER='-' THEN '1'
 when B.WM_NIAO_CHANGGUI_OTHER='异常' then '2'
 ELSE ''  END) as "axUrOther" --辅助检查-尿常规-其他
,(case when b.SYSJC_NCG_NAME is null then '' else b.SYSJC_NCG_NAME end) as "axUrConclusion" --辅助检查-尿常规-结论
,(case when b.WM_BLOOD_NAME is null then ''
 when b.WM_BLOOD_NAME='不详' or b.WM_BLOOD_NAME='未查' then '0'
 when b.WM_BLOOD_NAME='A型' then '1'
 when b.WM_BLOOD_NAME='B型' then '2'
 when b.WM_BLOOD_NAME='O型' then '3'
 when b.WM_BLOOD_NAME='AB型' then '4'
 else '' end) as "axAbo" --辅助检查-血型检查
,(CASE b.WM_XUEXING_RH_NAME WHEN '阳性' THEN '2' WHEN '阴性' THEN '1' when '不详' then '0' when '未查' then '0'  ELSE '' END) as "axRh" --辅助检查-RH 阴型检查
,(case when INSTR(b.WM_XUETANG, '未')>0 or b.WM_XUETANG='/' or b.WM_XUETANG='-' or b.WM_XUETANG='无' then '' else b.WM_XUETANG end) as "axBloodSugar" --辅助检查-血糖(mmol/L)
,(case when INSTR(b.WM_GANGONG_XQGBZAM, '未')>0 or b.WM_GANGONG_XQGBZAM='/' or b.WM_GANGONG_XQGBZAM='-' or b.WM_GANGONG_XQGBZAM='无' then '' else b.WM_GANGONG_XQGBZAM end) as "axLiverAst" --辅助检查-肝功能-血清谷丙转氨酶(U/L)
,(case when INSTR(b.WM_GANGONG_XQGCZAM, '未')>0 or b.WM_GANGONG_XQGCZAM='/' or b.WM_GANGONG_XQGCZAM='-' or b.WM_GANGONG_XQGCZAM='无' then '' else b.WM_GANGONG_XQGCZAM end) as "axLiverSgot" --辅助检查-肝功能-血清谷草转氨酶(U/L
,(case when INSTR(b.WM_GANGONG_BDB, '未')>0 or b.WM_GANGONG_BDB='/' or b.WM_GANGONG_BDB='-' or b.WM_GANGONG_BDB='无' then '' else b.WM_GANGONG_BDB end) as "axLiverAlb" --辅助检查-肝功能-白蛋白(g/L)
,(case when INSTR(b.WM_GANGONG_ZDHS, '未')>0 or b.WM_GANGONG_ZDHS='/' or b.WM_GANGONG_ZDHS='-' or b.WM_GANGONG_ZDHS='无' then '' else b.WM_GANGONG_ZDHS end) as "axLiverTbil" --辅助检查-肝功能-总胆红素(μmol/L)
,CASE 
    WHEN INSTR(b.WM_GANGONG_JHDHS, '未') > 0 
         OR b.WM_GANGONG_JHDHS = '/' 
         OR b.WM_GANGONG_JHDHS = '-' 
         OR b.WM_GANGONG_JHDHS = '无' 
    THEN NULL
		when regexp_like (trim(b.WM_GANGONG_JHDHS),'([^.0-9])+') then null 
    ELSE TO_NUMBER(b.WM_GANGONG_JHDHS)
END AS "axLiverDbil" --辅助检查-肝功能-结合胆红素(μmol/L)
,(case when b.WM_GANGONG_NAME='未见异常'  then '1'
when b.WM_GANGONG_NAME='异常'  then '2' 
when b.WM_GANGONG_NAME='未做'  then '0'
when b.WM_GANGONG_NAME is null then '0'
else b.WM_GANGONG_NAME end) as "axLiverConclusion" --辅助检查-肝功能-结论
,(case when INSTR(b.WM_SHENGONG_XQJG, '未')>0 or b.WM_SHENGONG_XQJG='/' or b.WM_SHENGONG_XQJG='-' or b.WM_SHENGONG_XQJG='无' then '' else TRIM(b.WM_SHENGONG_XQJG) end) as "axKidneyCre" --辅助检查-肾功能-血清肌酐(μmol/L)
,(case when INSTR(b.WM_SHENGONG_XNSD, '未')>0 or b.WM_SHENGONG_XNSD='/' or b.WM_SHENGONG_XNSD='-' or b.WM_SHENGONG_XNSD='无' then '' else TRIM(b.WM_SHENGONG_XNSD) end) as "axKidneyBun" --辅助检查-肾功能-血尿素氮(mmol/L)
,(case when b.WM_SHENGONG_NAME='未见异常'  then '1'
when b.WM_SHENGONG_NAME='异常'  then '2' 
when b.WM_SHENGONG_NAME='未做'  then '0'
when b.WM_SHENGONG_NAME is null then '0'
else b.WM_SHENGONG_NAME end) as "axKidneyConclusion" --辅助检查-肾功能-结论
,(CASE WHEN B.WM_HBSAG_NAME='阴性' THEN '-' WHEN B.WM_HBSAG_NAME IS NULL or INSTR(B.WM_HBSAG_NAME, '未')>0 THEN '0' WHEN B.WM_HBSAG_NAME='＋' THEN '+' WHEN B.WM_HBSAG_NAME='－' THEN '-' ELSE B.WM_HBSAG_NAME END) as "axHbHbsag" --辅助检查-乙型肝炎五项-乙型肝炎表面抗原
,(CASE WHEN B.WM_HBS_NAME='阴性' THEN '-' WHEN B.WM_HBS_NAME IS NULL or INSTR(B.WM_HBS_NAME, '未')>0  THEN '0' WHEN B.WM_HBS_NAME='＋' THEN '+' WHEN B.WM_HBS_NAME='－' THEN '-' ELSE B.WM_HBS_NAME END) as "axHbHbsab" --辅助检查-乙型肝炎五项-乙型肝炎表面抗体
,(CASE WHEN B.WM_HBEAG_NAME='阴性' THEN '-' WHEN B.WM_HBEAG_NAME IS NULL or INSTR(B.WM_HBEAG_NAME, '未')>0  THEN '0' WHEN B.WM_HBEAG_NAME='＋' THEN '+' WHEN B.WM_HBEAG_NAME='－' THEN '-' ELSE B.WM_HBEAG_NAME END) as "axHbHbeag" --辅助检查-乙型肝炎五项-乙型肝炎 e 抗原
,(CASE WHEN B.WM_HBE_NAME='阴性' THEN '-' WHEN B.WM_HBE_NAME IS NULL or INSTR(B.WM_HBE_NAME, '未')>0  THEN '0' WHEN B.WM_HBE_NAME='＋' THEN '+' WHEN B.WM_HBE_NAME='－' THEN '-' ELSE B.WM_HBE_NAME END) as "axHbHbeab" --辅助检查-乙型肝炎五项-乙型肝炎 e 抗体
,(CASE WHEN B.WM_HBC_NAME='阴性' THEN '-' WHEN B.WM_HBC_NAME IS NULL or INSTR(B.WM_HBC_NAME, '未')>0 THEN '0' WHEN B.WM_HBC_NAME='＋' THEN '+' WHEN B.WM_HBC_NAME='－' THEN '-' ELSE B.WM_HBC_NAME END) as "axHbHbcab" --辅助检查-乙型肝炎五项-乙型肝炎核心抗体
,to_char(B.WM_YG_DATE,'YYYY-MM-DD') as "axHbDate" --辅助检查-乙型肝炎五项-乙肝检查时间
,CALEYUNWEEK(a.WM_MENSTRUAL_LAST,B.WM_YG_DATE) as "axHbGest" --辅助检查-乙型肝炎五项-乙肝检查孕周
,(CASE WHEN B.WM_YDFMW_NAME IS NULL then ''
 when B.WM_YDFMW_NAME='未见异常' or B.WM_YDFMW_NAME='正常' then '0'
 when B.WM_YDFMW_NAME='滴虫' then '1'
 when B.WM_YDFMW_NAME='假丝酵母菌' then '2'
 when B.WM_YDFMW_NAME='其他' then '9'
 ELSE ''  END) as "axVaginalFluid" --辅助检查-阴道分泌物
,(CASE WHEN B.WM_YDQJD_NAME IS NULL then '' when B.WM_YDQJD_NAME='未查' or B.WM_YDQJD_NAME='不详' THEN '0'
when B.WM_YDQJD_NAME='I度' then '1'
when B.WM_YDQJD_NAME='II度' then '2'
when B.WM_YDQJD_NAME='III度' then '3'
when B.WM_YDQJD_NAME='IV度' then '4'
  ELSE '' END) as "axCleaningDegree" --辅助检查-阴道清洁度
,(CASE WHEN B.WM_MEIDU_NAME IS NULL or instr(B.WM_MEIDU_NAME,'未')>0 THEN '0'
WHEN B.WM_MEIDU_NAME='阳性' OR B.WM_MEIDU_NAME='＋' THEN '+'
WHEN B.WM_MEIDU_NAME='阴性' OR B.WM_MEIDU_NAME='－' THEN '-' ELSE B.WM_MEIDU_NAME END) as "axUsr" --辅助检查-梅毒血清学试验
,to_char(B.WM_MEIDU_DATE,'YYYY-MM-DD') as "axUsrdate" --辅助检查-梅毒检查时间
,CALEYUNWEEK(a.WM_MENSTRUAL_LAST,B.WM_YG_DATE) as "axGest" --辅助检查-梅毒-孕周
,(CASE when B.WM_HIV_ZX_CODE is not null THEN B.WM_HIV_ZX_CODE ELSE '' END) as "axHivConsult" --辅助检查-HIV 咨询
,(CASE WHEN B.WM_HIV_NAME IS NULL or instr(B.WM_HIV_NAME,'未')>0 THEN '0'
WHEN B.WM_HIV_NAME='阳性' OR B.WM_HIV_NAME='＋' THEN '+'
WHEN B.WM_HIV_NAME='阴性' OR B.WM_HIV_NAME='－' THEN '-'
ELSE B.WM_HIV_NAME  END) as "axHivResult" --辅助检查-HIV 检测结果
,to_char(B.WM_HIV_DATE,'YYYY-MM-DD') as "axHivDate" --辅助检查-HIV 检测时间
,CALEYUNWEEK(a.WM_MENSTRUAL_LAST,B.WM_HIV_DATE) as "axHivGest" --辅助检查-HIV 检查孕周
,(CASE WHEN B.WM_TSZHZSC_NAME IS NULL OR B.WM_TSZHZSC_NAME='未检查' THEN '未筛查' ELSE B.WM_TSZHZSC_NAME END) as "axSyndrome" --辅助检查-唐氏综合症筛查
,(CASE WHEN B.WM_BCHAO_name IS NULL OR B.WM_BCHAO_name='未检' OR B.WM_BCHAO_name='未查' THEN '0' WHEN B.WM_BCHAO_name='未见异常' or B.WM_BCHAO_name='正常' THEN '1' when B.WM_BCHAO_name='异常' then '2' ELSE '' END) as "axCt" --辅助检查-B 超
,
case when b.WM_BG_CODE = '0' then '1'
else b.WM_BG_CODE end as "axHcConsult" --辅助检查-丙肝-检测结果
,b.wm_bg_time as "axHcDate" --辅助检查-丙肝-检测时间
,b.wm_bg_week as "axHcGest" --辅助检查-丙肝-检测孕周
,'' as "axEpilepsyConsult" --辅助检查-孕期子癫前期检测结果
,to_char(b.check_date,'YYYY-MM-DD') as "otVisitDate" --其他-随访日期
,(case when b.yun_week is not null and b.yun_day is not null then b.yun_week||'+'||b.yun_day when b.yun_week is not null and b.yun_day is null then b.yun_week||'+'|| '0' 
WHEN CALEYUNWEEK(a.WM_MENSTRUAL_LAST,b.check_date) >= '13' then '12+6'
else CALEYUNWEEK(a.WM_MENSTRUAL_LAST,b.check_date) end) as "otVisitGestweek" --其他-随访孕周
,org_shoucha.xj_org_name as "otVisitNoworg" --其他-本次随访单位-名称
,org_shoucha.xj_org_code as "otVisitNoworgcode" --其他-本次随访单位-编码
,(case when sc.check_doctor_name is not null then sc.check_doctor_name else b.check_user_name end) as "otVisitNowdoctor" --其他-随访医师签名
,CALEYUNWEEK_WEEK(sc.check_time, SC.YUYUE_TIME) as "otVisitNextdateWeek" --其他-下次随访日期-周
,to_char(sc.YUYUE_TIME,'YYYY-MM-DD') as "otVisitNextdate" --其他-下次随访日期
,org_shoucha.jcjg as "otVisitOrg" --其他-检查机构
,'0' as "prsResult" --妊娠风险筛查-结果
,'select DISTINCT platform_code item from ZX_DIC_MAP_TAB where status=1 and TYPE_CODE=''hcg'' and instr('''||sc.ZHIDAO||''',local_code)>0' as "SArr:hcg" --保健指导
, sc.ZTPG_SM as "hcgOther" --保健指导-其他说明
,sc.ZTPGNEW as "overallEvaluation" --总体评估
,'' as "diagAnomaly" --诊断-本次异常
,'' as "diagSusDisease" --诊断-可疑病症
,'' as "diagAutoJudgment" --诊断-自动判断
,(case when sc.wm_diagnosis is null then '孕育知识宣教，按时产检' else to_char(sc.wm_diagnosis) end) as "diagDoctors" --诊断-医生诊断
,(CASE sc.IS_ZHUANZHEN_NAME   --sc.SFZZ
 when '有'  THEN '1'
 when '是'  THEN '1'
 when '无'  THEN '2'
 when '否'  THEN '2'
 ELSE ''  END) as "attendRefSuggestion" --处理-转诊建议
,sc.ZHUANZHEN_REASON as "attendRefCause" --处理-原因
,sc.ZHUANZHEN_ORG_NAME as "attendHos" --处理-机构及科室
,(CASE WHEN sc.XINLIZHIDAO_OTHER IS NOT NULL THEN 
to_char(sc.XINLIZHIDAO_OTHER)
ELSE '保持产妇居住环境的通风通气； 多听音乐，放松心情，对胎儿的正常发育有好处； 多散步，做简单的家务劳动； 多吃鱼； 多吃水果； 多吃高蛋白的禽蛋； 多吃高纤维食品； 多含铁食物，必要时补充铁剂，复查血常规' END) as "attendCareGuidance" --处理-保健指导
,'' as "attendHyGuidance" --处理-孕期贫血保健指导
,SC.zzzz as "attendRefIndication" --处理-转诊指征
,sc.xinlizhidao_other as "attendRemarks" --处理-备注
,a.YBJC_XL_NAME as "heartRate" --心率
,a.wm_tiwen as "temperature" --体温
,(case when SC.wm_gonggao='无' or sc.wm_gonggao='/' then '' else sc.wm_gonggao end) as "cervixHeight" --宫底高度(cm)
,(case when SC.yaowei='无' or SC.yaowei='/' then '' else SC.yaowei end) as "abdomenCircumf" --腹围(cm)
,sc.wm_taidong_name as "foetalMovement" --胎动
,(case 
 when sc.wm_taiwei1_name is null and sc.wm_taiwei1_code is null then '23'
 when instr(sc.wm_taiwei1_name,'枕左前')>0 or instr(sc.wm_taiwei1_name,'左枕前')>0 then '1'
 when instr(sc.wm_taiwei1_name,'枕右前')>0 or instr(sc.wm_taiwei1_name,'右枕前')>0 then '2'
 when instr(sc.wm_taiwei1_name,'枕左后')>0 or instr(sc.wm_taiwei1_name,'左枕后')>0 then '3'
 when instr(sc.wm_taiwei1_name,'枕右后')>0 or instr(sc.wm_taiwei1_name,'右枕后')>0 then '4'
 when instr(sc.wm_taiwei1_name,'枕左横')>0 or instr(sc.wm_taiwei1_name,'左枕横')>0 then '5'
 when instr(sc.wm_taiwei1_name,'枕右横')>0 or instr(sc.wm_taiwei1_name,'右枕横')>0 then '6'
 when instr(sc.wm_taiwei1_name,'颏左前')>0 or instr(sc.wm_taiwei1_name,'左颏前')>0 then '7'
 when instr(sc.wm_taiwei1_name,'颏右前')>0 or instr(sc.wm_taiwei1_name,'右颏前')>0 then '8'
 when instr(sc.wm_taiwei1_name,'颏左后')>0 or instr(sc.wm_taiwei1_name,'左颏后')>0 then '9'
 when instr(sc.wm_taiwei1_name,'颏右后')>0 or instr(sc.wm_taiwei1_name,'右颏后')>0 then '10'
 when instr(sc.wm_taiwei1_name,'颏左横')>0 or instr(sc.wm_taiwei1_name,'左颏横')>0 then '11'
 when instr(sc.wm_taiwei1_name,'颏右横')>0 or instr(sc.wm_taiwei1_name,'右颏横')>0 then '12'
 when instr(sc.wm_taiwei1_name,'骶左前')>0 or instr(sc.wm_taiwei1_name,'左骶前')>0 then '13'
 when instr(sc.wm_taiwei1_name,'骶右前')>0 or instr(sc.wm_taiwei1_name,'右骶前')>0 then '14'
 when instr(sc.wm_taiwei1_name,'骶左后')>0 or instr(sc.wm_taiwei1_name,'左骶后')>0 then '15'
 when instr(sc.wm_taiwei1_name,'骶右后')>0 or instr(sc.wm_taiwei1_name,'右骶后')>0 then '16'
 when instr(sc.wm_taiwei1_name,'骶左横')>0 or instr(sc.wm_taiwei1_name,'左骶横')>0 then '17'
 when instr(sc.wm_taiwei1_name,'骶右横')>0 or instr(sc.wm_taiwei1_name,'右骶横')>0 then '18'
 when instr(sc.wm_taiwei1_name,'肩左前')>0 or instr(sc.wm_taiwei1_name,'左肩前')>0 then '19'
 when instr(sc.wm_taiwei1_name,'肩右前')>0 or instr(sc.wm_taiwei1_name,'右肩前')>0 then '20'
 when instr(sc.wm_taiwei1_name,'肩左后')>0 or instr(sc.wm_taiwei1_name,'左肩后')>0 then '21'
 when instr(sc.wm_taiwei1_name,'肩右后')>0 or instr(sc.wm_taiwei1_name,'右肩后')>0 then '22'
 when TRIM(sc.wm_taiwei1_name)='胎位不清' or TRIM(sc.wm_taiwei1_name)='不详' or TRIM(sc.wm_taiwei1_name)='/' then '23'
else '' end) as "foetalPosition" --胎位
,(case when a.WM_FUZHONG_NAME='是' then '是' when a.WM_FUZHONG_NAME is not null then '否' else '' end) as "oedema" --浮肿
,(case when sc.wm_taixianlu1_name='足先露' or sc.wm_taixianlu2_name='足先露'
 or sc.wm_taixianlu3_name='足先露'
 or sc.wm_taixianlu4_name='足先露'
 then '是' else '' end) as "footPresentation" --足先露
,'' as "prsFactor" --妊娠风险筛查-阳性因素
,'' as "prsFactorOther" --妊娠风险筛查-阳性因素-补充
,org_shoucha.xj_org_name as "submitUnit" --填报单位
,(case when sc.check_doctor_name is not null then sc.check_doctor_name else b.check_user_name end) as "submitOperator" --填报人
,to_char(b.check_date,'YYYY-MM-DD') as "submitDate" --报告日期
,'select basic_id,"gradeCode","gradeUser","gradeDate" from V_UPLOAD_XINJIANG_FACTOR_SC where "riskStatus"=1 and "status"=1 and basic_id=' || a.id as "Array:highRiskFactors" --高危因素列表
from WCBJ_BASIC_INFO_TAB a
left join (select min(create_time) create_time_min,basic_id from WCBJ_FUZHU_CHECK_TAB where status=1 and XUEHONG_DANBAI is not null group by basic_id) b_min on a.id=b_min.basic_id
LEFT JOIN "WCBJ_FUZHU_CHECK_TAB" b ON  b.BASIC_ID = a.id and b.create_time=b_min.create_time_min and b.status=1 and b.xuehong_danbai is not null
left join WCBJ_SHOUCHA_RECORD_TAB sc on a.id=sc.basic_id and sc.status=1
left join (select min(check_date) check_date,basic_id,max(check_date) check_date_max from WCBJ_FUCHA_RECODE_TAB where status=1 group by basic_id) fucha_min on a.id=fucha_min.basic_id
left join (select basic_id,max("updateTime") update_time_max from V_UPLOAD_XINJIANG_FACTOR_SC group by basic_id) risk_max on sc.basic_id=risk_max.basic_id
inner JOIN UPLOAD_RECORD_XINJIANG R_BASIC ON R_BASIC."uploadType" = 'WM1002' AND R_BASIC."uploadRowId" = A.id and R_BASIC."outId" is not null
LEFT JOIN UPLOAD_RECORD_XINJIANG R ON R."uploadType" = 'WM2002' AND R."uploadRowId" = A.id
left join ZX_HOSPITAL_INFO_TAB org_shoucha on b.org_id=org_shoucha.org_id and org_shoucha.status=1
left join UPLOAD_DOCTOR_INFO doctor_info on sc.org_id = doctor_info.org_id and sc.CHECK_DOCTOR_NAME = doctor_info.DOCTOR_NAME and doctor_info.STATUS = 1
where 
 (greatest(nvl(a.update_time,to_date('1970-01-01 00:00:00','yyyy-MM-dd HH24:mi:ss')),nvl(b.update_time,to_date('1970-01-01 00:00:00','yyyy-MM-dd HH24:mi:ss')),nvl(sc.update_time,to_date('1970-01-01 00:00:00','yyyy-MM-dd HH24:mi:ss')))  > to_date(R."uploadTime",'YYYY-MM-DD hh24:mi:ss')
 or r."uploadTime" is null or INSTR(r."error",'无法连接到远程服务器')>0 or INSTR(r."error",'请求被中止: 操作超时。')>0
 -- 高危更新时间大于首查上传时间
  or risk_max.update_time_max>to_date(r."uploadTime",'YYYY-MM-DD HH24:MI:SS')
 )
and (b.create_time>=TO_DATE('2024-12-01', 'YYYY-MM-DD') or sc.create_time>=TO_DATE('2024-12-01','YYYY-MM-DD') or (fucha_min.check_date_max>=TO_DATE('2025-01-01','YYYY-MM-DD'))
	or a.wm_idno in ('65302120000906162X','653021199004062046','653001199303030428','653023199506040827','65302120000330116X','653024199706112086','653023199109010448','653022198904202325')
	)
-- and a.create_time>=TO_DATE('2025-01-01', 'YYYY-MM-DD')
and a.status=1 and b.status =1
and org_shoucha.license is not null
and (b.org_id!='100072' or (b.ORG_ID='100072' and CALEYUNWEEK_WEEK(a.WM_MENSTRUAL_LAST,sysdate) >= 13))
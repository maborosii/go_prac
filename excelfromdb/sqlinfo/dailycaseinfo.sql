SELECT
	date_format( crr.CREATE_DATE, '%Y-%m-%d' ) AS '评查时间',
CASE
		a.AREA_TYPE 
		WHEN 2 THEN
	NULL 
		WHEN 3 THEN
		a.AREA_NAME 
		WHEN 4 THEN
		b.AREA_NAME 
		WHEN 5 THEN
		c.AREA_NAME 
	END '市',
CASE
		a.AREA_TYPE 
		WHEN 2 THEN
	NULL 
		WHEN 3 THEN
	NULL 
		WHEN 4 THEN
		a.AREA_NAME 
		WHEN 5 THEN
		b.AREA_NAME 
	END '区县',
	ctd.dept_full_name AS '评查对象',
	cei.NAME AS '评查人',
CASE
		
		WHEN acc.CASE_LEVEL = 1 THEN
		'省级' 
		WHEN acc.CASE_LEVEL = 2 THEN
		'市级' 
		WHEN acc.CASE_LEVEL = 3 THEN
		'县（市、区）级' 
		WHEN acc.CASE_LEVEL = 4 THEN
		'乡镇（街道）级' 
	END AS '层级',
	acc.CASE_DEPT AS '案卷所属部门',
	acc.case_num AS '案号',
CASE
		
		WHEN acc.case_type = 1 THEN
		'行政处罚' 
		WHEN acc.case_type = 2 THEN
		'行政强制' 
		WHEN acc.case_type = 3 THEN
		'行政检查' 
		WHEN acc.case_type = 4 THEN
		'行政许可' 
	END AS '案卷类型',
	acc.PARTIES AS '当事人',
CASE
		
		WHEN crr.IS_QUALIFIED = 0 THEN
		'不合格' 
		WHEN crr.IS_QUALIFIED = 1 THEN
		'合格' 
	END AS '评查结果-合法性',
	crr.NORM_SCORE AS '评查结果-规范性',
CASE
		
		WHEN crr.GRADE = 1 THEN
		'优秀' 
		WHEN crr.GRADE = 2 THEN
		'良好' 
		WHEN crr.GRADE = 3 THEN
		'合格' 
		WHEN crr.GRADE = 4 THEN
		'不合格' 
	END AS '评定等次',
	acei.NAME AS '复核人员',
	date_format( crf.CREATE_DATE, '%Y-%m-%d' ) AS '复核时间',
CASE
		
		WHEN crf.IS_QUALIFIED = 0 THEN
		'不合格' 
		WHEN crf.IS_QUALIFIED = 1 THEN
		'合格' 
	END AS '复核结果-合法性',
	crf.NORM_SCORE AS '复核结果-规范性',
CASE
		
		WHEN crf.GRADE = 1 THEN
		'优秀' 
		WHEN crf.GRADE = 2 THEN
		'良好' 
		WHEN crf.GRADE = 3 THEN
		'合格' 
		WHEN crf.GRADE = 4 THEN
		'不合格' 
	END AS '复核结果-评定等次' 
FROM
	appr_case_choice acc
	INNER JOIN appr_case_review_task crk ON acc.task_id = crk.id
	INNER JOIN appr_case_task_dept ctd ON acc.dept_id = ctd.id
	INNER JOIN appr_case_review_record crr ON acc.id = crr.choice_id
	INNER JOIN appr_case_expert_info cei ON crr.expert_id = cei.id
	LEFT JOIN appr_case_review_form crf ON acc.id = crf.choice_id
	LEFT JOIN appr_case_expert_info acei ON crf.expert_id = acei.id
	LEFT JOIN c_sys_area a ON ctd.belong_area = a.CODE 
	AND a.del_flag = 0
	LEFT JOIN c_sys_area b ON b.CODE = a.parent_code 
	AND b.del_flag = 0
	LEFT JOIN c_sys_area c ON c.CODE = b.parent_code 
	AND c.del_flag = 0
	LEFT JOIN c_sys_area d ON d.CODE = c.parent_code 
	AND d.del_flag = 0 
WHERE
	crk.IS_DELETED = 0 
	AND crr.IS_DELETED = 0 
	AND crr.RECORD_STATUS != 0 
	AND crr.CREATE_DATE >= ? 
	AND crr.CREATE_DATE < ?
	AND crk.task_name LIKE '2020年度行政执法案卷评查%' 
ORDER BY
	cei.NAME,
	crr.CREATE_DATE
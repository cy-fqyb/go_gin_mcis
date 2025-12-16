SELECT DISTINCT
    CAST(l.lis_report_id AS VARCHAR(20)) + '_' + CAST(d.lis_detail_item_id AS VARCHAR(20)) AS sign,
    CAST(r.rowid AS VARCHAR(20))                                                          AS rypt_id,
    d.lis_detail_report_id,
    d.lis_detail_item_id,
    d.lis_detail_item_name,
    d.lis_detail_report_details,
    d.lis_detail_report_abnormal_sign,
    d.lis_detail_high_max_value,
    d.lis_detail_low_max_value,
    d.lis_detail_remark
FROM v_yyt_mz_lis_detail d
JOIN v_yyt_mz_lis l ON l.lis_report_id = d.lis_detail_report_id
JOIN gen_business_1570758835271_tab b ON l.lis_patient_id = b.patient_id
JOIN jj_upload_record r ON l.lis_report_id = r.rowid AND r.type = 'wcbj_mz_lis'
JOIN jj_upload_record s ON b.id = s.rowid AND s.type = 'wcbj_basic'
LEFT JOIN jj_upload_record rd 
       ON CAST(l.lis_report_id AS VARCHAR(20)) + '_' + CAST(r.rowid AS VARCHAR(20)) = CAST(rd.rowid AS VARCHAR(50))
      AND rd.type = 'wcbj_mz_lis_detail'
WHERE CAST(l.lis_report_date AS DATE) >= '2025-01-01'
  AND (ISNULL(r.lastSuccess, '') = '1' 
       OR CAST(l.lis_report_date AS DATE) >= CAST(r.uploadTime AS DATE))
  AND (r.id IS NULL 
       OR ISNULL(r.error, '') <> '该条记录已存在,无法进行任何操作');
//nomor 1a
SELECT
    COUNT(DISTINCT po.id) AS total_po,
    COUNT(DISTINCT cer.id) AS total_cer
FROM
    pr
    LEFT JOIN po ON pr.id = po.pr_id
    LEFT JOIN cer ON pr.id = cer.pr_id;

//nomor 1b
SELECT
    pr_id,
    CASE
        WHEN total_po = total_cer THEN 'Tally'
        ELSE 'Not Tally'
    END AS status
FROM
    (SELECT pr_id, SUM(quantity) AS total_po
    FROM po_line
    GROUP BY pr_id) as po_items
    FULL JOIN (SELECT pr_id, SUM(quantity) AS total_cer
    FROM cer_line
    GROUP BY pr_id) as cer_items
    ON po_items.pr_id = cer_items.pr_id;

//nomor 1c
SELECT
    pr_id,
    CASE
        WHEN COUNT(DISTINCT sequence, product_id) > 1 THEN 'Different'
        ELSE 'Same'
    END AS item_match_status
FROM
    (SELECT pr_id, sequence, product_id FROM (
        SELECT pr_id, sequence, product_id FROM po_line
        UNION ALL
        SELECT pr_id, sequence, product_id FROM cer_line
    )) as po_cer
GROUP BY
    pr_id;
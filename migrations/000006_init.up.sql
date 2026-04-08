CREATE OR REPLACE FUNCTION generate_item_mask()
RETURNS TRIGGER AS $$
DECLARE
    concatenated_mask TEXT;
BEGIN
    SELECT string_agg(
        pq.position || ':' || pqa.answer,
        '#' ORDER BY pq.position
    )
    INTO concatenated_mask
    FROM item_question_answers pqa
    JOIN item_questions pq
      ON pq.item_id = pqa.item_id
     AND pq.question_id = pqa.question_id
    WHERE pqa.item_id = NEW.item_id;

    INSERT INTO item_masks (
        item_id,
        mask,
        mask_hash,
        business_id,
        created_by,
        created_at
    )
    VALUES (
        NEW.item_id,
        concatenated_mask,
        substr(md5(concatenated_mask), 1, 8),
        'default_business',
        NEW.created_by,
        now()
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

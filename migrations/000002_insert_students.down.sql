BEGIN;

DELETE FROM 
    students 
WHERE 
    id IN (
        'bcff9f56-1ba6-4f92-9478-635c3f18e558',
        'dbf54856-9a98-4672-9c90-e9da71a1f893',
        '06b2ec25-3fe0-475e-9077-e77a113f4727',
        'e6e84c46-6ddf-4d9a-b27a-ddb74b4d63bb',
        '8e99273f-e566-4476-836e-048b1ecd9c4d'
    );

COMMIT;

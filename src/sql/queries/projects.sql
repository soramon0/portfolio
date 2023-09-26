-- name: ListProjects :many
SELECT
  p.id,
  p.client_name,
  p.name,
  p.description,
  p.live_link,
  p.code_link,
  p.start_date,
  p.end_date,
  p.created_at,
  p.updated_at,
  COALESCE(
    (
      SELECT 
        JSON_AGG(
          JSON_BUILD_OBJECT(
            'id', f.id,
            'url', f.url,
            'alt', f.alt,
            'name', f.name,
            'uploaded_at', f.uploaded_at,
            'type', f.type
          ) 
          ORDER BY f.id
        )
      FROM files AS f WHERE f.project_id = p.id
    )::json,
    '[]'::json
  ) AS gallery
FROM
  projects AS p
ORDER BY
  p.id;


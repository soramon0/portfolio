-- name: ListPublishedProjects :many
SELECT
  p.id,
  p.client_name,
  p.name,
  p.slug,
  p.subtitle,
  p.description,
  p.live_link,
  p.code_link,
  p.start_date,
  p.end_date,
  p.launch_date,
  p.created_at,
  p.updated_at,
  f.name as cover_image_name,
  f.url as cover_image_url,
  f.alt as cover_image_alt
FROM
  projects AS p
LEFT JOIN
  files as f
ON
  f.id = p.cover_image_id
WHERE
  p.published_at IS NOT NULL
ORDER BY
  p.id, p.created_at
LIMIT $1 OFFSET $2;

-- name: CountPublishedProjects :one
SELECT count(*) FROM projects
WHERE published_at IS NOT NULL;

-- name: GetPublishedProjectBySlug :one
SELECT
  p.id,
  p.client_name,
  p.name,
  p.slug,
  p.subtitle,
  p.description,
  p.live_link,
  p.code_link,
  p.start_date,
  p.end_date,
  p.launch_date,
  p.created_at,
  p.updated_at,
  f.name as cover_image_name,
  f.url as cover_image_url,
  f.alt as cover_image_alt,
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
LEFT JOIN
  files as f
ON
  f.id = p.cover_image_id
WHERE
  p.published_at IS NOT NULL AND slug = $1;

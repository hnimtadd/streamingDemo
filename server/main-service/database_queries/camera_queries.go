package dbqueries

const (
	CREATE_CAMERA = ` INSERT INTO video_analytics.cameras( camera_hls_streaming_endpoint, source_url)
		VALUES (?, ?)
	`

	GET_CAMERA = ` SELECT camera_id, camera_hls_streaming_endpoint, source_url
		FROM video_analytics.cameras WHERE 1
	`
	UPDATE_CAMERA = ` UPDATE video_analytics.cameras
		SET camera_id = ?, camera_hls_streaming_endpoint = ?, source_url = ?
		WHERE 1 `

	DELETE_CAMERA_BY_CAMERA_ID = ` DELETE FROM video_analytics.cameras WHERE camera_id = ?;`
)

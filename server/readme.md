Build Image:
	bash scripts/build-service.sh

Run Image:
	bash scripts/run-service.sh

apiEndpoint:
 GET 	http://localhost:10077/api/v1/event/get-cameras
 POST 	http://localhost:10077/api/v1/event/publish
    POST body:
    [
        {
        "camera_hls_streaming_endpoint": "sample",
        "source_url": "samplesource"
        }
    ]

example:
POST:
```
curl --location 'http://localhost:10077/api/v1/event/publish' \
--header 'Content-Type: application/json' \
--data '[
	{
	"camera_hls_streaming_endpoint": "sample",
	"source_url": "samplesource"
	}
]
'
```
GET:
```
curl --location 'http://localhost:10077/api/v1/event/get-cameras'
```
